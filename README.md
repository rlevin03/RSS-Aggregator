# RSS Scraper
The RSS scraper searches a given url for information in a specific format and then provides it to users that follow the feed. The scraper is part of a server that runs on the local machine so that when the feed is updated, so is the information that the user recieves.
## How To Use
The program can be started using the go command like so:
```bash
go run ./
```
After running the program, a server will run indefinitely and scrape the desired websites. At this point, certain api calls can be made to add users to the database, to have users follow a feed, and to add feeds to the database to be scraped. Other actions are also possible and will be addressed in a later section.
