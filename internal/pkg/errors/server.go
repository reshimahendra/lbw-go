/*
    package errors
    server.go
    - error constant for handling handler error
*/
package errors

const (
    // ErrServer is error code for server error
    // msg = "server fail to run"
    ErrServer = iota + 650
    
    // ErrServerMode is error code for server mode error
    // msg = "server mode is unknown"
    ErrServerMode

    // ErrServerHost is error code for server host error
    // msg = "server host is unknown"
    ErrServerHost

    // ErrServerPort is error code for server port error
    // msg = "server port is unknown or already used"
    ErrServerPort
)

const (
    // ErrServer is error code for server error
    // msg = "server fail to run"
    ErrServerMsg = "server fail to run"
    
    // ErrServerMode is error code for server mode error
    // msg = "server mode is unknown"
    ErrServerModeMsg = "server mode is unknown"

    // ErrServerHost is error code for server host error
    // msg = "server host is unknown"
    ErrServerHostMsg = "server host is unknown"

    // ErrServerPort is error code for server port error
    // msg = "server port is unknown or already used"
    ErrServerPortMsg = "server port is unknown or already used"
)
