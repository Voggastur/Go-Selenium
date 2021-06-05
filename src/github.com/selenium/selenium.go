package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

const (
	seleniumPath     = "deps/selenium-server-standalone-3.141.59.jar"
	chromeDriverPath = "deps/chromedriver"
	serverPort       = 4444
	webSite          = "https://www.google.com"
	// Input Parameters for Signup Page
	seed            = "200002022382"         // Skatteverket Test SSN
	email           = "rataveh652@lidte.com" // temp-mail.org
	phone           = "080101010"
	salary          = "123123"
	waitMillisecond = 150
)

var err error
var title string
var currentURL string

/* This test is dependent on the tebeka selenium package that translates Go for Selenium; use "go get github.com/tebeka/selenium"
A chromedriver for your Chrome version is needed from: - https://chromedriver.chromium.org/downloads
Webdrivers for Safari, Opera, IE and Firefox can also be used
A seleniun-server-standalone-3.141 is needed and can be downloaded at:
- https://www.selenium.dev/downloads/
The seleniumPath and chromeDriverPath should match the names of the files */
func main() {

	opts := []selenium.ServiceOption{
		selenium.ChromeDriver(chromeDriverPath),
		selenium.Output(os.Stderr),
	}

	selenium.SetDebug(true)
	service, err := selenium.NewSeleniumService(seleniumPath, serverPort, opts...)
	if err != nil {
		log.Fatal(err)
	}

	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", serverPort))
	if err != nil {
		log.Fatal(err)
	}

	if err := wd.Get(webSite); err != nil {
		log.Fatal(err)
	}

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Google" {
		log.Fatalf("FAIL: Unexpected Title, Got: %v, Expected: Google", title)
	}

	windowID, err := wd.CurrentWindowHandle()
	if err != nil {
		log.Fatal(err)
	}

	err = wd.MaximizeWindow(windowID)
	if err != nil {
		log.Fatal(err)
	}

	searchbar, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[3]/form/div[1]/div[1]/div[1]/div/div[2]/input")
	if err != nil {
		log.Fatal(err)
	}

	if err := searchbar.Click(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Millisecond * waitMillisecond)

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Bli medlem | Pensionera" {
		log.Fatalf("FAIL: Unexpected Title, Got: %v, Expected: Bli medlem | Pensionera", title)
	}

	currentURL, err = wd.CurrentURL()
	if err != nil {
		log.Fatal(err)
	}

	if currentURL != webSite+"/bli-medlem" {
		log.Fatalf("FAIL: Not the expected URL, got: %v wanted: %v/bli-medlem", currentURL, webSite)
	}

	//Personnummer input
	seedInput, err := wd.FindElement(selenium.ByID, "seed")
	if err != nil {
		log.Fatal(err)
	}

	//Email input
	emailInput, err := wd.FindElement(selenium.ByID, "email")
	if err != nil {
		log.Fatal(err)
	}

	//Phonenumber input
	phoneInput, err := wd.FindElement(selenium.ByID, "phone-mobile")
	if err != nil {
		log.Fatal(err)
	}

	//Salary input
	salaryInput, err := wd.FindElement(selenium.ByID, "salary")
	if err != nil {
		log.Fatal(err)
	}

	//Enter characters
	err = seedInput.SendKeys(seed)
	if err != nil {
		log.Fatal(err)
	}

	err = emailInput.SendKeys(email)
	if err != nil {
		log.Fatal(err)
	}

	err = phoneInput.SendKeys(phone)
	if err != nil {
		log.Fatal(err)
	}

	err = salaryInput.SendKeys(salary)
	if err != nil {
		log.Fatal(err)
	}

	button, err := wd.FindElement(selenium.ByCSSSelector, "button.btn.btn-sign")
	if err != nil {
		log.Fatal(err)
	}

	if err := button.Click(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Millisecond * waitMillisecond)
	fmt.Print("Signup process finished, redirecting to account/onboarding")

	// Click central button satt igang for quiz on risks
	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Välkommen! | Pensionera" {
		log.Printf("FAIL: Unexpected Title, Got: %v, Expected: Välkommen | Pensionera", title)
	}

	currentURL, err = wd.CurrentURL()
	if err != nil {
		log.Fatal(err)
	}

	if currentURL != webSite+"/account/onboarding" {
		log.Fatal("FAIL: Not the expected URL")
	}

	sattigang, err := wd.FindElement(selenium.ByCSSSelector, "#onboarding-page > div > div:nth-child(1) > div.risk-attitude-card > div.risk-attitude-form > button")
	if err != nil {
		log.Fatal(err)
	}

	err = sattigang.Click()
	if err != nil {
		log.Fatal(err)
	}

	// Satt igang modal opened, find radio buttons and click
	// Q1 Identify, Scroll and Click
	q1option1, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-body > div > div > div > ul:nth-child(1) > li > div:nth-child(2) > div > label")
	if err != nil {
		log.Fatal(err)
	}

	err = q1option1.Click()
	if err != nil {
		log.Fatal(err)
	}

	// Q2 Identify, Scroll and Click
	q2option2, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-body > div > div > div > ul:nth-child(2) > li > div:nth-child(3) > div > label")
	if err != nil {
		log.Fatal(err)
	}

	q2option2.LocationInView()
	err = q2option2.Click()
	if err != nil {
		log.Fatal(err)
	}

	// Q3 Identify, Scroll and Click
	q3option1, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-body > div > div > div > ul:nth-child(3) > li > div:nth-child(2) > div > label")
	if err != nil {
		log.Fatal(err)
	}

	q3option1.LocationInView()
	err = q3option1.Click()
	if err != nil {
		log.Fatal(err)
	}

	// Q4 Identify, Scroll and Click
	q4option1, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-body > div > div > div > ul:nth-child(4) > li > div:nth-child(2) > div > label")
	if err != nil {
		log.Fatal(err)
	}

	q4option1.LocationInView()
	err = q4option1.Click()
	if err != nil {
		log.Fatal(err)
	}

	// Q5 Identify, Scroll and Click
	q5option1, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-body > div > div > div > ul:nth-child(5) > li > div:nth-child(4) > div > label")
	if err != nil {
		log.Fatal(err)
	}

	q5option1.LocationInView()
	err = q5option1.Click()
	if err != nil {
		log.Fatal(err)
	}

	// Q6 Identify, Scroll and Click
	q6option2, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-body > div > div > div > ul:nth-child(6) > li > div:nth-child(3) > div > label")
	if err != nil {
		log.Fatal(err)
	}

	q6option2.LocationInView()
	err = q6option2.Click()
	if err != nil {
		log.Fatal(err)
	}

	// Q7 Identify, Scroll and Click - Last Q
	q7option4, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-body > div > div > div > ul:nth-child(7) > li > div:nth-child(5) > div > label")
	if err != nil {
		log.Fatal(err)
	}

	q7option4.LocationInView()
	err = q7option4.Click()
	if err != nil {
		log.Fatal(err)
	}

	// Press Spara in Modal
	savequiz, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-footer > div > button")
	if err != nil {
		log.Fatal(err)
	}

	savequiz.LocationInView()
	err = savequiz.Click()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Millisecond * waitMillisecond)

	// Confirm Risk Management Level 4 by Pressing Gå Vidare Button
	confirm, err := wd.FindElement(selenium.ByCSSSelector, "#risk-attitude-modal > div > div > div > div.new-modal-footer > div > div > button")
	if err != nil {
		log.Fatal(err)
	}

	if err = confirm.Click(); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Millisecond * waitMillisecond)

	// Identify and click each of 6 nav links and compare URL and title

	// Historik

	historik, err := wd.FindElement(selenium.ByCSSSelector, ".account-menu > div > div.menu > ul > li:nth-child(6) > a")
	if err != nil {
		log.Fatal(err)
	}

	if err := historik.Click(); err != nil {
		log.Fatal(err)
	}

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Historik | Pensionera" {
		fmt.Printf("FAIL: Unexpected Title, Got: %v, Expected: Historik | Pensionera", title)
	}

	currentURL, err = wd.CurrentURL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(currentURL)

	if currentURL != webSite+"/account/historik" {
		log.Fatalf("FAIL: Not the expected URL, got: %s", currentURL)
	}

	time.Sleep(time.Millisecond * waitMillisecond)

	// Prognos

	prognos, err := wd.FindElement(selenium.ByCSSSelector, ".account-menu > div > div.menu > ul > li:nth-child(5) > a")
	if err != nil {
		log.Fatal(err)
	}

	if err := prognos.Click(); err != nil {
		log.Fatal(err)
	}

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Pensionsprognos | Pensionera" {
		fmt.Printf("FAIL: Unexpected Title, Got: %v, Expected: Pensionsprognos | Pensionera", title)
	}

	currentURL, err = wd.CurrentURL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(currentURL)

	if currentURL != webSite+"/account/pensionsprognos" {
		log.Fatalf("FAIL: Not the expected URL, got: %v", currentURL)
	}

	time.Sleep(time.Millisecond * waitMillisecond)

	// Mina Abbonemang

	abbonemang, err := wd.FindElement(selenium.ByCSSSelector, ".account-menu > div > div.menu > ul > li:nth-child(4) > a")
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := abbonemang.Click(); err != nil {
		log.Fatal(err.Error())
	}

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Välkommen! | Pensionera" {
		fmt.Printf("FAIL: Unexpected Title, Got: %v, Expected: Välkommen! | Pensionera", title)
	}

	currentURL, err = wd.CurrentURL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(currentURL)

	if currentURL != webSite+"/account/onboarding" {
		fmt.Printf("FAIL: Not the expected URL, got: %v, wanted: %v/account/onboarding", currentURL, webSite)
	}

	time.Sleep(time.Millisecond * waitMillisecond)

	// Mina Tillgångar

	tillgangar, err := wd.FindElement(selenium.ByCSSSelector, ".account-menu > div > div.menu > ul > li:nth-child(3) > a")
	if err != nil {
		log.Fatal(err)
	}

	if err := tillgangar.Click(); err != nil {
		log.Fatalf(err.Error())
	}

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Mina Tillgångar | Pensionera" {
		log.Printf("FAIL: Unexpected Title, Got: %v, Expected: Mina Tillgångar | Pensionera", title)
	}

	currentURL, err = wd.CurrentURL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(currentURL)

	if currentURL != webSite+"/account/mina-tillgangar" {
		log.Fatal("FAIL: Not the expected URL")
	}

	time.Sleep(time.Millisecond * waitMillisecond)

	// Min Pension

	pension, err := wd.FindElement(selenium.ByCSSSelector, ".account-menu > div > div.menu > ul > li:nth-child(2) > a")
	if err != nil {
		log.Fatal(err)
	}

	if err := pension.Click(); err != nil {
		log.Fatal(err.Error())
	}

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Min Pension | Pensionera" {
		log.Printf("FAIL: Unexpected Title, Got: %v, Expected: Min Pension | Pensionera", title)
	}

	currentURL, err = wd.CurrentURL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(currentURL)

	if currentURL != webSite+"/account/min-pension" {
		log.Fatal("FAIL: Not the expected URL")
	}

	// Att Göra

	gora, err := wd.FindElement(selenium.ByCSSSelector, ".account-menu > div > div.menu > ul > li:nth-child(1) > a")
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := gora.Click(); err != nil {
		log.Fatalf(err.Error())
	}

	title, err = wd.Title()
	if err != nil {
		log.Fatal(err)
	}

	if title != "Välkommen! | Pensionera" {
		log.Printf("FAIL: Unexpected Title, Got: %v, Expected: Välkommen! | Pensionera", title)
	}

	currentURL, err = wd.CurrentURL()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(currentURL)

	if currentURL != webSite+"/account/attgora" {
		log.Fatal("FAIL: Not the expected URL")
	}

	time.Sleep(time.Millisecond * waitMillisecond)
	fmt.Print("Shutting Down Webdriver..")
	wd.Close()
}
