package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

func login() {

	const (
		seleniumPath     = "selenium-server-standalone-3.141.0.jar"
		chromeDriverPath = "chromedriver"
		port             = 4444
	)

	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(os.Stderr),
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, port, opts...)
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf(err.Error())
	}
	wd.Refresh()
	if err := wd.Get("http://localhost:8080/bli-medlem"); err != nil {
		log.Fatalf(err.Error())
	}
	url, err := wd.CurrentURL()
	if err != nil {
		log.Fatalf(err.Error())
	}
	if url != ("http://localhost:8080/bli-medlem") {
		wd.Get("https://pensionera.se/bli-medlem")
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	if err != nil {
		log.Fatalf(err.Error())
	}
	windowID, err := wd.CurrentWindowHandle()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = wd.MaximizeWindow(windowID)
	if err != nil {
		log.Fatalf(err.Error())
	}
	title, err := wd.Title()
	if err != nil {
		log.Fatalf(err.Error())
	}
	expected := string("Bli medlem | Pensionera")
	if string(title) != expected {
		log.Printf("FAIL: Unexpected Title, Got: %v, Expected: %v", title, expected)
	} else {
		log.Print("SUCCESS: Title matches expected")
	}
	if err != nil {
		log.Fatalf(err.Error())
	}
	//personnummer
	input1, err := wd.FindElement(selenium.ByID, "seed")
	if err != nil {
		log.Fatalf(err.Error())
	}
	//email
	input2, err := wd.FindElement(selenium.ByID, "email")
	if err != nil {
		log.Fatalf(err.Error())
	}
	//phonenumber
	input3, err := wd.FindElement(selenium.ByID, "phone-mobile")
	if err != nil {
		log.Fatalf(err.Error())
	}
	//salary
	input4, err := wd.FindElement(selenium.ByID, "salary")
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = input1.SendKeys(`
	1234
	`)
	if err != nil {
		log.Fatalf(err.Error())
	}
	time.Sleep(time.Millisecond * 300)
	err = input2.SendKeys(`
	a@hotmail.com
	`)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = input3.SendKeys(`
	08-00000000
	`)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = input4.SendKeys(`
	1000000
	`)
	if err != nil {
		log.Fatalf(err.Error())
	}
	button, err := wd.FindElement(selenium.ByCSSSelector, "button.btn.btn-sign")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err := button.Submit(); err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Print("Test run finished, Timeout 3 seconds")
	if err != nil {
		log.Fatalf(err.Error())
	}
	wd.Close()
}

func main() {
	login()
}
