# api-skeleton

## example post request
```
curl --location --request POST 'http://localhost:8080/example/createPizza' \
--header 'Content-Type: application/json' \
--data-raw '{
    "size": 12
}'
```