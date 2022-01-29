/*
    package domain
    user.status.go
    - containing user.status model and response dto struct
    - it only support "get" method since the values was preinstalled in the database
*/
package domain

// UserStatus is model for status of the user
type UserStatus struct {
    // ID is user status id which is its primary key
    ID          int     `json:"id"`

    // StatusName is the name of the status
    StatusName  string  `json:"status"`

    // Description is the short description of the status
    Description string  `json:"description"`
}
