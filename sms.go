package main

import (
  "fmt"
  "strings"
  "math/rand"
  "time"
  "net/http"
  "net/url"
  "encoding/json"
  // "io/ioutil"
)

// type Main struct {
// 	Name string `json:"name"`
// 	Abilities [][]string `json:"abilities"`
// }


func main() {

//   quoteurl := "https://en.wikiquote.org/w/api.php"

  // response, err := http.Get("https://en.wikiquote.org/w/api.php")
	
  //   if err != nil {
  //       fmt.Printf("The HTTP request failed with error %s\n", err)
  //   } else {
	// 	data, _ := ioutil.ReadAll(response.Body)
	// 	var responseObject Main
	// 	json.Unmarshal(data, &responseObject)
	// 	fmt.Println(responseObject.Name)
	// 	fmt.Println(len(responseObject.Abilities))
  //   }


  // My account
  accountSid := "Your SID"
  authToken := "Your token"
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

  // Create possible message bodies
  quote := [7]string{
    `- I love you the more in that I believe you had liked me for my own sake and for nothing else. John Keats`,
    "- But man is not made for defeat. A man can be destroyed but not defeated. Ernest Hemingway",
    "- When you reach the end of your rope, tie a knot in it and hang on. Franklin D. Roosevelt",
    "- There is nothing permanent except change. Heraclitus",
    "- You cannot shake hands with a clenched fist. Indira Gandhi",
    "- Let us sacrifice our today so that our children can have a better tomorrow. A. P. J. Abdul Kalam",
    "- Do not mind anything that anyone tells you about anyone else. Judge everyone and everything for yourself. Henry James",
  }

  //random numbers
  rand.Seed(time.Now().Unix())

  // Make message
  message := url.Values{}
  message.Set("To","+33 7 67 32 55 50")
  message.Set("From","+15057154914")
  message.Set("Body",quote[rand.Intn(len(quote))])
  messageReader := *strings.NewReader(message.Encode())

  // HTTP request client
  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, &messageReader)
  req.SetBasicAuth(accountSid, authToken)
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  // HTTP POST request 
  // Newdecoder() reads and decodes JSON values from an input.
  resp, _ := client.Do(req)
  // fmt.Println(resp.Status)
  if (resp.StatusCode < 200 && resp.StatusCode < 300) {
    var data map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode(&data)
    if (err == nil) {
      fmt.Println(data["works"])
    }
  } else {
    fmt.Println(resp.Status);
  }
}