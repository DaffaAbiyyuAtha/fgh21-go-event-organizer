# MyTick BACKEND

Welcome to the MyTick Backend Web Project! This backend is part of an online event ticket sales system. Built using Go (Gin framework) and PostgreSQL as the database. This backend provides integrated user, event, and ticket transaction management features.

## Features

- <b>User Management</b>: CRUD features for users, including login, registration, and profile management.
- <b>Event Management</b>: Create and view event details.
- <b>Authentication</b>: Using JWT (JSON Web Token) for user authentication.
- <b>Ticket Transactions</b>: Online ticket booking.

Built using

![JWT](https://img.shields.io/badge/JWT-black?style=for-the-badge&logo=JSON%20web%20tokens)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Google Chrome](https://img.shields.io/badge/Google%20Chrome-4285F4?style=for-the-badge&logo=GoogleChrome&logoColor=white)
![Visual Studio Code](https://img.shields.io/badge/Visual%20Studio%20Code-0078d7.svg?style=for-the-badge&logo=visual-studio-code&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

## 📌 Getting Started

1. Clone this repository
```sh
  git clone https://github.com/DaffaAbiyyuAtha/fgh21-go-event-organizer
  cd fgh21-go-event-organizer
```
2. Open in VSCode
```sh
  code .
```
3. Create Images Postgres
```sh
  docker build -t postgres .
```
4. Run Images Postgres
```sh
  docker run -e POSTGRES_PASSWORD=1 -p 5432:5432 -d postgres
```
5. Create Images Backend
```sh
  docker build -t backend .
```
6. Run Images Backend
```sh
  docker run -d -p 8080:8080 --name backend backend
```
7. You can run it in combination with my frontend, https://github.com/DaffaAbiyyuAtha/fgh21-react-event-organizer/

## API Reference

#### Login

```http
  POST auth/login
```
#### Register

```http
  POST auth/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `/users` | `GET` | Get a list of users data |
| `/users/:id` | `GET` | Get a detailed users data |
| `/users` | `POST` | Insert a users data |
| `/users/:id` | `PATCH` | Update a users data |
| `/users/password` | `PATCH` | Update a users password |
| `/users/:id` | `DELETE` | Delete a users data |
| `/profile/` | `GET` | Get a list of profile data |
| `/profile/update` | `PATCH` | Update a profile data |
| `/profile/picture` | `PATCH` | Update image profile |
| `/events` | `GET` | Get a list of events data |
| `/events/:id` | `GET` | Get a detailed events data |
| `/events/section/:id` | `GET` | Get a detailed price events data |
| `/events/update` | `POST` | Create events by user |
| `/events/see_one_event` | `GET` | Get list Create events by user |
| `/events/:id` | `PATCH` | Update a event data |
| `/events/:id` | `DELETE` | Delete a events data |
| `/transactions` | `GET` | Get list transactions by user |
| `/transactions` | `POST` | Create transactions by user |
| `/transactions/payment` | `GET` | Get transactions by user |
| `/locations` | `GET` | Get list locations |
| `/partner` | `GET` | Get list our partner |
| `/wishlist` | `GET` | Get list wishlist by user |
| `/wishlist/:id` | `GET` | Get list wishlist by id |
| `/wishlist/:id` | `POST` | Create wishlist by user |
| `/wishlist/:id` | `DELETE` | Delete wishlist by user |
| `/nationalities` | `GET` | Get list nationalities |
