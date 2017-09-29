package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/nezorflame/ninegago"
)

var (
	username, password string

	apiClient *ninegago.APIClient
)

func init() {
	flag.StringVar(&username, "u", "", "9GAG username")
	flag.StringVar(&password, "p", "", "9GAG password")
	flag.Parse()

	if username == "" || password == "" {
		username, password = flag.Arg(0), flag.Arg(1)
	}

	if username == "" || password == "" {
		flag.Usage()
		log.Fatal("Wrong credentials")
	}
}

func main() {
	apiClient = ninegago.NewAPIClient()
	if err := apiClient.Login(username, password); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Welcome, %s\n", apiClient.User.FullName)

	if posts, err := apiClient.GetHotPosts("hot", 10); err == nil {
		for _, p := range posts {
			fmt.Println(p.Title, p.URL, p.Images.Image700.URL)
		}
	} else {
		log.Fatal(err)
	}
}
