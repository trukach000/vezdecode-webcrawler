package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPSearch(t *testing.T) {
	tests := []struct {
		Title           string
		Data            string
		expectedResults []string
	}{
		{
			Title:           "Simple text with only IP address",
			Data:            `255.0.255.1`,
			expectedResults: []string{"255.0.255.1"},
		},
		{
			Title: "Text with some IP addresses",
			Data: `Hello there! 255.255.255.255 
								General Kenobi! 1.1.1.1`,
			expectedResults: []string{"255.255.255.255", "1.1.1.1"},
		},
		{
			Title: "Html with some IP addresses",
			Data: `<body><p>Proxy Port Last Check Proxy Speed Proxy Country Anonymity 118.99.81.204
					<div>118.99.81.204</div> 8080 34 sec Indonesia - Tangerang Transparent 2.33 Transparent<p></body>`,
			expectedResults: []string{"118.99.81.204", "118.99.81.204"},
		},
	}

	for _, test := range tests {

		ipSearch, err := NewSearchIP()
		assert.NoError(t, err)

		result, err := ipSearch.Search(test.Data)
		assert.NoError(t, err)

		t.Log(test.Title)
		assert.Equal(t, test.expectedResults, result)
	}
}
