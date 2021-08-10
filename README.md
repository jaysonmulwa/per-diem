## Per-Diem Backend Challenge

### Introduction

Per-Diem (Latin for "per day" or "for each day") - This is a is a substitute for using an actual expense reimbursement method.

The project in this repository is the backend api for the per-diem backend challenge. Which - as is related to its name - involves a scheduled way for customers to schedule orders they would like fulfilled over a period of time.

The features of the API satisy the given requirements, they include:
- 4 API endpoints.
- JWT Authentication.
- Docker.
- Tests.
- Inline and external documentation of the repo (This readme file).

### Architecture

Here is a representation of the architecture.

![alt text](https://github.com/jaysonmulwa/per-diem/blob/main/arch.png?raw=true)

### Install

1. Clone this repositry.
```bash
$ https://github.com/jaysonmulwa/per-diem.git
```
2. Install docker to your machine.

3. Run the following commands to start up the instance.
```bash
$ docker build -t per .
```
```bash
$ docker run -p 3000:3000 per
```

4. Request for the jwt token from the endpoint.
```bash
$ https://localhost:3000/jwt
```
5. Add the token your request header for other requests.
```bash
$ Authorization : Bearer <token>
```

6. Access other endpoints. eg:
```bash
$ https://localhost:3000/roles
```


### Endpoints

### Tests

- Run main_test.go


