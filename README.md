[![Build status](https://badge.buildkite.com/acca603faab8c27aa01546a4254e6672bff7e5786724c2c122.svg)](https://buildkite.com/myob/martyna-tech-test)

# DevOps Technical Test

This is a simple implementation of the devops technical test.
This project contains source code and supporting files for a simple api server which accepts http GET requests only

## Installation

### System Requirements

- A command line interface (CLI) such as Command Prompt for Windows or Terminal for macOS
- IDE of your choice

### Clone

- Clone this repo to your local machine and in the CLI, navigate into the folder containing the solution

### Running the application locally

Run the below script to build and run the application at localhost:8080

```bash
./scripts/build.sh
```

### Running tests

Run the below script to run tests and HTML output of the coverage profile

```bash
./scripts/test.sh
```

## Usage

The application will only respond to GET requests to the below endpoints:

- `/` returns "hello world".
- `/health` returns a response code of 200.
- `/metadata` returns the application description, version and the latest github commit sha as json.

```json
{
  "myapplication": [
    {
      "version": "v1.0.0-23-g2e298b4",
      "description": "pre-interview technical test",
      "lastcommitsha": "c6d48a1d2fb06b549c6f19bf7d129ce987050dcd"
    }
  ]
}
```
