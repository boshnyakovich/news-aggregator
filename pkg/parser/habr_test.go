package parser

import (
	"github.com/boshnyakovich/news-aggregator/internal/models"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/export"
	"github.com/prometheus/common/log"
	"github.com/stretchr/testify/assert"
	"html/template"
	"io/ioutil"
	"net/http"
	"testing"
)

type TestExporter struct {
	news []models.HabrNews
}

func (t *TestExporter) Export(c chan interface{}) {
	for data := range c {
		t.news = append(t.news, data.(models.HabrNews))
	}
}

func readHabrFile() []byte {
	data, _ := ioutil.ReadFile("test_files/habr.html")
	return data
}

func TestHabrParser_MakeUrl(t *testing.T) {
	parser := &HabrParser{}
	url, err := parser.MakeUrl(models.HabrCriteria{})

	assert.NoError(t, err)
	assert.Equal(t, "https://habr.com/ru/news/", url)

	url, err = parser.MakeUrl(models.HabrCriteria{
		Articles: true,
		All:      true,
		Rating:   uint64(models.TEN),
	})

	assert.NoError(t, err)
	assert.Equal(t, "https://habr.com/ru/all/top10/", url)
}

func TestHabrParser_Parse(t *testing.T) {

	testExporter := &TestExporter{}

	go func() {
		http.HandleFunc("/test", habrPageTemplate)
		err := http.ListenAndServe("localhost:9098", nil)
		assert.NoError(t, err)
	}()

	parser := &HabrParser{}
	gz := geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"http://localhost:9098/test"},
		ParseFunc: parser.parse,
		Exporters: []export.Exporter{testExporter},
	})

	gz.Start()

	assert.NotEmpty(t, testExporter.news)
}

func habrPageTemplate(w http.ResponseWriter, r *http.Request) {
	var homeTemplate = template.Must(template.New("").Parse(string(readHabrFile())))
	if err := homeTemplate.Execute(w, ""); err != nil {
		log.Warn("cannot close rows")
	}
}