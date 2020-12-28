# Scrapper microservice application demo
This is demo application for showing purposes. Application is consisted of 2 microservice apps: scrapper and API.

# Technical part
1. Download and install "migrate" CLI for runngin migrations (https://github.com/golang-migrate/migrate). 
2. After running docker-compose build/up - go into scrapper and api folders and run "make migrateup" to run migrations

# Scrapper 
Scrapper service is responsible for scrapping posts every "x" seconds from websites (https://www.autoblog.com, https://www.buzzfeed.com). 
This service has it's own database for saving posts.

# api
API is a service for providing REST API to end users. Also this service has it's own database. This service will have basic functionalities such as Login, Registration, getting posts from scrapper service using GRPC Protobuffer and providing them to the dashboard.

