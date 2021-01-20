from selenium import webdriver
import sys, time


# Create an instance of Chrome WebDriver
driver = webdriver.Chrome(executable_path='/webdriver/chromedriver')
driver.maximize_window()

# Passed Tests and Total Tests counter variables
passed = 0
total = 0

driver.get('https://pensionera.se/bli-medlem')

# Check the title of the page
if(driver.title=="Bli medlem | Pensionera"):
    print ("Expected title found")
    passed += 1
else:
    print ("Unexpected title found")
    total += 1


# Identifiers

# Navbar 1 - Prisplaner
nav1 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[1]/a")
# Navbar 2 - FAQ
nav2 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[2]/a")
# Navbar Dropdown 3 - Om Pensionera
nav3 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]")

# Dropdown 1 - Om Pensionera
nav31 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[1]/a")
# Dropdown 2 - Pensioneras Affärsmodell
nav32 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[2]/a")
# Dropdown 3 - Avgifter och Provision
nav33 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[3]/a")
# Dropdown 4 - Lediga Tjänster
nav34 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[4]/a")
# Dropdown 5 - Media
nav35 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[5]/a")
# Dropdown 6 - Användarvillkor
nav36 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[6]/a")
# Dropdown 7 - Integritetspolicy
nav37 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[7]/a")
# Dropdown 8 - Cookies
nav38 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[8]")
# Dropdown 9 - Kundnöjdhet
nav39 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[9]/a")
# Dropdown 10 - PensioneraScore
nav310 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]/ul/li[10]/a")

# Navbar 4 - Kontakta Oss
nav4 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[4]/a")
# Navbar 5 - Öppna Sparkonto
nav5 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[5]/a")
# Navbar 6 - Teckna Livförsäkring
nav6 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[6]/a")
# Navbar 7 - Nyheter
nav7 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[7]/a")
# Navbar 8 - Bli Medlem
nav8 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[8]/a")
# Navbar 9 - Logga In
nav9 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[9]/a")


# Nav 1 test
nav1.click()
time.sleep(2)
if (driver.current_url=='https://pensionera.se/prisplaner'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)


# Nav 2 test
nav2.click()
time.sleep(1)
if (driver.current_url=='https://pensionera.se/faq'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)


# Nav 3 Dropdown 1
nav3.click()  
time.sleep(1)
nav31.click()
if (driver.current_url=='https://pensionera.se/om-pensionera'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 2
nav3.click()  
time.sleep(1)
nav32.click()
if (driver.current_url=='https://pensionera.se/affarsmodell'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 3
nav3.click()  
time.sleep(1)
nav33.click()
if (driver.current_url=='https://pensionera.se/avgifter-och-provision'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 4
nav3.click()  
time.sleep(1)
nav34.click()
if (driver.current_url=='https://pensionera.se/lediga-tjanster'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 5
nav3.click()  
time.sleep(1)
nav35.click()
if (driver.current_url=='https://pensionera.se/media'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 6
nav3.click()  
time.sleep(1)
nav36.click()
if (driver.current_url=='https://pensionera.se/anvandarvillkor'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 7
nav3.click()  
time.sleep(1)
nav37.click()
if (driver.current_url=='https://pensionera.se/integritetspolicy'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 8
nav3.click()  
time.sleep(1)
nav38.click()
if (driver.current_url=='https://pensionera.se/cookies'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 9
nav3.click()  
time.sleep(1)
nav39.click()
if (driver.current_url=='https://pensionera.se/kundnojdhet'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

# Nav 3 Dropdown 10
nav3.click()  
time.sleep(1)
nav310.click()
if (driver.current_url=='https://pensionera.se/pensionera-score'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

#End of Dropdown tests


# Nav 4 test
nav4.click()
time.sleep(1)
if (driver.current_url=='https://pensionera.se/kontakt'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)


# Nav 5 test
nav5.click()
time.sleep(1)
if (driver.current_url=='https://pensionera.se/om-sparkonto-isk'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)


# Nav 6 test
nav6.click()
time.sleep(1)
if (driver.current_url=='https://liv.pensionera.se/'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)


# Nav 7 test
nav7.click()
time.sleep(1)
if (driver.current_url=='https://pensionera.se/guide'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)


# Nav 8 test
nav8.click()
time.sleep(1)
if (driver.current_url=='https://pensionera.se/bli-medlem'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1


# Nav 9 test
nav9.click()
time.sleep(1)
if (driver.current_url=='https://pensionera.se/logga-in'):
    print("Success")
    passed += 1
else:
    print("Failure")
    total += 1

# Return
driver.get('https://pensionera.se/bli-medlem')
time.sleep(1)

#Navbar test finished
print("Navbar test finished")


# Input Element Identifiers
inputseed = driver.find_element_by_id("seed")
inputemail = driver.find_element_by_id("email")
inputphone = driver.find_element_by_id("phone-mobile")
inputsalary = driver.find_element_by_id("salary")
sendbutton = driver.find_element_by_xpath("/html/body/div[1]/div/div[1]/div/div[2]/div[1]/div/button")

inputseed.send_keys("Donald Duck")
# temp email
inputemail.send_keys("wahekon130@cocyo.com")
sendbutton.click()
time.sleep(1)




print("Number of passed tests: " + passed + ", Number of total tests: " + total)
time.sleep(3)

driver.close() 
