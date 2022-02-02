# ACCOUNT APP

### 1. Quick Review

Account app will hold user related operation including :
1. CRUD for user accout
2. CRUD for user.role
3. Read only for user.status (since the status is fix)
4. Auth for user account (login, signin, signout)

### 2. Directory Structure

```bash
|-- account/
|-- |-- datastore/
|-- |-- |-- user.go
|-- |-- |-- user.role.go
|-- |-- |-- user.role_test.go
|-- |-- |-- user.status.go
|-- |-- |-- user.status_test.go
|-- |-- |-- user_test.go
|-- |-- handler/
|-- |-- |-- user.go
|-- |-- |-- user.role.go
|-- |-- |-- user.role_test.go
|-- |-- |-- user.status.go
|-- |-- |-- user.status_test.go
|-- |-- |-- user_test.go
|-- |-- service/
|-- |-- |-- user.go
|-- |-- |-- user.role.go
|-- |-- |-- user.role_test.go
|-- |-- |-- user.status.go
|-- |-- |-- user.status_test.go
|-- |-- |-- user_test.go
|-- |-- README.md
|-- |-- router.go

```
