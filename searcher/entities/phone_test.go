package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhoneSearch(t *testing.T) {
	tests := []struct {
		Title           string
		Data            string
		expectedResults []string
	}{
		{
			Title:           "Simple text with only phoe",
			Data:            `+79131539999`,
			expectedResults: []string{"+79131539999"},
		},
		{
			Title:           "Text with some phones",
			Data:            `Please don't be so ugly +79131539999, it's not your guilty. 8 908 800 03 03`,
			expectedResults: []string{"+79131539999", "8 908 800 03 03"},
		},
		{
			Title: "Html with some phones",
			Data: `<body><p>Proxy Port Last Check Proxy Speed Proxy Country Anonymity +79131539999
					<div>8 908 800 03 03</div> 8080 34 sec Indonesia <input value="8 908-800-44-44"> - Tangerang Transparent 2.33 Transparent<p></body>`,
			expectedResults: []string{"+79131539999", "8 908 800 03 03", "8 908-800-44-44"},
		},
	}

	for _, test := range tests {

		search, err := NewSearchPhone()
		assert.NoError(t, err)

		result, err := search.Search(test.Data)
		assert.NoError(t, err)

		t.Log(test.Title)
		assert.Equal(t, test.expectedResults, result)
	}
}
