package main

import (
        "fmt"
        "os"
        "log"
        "github.com/tebeka/selenium"
)


func main() {

        const (
                seleniumPath    = "selenium-server-standalone-3.141.59.jar"
                chromeDriverPath = "chromedriver"
                port            = 4444
        )

        opts := []selenium.ServiceOption{
                selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
                selenium.ChromeDriver(chromeDriverPath), // Specify the path to ChromeDriver in order to use Chrome.
                selenium.Output(os.Stderr),            // Output debug information to STDERR.
        }

        selenium.SetDebug(true)
        service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
        if err != nil {
                log.Fatalf(err.Error())
        }

        defer service.Stop()

        caps := selenium.Capabilities{"browserName": "Chrome"}
        wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))

        if err != nil {
                log.Fatalf(err.Error())
        }

        wd.Get("https://pensionera.se/bli-medlem")
        title, err := wd.Title()
        if err != nil {
                log.Fatalf(err.Error())
        }
        expected := string("Bli medlem | Pensionera")

        if string(title) != expected {
                log.Printf("Fail: Unexpected title found, got: %v, expected: %v", title, expected)
        }

        wd.Quit()
}
