package main

import (
	"fmt"
	"os"
	"log"
	"github.com/tebeka/selenium"
)


func main() {

	const (
		seleniumPath    = "vendor/selenium-server-standalone-3.141.59.jar"
		geckoDriverPath = "vendor/geckodriver-v0.28.0-win64"
		port            = 8080
	)

	opts := []selenium.ServiceOption{
		selenium.StartFrameBuffer(),           // Start an X frame buffer for the browser to run in.
		selenium.GeckoDriver(geckoDriverPath), // Specify the path to GeckoDriver in order to use Firefox.
		selenium.Output(os.Stderr),            // Output debug information to STDERR.
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		log.Fatalf(err)
	}

	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))

	if err != nil {
		log.Fatalf(err)
	}

	wd.get("https://pensionera.se/bli-medlem")
	title, err := wd.getTitle()
	if err != nil {
		log.Fatalf(err)
	}
	expected := string("Bli medlem | Pensionera")

	if string(title) != expected {
		log.Println("Fail: Unexpected title found, got: %v, expected: %v", title, expected)
	}

	wd.Quit()
}
