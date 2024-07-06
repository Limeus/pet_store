CREATE TABLE pets (
                      id SERIAL PRIMARY KEY,
                      age INT NOT NULL,
                      type VARCHAR(50) NOT NULL,
                      price NUMERIC(10, 2) NOT NULL,
                      sold BOOLEAN DEFAULT FALSE
);

-- Дополнительные таблицы и настройки, если необходимо
