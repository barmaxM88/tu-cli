# TU — Terminal Utility

Simple terminal utility written in Go.

The program allows you to create, change, delete and list text records.
All data is stored locally in a JSON file.

## Features

- create — create a new record
- change — update an existing record by id
- delete — delete a record by id
- list — show all records

## Usage

### Build the program

Windows:
go build -o tu.exe main.go

Mac/Linux:
go build -o tu main.go

### Run commands

Windows:
.\tu.exe create "My first record"

.\tu.exe change 1 "Updated text"

.\tu.exe delete 1

.\tu.exe list

Mac/Linux:
./tu create "My first record"

./tu change 1 "Updated text"

./tu delete 1

./tu list

## Notes

- Records are stored in records.json in the same folder as the program.
- Make sure to run the program in the folder where main.go or the compiled binary is located.
- Use the list command to see all existing records with their IDs.

