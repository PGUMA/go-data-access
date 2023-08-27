CREATE TABLE mydata (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  mail TEXT,
  age INTEGER
);

INSERT INTO mydata (name, mail, age) VALUES ('Taro', 'taro@yamada', 39);
INSERT INTO mydata (name, mail, age) VALUES ('Hanako', 'hanako@flower', 28);
INSERT INTO mydata (name, mail, age) VALUES ('Sachiko', 'sachiko@happy', 17);
INSERT INTO mydata (name, mail, age) VALUES ('Jiro', 'jiro@change', 6);
