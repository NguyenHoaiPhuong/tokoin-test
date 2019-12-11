# tokoin-simple-test

## Introduction

A simple code challenge from tokoin.

2 Solutions are planned to provide: one programmed by golang and another one programmed by python.

## Usage

### Golang

#### Prerequisite

golang and dep (dependency management) must be installed.

#### How to run the app

- Access to golang directory:

```
cd golang
```

- In order to run the app without building the binary file, please run following commands:

```
dep ensure
go run *.go
```

- If you want to run the app by building binary file, please run following commands:

```
dep ensure
go build -o app
./app
```

### Python

TO UPDATE

## Source code structure explanation

### Golang

The app is splitted into serveral packages:

- **resources**: it is the place where we store the default configurations ```default.toml``` which need to be parsed at the beginning. Furthermore, user, ticket and organization data is also stored in json files in resources/data directory.

- **config**: this package is responsible for parsing configurations set by environment varibles or default config files (extension can be 'json', 'toml', 'yml', etc). A singleton GlobalConfig is provided for handling app configurations and it is unique.

- **model**: contains structures defining objects used in the app, such as User, Ticket and Organization.

- **jsonfunc**: contains useful json functions (parsing json object, reading json file, etc)

- **utils**: contains other useful functions, sucn as CheckError, PrintPtrStructObject, etc

- **main**: this is the "entrypoint" package where go cmd will run the whole app. There are 4 go files in this package:
    - main.go: includes function main() which is the entrypoint of *go run* command.
    - user.go: includes all functions related to user object
    - ticket.go: includes all functions related to ticket object
    - organization.go: includes all functions related to organization object

Workflow:

1- Read read json files, parse data and save all objects to slices (arrays) of users, tickets and organizations

2- Consolidate data: since we must be able to search all fields of user / ticket / organization struct fast, O(1) time - complexity solution is provided. However, in order to obtain O(1) time - complexity, the trade-off is that a lot of consolidated maps are used. This leads to "heavily" loading data at the beginning.

### Python

TO UPDATE

## Difficulties during the assignments

A binary file named **challenge** was provided together with the input data. However, this binary file was written in Mac OS and it cannot be run in linux OS. Therefore, the source code was written completely based on instructions **instruction.pdf** file.

Moreover, in case of a field compose of an array of values like *tag* in a user object, what should the searching work do? The searching work will return the object if the whole array is matched, won't it? Or will it return objects whose field array includes searching values?

```
"tags": [
    "Cawood",
    "Disautel",
    "Boling",
    "Southview"
],
```

## What have been done and what haven't been done

1- Available to **search user** by any fields, except **tags**. Reason is explained in the [section](#Difficulties-during-the-assignments) above.

2- Available to **search ticket** by fields listed below. Searching other fields are similar to the way used in Searching user, so repeating the work is not necessary.
- _id
- external_id
- subject
- organization_id
- submitter_id    
- assignee_id

3- Available to **search organization** by fields listed below.
- _id
- external_id
- name