# phonebook-db

## Purpose

This is my first application written with Go. My main objective in creating this simple project was to get a feel for the language and syntax.
As such, not everything is perfectly organized or optimized. For the purposes of this project, I prioritized experimentation, learning, and speed 
over building a flawless and complete application.

## Description

This is a simple CRUD application designed to simulate a phonebook. For data persistence, rather than using an actual database, this app
simply reads and writes from a text file titled "database.txt". This application is designed to be run in the console.

The CRUD commands implemented are:

- Get all phonebook records
- Find a phonebook record by ID
- Create a phonebook record
- Modify an existing phonebook record
- Delete a phonebook record

For convenience's sake, I also implemented these commands:

- A command to wipe the "database"
- A command to seed the "database"
