# snippetbox
---
《Let’s Go: Learn to build professional web applications with Go》

https://lets-go.alexedwards.net/

Contents: https://lets-go.alexedwards.net/sample/00.01-contents.html

---
《Let's Go Further: Advanced patterns for building APIs and web applications in Go 》

https://lets-go-further.alexedwards.net/

Contents: https://lets-go-further.alexedwards.net/sample/00.01-contents.html


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

- 4.4 Designing a Database Model``

```bash
curl -iL -X POST http://localhost:4000/snippet/create
select id, title, created, expires from snippets;
```

http://localhost:4000/snippet?id=1


## Chapter 10

- 10.1 Generating a Self-Signed TLS Certificate

```bash
mkdir tls && cd tls
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

```

## Chapter 11 User Authentication

- 11.2 Creating a Users Model

```bash
USE snippetbox;

CREATE TABLE users(
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	hashed_password CHAR(60) NOT NULL,
	created DATETIME NOT NULL,
	active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
```

## Chapter 12 Request Context for Authentication/Authorization

```bash
USE snippetbox;
UPDATE users SET active = false WHERE email="able";
```



## Chapter 13 Testing

### 13.1 Unit Testing and Sub-Tests

```bash
go test -v ./cmd/web
```

### 13.2 Testing HTTP Handlers

Running Specific Tests
```bash
go test -v -run="^TestPing$" ./cmd/web

go test -v -run="^TestHumanData$/^UTC|CET$" ./cmd/web
```


### 13.3 End-To-End Testing

### 13.4.Mocking Dependencies

- Mocking the Database Models

### 13.5.Testing HTML Forms

You man get the same CSRF token because the test results have been cached.
If you want, you can force the test to run again by using the -count=1 flag like so:

```bash
go test -v -run "TestSignupUser" ./cmd/web -count=1
or
go clean -testcache
```

### 13.6.Integration Testing

- Test Database Setup and Teardown

```bash
CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'test_web'@'localhost' IDENTIFIED BY 'password';
GRANT CREATE, DROP,ALTER, INDEX,SELECT,INSERT,UPDATE,DELETE ON test_snippetbox.* TO 'test_web'@'localhost';

mkdir pkg/models/mysql/testdata
touch pkg/models/mysql/testdata/setup.sql
touch pkg/models/mysql/testdata/teardown.sql


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

CREATE TABLE users(
	id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL,
	hashed_password CHAR(60) NOT NULL,
	created DATETIME NOT NULL,
	active BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);

INSERT INTO users(name,email, hashed_password, created) VALUES (
	'Alice Jones',
	'alice@example.com',
	'$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
	'2018-12-23 17:25:22',
);


DROP TABLE users;
DROP TABLE snippets;

```

### 13.7 Profiling Test Coverage

```bash
go test -cover ./...

go test -coverprofile=profile.out ./...

go tool cover -html=profile.out

go test --covermode=count -coverprofile=profile.out ./...

go tool cover -html=profile.out

```

### Chapter 16.1 Add an About Page to the Application

### Chapter 16.2 Test the CreateSnippetForm Handler

```sh
go test -v -run=TestCreateSnippetForm ./cmd/web
```

### Chapter 16.3 Add a User Profile Page to the Application

### Chapter 16.4 Implement a Change Password Feature

### Chapter 16.5 Redirect User Appropriately after Login

### Chapter 16.6 Add a Debug Mode

```bash
go run ./cmd/web -debug
```