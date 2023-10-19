# GreenScraper
***
GreenScraper is a short program that will gather CNCF Conferences talks and sessions. The principal design was geared towards Sustainability specific topics. However, the introduction of a ```keywords.txt``` file, allows for additional use cases.

## Usage

* clone this repository and ```cd``` into the root directory.

- Build GreenScraper: ```go build```.
- Keywords are read by GreenScraper from file ```keywords.txt```. Add the desired keywords, one by line.
- Add urls to scrape to file ```urls.txt```
- Run GreenScraper ```./greenscraper```.
- A list of talks and links to it will be outputted to STDOUT

## Inspiration

GreenScraper is inspired by the work of Al-Hussein Hameed Jasim that can be found at:
https://github.com/Al-HusseinHameedJasim/green-talks-scraper/tree/main#green-talks-scraper

***
**Note**: GreenScraper aligns to the requirements and recommendations of Green Software Development.