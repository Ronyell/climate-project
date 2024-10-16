package repositories

import (
	"api/src/models"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
)

// Represent a city repository
type EventsRepository struct {
	db *sql.DB
}

// Get to return repository
func GetEventsRepository(db *sql.DB) *EventsRepository {
	return &EventsRepository{db}
}

// Create a event
func (eventsRepository EventsRepository) CreateEvent(requestBody []byte) (uint64, error) {
	var eventObject models.Event
	if erro := json.Unmarshal(requestBody, &eventObject); erro != nil {
		return 0, erro
	}

	eventType := eventObject.EventType

	switch eventType {
	case DRY:
		var eventDry models.EventDry
		return genericCreateEvent(&eventDry, requestBody, eventsRepository)
	case BURN:
		var eventBurn models.EventBurn
		return genericCreateEvent(&eventBurn, requestBody, eventsRepository)
	case HOT:
		var eventHot models.EventHot
		return genericCreateEvent(&eventHot, requestBody, eventsRepository)
	case COLD:
		var eventCold models.EventCold
		return genericCreateEvent(&eventCold, requestBody, eventsRepository)
	case FLOOD:
		var eventFlood models.EventFlood
		return genericCreateEvent(&eventFlood, requestBody, eventsRepository)
	case SLIDE:
		var eventSlide models.EventSlide
		return genericCreateEvent(&eventSlide, requestBody, eventsRepository)
	}

	return 0, errors.New("type not permited")
}

// Get all events by type
func (eventsRepository EventsRepository) GetEventByType(eventType string) (any, error) {
	if eventType == "" {
		var resultChannels []chan []models.EventDescriber
		var waitGroup sync.WaitGroup

		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, DRY, SQL_SELECT_EVENTS))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, BURN, SQL_SELECT_EVENTS))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, HOT, SQL_SELECT_EVENTS))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, COLD, SQL_SELECT_EVENTS))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, FLOOD, SQL_SELECT_EVENTS))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, SLIDE, SQL_SELECT_EVENTS))

		waitGroup.Wait()

		return eventsRepository.closeAndThreatingData(resultChannels)
	}

	var resultChannels chan []models.EventDescriber
	var waitGroup sync.WaitGroup
	resultChannels = eventsRepository.startGetGenericEvent(&waitGroup, eventType, SQL_SELECT_EVENTS)
	waitGroup.Wait()

	return <-resultChannels, nil

}

// Get all events by type and city uf
func (eventsRepository EventsRepository) GetEventByTypeAndUf(eventType string, cityUf string) (any, error) {
	if eventType == "" {
		var resultChannels []chan []models.EventDescriber
		var waitGroup sync.WaitGroup

		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, DRY, SQL_SELECT_EVENTS_FILTER_UF, cityUf))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, BURN, SQL_SELECT_EVENTS_FILTER_UF, cityUf))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, HOT, SQL_SELECT_EVENTS_FILTER_UF, cityUf))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, COLD, SQL_SELECT_EVENTS_FILTER_UF, cityUf))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, FLOOD, SQL_SELECT_EVENTS_FILTER_UF, cityUf))
		resultChannels = append(resultChannels, eventsRepository.startGetGenericEvent(&waitGroup, SLIDE, SQL_SELECT_EVENTS_FILTER_UF, cityUf))

		waitGroup.Wait()

		return eventsRepository.closeAndThreatingData(resultChannels)
	}

	var resultChannels chan []models.EventDescriber
	var waitGroup sync.WaitGroup
	resultChannels = eventsRepository.startGetGenericEvent(&waitGroup, eventType, SQL_SELECT_EVENTS_FILTER_UF, cityUf)
	waitGroup.Wait()

	return <-resultChannels, nil

}

func genericCreateEvent[T models.EventDescriber](eventEntry T, requestBody []byte, eventsRepository EventsRepository) (uint64, error) {
	eventEntry.GetUnmarshalObject(requestBody)
	fieldName, table := eventEntry.GetFieldAndTableName()
	value := eventEntry.GetValueByFieldName(fieldName)
	idEventDry, erro := eventsRepository.create(eventEntry.GetEvent(), fieldName, table, value)
	if erro != nil {
		return 0, erro
	}
	return idEventDry, nil
}

func (eventsRepository EventsRepository) create(eventObj models.Event, fieldName string, tableName string, value interface{}) (uint64, error) {

	// Start transaction
	tx, erro := eventsRepository.db.Begin()
	if erro != nil {
		return 0, erro
	}
	defer eventsRepository.db.Close()

	// Create generic event
	statement, erro := tx.Prepare(SQL_INSERT_EVENT)
	if erro != nil {
		return 0, nil
	}

	result, erro := statement.Exec(eventObj.EventType, eventObj.InitialDate, eventObj.FinalDate, eventObj.City.ID)
	if erro != nil {
		// Se ocorrer um erro, reverter a transação
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("erro ao reverter a transação: %v", rollbackErr)
		}
		return 0, erro
	}

	eventId, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	// Create specific event
	statement2, erro := tx.Prepare(fmt.Sprintf(SQL_INSERT_ESPECIFIC_EVENT, tableName, "eventId", fieldName))
	if erro != nil {
		return 0, erro
	}

	result2, erro := statement2.Exec(eventId, value)
	if erro != nil {
		// Se ocorrer um erro, reverter a transação
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("erro ao reverter a transação: %v", rollbackErr)
		}
		return 0, erro
	}

	lastID2, erro := result2.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	// Confirm transaction
	if err := tx.Commit(); err != nil {
		log.Fatalf("erro ao confirmar a transação: %v", err)
	}

	return uint64(lastID2), nil
}

func (eventsRepository EventsRepository) startGetGenericEvent(waitGroup *sync.WaitGroup, eventType string, mapSelect map[string]string, params ...string) chan []models.EventDescriber {
	results := make(chan []models.EventDescriber, 1)
	waitGroup.Add(1)
	go eventsRepository.getGenericEvent(waitGroup, results, eventType, mapSelect, params...)
	return results
}

func (eventsRepository EventsRepository) closeAndThreatingData(resultChanel []chan []models.EventDescriber) (any, error) {
	var finalResult []models.EventDescriber
	for _, rChan := range resultChanel {
		close(rChan)
		finalResult = append(finalResult, <-rChan...)
	}

	return finalResult, nil
}

func (eventsRepository EventsRepository) getGenericEvent(waitGroup *sync.WaitGroup, results chan<- []models.EventDescriber, eventType string, mapSelect map[string]string, params ...string) {
	var rows *sql.Rows
	var erro error

	defer waitGroup.Done()

	if params != nil {
		rows, erro = eventsRepository.db.Query(mapSelect[eventType], params[0])

	} else {
		rows, erro = eventsRepository.db.Query(mapSelect[eventType])
	}

	if erro != nil {
		results <- nil
		return
	}

	defer rows.Close()
	var eventsResp []models.EventDescriber

	for rows.Next() {
		var eventObj models.EventDescriber
		switch eventType {
		case DRY:
			eventObj, erro = eventsRepository.getEventsDrySQL(rows)
		case BURN:
			eventObj, erro = eventsRepository.getEventsBurnSQL(rows)
		case HOT:
			eventObj, erro = eventsRepository.getEventsHotSQL(rows)
		case COLD:
			eventObj, erro = eventsRepository.getEventsColdSQL(rows)
		case FLOOD:
			eventObj, erro = eventsRepository.getEventsFloodSQL(rows)
		case SLIDE:
			eventObj, erro = eventsRepository.getEventsSlideSQL(rows)
		}

		if erro != nil {
			results <- nil
		}
		eventsResp = append(eventsResp, eventObj)
	}
	results <- eventsResp
}

func (eventsRepository EventsRepository) getEventsDrySQL(rows *sql.Rows) (models.EventDescriber, error) {
	var eventObj models.EventDry

	if erro := rows.Scan(
		&eventObj.ID,
		&eventObj.EventType,
		&eventObj.InitialDate,
		&eventObj.FinalDate,
		&eventObj.RelativeHumidity,
		&eventObj.City.Name,
		&eventObj.City.UF,
	); erro != nil {
		return nil, erro
	}
	return &eventObj, nil
}

func (eventsRepository EventsRepository) getEventsBurnSQL(rows *sql.Rows) (models.EventDescriber, error) {
	var eventObj models.EventBurn

	if erro := rows.Scan(
		&eventObj.ID,
		&eventObj.EventType,
		&eventObj.InitialDate,
		&eventObj.FinalDate,
		&eventObj.IsConservationArea,
		&eventObj.City.Name,
		&eventObj.City.UF,
	); erro != nil {
		return nil, erro
	}
	return &eventObj, nil
}

func (eventsRepository EventsRepository) getEventsHotSQL(rows *sql.Rows) (models.EventDescriber, error) {
	var eventObj models.EventHot

	if erro := rows.Scan(
		&eventObj.ID,
		&eventObj.EventType,
		&eventObj.InitialDate,
		&eventObj.FinalDate,
		&eventObj.Temperature,
		&eventObj.City.Name,
		&eventObj.City.UF,
	); erro != nil {
		return nil, erro
	}
	return &eventObj, nil
}

func (eventsRepository EventsRepository) getEventsColdSQL(rows *sql.Rows) (models.EventDescriber, error) {
	var eventObj models.EventCold

	if erro := rows.Scan(
		&eventObj.ID,
		&eventObj.EventType,
		&eventObj.InitialDate,
		&eventObj.FinalDate,
		&eventObj.Temperature,
		&eventObj.City.Name,
		&eventObj.City.UF,
	); erro != nil {
		return nil, erro
	}
	return &eventObj, nil
}

func (eventsRepository EventsRepository) getEventsFloodSQL(rows *sql.Rows) (models.EventDescriber, error) {
	var eventObj models.EventFlood

	if erro := rows.Scan(
		&eventObj.ID,
		&eventObj.EventType,
		&eventObj.InitialDate,
		&eventObj.FinalDate,
		&eventObj.RainPrecipitation,
		&eventObj.City.Name,
		&eventObj.City.UF,
	); erro != nil {
		return nil, erro
	}
	return &eventObj, nil
}

func (eventsRepository EventsRepository) getEventsSlideSQL(rows *sql.Rows) (models.EventDescriber, error) {
	var eventObj models.EventSlide

	if erro := rows.Scan(
		&eventObj.ID,
		&eventObj.EventType,
		&eventObj.InitialDate,
		&eventObj.FinalDate,
		&eventObj.HousesAffected,
		&eventObj.City.Name,
		&eventObj.City.UF,
	); erro != nil {
		return nil, erro
	}
	return &eventObj, nil
}
