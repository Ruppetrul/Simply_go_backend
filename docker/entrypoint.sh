#!/bin/bash

while ! nc -z db 3306; do
    echo "Waiting for MySQL to be available..."
    sleep 1
done

./myapp