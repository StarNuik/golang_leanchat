-- Write your migrate up statements here
insert into channels
	(chan_id, chan_name)
values
	('00000000-0000-0000-0000-000000000000', '/dev/null'),
	('ffffffff-ffff-ffff-ffff-ffffffffffff', '/dev/max');

insert into messages
	(msg_id, chan_id, user_name, content)
values
	('00000000-0000-0000-0000-000000000000', '00000000-0000-0000-0000-000000000000', 'Leanchat', 'Welcome to Leanchat!'),
	('ffffffff-ffff-ffff-ffff-ffffffffffff', 'ffffffff-ffff-ffff-ffff-ffffffffffff', 'Leanchat', 'Welcome to Leanchat!');
---- create above / drop below ----

delete from messages
where
	msg_id = '00000000-0000-0000-0000-000000000000';

delete from channels
where
	chan_id = '00000000-0000-0000-0000-000000000000';

-- Write your migrate down statements here. If this migration is irreversible
-- Then delete the separator line above.
