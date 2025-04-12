# TableTap API

## Running Locally

Clone the project

```bash
git clone git@github.com:Tech-Notes/table-tap.api.git
```

Go to the project directory

```bash
cd table-tap.api
```

Remove remote origin

```bash
git remote remove origin
```

Install dependencies

```bash
go mod tidy
```

Copy .env.sample

```bash
cd server && cp .env.sample .env
```

Other steps
   - set up postgresql database
   - create a bucket on amazon s3
   - and update .env keys

Start the server

```bash
go run .
```

## Database tables migration guide
   First a all, you have to install [golang-migrate cli](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) in your system and *cd to table-tap.api*.

 **Install dotenv-cli**

    npm install -g dotenv-cli

#### Database migration is simply done by running **'make'** commands from cli.

 **migration up:**
     
     dotenv make migration_up 

 **migration down:**
 
     dotenv make migration_down

#### Or if your system can't use **'make'** command,

 **migration up:**
 
     migrate -path ./migrations/ -database "postgresql://username:password@localhost:5432/database_name?sslmode=disable" -verbose up

 **migration down:**
 
     migrate -path ./migrations/ -database "postgresql://username:password@localhost:5432/database_name?sslmode=disable" -verbose down

 
## Project structure

There are seven packages.

   ### 1. main ( folder:  server )
   *This package have main router setup and handlers for admin and kitchen dashboards.* 
     
   ### 2. db ( folder: db )
   *Database functions which is used by both main and shopper packages' handlers.*
     
   ### 3. shopper ( folder: shopper )
   *This package is for handling clients (shoppers)'s requests like placing orders.*
   
   ### 4. httphelper (folder: internal/httphelper )
   *http reader and writer methods for handlers.*

   ### 5. types ( folder: internal/types )
   *common types used by all packages.*
   
   ### 6. utils ( folder: internal/utils )
   *common methods for all packages like uploading photo to s3.*

   ### 7. notifications ( folder: notifications )
   *notifications system.*

### Other folders

   **1. migrations** <br>
   *Database migrations files*
      
     
