#!/usr/bin/env bash

source .env;

cd ./osm-export

chmod +x ./wait-for-it.sh;

./wait-for-it.sh postgres:5432 -s -t 200 -- ogr2ogr -f PostgreSQL PG:"host=postgres port=5432 user=${POSTGRES_USER} dbname=${POSTGRES_DB} password=${POSTGRES_PASSWORD}" -update -overwrite -select name,name:en -lco GEOMETRY_NAME=the_geom -lco FID=gid -nln districts districts.geojson