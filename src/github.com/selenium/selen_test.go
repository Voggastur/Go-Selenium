package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

const (
	testPath         = "deps/"
	seleniumPath     = testPath + "selenium-server-standalone-3.141.59.jar"
	chromeDriverPath = testPath + "chromedriver"
	serverPort       = 4444
	webSite          = "https://www.google.com"
	impTimeout       = 60
	// Input Parameters for Signup Page
	seed   = "200002022382"         // Skatteverket Test SSN
	email  = "rataveh652@lidte.com" // temp-mail.org
	phone  = "080101010"
	salary = "123123"
)

var (
	// declare the variable at package level (outside functions) so that all tests can share the same one
	service       *selenium.Service
	err           error
	title         string
	expectedTitle string
	currentURL    string
	selected      bool
)

// init the driver once at the start
func setup() {
	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(os.Stderr),
	}

	selenium.SetDebug(true)
	service, err = selenium.NewSeleniumService(seleniumPath, serverPort, opts...)
	if err != nil {
		log.Fatal(err)
	}
}

func shutdown() {
	// Stop service when called
	service.Stop()
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	shutdown()
	os.Exit(code)
}

/* This test is dependent on the tebeka selenium package for Go Selenium; "go get github.com/tebeka/selenium"
And a chromedriver for your Chrome version from: - https://chromedriver.chromium.org/downloads
Webdrivers for Safari, Opera, IE and Firefox can also be used.
A GRID selenium-server-standalone-3.141.59.jar from: - https://www.selenium.dev/downloads/
The seleniumPath and chromeDriverPath should match the names of the files, paths can be used */

// Main Run Below
func Test_Practice(t *testing.T) {
	caps := selenium.Capabilities{"browserName": "chrome", "acceptInsecureCerts": true}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", serverPort))
	if err != nil {
		log.Fatal(err)
	}

	// By default wait for up to a minute, can be fine-tuned in every way
	// WaitWithTimeout(condition Condition, timeout time.Duration) error
	err = wd.SetImplicitWaitTimeout(time.Second * impTimeout)
	if err != nil {
		log.Fatal(err)
	}

	if err := wd.Get(webSite + "/search?q=WWE+The+Undertaker"); err != nil {
		log.Fatal(err)
	}

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "WWE The Undertaker - Google Search" {
		log.Fatalf("FAIL: Got unexpected title: %v", title)
	} else {
		log.Print("Matching Title..")
	}

	windowID, err := wd.CurrentWindowHandle()
	if err != nil {
		log.Fatal(err)
	}

	err = wd.MaximizeWindow(windowID)
	if err != nil {
		log.Fatal(err)
	}

	acceptAlert, err := wd.FindElement(selenium.ByCSSSelector, "#L2AGLb > div")
	if err != nil {
		log.Fatal(err)
	}

	_, err = acceptAlert.LocationInView()
	if err != nil {
		log.Fatal(err)
	}

	if err = acceptAlert.Click(); err != nil {
		log.Fatal(err)
	}

	images, err := wd.FindElement(selenium.ByCSSSelector, "#hdtb-msb > div:nth-child(1) > div > div:nth-child(2) > a")
	if err != nil {
		log.Fatal(err)
	}

	if err = images.Click(); err != nil {
		log.Fatal(err)
	}

	firstImage, err := wd.FindElement(selenium.ByCSSSelector, "#islrg > div.islrc > div:nth-child(1) > a.wXeWr.islib.nfEiy.mM5pbd > div.bRMDJf.islir > img")
	if err != nil {
		log.Fatal(err)
	}

	srcOne, err := firstImage.GetAttribute("src")
	if err != nil {
		log.Fatal(err)
	}

	secondImage, err := wd.FindElement(selenium.ByCSSSelector, "#islrg > div.islrc > div:nth-child(2) > a.wXeWr.islib.nfEiy.mM5pbd > div.bRMDJf.islir > img")
	if err != nil {
		log.Fatal(err)
	}

	srcTwo, err := secondImage.GetAttribute("src")
	if err != nil {
		log.Fatal(err)
	}

	if err := wd.Get(srcOne); err != nil {
		log.Fatal(err)
	}

	imgBytes, err := wd.Screenshot()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("undertaker.jpg", imgBytes, 0644)
	if err != nil {
		log.Fatal("Could not write image: ", err)
	}

	if err := wd.Get(srcTwo); err != nil {
		log.Fatal(err)
	}

	imgBytes, err = wd.Screenshot()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("undertaker2.jpg", imgBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = wd.Get("https://voggastur.github.io/user-centric-frontend-project/")
	if err != nil {
		log.Fatal(err)
	}

	skillsNav, err := wd.FindElement(selenium.ByCSSSelector, "#navbarSupportedContent > ul > li:nth-child(3) > a")
	if err != nil {
		log.Fatal(err)
	}

	err = skillsNav.Click()
	if err != nil {
		log.Fatal(err)
	}

	skill, err := wd.FindElement(selenium.ByCSSSelector, "#skill-section > div > div:nth-child(2) > h3")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = skill.LocationInView(); err != nil {
		log.Fatal(err)
	}

	copy, err := skill.Text()
	if err != nil {
		log.Fatal(err)
	}

	// Document is a list of strings that will append found skills
	var document []string
	document = append(document, copy)

	// second
	skill, err = wd.FindElement(selenium.ByCSSSelector, "#skill-section > div > div:nth-child(3) > h3")
	if err != nil {
		log.Fatal(err)
	}

	copy, err = skill.Text()
	if err != nil {
		log.Fatal(err)
	}

	document = append(document, copy)

	// third
	skill, err = wd.FindElement(selenium.ByCSSSelector, "#skill-section > div > div:nth-child(4) > h3")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = skill.LocationInView(); err != nil {
		log.Fatal(err)
	}

	copy, err = skill.Text()
	if err != nil {
		log.Fatal(err)
	}

	document = append(document, copy)

	// fourth
	skill, err = wd.FindElement(selenium.ByCSSSelector, "#skill-section > div > div:nth-child(5) > h3")
	if err != nil {
		log.Fatal(err)
	}

	copy, err = skill.Text()
	if err != nil {
		log.Fatal(err)
	}

	document = append(document, copy)

	// fifth
	skill, err = wd.FindElement(selenium.ByCSSSelector, "#skill-section > div > div:nth-child(6) > h3")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = skill.LocationInView(); err != nil {
		log.Fatal(err)
	}

	copy, err = skill.Text()
	if err != nil {
		log.Fatal(err)
	}

	document = append(document, copy)

	// sixth
	skill, err = wd.FindElement(selenium.ByCSSSelector, "#skill-section > div > div:nth-child(7) > h3")
	if err != nil {
		log.Fatal(err)
	}

	copy, err = skill.Text()
	if err != nil {
		log.Fatal(err)
	}

	document = append(document, copy)

	// seventh
	skill, err = wd.FindElement(selenium.ByCSSSelector, "#skill-section > div > div:nth-child(8) > h3")
	if err != nil {
		log.Fatal(err)
	}

	if _, err = skill.LocationInView(); err != nil {
		log.Fatal(err)
	}

	copy, err = skill.Text()
	if err != nil {
		log.Fatal(err)
	}

	document = append(document, copy)

	// eight
	skill, err = wd.FindElement(selenium.ByCSSSelector, "#skill-section > div > div:nth-child(9) > h3")
	if err != nil {
		log.Fatal(err)
	}

	copy, err = skill.Text()
	if err != nil {
		log.Fatal(err)
	}

	document = append(document, copy)

	t.Log(document)

	err = ioutil.WriteFile("skillz.txt", []byte("Johans Skills:\n\n"+strings.Join(document, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}

	t.Log("Shutting Down Webdriver..")
	wd.Quit()
}
