# Pexels API Go Client
This Go application demonstrates how to use the Pexels API to search for photos and videos. It provides functionalities to search for photos and videos based on a query, retrieve random curated photos and popular videos, and get information about a specific photo or video.
# Features
1. Search for Photos: Search for photos by a query, with options for specifying the number of results per page and the page number.
2. Retrieve Curated Photos: Retrieve a list of curated photos.
3. Get a Random Curated Photo: Get a random photo from the curated collection.
4. Search for Videos: Search for videos by a query, with options for specifying the number of results per page and the page number.
5. Retrieve Popular Videos: Retrieve a list of popular videos.
6. Get a Random Popular Video: Get a random popular video.
7. Get Information About a Specific Photo or Video: Retrieve detailed information about a specific photo or video by its ID.
# Prerequisites
Before you begin, ensure you have the following:
1. Go installed on your local machine.
2. A Pexels API key. You can obtain one by signing up on the [Pexels API website.](https://www.pexels.com/api/)
# Installation
1. Clone the repository to your local machine:
git clone https://github.com/bakhtybayevn/pexels-api-go-client.git
2. Navigate to the project directory:
cd pexels-api-go-client
3. Set your Pexels API key as an environment variable. Replace your_api_key with your actual API key:
export PEXELS_API_KEY=your_api_key
Alternatively, you can set your API key in a .env file in the project directory.
4. Build and run the Go application:
   go build -o pexels-app
   ./pexels-app
This will compile and execute the application, making requests to the Pexels API.
# Usage
The application provides the following features:
1. Search for photos and videos based on a query.
2. Retrieve curated photos and popular videos.
3. Get a random curated photo or popular video.
4. Get information about a specific photo or video.
To use these features, you can refer to the code in main.go for examples. Here's a brief overview:
# Search for Photos:
res, err := c.SearchPhotos("nature", 15, 1)
# Retrieve Curated Photos:
res, err := c.CuratedPhotos(15, 1)
# Get a Random Curated Photo:
photo, err := c.GetRandomPhoto()
# Search for Videos:
res, err := c.SearchVideo("nature", 15, 1)
# Retrieve Popular Videos:
res, err := c.PopularVideo(15, 1)
# Get a Random Popular Video:
video, err := c.GetRandomVideo()
# Get Information About a Specific Photo or Video:
photo, err := c.GetPhoto(12345) // Replace 12345 with the actual photo ID
video, err := c.GetVideo(54321) // Replace 54321 with the actual video ID
# Contributing
Contributions are welcome! If you have suggestions for improvements or new features, please create a pull request.
