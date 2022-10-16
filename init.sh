#!/bin/bash
curl --request POST \
    --url http://localhost:5000/keyboards \
    --header 'Content-Type: application/json' \
    --data '{ "name": "HyperX Alloy Red V4 ",
    "url_photo": "https://cdn.thewirecutter.com/wp-content/media/2021/05/mechanicalkeyboards-2048px-0032.jpeg?auto=webp&quality=75&width=768&dpr=1.5"}' \
    -: --request POST \
    --url http://localhost:5000/keyboards \
    --header 'Content-Type: application/json' \
    --data '{
    "name": "GK61 Red Mamba Ultra",
    "url_photo": "https://cdn.thewirecutter.com/wp-content/media/2021/05/mechanicalkeyboards-2048px-0018.jpeg?auto=webp&quality=75&width=768&dpr=1.5" }'

curl --request POST \
    --url http://localhost:5000/register \
    --header 'Content-Type: application/json' \
    --data '{ "username": "myge",
        "password": "myge",
        "email": "myge@mail.udp.cl",
        "name": "Miguel" }' \
    -: --request POST \
    --url http://localhost:5000/register \
    --header 'Content-Type: application/json' \
    --data '{ "username": "naikelin",
        "password": "naikelin",
        "email": "naikelin@mail.udp.cl",
        "name": "Nicolas Nunez" }'