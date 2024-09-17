# FG21 Go Event Organizer

## Description

Proyek ini adalah backend service yang dibangun menggunakan Go (Golang). Tujuan dari proyek ini adalah untuk menyediakan API yang kuat dan scalable untuk aplikasi Event Ogranizer (FindYourEvent).

## Features

- Authentication JWT token
- Transaction
- Filter Event by category
- Update Profile
- Create Event

## Installation

1. **Clone repository**

```sh
git clone https://github.com/ashari-dev/fgh21-go-event-organizer.git
cd fgh21-go-event-organizer
```

2. **Install Dependensi**

```sh
go mod tidy
```

3. **Open VScode**

```sh
Code .
```

4. **Migrate Database**

```sh
    make migrate:reset
```

5. **Config DB**

open file lib/db.lib

```go
    host := `hostname`
	port := `port`
	user := `userDB`
	pass := `passDB`
	db := `DB Name`
```

6. **Run**

```sh
go run main.go
```

