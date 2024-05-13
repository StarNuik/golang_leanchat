create table channels (
	chan_id uuid default gen_random_uuid() primary key,
	chan_name varchar(63) not null,
	chan_created timestamp default current_timestamp
);

create table messages (
	msg_id uuid default gen_random_uuid() primary key,
	chan_id uuid,
	user_name varchar(63) not null,
	content text,
	msg_created timestamp default current_timestamp
);

---- create above / drop below ----

drop table channels;
drop table messages;
