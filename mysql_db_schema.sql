-- DROP DATABASE LearningResourcesTracker;
CREATE DATABASE IF NOT EXISTS LearningResourcesTracker;

USE LearningResourcesTracker;

CREATE TABLE IF NOT EXISTS State(
	state_id INT AUTO_INCREMENT PRIMARY KEY,
    label VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS Category(
	category_id INT AUTO_INCREMENT PRIMARY KEY,
    label VARCHAR(100) NOT NULL,
    state_id INT NOT NULL,
    parent_id INT,
    FOREIGN KEY (state_id)
		REFERENCES State(state_id)
			ON UPDATE RESTRICT
            ON DELETE RESTRICT,
	FOREIGN KEY (parent_id)
		REFERENCES Category(category_id)
			ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS Priority(
	priority_id INT AUTO_INCREMENT PRIMARY KEY,
    label VARCHAR(100) NOT NULL,
    worth INT NOT NULL
);

CREATE TABLE IF NOT EXISTS Topic(
	topic_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    summary TEXT,
    category_id INT,
    state_id INT NOT NULL,
    priority_id INT NOT NULL,
    FOREIGN KEY (category_id)
		REFERENCES Category(category_id)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
	FOREIGN KEY (state_id)
		REFERENCES State(state_id)
			ON UPDATE RESTRICT
			ON DELETE RESTRICT,
	FOREIGN KEY (priority_id)
		REFERENCES Priority(priority_id)
			ON UPDATE RESTRICT
			ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS ResourceItem(
	resource_item_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    URL VARCHAR(2083) NOT NULL, -- According to SO, this is a good number
    summary TEXT,
    topic_id INT NOT NULL,
    state_id INT NOT NULL,
    FOREIGN KEY (topic_id)
		REFERENCES Topic(topic_id)
			ON UPDATE CASCADE
            ON DELETE CASCADE,
	FOREIGN KEY (state_id)
		REFERENCES State(state_id)
			ON UPDATE RESTRICT
            ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS Tag(
	tag_id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS CategoryTag(
	category_id INT,
    tag_id INT,
    PRIMARY KEY (category_id, tag_id),
    FOREIGN KEY (category_id)
		REFERENCES Category(category_id)
			ON UPDATE CASCADE
            ON DELETE CASCADE,
	FOREIGN KEY (tag_id)
		REFERENCES Tag(tag_id)
			ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS TopicTag(
	topic_id INT,
    tag_id INT,
    PRIMARY KEY (topic_id, tag_id),
    FOREIGN KEY (topic_id)
		REFERENCES Topic(topic_id)
			ON UPDATE CASCADE
            ON DELETE CASCADE,
	FOREIGN KEY (tag_id)
		REFERENCES Tag(tag_id)
			ON UPDATE CASCADE
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS ResourceTag(
	resource_item_id INT,
    tag_id INT,
    PRIMARY KEY (resource_item_id, tag_id),
    FOREIGN KEY (resource_item_id)
		REFERENCES ResourceItem(resource_item_id)
			ON UPDATE CASCADE
            ON DELETE CASCADE,
	FOREIGN KEY (tag_id)
		REFERENCES Tag(tag_id)
			ON UPDATE CASCADE
            ON DELETE CASCADE
);
