--postgresql

CREATE DATABASE IF NOT EXISTS camdeyeGO;

CREATE table users(
    id SERIAL,
    username text NOT NULL,     
    UNIQUE(email varchar(254) NOT NULL),   
    UNIQUE(token varchar(64) NOT NULL) ,    
    password varchar(60) NOT NULL, 
    public_key integer NOT NULL,   -- generate with the unix time or something random with crypt/rand.
);
create table camera(
    id SERIAL,
    name varchar(254) NOT NULL,       -- crypt/rand.
    owner_id integer not null,        -- you are the one that create a new camera id.
    public_key integer not null,      -- generate a new one when you log in with crypt/rand.
    UNIQUE(token varchar(64) not null) ,       -- generate a new one when you log  .
    private_key varchar(60) not null, -- sha256(password+public_key).
    -- how did you generate all this stuff, well in the case that you want to use it
    -- the only thing that you need to use is log in, in our page and then go to the part of camera
    -- and then click in "new camera", and the it will appear you a token that you need to put in the node administration page
    -- "http://camdeye.local".

);