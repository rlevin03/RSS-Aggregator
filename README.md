# RSS Scraper
The RSS scraper searches a given url for information in a specific format and then provides it to users that follow the feed. The scraper is part of a server that runs on the local machine so that when the feed is updated, so is the information that the user recieves.
## How To Use
The program can be started using the go command like so:
```bash
go run ./
```
After running the program, a server will run indefinitely and scrape the desired websites. At this point, certain api calls can be made to add users to the database, to have users follow a feed, and to add feeds to the database to be scraped. Other actions are also possible and will be addressed in a later section.
## Parts of the Project
### Database
#### Users
The users table stores user information like unique identifiers, creation timestamps, and names. Each user has an api key for authorization. This key is produced using SHA-256 hashing and encoding the hash in hexidecimal format
#### Feeds
The feeds table contains the information of the feeds that needs to be scraped. The url of the feed is stored here and used for scraping. Users can follow feeds which creates a many to many relationship
#### Feed Follows
Connects users to the feeds they follow. Cannot have duplicate feed follows since users cannot follow a feed twice.
#### Posts
Each feed has a number of posts associated with it. Posts have descriptions of the information they contain and this is the information that users will be looking for.
### API Calls
The router supports many api calls to access all the desired functions of the scraper. The url for the router starts with http://localhost:8080/v1 with the following calls supported:
![image](https://github.com/rlevin03/Golang-RSS-Scraper/assets/97414525/5f5ea57a-da85-48de-b182-595cfa41a913)
Some calls have a middleware authentication called on the input variable. This middleware is designed to authenticate incoming requests by extracting the API key from the request headers, retrieving the corresponding user from the database, and then passing control to the actual handler function along with the authenticated user information. This makes sure that no other user can get access to information about another user. When using functions that require verification, the credentials can be passed in to the Headers tab by selecting the Authorization header and typing in the sequence "ApiKey: *user's key*"

All the post requests also require information in the body section of the api call. For the feeds post you will need a name and a valid url:
```bash
{
  "name": "*name*"
  "url": "*valid url*"
}
```
users:
```bash
{
  "name": "*name*"
}
```
feed follows(also requires authentication):
```bash
{
  "feed_id": "*id*"
}
```

In the future, I plan to make a front end where users can sign up, subscribe to feeds, and recieve a summarized list of posts as a source of information




