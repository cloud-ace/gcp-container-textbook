 * create
```sql
CREATE TABLE mydb.chatlog (
id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
name VARCHAR(16) NOT NULL,
message VARCHAR(256) NOT NULL
);
```

 * insert
```sql
INSERT chatlog VALUES(0,NULL,"0delta","hello")
```