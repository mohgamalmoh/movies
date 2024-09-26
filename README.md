# Build an End-to-End Best Movie of All Time App

Building a RESTful CRUD (Create, Read, Update, Delete) API with Golang, Gin, and Gorm.

Create a CRUD (Create, Read, Update, Delete) application that utilizes a CSV file and consumes
the TMDB APIs. This app is to seed, store, and sync the data in the attached CSV into the backend
application database and create a way to interact with it along with the IMDB APIs.

to run the app execute this command: docker-compose up
to run DB migrations get into the app container using this command: docker exec -it {container_name} sh
then, run this comman: go run main.go migrations up
then, execute the import endpoint to seed the data into the DB (the endpoint is documented using swagger in the docs folder)
then, you can play aroubd with the app using the api endpoints documented in "docs" folder