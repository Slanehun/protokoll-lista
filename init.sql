USE protokoll;

CREATE TABLE magyar (
                        ID INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
                        Name VARCHAR(255) NOT NULL,
                        Government VARCHAR(255) NOT NULL,
                        Event VARCHAR(255) NOT NULL,
                        Position VARCHAR(255) NOT NULL,
                        Address VARCHAR(255) NOT NULL,
                        Email VARCHAR(255) NOT NULL,
                        Phone VARCHAR(255) NOT NULL,
                        Comment VARCHAR(255) NOT NULL
);

LOAD DATA INFILE '/var/lib/mysql-files/data.csv'
    INTO TABLE magyar
    FIELDS TERMINATED BY ','
    ENCLOSED BY '"'
    LINES TERMINATED BY '\n'
    IGNORE 1 ROWS;

ALTER TABLE magyar MODIFY ID INT(11) NOT NULL AUTO_INCREMENT;