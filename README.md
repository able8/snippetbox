# snippetbox

《Let’s Go: Learn to build professional web applications with Go》


## Chapter 2

http://localhost:4000/snippet?id=1

```
curl -i -X POST http://localhost:4000/snippet/create
curl -i -X PUT http://localhost:4000/snippet/create

mkdir -p cmd/web pkg ui/html ui/static

go run ./cmd/web

curl https://www.alexedwards.net/static/sb.v130.tar.gz | tar -xvz -C ./ui/static
```

> Range requests are fully supported. This is great if your application is servinglarge files and you want to support resumable downloads. You can see thisfunctionality in action if you use curl to request bytes 100-199 of the logo.png file,like so:

curl -i -H "Range: bytes=100-199" --output - http://localhost:4000/static/img/logo.png

## Chapter 3 Managing Configuration Settings

- Custom loggers created by log.New() are concurrency-safe. You can share asingle logger and use it across multiple goroutines and in your handlers withoutneeding to worry about race conditions.



## Chapter 4 Setting Up MySQL

https://blog.csdn.net/w605283073/article/details/80417866

```bash
brew install mysql
mysql.server start 
mysql_secure_installation
mysql -u root -p

CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE snippetbox;
# Create a snippet table.
CREATE TABLE snippets (
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(100) NOT NULL,
	content TEXT NOT NULL,
	created DATETIME NOT NULL,
	expires DATETIME NOT NULL
);
# Add an index on the created column.
CREATE INDEX idx_snippets_created ON snippets(created);

# Add some dummy records(which we'll use in the next couple of chapers).
INSERT INTO snippets (title, content, created, expires) VALUES (
	'An old silent pond',
	'An old silent pnd...\nA frog jumps into the pond,\nsplash! Silence again.\n\n- Matsuo Basho',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
	'Over the wintry forest',
	'AOver the wintry\forest, winds howl in rage\n with no leaves to blow.\n\n- Natsume Soseki',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
	'First autumn morning',
	'First autumn morning\n the mirror I stare into \n shows my father''s face.\n\n- Murakami Kijo',
	UTC_TIMESTAMP(),
	DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

# Creating a new user
CREATE USER 'web'@'localhost' IDENTIFIED BY 'password';
GRANT SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';

# Test the new user
mysql -D snippetbox -u web -p
SELECT id, title, created, expires FROM snippets;
DROP TABLE snippets;
```

- 4.2 Installing a Database Driver

```bash
go get github.com/go-sql-driver/mysql@v1
```

- 4.4 Designing a Database Model
