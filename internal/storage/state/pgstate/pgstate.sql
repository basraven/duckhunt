CREATE TABLE parties
(
    pId SERIAL  NOT NULL,
    uId INTEGER NOT NULL,
    timestamp TIMESTAMP  default current_timestamp,
    CONSTRAINT parties_pkey PRIMARY KEY (uId, pId)
)
WITH (OIDS=FALSE);


CREATE TABLE parties_name
(
    pId INTEGER NOT NULL,
    pName VARCHAR(140) NOT NULL,
    CONSTRAINT parties_name_pkey PRIMARY KEY (pId)
)
WITH (OIDS=FALSE);



INSERT INTO parties (pId, uId)
VALUES (1, 8);
INSERT INTO parties_name (pId, pName)
VALUES (1, 'Init Party');

INSERT INTO parties (pId, uId)
VALUES (2, 8);
INSERT INTO parties (pId, uId)
VALUES (5, 8);

