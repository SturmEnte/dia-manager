# Dia Manager

## Environment Variables

| Name        | Optional | Description                                                |
| ----------- |:--------:| ---------------------------------------------------------- |
| PORT        | ×        | The port the website should run on                         |
| PG_HOST     | ×        | The host of the PostgreSQL database                        |
| PG_DATABASE | ×        | The name of the database you want to let the server run on |
| PG_USERNAME | ×        | The username used to connect to the database server        |
| PG_PASSWORD | ×        | The password used to connect to the database server        |

## Development progess / tasks

[here](https://github.com/users/SturmEnte/projects/2/views/1)

## Languages

| Language | Availability              |
| -------- | ------------------------- |
| English  | Default                   |
| German   | 100% (if release version) |

## Repositories

[GitHub](https://github.com/SturmEnte/dia-manager)
[GitLab](https://gitlab.com/SturmEnte/dia-manager)

## How to develop the project

### Backend

You have to have **Node.js** and **npm** installed to develope on the backend

1. Run `npm i`
2. Create a `.env` file and add the in the `Environment variables` heading mentioned variables
3. Run `npm test` to start the backend server for development

### Frontend

You have to have **Node.js**, **npm**, **Python 3** and **pip** installed in order to develope on the frontend

1. Run `npm i`

2. Run `pip install colorama`

3. Run `py buildOnChange.py` or `python buildOnChange.py` depending on your operation system and the way you installed Python

4. Start the backend to see the changes on the site and make it work corectly
