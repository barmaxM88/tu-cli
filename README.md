TU — Terminal Utility

Simple terminal utility written in Go.

The program allows you to create, change, delete and list text records.
All data is stored locally in a JSON file.

Features

create — create a new record

change — update an existing record by id

delete — delete a record by id

list — show all records

Usage

Build the program:

go build -o tu

Create a record:

./tu create "My first record"

Change a record:

./tu change 1 "Updated text"

Delete a record:

./tu delete 1

List records:

./tu list