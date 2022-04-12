CREATE TABLE events
(
	id          serial       not null unique,
	title       varchar(255) not null,
	description varchar(255),
	date        timestamp    not null,
	done        boolean      not null default false
);
