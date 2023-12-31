CREATE DATABASE OPERATIONS;

CREATE TABLE ASSIGNEE(
    ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    NAME VARCHAR(50) NOT NULL,
    REGISTERED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE TASK_STATUS(
    ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    STATUS VARCHAR(30) NOT NULL
);

CREATE TABLE TASK(
    ID INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    TASK_NAME VARCHAR(50) NOT NULL,
    DESCRIPTION VARCHAR(300) NULL,
    ASSIGNEE_ID INT NOT NULL,
    STATUS_ID INT NOT NULL,
    DUE_DATE DATE NULL,
    CREATED_AT TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UPDATED_AT TIMESTAMP NULL,
    FOREIGN KEY (ASSIGNEE_ID) REFERENCES ASSIGNEE(ID),
    FOREIGN KEY (STATUS_ID) REFERENCES TASK_STATUS(ID)
);