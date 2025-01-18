from requests_html import HTMLSession

session = HTMLSession()

request = session.get("https://www.amazon.com/dp/B08DWD38VX")

request.html.render(sleep=3, timeout = 20)

price = request.html.find("#a-price-whole")[0].text

print(price)