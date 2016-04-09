package aimless 

// native packages 
import (
    "os"
    "encoding/json"
    "net/url"
    "log"
    "strings"
    "strconv"
    //"fmt"
)

// external packages
import (
   "github.com/ChimeraCoder/anaconda"
   "github.com/kwkroll32/aimless-apparition/Settings"
)

// LaunchMyAPI will grab tokens from the secret json config and return a pointer to the api instance
func LaunchMyAPI() *anaconda.TwitterApi {
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
    return api
}

// Search will implement a search in the given api for the requested number of tweets matching a keyword
func Search(api *anaconda.TwitterApi ,term string, count int) []string {
    v := url.Values{}
    v.Set("count",strconv.Itoa(count))
    searchRes, err := api.GetSearch("tea", v)
    if err != nil{
        log.Fatal(err)
    }
    var text string
    var outlist []string
    for _,tweet := range(searchRes.Statuses) {
        text = tweet.Text
        for _,res := range(ExtractWordFromTweet(term, text)) {
            outlist = append(outlist, res)
        }
    }
    return outlist
}

// ExtractWordFromTweet extracts the word that comes after the search term
func ExtractWordFromTweet(word, tweet string) []string {
    // init a nil string incase the word cannot be found in the tweet 
    var outstrings []string 
    if strings.Contains( strings.ToLower(tweet), strings.ToLower(word)) {
        // check all words in the tweet for matches (case insensitive)
        // append the next word to the out array
        tweetArray := strings.Fields(tweet)
        for tweetPos,tweetWord := range(tweetArray) {
            if (strings.EqualFold(word,tweetWord)) && (tweetPos < len(tweetArray)-1) {
                outstrings = append(outstrings, tweetArray[tweetPos+1])
            }
        }     
    }
    
    return outstrings
}
