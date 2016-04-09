package main 

// native packages 
import (
    "fmt"
    "os"
    "encoding/json"
    "net/url"
    "log"
)

// external packages
import (
   "github.com/ChimeraCoder/anaconda"
   "github.com/kwkroll32/aimless-apparition/Settings"
)

func main() {
    twitterSettingsFile, err := os.Open("../secret_keys.json")
    if err != nil {
        log.Fatal("cannot open the twitter settings json file: ../secret_keys.json")
    }
    twitterSettings := TwitterSettings.TwitterSettings{}
    jsonParser := json.NewDecoder(twitterSettingsFile)
    if err = jsonParser.Decode(&twitterSettings); err != nil {
        log.Fatal("cannot parse the twitter settings json file: ../secret_keys.json")
    }
    
    anaconda.SetConsumerKey(twitterSettings.ConsumerKey)
    anaconda.SetConsumerSecret(twitterSettings.ConsumerSecret)
    api := anaconda.NewTwitterApi(twitterSettings.AccessToken, twitterSettings.AccessTokenSecret)
    
    v := url.Values{}
    v.Set("count","300")
    searchRes, err := api.GetSearch("tea", v)
    if err != nil{
        log.Fatal(err)
    }
    for _,tweet := range(searchRes.Statuses) {
        fmt.Println(tweet.User.Name)
        fmt.Println(tweet.CreatedAt)
        fmt.Println(tweet.Text )
        fmt.Println()
    }       
}