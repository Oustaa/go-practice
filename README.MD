# this project is the result of what i have learn so far:
## Descriptio:
- the project is a tasks/habbit tracking backend server, it provide functionalities such as adding tasks and organazing them by category, status, priority, and it should contain auth so each user should be able to manage his tasks. 
(!!! OPTIONAL/ADD LATER: functionalities to add collaboration)

Version One: Browser Based ( Golang, MySQL )
Version Two: API BASED ( Golang, MySQL, React )
  
## Setup
- [x] GIT/GitHub
- [x] Setup Routing
- [x] Docker ( Done for now ) 
    - [x] MySQL ( I should add volumes )
    - [ ] GO! ( if added i should add network )
- [x] Connect to database
- [x] Setup migrations


## Functionalities
- [x] Tasks CRUD:
    - [x] Read.
    - [x] Create.
    - [x] Update.
    - [x] Delete.
- [ ] Make tasks as completed
- [ ] Administration:
    - [ ] Manage Account ?
    - [ ] tasts Category CRUD:
        - [ ] Read.
        - [ ] Create.
        - [ ] Update.
        - [ ] Delete.
    - [ ] Tasks Status CRUD:
        - [ ] Read.
        - [ ] Create.
        - [ ] Update.
        - [ ] Delete.
- [ ] Auth:
    - [ ] Login
    - [ ] Singin
    - [ ] Forgget Password
    - [ ] Delete Account ( Delete all my tasks ?????????? )

## Time Constraint
* Start at: 2025-05-24 / 2025-05-25 
* Shoud end at: 2025-06-07


## Technologies
- Go, Mysql/Sqlit, Docker, HTML, CSS, !testing

## Build For the API, and WEB
- Routes Groups: /api and /nothig
- Handlers:
    - For the web: Returns View ( html, css, js )
    - For the api: Returns json ( json )
- Services:
    - Encapsulate the business logic (e.g Create A Task, Delete a Task)
