--in case that you want to  setup the database , execute this command
--cat once.sql |sqlite3 database.db

CREATE TABLE cameras(
    domain TEXT NOT NULL UNIQUE,    
    last_time_up integer, -- unix time
)

