-- Set params 
set session my.number_of_albums = '5';
set session my.number_of_artists = '10';

-- Filling of album table
INSERT INTO album
select id, 
	concat('Album ', id),  
	floor(random() * (current_setting('my.number_of_artists')::int) + 1)::int,
	round(CAST(float8 (random() * 10000) as numeric), 3)
FROM GENERATE_SERIES(1, current_setting('my.number_of_albums')::int) as id;
