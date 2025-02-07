package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"os"
)


type GithubUser struct {
 type string `json "type"`
	Repo struct {
		Name string `json: "name"`
	} `json: "repo"`
}


func main () {
	user := flag.String ("u", "", "Github user")
	flag.Parse()

	if user == "" {
		fmt.Println("Github user u")
		os.Exit(1)
	}
	
	fmt.Println("user dont type right user", user)
}

url := fmt.Sprintf("https://api.github.com", user)
resp, err := http.Get(url)
if err != nil {
	fmt.Println("cant get api", err)
	os.Exit(1)
}
defer resp.Body.Close()

if repo.StatusCode != 200 {
	fmt.Println("not found the api")
	os.Exit(1)
}

body, err := ioutil.ReadALL(resp.Body)
if err != nil {
	log.Fatal(err)
}
var events []GithubUser 
json.Unmarshal(body, &events)

fmt.Println("user get api", user)
for, events := range events {
	fmt.Printf("%s in %s/n", event.type, event.Repo.Name)
}