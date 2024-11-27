CREATE TABLE authors (
    id SERIAL PRIMARY KEY ,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
	password BYTEA NOT NULL
);

CREATE TABLE categories (
    id SERIAL PRIMARY KEY ,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE OR REPLACE FUNCTION update_updated_at() --для изменений в комментах и новостях
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$
 LANGUAGE plpgsql;


CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    news_content TEXT NOT NULL,
    author_id INT,
    category_id INT,
    published_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, --работает автоматически при вставке
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, --работает автоматически при вставке и обновлении
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE news_comments (
    id SERIAL PRIMARY KEY,
    news_id INT,
    author_id INT,
    comment_content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (news_id) REFERENCES news(id),
	FOREIGN KEY (author_id) REFERENCES authors(id)
);

CREATE TRIGGER set_timestamp_news
BEFORE UPDATE ON news
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

CREATE TRIGGER set_timestamp_comments
BEFORE UPDATE ON news_comments
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

INSERT INTO authors (name, email, password) VALUES
('Иван Иванов', 'ivan.ivanov@example.com', 'password1'),
('Мария Петрова', 'maria.petrova@example.com', 'password2'),
('Сергей Симонов', 'sergey.simonov@example.com', 'password3');

INSERT INTO categories (name) VALUES
('Технологии'),
('Спорт'),
('Наука'),
('Культура');

INSERT INTO news (title, news_content, author_id, category_id) VALUES
('Новые технологии в 2024 году', 'Обзор новых технологий, которые изменят мир.', 1, 1),
('Спортивные достижения России', 'Россия завоевала медали на международных соревнованиях.', 2, 2),
('Научные открытия в области медицины', 'Недавние открытия в медицине, которые могут спасти жизни.', 3, 3);

INSERT INTO news_comments (news_id, author_id, comment_content) VALUES
(1, 2, 'Отличная статья! Жду новых технологий.'),
(1, 3, 'Интересно, как это повлияет на нашу жизнь.'),
(2, 1, 'Поздравляю спортсменов с победой!'),
(3, 2, 'Научные исследования — это очень важно для нашего будущего.');