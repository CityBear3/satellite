ALTER TABLE client
    CHANGE COLUMN secrets secret VARCHAR(255) NOT NULL;
ALTER TABLE device
    CHANGE COLUMN secrets secret VARCHAR(255) NOT NULL;