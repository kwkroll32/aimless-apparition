package main 

// native packages 
import (
    //"fmt"
    "os"
    "testing"
    "encoding/json"
)

// external packages
import (
   "github.com/ChimeraCoder/anaconda"
   "github.com/kwkroll32/aimless-apparition/Settings"
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
func TestMain(m *testing.M) {
    os.Exit(m.Run())
}