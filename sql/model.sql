CREATE DATABASE IF NOT EXISTS climate;
USE climate;

DROP TABLE IF EXISTS cities;

CREATE TABLE cities(
    cityId int auto_increment primary key,
    cityName varchar(100) not null,
    cityUf varchar(2) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;


CREATE TABLE events(
    eventId int auto_increment primary key,
    eventType varchar(100) not null,
    eventInitialDate timestamp not null,
    eventFinalDate timestamp not null,
    cityId int not null,
    createdAt timestamp default current_timestamp(),
    FOREIGN KEY (cityId) REFERENCES cities(cityId)
) ENGINE=INNODB;


CREATE TABLE eventsdry(
    eventDryId int auto_increment primary key,
    relativeHumidity int not null,
    eventId int not null,
    FOREIGN KEY (eventId) REFERENCES events(eventId) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE eventsburn(
    eventBurnId int auto_increment primary key,
    isConservationArea BOOLEAN not null,
    eventId int not null,
    FOREIGN KEY (eventId) REFERENCES events(eventId) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE eventshot(
    eventHotId int auto_increment primary key,
    temperature float not null,
    eventId int not null,
    FOREIGN KEY (eventId) REFERENCES events(eventId) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE eventscold(
    eventColdId int auto_increment primary key,
    temperature float not null,
    eventId int not null,
    FOREIGN KEY (eventId) REFERENCES events(eventId) ON DELETE CASCADE
) ENGINE=INNODB;


CREATE TABLE eventsflood(
    eventFloodId int auto_increment primary key,
    precipitation float not null,
    eventId int not null,
    FOREIGN KEY (eventId) REFERENCES events(eventId) ON DELETE CASCADE
) ENGINE=INNODB;

CREATE TABLE eventsslide(
    eventSlideId int auto_increment primary key,
    housesAffected int not null,
    eventId int not null,
    FOREIGN KEY (eventId) REFERENCES events(eventId) ON DELETE CASCADE
) ENGINE=INNODB;