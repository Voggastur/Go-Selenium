package main

import (
	"fmt"
	"os"
	"log"
	"github.com/tebeka/selenium"
)


// If you want to actually run this example:
//
//   1. Ensure the file paths at the top of the function are correct.
//   2. Remove the word "Example" from the comment at the bottom of the
//      function.
//   3. Run:
//      go test -test.run=Example$ github.com/tebeka/selenium
func main() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		seleniumPath    = "vendor/selenium-server-standalone-3.4.jar"
		geckoDriverPath = "vendor/geckodriver-v0.18.0-linux64"
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
