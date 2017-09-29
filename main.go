package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"strconv"
)

// API URLs
const (
	AdURL         = "http://ad.9gag.com"
	AdminURL      = "http://admin.9gag.com"
	APIURL        = "http://api.9gag.com"
	CommentURL    = "http://comment.9gag.com"
	CommentCDNURL = "http://comment-cdn.9gag.com"
	NotifyURL     = "http://notify.9gag.com"
)

// API endpoint paths
const (
	LoginPath    = "/v2/user-token/"
	PostListPath = "/v2/post-list/"
)

// App consts
const (
	appID      = "com.ninegag.android.app"
	deviceType = "android"
	bucketName = "MAIN_RELEASE"
	commentCDN = "a_dd8f2b7d304a10edaf6f29517ea0ca4100a43d1b"
)

var (
	appToken, deviceUUID string

	userData *UserData
)

func init() {
	appToken = randomSHA1HexStr()
	deviceUUID = randomUUIDHexStr()
}

func main() {
	login("someusername", "somepassword")
	fmt.Printf("Welcome, %s\n", userData.FullName)
	getHot(10)
}

func login(username, password string) {
	args := map[string]string{
		"loginMethod": "9gag",
		"loginName":   username,
		"password":    getMD5HexStr(password),
		"pushToken":   randomSHA1HexStr(),
		"language":    "en_US",
	}
	respBody, err := requestGET(APIURL+LoginPath+urlArgsStr(args), true)
	if err != nil {
		log.Fatal(err)
	}

	resp := LoginResponse{}
	if err = json.Unmarshal(respBody, &resp); err != nil {
		log.Fatal(err)
	}

	appToken = resp.Data.UserToken
	userData = &resp.Data.User
}

func getHot(count int) {
	args := map[string]string{
		"group":      "1",
		"type":       "hot",
		"itemCount":  strconv.Itoa(count),
		"entryTypes": "animated,photo,video,album",
		"offset":     "10",
	}
	respBody, err := requestGET(APIURL+PostListPath+urlArgsStr(args), true)
	if err != nil {
		log.Fatal(err)
	}

	resp := PostListResponse{}
	if err = json.Unmarshal(respBody, &resp); err != nil {
		log.Fatal(err)
	}

	for _, p := range resp.Data.Posts {
		fmt.Println(html.UnescapeString(p.Title), p.URL, p.Images.Image700.URL)
	}
}
