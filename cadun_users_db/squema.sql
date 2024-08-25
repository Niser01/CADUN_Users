USE cadun_users_db;

DROP TABLE IF EXISTS REQUEST;
DROP TABLE IF EXISTS REQUEST_TYPES;
DROP TABLE IF EXISTS USERS_ELEMENTS_FOR_QUOTATION;
DROP TABLE IF EXISTS USERS_PROFILE;

CREATE TABLE REQUEST_TYPES (
    id INT NOT NULL AUTO_INCREMENT,
    Status VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE USERS_PROFILE (
    id INT NOT NULL AUTO_INCREMENT,
    names VARCHAR(100) NOT NULL,
    lastNames VARCHAR(100) NOT NULL,
    alias VARCHAR(100) NOT NULL,
    password VARCHAR(64) NOT NULL,
    eMail VARCHAR(100) NOT NULL,
    phoneNumber VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE REQUEST (
    id INT NOT NULL AUTO_INCREMENT,
    idUser INT NOT NULL,
    request_status INT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (idUser) REFERENCES USERS_PROFILE(id),
    FOREIGN KEY (request_status) REFERENCES REQUEST_TYPES(id)
);

CREATE TABLE USERS_ELEMENTS_FOR_QUOTATION (
    id INT NOT NULL AUTO_INCREMENT,
    idUser INT NOT NULL,
    IAM_URL VARCHAR(255) NOT NULL,
    PDF_URL VARCHAR(255) NOT NULL,
    QUOTE_PDF_URL VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (idUser) REFERENCES USERS_PROFILE(id)
);

INSERT INTO REQUEST_TYPES (id, Status)
VALUES (1, "Finalizado"),
       (2, "Cotización"),
       (3, "Rechazado"),
       (4, "Aceptado");

INSERT INTO USERS_PROFILE (names, lastNames, alias, password, eMail, phoneNumber, country) 
VALUES ("Pepe", "Rodriguez", "Pepe123", SHA2('password123', 256), "pepe@unal.edu.co", "3150000000", "Colombia"),
       ("Juan", "Torres", "Juan123", SHA2('password456', 256), "juan@unal.edu.co", "3150000001", "Colombia"),
       ("Ana", "Cortez", "Ana123", SHA2('password789', 256), "ana@unal.edu.co", "3150000002", "Colombia");

INSERT INTO REQUEST (idUser, request_status)
VALUES (1, 2), -- Request para Pepe, tipo "Cotización"
       (2, 3), -- Request para Juan, tipo "Rechazado"
       (3, 1); -- Request para Ana, tipo "Finalizado"
