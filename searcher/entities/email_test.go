package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailSearch(t *testing.T) {
	tests := []struct {
		Title           string
		Data            string
		expectedResults []string
	}{
		{
			Title:           "Simple text with only email address",
			Data:            `tr@tr.com`,
			expectedResults: []string{"tr@tr.com"},
		},
		{
			Title:           "Text with some emails",
			Data:            `Please don't be so ugly abcd@gmailyahoo.ru, it's not your guilty. tr@ggg.ry1`,
			expectedResults: []string{"abcd@gmailyahoo.ru", "tr@ggg.ry"},
		},
		{
			Title: "Html with some emails",
			Data: `<body><p>Proxy Port Last Check Proxy Speed Proxy Country Anonymity truk@gmail.com
					<div>admin@google.com</div> 8080 34 sec Indonesia - Tangerang Transparent 2.33 Transparent<p></body>`,
			expectedResults: []string{"truk@gmail.com", "admin@google.com"},
		},
		{
			Title: "Html with cyrillic emails",
			Data: `<body><p>Proxy Port Last Check Proxy Speed Proxy Country Anonymity
					<div>адми@очта.рф</div><p></body>`,
			expectedResults: []string{"адми@очта.рф"},
		},
	}

	for _, test := range tests {

		search, err := NewSearchEmail()
		assert.NoError(t, err)

		result, err := search.Search(test.Data)
		assert.NoError(t, err)

		t.Log(test.Title)
		assert.Equal(t, test.expectedResults, result)
	}
}
