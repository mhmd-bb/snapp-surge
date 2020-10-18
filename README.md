# Surge
Increase price as demand increases.
# Features!

  - User management (create, update, login)
  - Rule management (threshold, request pairs)
  - Price coefficient calculation


### Tech



* [Go](https://golang.org/) -  an open source programming language that makes it easy to build simple, reliable, and efficient software.
* [Gin](https://github.com/gin-gonic/gin) - a web framework written in Go (Golang).
* [Gorm](https://gorm.io/) - The fantastic ORM library for Golang.
* [OpenStreetMap](https://openstreetmap.org/) - OpenStreetMap is the free wiki world map.
* [PostgreSQL](https://www.postgresql.org/) - PostgreSQL is a powerful, open source object-relational database system.
* [PostGIS](https://postgis.net/) - PostGIS is a spatial database extender for PostgreSQL.
* [Docker](https://www.docker.com/) - for containerization.
* [Logrus](https://github.com/sirupsen/logrus) - Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger.
* [JWT](https://jwt.io/) - JSON Web Tokens are an open, industry standard RFC 7519 method for representing claims securely between two parties.
* [Swagger](https://swagger.io/) - Simplify API development for users, teams, and enterprises with the Swagger open source and professional toolset.


### How to run

Surge requires [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/install/) to run.


```sh
$ git clone https://github.com/mhmd-bb/snapp-surge.git
$ cd snapp-surge
$ cp sample.env .env
$ cp ./osm-export/sample.env ./osm-export/.env
$ docker-compose up -d --build
```
## Note
If you have timeout problem while downloading project dependencies use a VPN. (Sanctions...)

## How to use
* After all services are up and running you can open **http://localhost:8080/docs/index.html** to see Swagger documentation.
* A Default user is created for you on application startup (you can change the user pass in .env file) so you can login.
* Don't forget to create rules (threshold, coefficient pairs).

You can easily use the Swagger Documentation; But I'm providing some curl commands as well:

### User
Login:
```sh
$ curl -X POST "http://localhost:8080/users/login" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{  \"password\": \"admin\",  \"username\": \"admin\"}"
```

Update Password:
```sh
$ curl -X PATCH "http://localhost:8080/users/password" -H  "accept: application/json" -H  "Authorization: Bearer {Your Token}" -H  "Content-Type: application/json" -d "{  \"password\": \"admin\"}"
```

Create New User:
```sh
$ curl -X POST "http://localhost:8080/users/register" -H  "accept: application/json" -H  "Authorization: Bearer {Your Token}" -H  "Content-Type: application/json" -d "{  \"password\": \"mohammad\",  \"username\": \"mohamamd\"}"
```

### Rule

Get All Rules:
```sh
$ curl -X GET "http://localhost:8080/rules" -H  "accept: application/json" -H  "Authorization: Bearer {Your Token}"
```

Create New Rule:
```sh
$ curl -X POST "http://localhost:8080/rules" -H  "accept: application/json" -H  "Authorization: Bearer {Your Token}" -H  "Content-Type: application/json" -d "{  \"coefficient\": 1.1,  \"threshold\": 10}"
```

Delete Rule:
```sh
$ curl -X DELETE "http://localhost:8080/rules/1" -H  "accept: application/json" -H  "Authorization: Bearer {Your Token}"
```

### And most importantly:
Ride Request which returns coefficient and increases district request counter:
```sh
$ curl -X POST "http://localhost:8080/surge/ride" -H  "accept: application/json" -H  "Content-Type: application/json" -d "{  \"lat\": 51.13199462890625,  \"lon\": 35.73425097869431}"
```




## How the project works

#### How the map data is collected and stored:
* [District Polygons](https://github.com/mhmd-bb/snapp-surge/blob/main/osm-export/districts.geojson) of Tehran city are exported using [overpass-turbo](https://overpass-turbo.eu) with geojson format and stored in [osm-export](https://github.com/mhmd-bb/snapp-surge/tree/main/osm-export) folder.
* As you start the services using docker-compose, the [gdal service](https://github.com/mhmd-bb/snapp-surge/blob/main/docker-compose.yml#L6-L17) exports the geojson file to PostGIS.

This way i can do spatial queries on the database and find the requested location's district.




#### How request counters are stored:

Instead of reseting the counter of a district every two hours and losing data on each reset i used the idea below:

The overal idea is that we have a moving window on small buckets (like a 2 hour Moving window on 10 minute Buckets).

Each Bucket has:
* counter
* expiration time (10 minutes for example)
* creation date
* district id

When a district's Bucket expires a new Bucket is created with counter value of 1.

The Moving window sums all Bucket counters of a district (in last 2 hours for example).
This way we can have count of requests in the last two hours with +-10 minute accuracy.





