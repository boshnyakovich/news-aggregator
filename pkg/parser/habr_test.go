package parser

import (
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

type TestExporter struct {
	news []models.HabrNews
}

func (t *TestExporter) Export (c chan interface{}) {
	for data := range c {
		t.news = append(t.news, data.(models.HabrNews))
	}
}

func reaHabrFile() []byte {
	data, _ := ioutil.ReadFile("test_files/habr.html")
	return data
}

func TestHabrParser_MakeUrl(t *testing.T) {
	parser := &HabrParser{}
	url, err := parser.MakeUrl(models.HabrCriteria{})

	assert.NoError(t, err)
	assert.Equal(t,"https://habr.com/ru/news/", url)

	url, err = parser.MakeUrl(models.HabrCriteria{
		Articles: true,
		All: true,
		Rating: uint64(models.TEN),
	})

	assert.NoError(t, err)
	assert.Equal(t,"https://habr.com/ru/all/top10/", url)
}
