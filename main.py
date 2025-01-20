from http.client import HTTPResponse
import urllib.request
from bs4 import BeautifulSoup
from parser import parse_recipes

def fetch_and_clean_webpage(url):
    # Create headers to mimic a browser request
    headers = {
        'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3'
    }
    
    # Create a Request object with headers
    req = urllib.request.Request(url, headers=headers)
    
    # Fetch the webpage with headers
    response: HTTPResponse = urllib.request.urlopen(req)
    
    # Parse the HTML
    soup = BeautifulSoup(response.read(), 'html.parser')
    
    # Remove all script and style tags
    for script in soup.find_all(['script', 'style', 'link', 'svg', 'img', 'form', 'noscript']):
        script.decompose()
    
    return soup.prettify()

def main():
    try:
        url = input("Enter the webpage URL: ")
        cleaned_html = fetch_and_clean_webpage(url)
        recipe = parse_recipes(cleaned_html)
        print(recipe)
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    main()
