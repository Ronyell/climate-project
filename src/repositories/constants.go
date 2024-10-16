package repositories

const (
	DRY   = "SECA"
	BURN  = "INCENDIO"
	HOT   = "CALOR"
	COLD  = "FRIO"
	FLOOD = "INUNDACAO"
	SLIDE = "DESLIZAMENTO"
)

const SQL_INSERT_EVENT = "insert into events (eventType, eventInitialDate, eventFinalDate, cityId) values (?, ?, ?, ?)"
const SQL_INSERT_ESPECIFIC_EVENT = "insert into %s (%s, %s) values (?, ?)"

var SQL_SELECT_EVENTS = map[string]string{
	"SECA":         "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsdry.relativeHumidity, cities.cityName, cities.cityUf, cities.cityId from events inner join eventsdry on eventsdry.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"INCENDIO":     "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsburn.isConservationArea, cities.cityName, cities.cityUf, cities.cityId from events inner join eventsburn on eventsburn.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"CALOR":        "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventshot.temperature, cities.cityName, cities.cityUf, cities.cityId from events inner join eventshot on eventshot.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"FRIO":         "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventscold.temperature, cities.cityName, cities.cityUf, cities.cityId from events inner join eventscold on eventscold.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"INUNDACAO":    "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsflood.rainPrecipitation, cities.cityName, cities.cityUf, cities.cityId from events inner join eventsflood on eventsflood.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"DESLIZAMENTO": "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsslide.housesAffected, cities.cityName, cities.cityUf, cities.cityId from events inner join eventsslide on eventsslide.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
}

var SQL_SELECT_EVENTS_FILTER_UF = map[string]string{
	"SECA":         "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsdry.relativeHumidity, cities.cityName, cities.cityUf, cities.cityId from events inner join eventsdry on eventsdry.eventId = events.eventId inner join cities on cities.cityId = events.cityId where cities.cityUf = ?",
	"INCENDIO":     "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsburn.isConservationArea, cities.cityName, cities.cityUf, cities.cityId from events inner join eventsburn on eventsburn.eventId = events.eventId inner join cities on cities.cityId = events.cityId where cities.cityUf = ?",
	"CALOR":        "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventshot.temperature, cities.cityName, cities.cityUf, cities.cityId from events inner join eventshot on eventshot.eventId = events.eventId inner join cities on cities.cityId = events.cityId where cities.cityUf = ?",
	"FRIO":         "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventscold.temperature, cities.cityName, cities.cityUf, cities.cityId from events inner join eventscold on eventscold.eventId = events.eventId inner join cities on cities.cityId = events.cityId where cities.cityUf = ?",
	"INUNDACAO":    "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsflood.rainPrecipitation, cities.cityName, cities.cityUf, cities.cityId from events inner join eventsflood on eventsflood.eventId = events.eventId inner join cities on cities.cityId = events.cityId where cities.cityUf = ?",
	"DESLIZAMENTO": "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsslide.housesAffected, cities.cityName, cities.cityUf, cities.cityId from events inner join eventsslide on eventsslide.eventId = events.eventId inner join cities on cities.cityId = events.cityId where cities.cityUf = ?",
}
