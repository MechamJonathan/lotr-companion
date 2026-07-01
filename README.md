# LOTR Companion CLI
<p align="center">
  <img src="https://github.com/MechamJonathan/lotr-companion-app/blob/main/img/Palantir.png" width="350" title="logo">
</p>

# Palantír CLI

![badge](https://github.com/MechamJonathan/lotr-companion-app/actions/workflows/ci.yml/badge.svg?event=pull_request)

A command-line interface (CLI) application written in Go that functions as a companion app. This application allows users to retrieve information about Lord of the Rings movies, books, and characters using The-One-Api (https://the-one-api.dev/).

> In J.R.R. Tolkien's The Lord of the Rings, a Palantír is a powerful, indestructible crystal ball used for communication and seeing events in distant places or the past, often referred to as "seeing stones"

## Why?
This unofficial CLI application is designed for fans of The Lord of the Rings who want fun, quick, and easy access to information about the books, movies, and characters. Whether you're looking up quotes, exploring Middle-earth lore, or just need a reference guide, this tool provides an efficient way to find details without searching through multiple sources.

## Features 
- Retrieve a a list of _Lord of the Rings_ books, movies, characters.
- Get details about specific books, movies, characters.
- Browse character quotes with pagination
- Cache frequently accessed data for faster subsequent retrievals.

## Requirements
- Go (version 1.18 or later)
- Internet connection (to fetch data from The-One-Api)
- Register at https://the-one-api.dev/ to get an access token

## Contributing
### Installation 
1. Clone the Repository:
```
git clone https://github.com/MechamJonathan/palintir-cli
cd palintir-cli
```
2. Create a .env file at the root of the project with the following:
```
API_KEY = "{Your Access Token}"
```

### Run Application
Build the application and then run the executable:
```
go build
go run .
```
# Commands

| Command     | Description                                        |
| ----------- | -----------                                        |
| books       | Lists all books                                    |
| characters  | Lists all characters or group of characters        |
| details     | Return details about specific character, movie, or book     |
| exit        | Exit the program                                   |
| help        | Display help message and all available commands    |
| movies      | List all movies                                    |
| quotes      | View next page of a character's quotes             |
| quotesb     | View previous page of a character's quotes         |

## Demo
<p align="left">
  <img src="https://github.com/MechamJonathan/Palantir-cli/blob/main/demo.gif" width="800" title="logo">
</p>
