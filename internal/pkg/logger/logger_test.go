/*
   package logger
   logget_test.go
   - test logger behaviour
*/
package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// since most of the test is automatic loaded due to the module use 'init' function
// so it is automaticaly run by init

// TestLogger is for testing logger modul
func TestLogger(t *testing.T) {
    // EXPECT FAIL to get/ create writer stdout. Simulated by inputing ilegal
    // parameter (ilegal file name)
    t.Run("EXPECT FAIL writer file not found", func(t *testing.T){
        // mock go os var and change its value to "windows" to simulate test
        var aGoOS = goOS
        goOS = "windows"

        // actual test
        _ = getWriter("/as/atest.xyz")

        // return goOS value back to before
        goOS = aGoOS
    })

    // EXPECT SUCCESS DEBUG. Simulate debug log creation
    t.Run("EXPECT SUCCESS DEBUG", func(t *testing.T){
        // set loglevel to debug
        el(logrus.DebugLevel)

        // actual test
        Debugf("LOG DEBUG: %v\n", "is running debug")
    })

    // EXPECT SUCCESS INFO. Simulate info log creation
    t.Run("EXPECT SUCCESS INFO", func(t *testing.T){
        // set loglevel to info
        el(logrus.InfoLevel)

        // actual test
        Infof("LOG INFO: %v\n", "is running info")
    })

    // EXPECT SUCCESS WARNING. Simulate warning log creation
    t.Run("EXPECT SUCCESS WARNING", func(t *testing.T){
        // set loglevel to warning
        el(logrus.WarnLevel)

        // actual test
        Warnf("LOG WARNING: %v\n", "is running warning")
    })

    // EXPECT SUCCESS ERROR. Simulate error log creation
    t.Run("EXPECT SUCCESS ERROR", func(t *testing.T){
        // set loglevel to error
        el(logrus.ErrorLevel)

        // actual test
        Errorf("LOG ERROR: %v\n", "is running error")
    })

    // EXPECT SUCCESS FATAL. Simulate fatal log creation
    // t.Run("EXPECT SUCCESS FATAL", func(t *testing.T){
        // we need to recover the process since we testing 'panic' here
        // IT IS NOT HELP SINCE FATAL CALL os.exit
        // defer func() {
        //     if r := recover(); r != nil {
        //         t.Logf("Recovered in NewTest: %v", r)
        //     }
        // }()

        // set loglevel to fatal
        // el(logrus.FatalLevel)

        // actual test
        // Fatalf("LOG FATAL: %v\n", "is running fatal")
    // })
}
