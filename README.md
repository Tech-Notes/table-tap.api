# TableTap API

## Database tables migration guide
   First a all, you have to install [golang-migrate cli](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) in your system and *cd to table-tap.api *.

## Install dotenv-cli

    npm install -g dotenv-cli

### Database migration is simply done by running **'make'** command from cli.

 **migration up:**
     
     dotenv make migration_up 

 **migration down:**
 
     dotenv make migration_down

### Or if your system can't use **'make'** command,

 **migration up:**
 
     migrate -path ./migrations/ -database "postgresql://username:password@localhost:5432/database_name?sslmode=disable" -verbose up

 **migration down:
 
     migrate -path ./migrations/ -database "postgresql://username:password@localhost:5432/database_name?sslmode=disable" -verbose down
 
