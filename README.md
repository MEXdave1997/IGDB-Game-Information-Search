# IGDB Search

This program was created with the intent of searching and displaying information for a game from the __IGDB__.

This project was created in [Go](https://golang.org) using the [IGDB Package](https://github.com/Henry-Sarabia/igdb).

## Purpose
This tool is designed with the needs of game reviewers in mind. It will help
someone compile information for a game that they get that would be necessary,
like a brief summary of the plot, what platforms the game was released on as
well as the dates, and more.

## Compiling From Source

If Go isn't already install, you can find installation instructions [here](https://golang.org/doc/install). 

### **NOTE**: This Project has dependencies that require Go Version 1.9 or later

After installation, you can grab the dependencies using:
```
go get github.com/Henry-Sarabia/igdb
```

From here, you can compile by running:

```
go build
```
or
```
go build search.go
```
This will create a binary which you can then run from the command line for testing.

## Usage
Once compiled, you should have an executable file located in the project directory, named according to how you decided to build the project (more on that found [here](https://golang.org/pkg/go/build/)).

The executable has 2 flags required in order to run:

```
    -g: Name of the game to search (in quotations if it has multiple words)
    
    -k: IGDB API-Key that is used to communicate with the database
```

To run, assuming you used the second command to build:

```
./search -g <NAME_OF_GAME> -k <YOUR_API_KEY>
```
