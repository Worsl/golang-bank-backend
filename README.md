# golang-bank-backend

## Current Schema
![Alt text](/static/image.png)


## [Before you begin]
1. ```docker start postrgres12```,  ```make createdb```, ```make migrateup``` if u did not done so previously

## [3]
 
1. ```history | grep "docker run"``` is extremely useful to view docker history, apparently it does show that some of my docker images were removed somehow. Not sure what happened exactly
<br>

2. MakeFile is useful to run short hand commands like ```make migrateup```. it is actually short for 
   ```migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up```

        
## [4]
1. SQLC is recommended for GOlang CRUD operations.

## [5]
1. Convention for unit testing in GoLang : testfile is in the same folder as the actual code.
<br>

2. Convention for unit testing in Golang : main_test is the entry point of all unit tests inside a specific golang package