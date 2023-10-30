# golang-bank-backend

## Current Schema
![Alt text](/static/image.png)

## [3]
 
1. ```history | grep "docker run"``` is extremely useful to view docker history, apparently it does show that some of my docker images were removed somehow. Not sure what happened exactly
<br>

2. MakeFile is useful to run short hand commands like ```make migrateup```. it is actually short for 
   ```migrate -path db/migration/ -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up```

        