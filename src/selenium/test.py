from selenium import webdriver
import time

# Create an instance of Firefox WebDriver
driver = webdriver.Chrome()

# KEY POINT: The driver.get method will navigate to a page given by the URL
driver.get('https://pensionera.se/bli-medlem')

# Check the title of the page
if(driver.title=="Bli medlem | Pensionera"):
    print ("Expected title found")
else:
    print ("Unexpected title found") 


# Identifiers for Nav Items

# Navbar 1 - Prisplaner
nav1 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[1]/a")
# Navbar 2 - FAQ
nav2 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[2]/a")
# Navbar Dropdown 3 - Om Pensionera
nav3 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[3]")

# Dropdown items
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
nav8 = driver.find_element_by_xpath("/html/body/div[1]/div/nav/div/div[2]/div/ul/li[9]/a")

nav1.click()
# Wait for load
time.sleep(2)

driver.close() 
