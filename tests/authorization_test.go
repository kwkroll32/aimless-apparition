package main 

// native packages 
import (
    //"fmt"
    "os"
    "testing"
    "encoding/json"
    //"strconv"
    "reflect"
)

// external packages
import (
   "github.com/ChimeraCoder/anaconda"
   "github.com/kwkroll32/aimless-apparition/Settings"
   "github.com/kwkroll32/aimless-apparition/aimless"
)

func TestParseTwitterSettingsJson(t *testing.T) {
    twitterSettingsFile, err := os.Open("../secret_keys.json")
    if err != nil {
        t.Errorf("cannot open the twitter settings json file: ../secret_keys.json")
    }
    twitterSettings := TwitterSettings.TwitterSettings{}
    jsonParser := json.NewDecoder(twitterSettingsFile)
    if err = jsonParser.Decode(&twitterSettings); err != nil {
        t.Errorf("cannot parse the twitter settings json file: ../secret_keys.json")
    }
    if len(twitterSettings.AccessToken) == 0 {
        t.Errorf("twitter access token is empty!")
    }
    if len(twitterSettings.AccessTokenSecret) == 0 {
        t.Errorf("twitter access token secret is empty!")
    }
    if len(twitterSettings.ConsumerKey) == 0 {
        t.Errorf("twitter consumer key is empty!")
    }
    if len(twitterSettings.ConsumerSecret) == 0 {
        t.Errorf("twitter consumer secret is empty!")
    }
}

func TestLaunchNewAPI(t *testing.T) {
    twitterSettingsFile, err := os.Open("../secret_keys.json")
    if err != nil {
        t.Errorf("cannot open the twitter settings json file: ../secret_keys.json")
    }
    defer twitterSettingsFile.Close()
    twitterSettings := TwitterSettings.TwitterSettings{}
    jsonParser := json.NewDecoder(twitterSettingsFile)
    if err = jsonParser.Decode(&twitterSettings); err != nil {
        t.Errorf("cannot parse the twitter settings json file: ../secret_keys.json")
    }
    
    anaconda.SetConsumerKey(twitterSettings.ConsumerKey)
    anaconda.SetConsumerSecret(twitterSettings.ConsumerSecret)
    api := anaconda.NewTwitterApi(twitterSettings.AccessToken, twitterSettings.AccessTokenSecret)
    ok, err := api.VerifyCredentials()
    if err != nil {
        t.Errorf("could not verify twitter credentials")
    } else if !ok {
        t.Errorf("twitter credentials invalid")
    }
}

func TestTweetParserThatPulls1WordAfterASearchTerm(t *testing.T) {
    var res []string
    var cases [][2]string
    var answers [][]string
    cases = append(cases, [2]string{"hello", "hello world"})
    answers = append(answers, []string{"world"})
    cases = append(cases, [2]string{"hello", "hello world hello"})
    answers = append(answers, []string{"world"})
    cases = append(cases, [2]string{"hello", "hello world hello world"})
    answers = append(answers, []string{"world","world"})
    cases = append(cases, [2]string{"Hello", "hello hello world"})
    answers = append(answers, []string{"hello","world"})
    
    for counter,pair := range(cases) {
        res = aimless.ExtractWordFromTweet(pair[0], pair[1])
        if !reflect.DeepEqual(res,answers[counter]) {
            var errorWords string 
            for _,wrongWord := range(res) {
                errorWords+=wrongWord + " "
            }
            t.Errorf("word extraction failed; expected " + "world" + " but got " + errorWords)
        }  
    }
    
    /* doesnt actually test the aimless.Search function; can't inject fake tweets into the api to try it 
    api := aimless.LaunchMyAPI()
    searchRes := aimless.Search(api, "term", 10)
    for i,word := range(searchRes) {
        fmt.Println(strconv.Itoa(i), word)
    }
    */
    
}

func TestMain(m *testing.M) {
    res := m.Run()
    os.Exit(res)
}