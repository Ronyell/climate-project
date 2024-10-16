package repositories

import (
	"api/src/models"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
	statement, erro := tx.Prepare("insert into events (eventType, eventInitialDate, eventFinalDate, cityId) values (?, ?, ?, ?)")
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
	statement2, erro := tx.Prepare(fmt.Sprintf("insert into %s (%s, %s) values (?, ?)", tableName, "eventId", fieldName))
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

// Get all events
func (eventsRepository EventsRepository) GetEventByType(eventType string) (any, error) {
	switch eventType {
	case DRY:
		return eventsRepository.getDryEvents()
		// case "INCENDIO":
		// 	var eventBurn models.EventBurn
		// 	return genericCreateEvent(&eventBurn, requestBody, eventsRepository)
		// case "CALOR":
		// 	var eventHot models.EventHot
		// 	return genericCreateEvent(&eventHot, requestBody, eventsRepository)
		// case "FRIO":
		// 	var eventCold models.EventCold
		// 	return genericCreateEvent(&eventCold, requestBody, eventsRepository)
		// case "INUNDACAO":
		// 	var eventFlood models.EventFlood
		// 	return genericCreateEvent(&eventFlood, requestBody, eventsRepository)
		// case "DESLIZAMENTO":
		// 	var eventSlide models.EventSlide
		// 	return genericCreateEvent(&eventSlide, requestBody, eventsRepository)
	}
	return eventsRepository.getDryEvents()
}

func (eventsRepository EventsRepository) getDryEvents() ([]models.EventDry, error) {
	rows, erro := eventsRepository.db.Query(SQL_SELECT_EVENTS[DRY])
	if erro != nil {
		return nil, erro
	}

	defer rows.Close()
	var eventsDry []models.EventDry

	for rows.Next() {
		var eventDry models.EventDry

		if erro = rows.Scan(
			&eventDry.ID,
			&eventDry.EventType,
			&eventDry.InitialDate,
			&eventDry.FinalDate,
			&eventDry.RelativeHumidity,
			&eventDry.City.Name,
			&eventDry.City.UF,
		); erro != nil {
			return nil, erro
		}
		eventsDry = append(eventsDry, eventDry)
	}
	return eventsDry, nil
}

func genericGetEvents[T models.EventDescriber](eventEntries []T, eventsRepository EventsRepository) error {

	return nil
}

//Get city by id
// func (eventsRepository EventsRepository) GetEventById(id uint64) (models.Event, error) {
// 	rows, erro := eventsRepository.db.Query("select * from events where cityId = ?", id)
// 	if erro != nil {
// 		return models.Event{
// 			ID:        0,
// 			Name:      "",
// 			UF:        "",
// 			CreatedAt: time.Time{},
// 		}, erro
// 	}

// 	defer rows.Close()
// 	var eventObj models.Event

// 	for rows.Next() {
// 		if erro = rows.Scan(
// 			&eventObj.ID,
// 			&eventObj.Name,
// 			&eventObj.UF,
// 			&eventObj.CreatedAt,
// 		); erro != nil {
// 			return models.Event{
// 				ID:        0,
// 				Name:      "",
// 				UF:        "",
// 				CreatedAt: time.Time{},
// 			}, erro
// 		}
// 	}
// 	return eventObj, nil
// }

// func (eventsRepository EventsRepository) UpdateEventById(id uint64, eventObj models.Event) error {

// 	statement, erro := eventsRepository.db.Prepare("update  events set cityName = ?, cityUf = ? where cityId = ?")
// 	if erro != nil {
// 		return erro
// 	}
// 	defer statement.Close()

// 	if _, erro = statement.Exec(eventObj.Name, eventObj.UF, id); erro != nil {
// 		return erro
// 	}
// 	return nil
// }
