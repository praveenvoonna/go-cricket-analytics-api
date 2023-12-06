# Cricket Data REST API - Golang Coding Assignment

## Overview
This Golang program serves a JSON REST API locally, utilizing data from the provided "ODI Data.csv" file, which can be downloaded from [Kaggle](https://www.kaggle.com/datasets/mahendran1/icc-cricket/data) after creating a free account. The CSV file contains columns such as "Player Span, Played Matches, Innings, Number of NotOuts, Number of Runs, Highest Score, Total Average, Balls Faced, Strike Rate, Total number of 100's, Total number of 50's, Total number of ducks".

### Task Description
The goal is to create a Golang program that loads the .csv file into memory, provides a REST API, and responds to the following queries:
- For players who finished their career in a given year, determine which player scored the most runs over their career.
- For a given year, identify which players were active.

## Instructions
1. Download the "ODI Data.csv" file from [Kaggle](https://www.kaggle.com/datasets/mahendran1/icc-cricket/data).
2. Write a Golang program that loads the CSV data into memory and serves a JSON REST API locally.
3. Implement API endpoints to handle queries for player analysis based on career end year and active players in a given year.
4. Include proper documentation and tests within your solution.

### File Structure
- main.go            // Main application entry point
- data/              // Folder to store ODI Data.csv
  - ODI Data.csv     // Cricket data CSV file
- csv/               // Package for CSV handling
  - reader.go        // Code for reading and parsing CSV data
  - reader_test.go   // Test file for CSV reader
- handler/           // Package for HTTP handlers
  - handler.go       // Implement API endpoints
  - handler_test.go  // Test file for handlers
- model/             // Package for defining data models
  - player.go        // Player model definition
- tests/             // Folder for additional test files
- README.md          // Instructions and details



## Implementation Details
- The Golang program should utilize CSV parsing to load data from "ODI Data.csv".
- Implement API endpoints to handle queries for player analysis based on the specified requirements.
- Ensure proper error handling and robust code structure.
- Document the API endpoints, their functionality, input/output formats, and examples.

## Running the Program
1. Clone or Download the Repository: Obtain the source code by either cloning the repository or downloading it as a ZIP file.
2. Initialize the Go Module: If your project is not yet a Go module, initialize it using the following command in your project directory:

``` bash
  go mod init yourmodulepath
```
Replace yourmodulepath with the module path you desire.

3. Install Dependencies: If the project contains external dependencies, ensure they are imported using:

```bash
go mod tidy
```
4. Run the Application: Use the go run command to execute the main Go file (main.go in this case):

```bash
go run main.go
```

## API Endpoints
### `GET /players/most_runs?year={year}`
- **Description**: Retrieves the player who scored the most runs in their career for a given year.
- **Parameters**:
  - `year`: Year indicating the end of player careers.
- **Example**: `/players/most_runs?year=2012`

### `GET /players/active?year={year}`
- **Description**: Retrieves the list of players active in a given year.
- **Parameters**:
  - `year`: Year for which active players are to be queried.
- **Example**: `/players/active?year=2000`

## Testing
- Include comprehensive tests covering edge cases, API endpoints, and data loading functionalities.
- Ensure test coverage and reliability of the codebase.