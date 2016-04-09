package main 
// native packages
import (
    "fmt"
    "net/url"
    "log"
    "strings"
    "strconv"
)

// external packages
import (
   "github.com/kwkroll32/aimless-apparition/aimless"
)

func prettyPrintMap(m map[string]int) {
    for key, value := range(m) {
        fmt.Println(strconv.Itoa(value) + "\t" + key)
    }
}

func main() {
    api := aimless.LaunchMyAPI()
    v := url.Values{}
    v.Set("count","300")
    searchRes, err := api.GetSearch("much", v)
    if err != nil{
        log.Fatal(err)
    }
    wordResCounts := make(map[string]int)
    for _,tweet := range(searchRes.Statuses) {
        //fmt.Println(tweet.Text )
        for _,word := range(aimless.ExtractWordFromTweet("much", tweet.Text)) {
            word = strings.ToLower(word)
            if wordResCounts[word] == 0 {
                wordResCounts[word] = 1
            } else {
                wordResCounts[word]++
            }
        }
    }
    
    prettyPrintMap(wordResCounts)
}