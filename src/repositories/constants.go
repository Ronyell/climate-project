package repositories

const (
	DRY   = "SECA"
	BURN  = "INCENDIO"
	HOT   = "CALOR"
	COLD  = "FRIO"
	FLOOD = "INUNDACAO"
	SLIDE = "DESLIZAMENTO"
)

var SQL_SELECT_EVENTS = map[string]string{
	"SECA":         "select events.eventId, events.eventType, events.eventInitialDate, events.eventFinalDate, eventsdry.relativeHumidity, cities.cityName, cities.cityUf from events inner join eventsdry on eventsdry.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"INCENDIO":     "select * from events inner join eventsburn on eventsburn.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"CALOR":        "select * from events inner join eventshot on eventshot.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"FRIO":         "select * from events inner join eventscold on eventscold.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"INUNDACAO":    "select * from events inner join eventsflood on eventsflood.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
	"DESLIZAMENTO": "select * from events inner join eventsslide on eventsslide.eventId = events.eventId inner join cities on cities.cityId = events.cityId",
}

var SQL_SELECT_EVENTS_FILTER_UF = map[string]string{
	"SECA":         "select * from events inner join eventsdry on eventsdry.eventId = events.eventId inner join cities on cities.cityId = events.cityId where events.uf = ?",
	"INCENDIO":     "select * from events inner join eventsburn on eventsburn.eventId = events.eventId inner join cities on cities.cityId = events.cityId where events.uf = ?",
	"CALOR":        "select * from events inner join eventshot on eventshot.eventId = events.eventId inner join cities on cities.cityId = events.cityId where events.uf = ?",
	"FRIO":         "select * from events inner join eventscold on eventscold.eventId = events.eventId inner join cities on cities.cityId = events.cityId where events.uf = ?",
	"INUNDACAO":    "select * from events inner join eventsflood on eventsflood.eventId = events.eventId inner join cities on cities.cityId = events.cityId where events.uf = ?",
	"DESLIZAMENTO": "select * from events inner join eventsslide on eventsslide.eventId = events.eventId inner join cities on cities.cityId = events.cityId where events.uf = ?",
}
