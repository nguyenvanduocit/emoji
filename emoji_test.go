package emoji

import (
	"testing"
)


func TestFindEmoji(t *testing.T) {
	_, err := FindEmoji("I want to drink")
	if err != nil {
		t.Error("Result must not be error")
	}
}
