CREATE TABLE tasks
(
    id          INT             NOT NULL AUTO_INCREMENT,
    user_id     INT             NOT NULL,
    title       VARCHAR(255)    NOT NULL,
    description TEXT            NOT NULL,
    status      VARCHAR(50)     DEFAULT 'pending',
    created_at  DATETIME  DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB;