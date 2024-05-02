#!/bin/sh

echo "Waiting for 10 seconds before starting..."
sleep 10

# Loop through each IP address obtained from DNS lookup
for IP_ADDRESS in $(dig +short +search elasticsearch); do
    echo "Checking Elasticsearch instance at $IP_ADDRESS..."
    # Check if the Elasticsearch instance is ready
    while true; do
        status_code=$(curl -s -o /dev/null -w "%{http_code}" "$IP_ADDRESS":9200)
        if [ "$status_code" -eq 200 ]; then
            break
        else
            echo "Elasticsearch instance at $IP_ADDRESS is not ready yet. Retrying in 10 seconds..."
            sleep 10
        fi
    done
    echo "Elasticsearch instance at $IP_ADDRESS is ready."

    # Insert data into the Elasticsearch instance
    echo "Inserting data into Elasticsearch instance at $IP_ADDRESS..."
    curl -X POST "$IP_ADDRESS:9200/_bulk" -H 'Content-Type: application/json' --data-binary @movies.json
    echo "Data insertion into Elasticsearch instance at $IP_ADDRESS complete."
done

echo "done"