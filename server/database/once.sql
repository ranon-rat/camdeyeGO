--postgresql

CREATE DATABASE IF NOT EXISTS camdeyeGO;




CREATE table users(
    id SERIAL,
    username text NOT NULL,     
    gmail varchar(254) NOT NULL,--pure
    token varchar(64) NOT NULL,--  sha256(public_key+username+password) then i hash it again 
    password varchar(64) NOT NULL,-- sha256(password+public_key)
    public_key integer NOT NULL,-- generate with the unix time or something random with cryp/rand
);