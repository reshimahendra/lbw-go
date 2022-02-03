# DATABASE

Database folder containing database connection function and method to our postgresql server. To connect with the database, we are using [pgx][1] to get best performance. The `db.go` file contain database pool connection preparation and validation for the established pool connection.

### File Structure
```bash
|-- database
|-- |-- sql
|-- |-- |-- user.sql
|-- |-- db.go
|-- |-- README.md
```

[1]:https://github.com/jackc/pgx
