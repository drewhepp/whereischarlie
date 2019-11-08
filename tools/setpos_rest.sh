#!/bin/bash

curl -i -X POST -H 'Content-Type: application/json' -d '{"lat": '"$1"',"lng": '"$2"'}' http://whereischarlie.org/position
