# db-docu-to-markdown

[![License](https://img.shields.io/github/license/PDCMFinder/db-docu-to-markdown)](https://github.com/PDCMFinder/db-docu-to-markdown/blob/main/LICENSE)

## Overview

`db-docu-to-markdown` is a tool for creating markdown files with basic information about a database schema.

## Features

- Retrieve information about tables, columns, and their attributes from a database schema.
- Generate structured Markdown documentation with details such as column name, data type, comments, and more.

## Installation

To install db-docu-to-markdown, use the following command:

```bash
go install github.com/PDCMFinder/db-docu-to-markdown/cmd/db-docu-to-markdown@latest
```

## Usage
```bash

USAGE:
   db-docu-to-markdown [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value, -H value              database host (default: "localhost")
   --port value, -P value              database port (default: 8080)
   --user value, -u value              database user (default: "admin")
   --password value, -p value          database password (default: "password")
   --name value, -n value              database name (default: "test")
   --schemas value, -s value           comma-separated list of schemas to describe (default: "public")
   --dbtype value, --dt value          specify the database type (default: "postgres")
   --help, -h 
```
## Output
The command generates a folder called `output` with one Markdown file per schema provided with the `schemas` flag.

## License
This project is licensed under the Apache License 2.0.
