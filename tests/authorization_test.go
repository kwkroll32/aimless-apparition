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
   // "github.com/ChimeraCoder/anaconda"
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
    
}
 
func TestMain(m *testing.M) {
    os.Exit(m.Run())
}