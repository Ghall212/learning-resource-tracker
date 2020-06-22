# Learning Resource Tracker

## Summary

This toy application is an API for a Learning Resource Tracker which allows consumers to store resources for learning different topics in different categories. Resources belong to particular topics and topics themselves can optionally belong to a category. Categories can be nested within other categories with a depth limit of 4. At the moment, this application was not created to be robust, but rather to familiarize myself more with the tech stack.

## Tech Stack

This application is written in Go and uses a MySQL database. The [gorilla mux router](https://github.com/gorilla/mux) is used, along with the [Go MySQL Driver](https://github.com/go-sql-driver/mysql) and the [godotenv package](https://github.com/joho/godotenv) to load the `.env` file. I am developing this application on Ubuntu 20.04 running in WSL2 on Windows 10.

## Database

MySQL is used as the database in this application. In the project root, one will find two `.sql` files. `mysql_db_schema.sql` contains the script to create the database and tables. `mysql_db_dummy_data.sql` contains the script to populate the existing database with dummy data.

## Environment Variable

This application can optionally read its environment variables from a `.env` file in the project root. Simple add a `.env` file to the root of the project and add the following database connection string template for your local MySQL database:

```
DB_CONNECTION_STRING=username:password@tcp(localhost:3306)/LearningResourceTracker
```

If you opt not to use the `.env` file, this environment variable must set in a different way.
