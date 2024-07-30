# profile-test-service
A test server for managing user profiles.

### Setup database:
Run docker container
`docker run --name local-mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=testdb -p 3306:3306 -d mysql:9.0.1`

Create schema:
`go build -o migrate.exe`

`.\migrate.exe --dbUser=root --dbPass=root --dbHost=localhost --dbPort=3306 --sqlFile=./migrations.sql`


### Run application:
`go build  -o .\bin\mytestapi.exe`

`.\bin\mytestapi.exe --port=8080 --dbUser=root --dbPass=root --dbHost=localhost --dbPort=3306`