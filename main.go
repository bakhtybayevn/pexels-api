package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Constants for Pexels API endpoints.
const (
	PhotoAPI = "https://api.pexels.com/v1/"
	VideoAPI = "https://api.pexels.com/videos/"
)

// Struct representing a Pexels API client.
type Client struct {
	Token          string
	HTTPClient     http.Client
	RemainingTimes int32
}

// Create a new Pexels API client with a given API token.
func NewClient(token string) *Client {
	return &Client{
		Token:      token,
		HTTPClient: http.Client{},
	}
}

// Struct for search results of photos.
type SearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Photos       []Photo `json:"photos"`
}

// Struct representing a photo.
type Photo struct {
	ID              int32       `json:"id"`
	Width           int32       `json:"width"`
	Height          int32       `json:"height"`
	Url             string      `json:"url"`
	Photographer    string      `json:"photographer"`
	PhotographerURL string      `json:"photographer_url"`
	Src             PhotoSource `json:"src"`
}

// Struct representing photo sources.
type PhotoSource struct {
	Original  string `json:"original"`
	Large     string `json:"large"`
	Large2x   string `json:"large2x"`
	Medium    string `json:"medium"`
	Small     string `json:"small"`
	Portrait  string `json:"portrait"`
	Square    string `json:"square"`
	Landscape string `json:"landscape"`
	Tiny      string `json:"tiny"`
}

// Struct for curated photo results.
type CuratedResult struct {
	Page     int32   `json:"page"`
	PerPage  int32   `json:"per_page"`
	NextPage string  `json:"next_page"`
	Photos   []Photo `json:"photos"`
}

// Struct for search results of videos.
type VideoSearchResult struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	NextPage     string  `json:"next_page"`
	Videos       []Video `json:"videos"`
}

// Struct representing a video.
type Video struct {
	ID            int32          `json:"id"`
	Width         int32          `json:"width"`
	Height        int32          `json:"height"`
	Url           string         `json:"url"`
	Image         string         `json:"image"`
	FullRes       interface{}    `json:"full_res"`
	Duration      float64        `json:"duration"`
	VideoFiles    []VideoFiles   `json:"video_files"`
	VideoPictures []VideoPicture `json:"video_pictures"`
}

// Struct for popular video results.
type PopularVideos struct {
	Page         int32   `json:"page"`
	PerPage      int32   `json:"per_page"`
	TotalResults int32   `json:"total_results"`
	Url          string  `json:"url"`
	Videos       []Video `json:"videos"`
}

// Struct representing video files.
type VideoFiles struct {
	ID       int32  `json:"id"`
	Quality  string `json:"quality"`
	FileType string `json:"file_type"`
	Width    int32  `json:"width"`
	Height   int32  `json:"height"`
	Link     string `json:"link"`
}

// Struct representing video pictures.
type VideoPicture struct {
	ID      int32  `json:"id"`
	Picture string `json:"picture"`
	Nr      int32  `json:"nr"`
}

// Perform an HTTP request with authentication headers.
func (c *Client) requestDoWithAuth(method, url string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.Token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	times, err := strconv.Atoi(res.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return nil, err
	}

	c.RemainingTimes = int32(times)

	return res, nil
}

// Search for photos based on a query, page, and perPage parameters.
func (c *Client) SearchPhotos(query string, perPage, page int) (*SearchResult, error) {
	url := fmt.Sprintf(PhotoAPI+"search?query=%s&per_page=%d&page=%d", query, perPage, page)
	res, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result SearchResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Retrieve curated photos based on page and perPage parameters.
func (c *Client) CuratedPhotos(perPage, page int) (*CuratedResult, error) {
	url := fmt.Sprintf(PhotoAPI+"curated?per_page=%d&page=%d", perPage, page)
	res, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result CuratedResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Retrieve photo information by ID.
func (c *Client) GetPhoto(id int32) (*Photo, error) {
	url := fmt.Sprintf(PhotoAPI+"photos/%d", id)
	res, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result Photo
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Get a random curated photo.
func (c *Client) GetRandomPhoto() (*Photo, error) {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1000)
	res, err := c.CuratedPhotos(1, randNum)
	if err != nil {
		return nil, err
	}

	return &res.Photos[0], nil
}

// Search for videos based on a query, page, and perPage parameters.
func (c *Client) SearchVideo(query string, perPage, page int) (*VideoSearchResult, error) {
	url := fmt.Sprintf(VideoAPI+"search?query=%s&per_page=%d&page=%d", query, perPage, page)
	res, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result VideoSearchResult
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Retrieve popular videos based on page and perPage parameters.
func (c *Client) PopularVideo(perPage, page int) (*PopularVideos, error) {
	url := fmt.Sprintf(VideoAPI+"popular?per_page=%d&page=%d", perPage, page)
	res, err := c.requestDoWithAuth("GET", url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var result PopularVideos
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Get a random popular video.
func (c *Client) GetRandomVideo() (*Video, error) {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(1000)
	res, err := c.PopularVideo(1, randNum)
	if err != nil {
		return nil, err
	}

	return &res.Videos[0], nil
}

// Get the remaining API requests available for this month.
func (c *Client) GetRemainingRequestsInThisMonth() int32 {
	return c.RemainingTimes
}

func main() {
	os.Setenv("PEXELS_API_KEY", "yjvbIk7SvHQEJfYIPWkDA9TcI7GYj5IfHpB4UxrVSDkWb2UsgtUvl8Z2")
	var Token = os.Getenv("PEXELS_API_KEY")

	var c = NewClient(Token)

	res, err := c.SearchVideo("nature", 15, 1)
	if err != nil {
		fmt.Errorf("Error: %v", err)
	}
	if res.Page == 0 {
		fmt.Println("Error: Page is 0")
	}

	fmt.Println(res)
}
