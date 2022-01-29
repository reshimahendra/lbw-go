/*
    package errors
    server.go
    - error constant for handling handler error
*/
package errors

const (
    // ErrServer is error code for server error
    ErrServer = iota + 650
    
    // ErrServerMode is error code for server mode error
    ErrServerMode

    // ErrServerHost is error code for server host error
    ErrServerHost

    // ErrServerPort is error code for server port error
    ErrServerPort
)

const (
    // ErrServer is error code for server error
    ErrServerMsg = "server fail to run"
    
    // ErrServerMode is error code for server mode error
    ErrServerModeMsg = "server mode is unknown"

    // ErrServerHost is error code for server host error
    ErrServerHostMsg = "server host is unknown"

    // ErrServerPort is error code for server port error
    ErrServerPortMsg = "server port is unknown or already used"
)
