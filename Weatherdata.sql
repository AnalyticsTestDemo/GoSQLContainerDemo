Create Database WeatherDB; 

CREATE TABLE CountryMaster
(
    CountryID bigint ,
    CountryName varchar(100)    
);

INSERT INTO CountryMaster(CountryID, CountryName) VALUES
 (1, 'USA'),
 (2, 'Canada');
