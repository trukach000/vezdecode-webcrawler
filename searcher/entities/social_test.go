package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSocialSearch(t *testing.T) {
	tests := []struct {
		Title           string
		Data            string
		expectedResults []string
	}{
		{
			Title:           "Simple text with only vk link",
			Data:            `https://vk.com/try_catch_finally`,
			expectedResults: []string{"https://vk.com/try_catch_finally"},
		},
		{
			Title: "Text with some vk",
			Data: `Hello there! vk.com/ddd 
								General Kenobi! vk.com/id1`,
			expectedResults: []string{"vk.com/ddd", "vk.com/id1"},
		},
		{
			Title: "Html with some facebook addresses",
			Data: `<body><p>Proxy Port Last Check Proxy Speed Proxy Country Anonymity http://www.facebook.com/someusername
					<div>http://www.facebook.com/profile.php?id=123456789</div> 8080 34 sec Indonesia - Tangerang Transparent 2.33 Transparent<p></body>`,
			expectedResults: []string{"http://www.facebook.com/someusername", "http://www.facebook.com/profile.php?id=123456789"},
		},
	}

	for _, test := range tests {

		ipSearch, err := NewSearchSocial()
		assert.NoError(t, err)

		result, err := ipSearch.Search(test.Data)
		assert.NoError(t, err)

		t.Log(test.Title)
		assert.Equal(t, test.expectedResults, result)
	}
}
