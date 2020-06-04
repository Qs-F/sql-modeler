run:
	sqlite3 /tmp/models.sqlite < schema.sql
	SQLITE3_DBNAME="/tmp/models.sqlite" sqlboiler sqlite3 --output models --wipe 
	rm /tmp/models.sqlite

conv:
	mysql2sqlite schema.sql > schema2.sql
