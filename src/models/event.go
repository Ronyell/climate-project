package models

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/guregu/null"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// describe is a common behavior that Shape and its sub-classes
type EventDescriber interface {
	GetFieldAndTableName() (string, string)
	GetUnmarshalObject(requestBody []byte) error
	GetValueByFieldName(fieldName string) interface{}
	GetEvent() Event
}

func getValueByFieldName(obj any, fieldName string) interface{} {
	// Obtendo o valor do objeto usando reflection
	fieldName = cases.Title(language.BrazilianPortuguese, cases.NoLower).String(fieldName)
	v := reflect.ValueOf(obj)

	// Certificar que estamos lidando com um ponteiro para a struct (ou uma struct direta)
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Desreferenciar o ponteiro
	}

	field := reflect.Indirect(v).FieldByName(fieldName)
	return field.Interface()
}

// Represent the event
type Event struct {
	ID          uint64     `json:"id,omitempty"`
	City        City       `json:"city,omitempty"`
	EventType   string     `json:"eventType,omitempty" validate:"required,oneof=SECA INCENDIO CALOR FRIO INUNDACAO DESLIZAMENTO"`
	InitialDate time.Time  `json:"initialDate,omitempty"`
	FinalDate   *null.Time `json:"finalDate,omitempty"`
	CreatedAt   *null.Time `json:"createdAt,omitempty"`
}

func (e *Event) GetFieldAndTableName() (string, string) {
	return "", ""
}

func (e *Event) GetValueByFieldName(fieldName string) interface{} {
	return getValueByFieldName(e, fieldName)
}

func (e *Event) GetEvent() Event {
	return *e
}

// Represent the event dry
type EventDry struct {
	Event
	RelativeHumidity uint64 `json:"relativeHumidity,omitempty" validate:"required"`
}

func (e *EventDry) GetUnmarshalObject(requestBody []byte) error {
	if erro := json.Unmarshal(requestBody, &e); erro != nil {
		return erro
	}
	return nil
}

func (e *EventDry) GetFieldAndTableName() (string, string) {
	return "relativeHumidity", "eventsdry"
}

func (e *EventDry) GetValueByFieldName(fieldName string) interface{} {
	return getValueByFieldName(&e, fieldName)
}

func (e *EventDry) GetEvent() Event {
	return e.Event
}

// Represent the event burn
type EventBurn struct {
	Event
	IsConservationArea bool `json:"isConservationArea,omitempty" validate:"required"`
}

func (e *EventBurn) GetUnmarshalObject(requestBody []byte) error {
	if erro := json.Unmarshal(requestBody, &e); erro != nil {
		return erro
	}
	return nil
}

func (e *EventBurn) GetFieldAndTableName() (string, string) {
	return "isConservationArea", "eventsburn"
}

func (e *EventBurn) GetValueByFieldName(fieldName string) interface{} {
	return getValueByFieldName(&e, fieldName)
}

func (e *EventBurn) GetEvent() Event {
	return e.Event
}

// Represent the event hot
type EventHot struct {
	Event
	Temperature float32 `json:"temperature,omitempty" validate:"required"`
}

func (e *EventHot) GetUnmarshalObject(requestBody []byte) error {
	if erro := json.Unmarshal(requestBody, &e); erro != nil {
		return erro
	}
	return nil
}

func (e *EventHot) GetFieldAndTableName() (string, string) {
	return "temperature", "eventshot"
}

func (e *EventHot) GetValueByFieldName(fieldName string) interface{} {
	return getValueByFieldName(&e, fieldName)
}

func (e *EventHot) GetEvent() Event {
	return e.Event
}

// Represent the event cold
type EventCold struct {
	Event
	Temperature float32 `json:"temperature,omitempty" validate:"required"`
}

func (e *EventCold) GetUnmarshalObject(requestBody []byte) error {
	if erro := json.Unmarshal(requestBody, &e); erro != nil {
		return erro
	}
	return nil
}

func (e *EventCold) GetFieldAndTableName() (string, string) {
	return "temperature", "eventscold"
}

func (e *EventCold) GetValueByFieldName(fieldName string) interface{} {
	return getValueByFieldName(&e, fieldName)
}

func (e *EventCold) GetEvent() Event {
	return e.Event
}

// Represent the event flood
type EventFlood struct {
	Event
	RainPrecipitation float32 `json:"rainPrecipitation,omitempty" validate:"required"`
}

func (e *EventFlood) GetFieldAndTableName() (string, string) {
	return "rainPrecipitation", "eventsflood"
}

func (e *EventFlood) GetUnmarshalObject(requestBody []byte) error {
	if erro := json.Unmarshal(requestBody, &e); erro != nil {
		return erro
	}
	return nil
}

func (e *EventFlood) GetValueByFieldName(fieldName string) interface{} {
	return getValueByFieldName(&e, fieldName)
}

func (e *EventFlood) GetEvent() Event {
	return e.Event
}

// Represent the event slide
type EventSlide struct {
	Event
	HousesAffected uint64 `json:"housesAffected,omitempty" validate:"required"`
}

func (e *EventSlide) GetUnmarshalObject(requestBody []byte) error {
	if erro := json.Unmarshal(requestBody, &e); erro != nil {
		return erro
	}
	return nil
}

func (e *EventSlide) GetFieldAndTableName() (string, string) {
	return "housesAffected", "eventsslide"
}

func (e *EventSlide) GetValueByFieldName(fieldName string) interface{} {
	return getValueByFieldName(&e, fieldName)
}

func (e *EventSlide) GetEvent() Event {
	return e.Event
}
