USE LearningResourceTracker;

INSERT INTO State (label)
VALUES
	('Unchecked'),
    ('Checked'),
    ('Deleted'),
    ('Category')
;

INSERT INTO Category (label, state_id, parent_id)
VALUES
	('Computer Science', 4, null),
    ('Yiddish', 4, null)
;

INSERT INTO Category (label, state_id, parent_id)
VALUES
	('Golang', 4, 1),
    ('Linux', 4, 1)
;

INSERT INTO Category (label, state_id, parent_id)
VALUES
	('GraphQL', 4, 3)
;

INSERT INTO Priority (label, worth)
VALUES
	('Low', 1),
    ('Medium', 2),
    ('High', 3)
;

INSERT INTO Topic (title, summary, category_id, state_id, priority_id)
VALUES
	('SQL Driver Best Practices', null, 3, 1, 2),
    ('GraphQL', 'Is this even feasible in go at the moment???', 3, 1, 3),
    ('apt vs apt-get', 'Is there a difference or are they aliases?', 4, 2, 1),
    ('Plural of alias', 'Sometimes, I think that I can\'t actually speak English...', null, 1, 1),
    ('Resolvers', 'How do resolvers work?', 5, 1, 3)
;

INSERT INTO ResourceItem (title, URL, summary, topic_id, state_id)
VALUES
	('Thomas\' Reddit Thread', 'https://www.reddit.com/r/golang/comments/fnygei/golang_and_sqlite_when_to_open_connection/', 'Thomas sent me this reddit thread which might help', 1, 2),
    ('Official Godoc', 'https://golang.org/pkg/database/sql/#DB.Query', null, 1, 1),
    ('Tutorial Edge', 'https://tutorialedge.net/golang/golang-mysql-tutorial/#performing-basic-sql-commands', 'Scottish gopher could help', 1, 2),
    ('GraphQL Site', 'https://graphql.org/code/#go', 'There are a list of implementations to look at here', 2, 1),
    ('The actual dictionary', 'https://www.dictionary.com/browse/alias#:~:text=noun%2C%20plural%20a%C2%B7li%C2%B7,is%20an%20alias%20for%20Simpson.', 'I was correct', 4, 2)
;

INSERT INTO Tag (title)
VALUES
	('Interesting'),
    ('Fun')
;

INSERT INTO CategoryTag (category_id, tag_id)
VALUES
	(2, 1),
    (4, 1),
    (2, 2)
;

INSERT INTO TopicTag (topic_id, tag_id)
VALUES
	(2, 1),
    (3, 1),
    (4, 2)
;

INSERT INTO ResourceTag (resource_item_id, tag_id)
VALUES
	(3, 1),
    (4, 1),
    (4, 2)
;
