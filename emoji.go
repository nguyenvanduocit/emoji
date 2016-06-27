package emoji

import (
	"net/http"
	"encoding/json"
	"net/url"
)

type Emoji struct{
	Score float32 `json:"score"`
	Text string `json:"text"`
}

type Result struct {
	Results []Emoji `json:"results"`
}

//Find emoji
func FindEmoji(text string)([]Emoji, error){
	encodedText, err := UrlEncoded(text)
	if err != nil {
		return nil, err
	}
	resp, err := http.Get("http://emoji.getdango.com/api/emoji?q=" + encodedText)
	if err != nil {
		return nil, err
	}
	var emojiResult Result
	err = json.NewDecoder(resp.Body).Decode(&emojiResult)
	if err != nil {
		return nil, err
	}
	return emojiResult.Results, nil
}

// Encode string
func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
