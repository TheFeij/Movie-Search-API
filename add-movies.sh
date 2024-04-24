#!/bin/sh

curl -X POST "http://es8:9200/_bulk" -H 'Content-Type: application/json' --data-binary @movies.json