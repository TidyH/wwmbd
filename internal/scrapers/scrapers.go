package scrapers

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// This is the current Path to the Statistics Table in Yahoo Finance
const XPATH = "/html/body/div[2]/main/section/section/section/article/section[2]/div[2]/table[1]"
const COMPANYURL = "https://finance.yahoo.com/quote/<COMPANY TICKER>/key-statistics/"

// Comment representation of expected table
// currently all strings to make initial work easier
// TODO: correctly convert all strings to appropriate values
// +------------------------+---------+---------+---------+---------+----------+---------+
// |Index                   | 0       | 1       | 2       | 3       | 4        | 5       |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | Titles                 | Current | 9/30/24 | 6/30/24 | 3/31/24 | 12/31/23 | 9/30/23 |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | Market Cap             | string  |  string |  string | string  | string   | string  |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | Enterprise Value       | string  |  string |  string | string  | string   | string  |
// +-------------------- ---+---------+---------+---------+---------+----------+---------+
// | TrailingPE             | string  |  string |  string | string  | string   | string  |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | ForwardPE              | string  |  string |  string | string  | string   | string  |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | PEGRatio               | string  |  string |  string | string  | string   | string  |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | PriceSales             | string  |  string |  string | string  | string   | string  |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | PriceBook              | string  |  string |  string | string  | string   | string  |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | EnterpriseValueRevenue | string  |  string |  string | string  | string   | string  |
// +------------------------+---------+---------+---------+---------+----------+---------+
// | EnterpriseValueEBITDA  | string  |  string |  string | string  | string   | string  |
// +------------------------+---------+---------+---------+---------+----------+---------+

type StatisticRow struct {
	MarketCap              string
	EnterpriseValue        string
	TrailingPE             string
	ForwardPE              string
	PEGRatio               string
	PriceSales             string
	PriceBook              string
	EnterpriseValueRevenue string
	EnterpriseValueEBITDA  string
}

type StatisticsTable struct {
	Symbol       string
	ColumnTitles []string
	Rows         []StatisticRow
}

// Init function
func NewStatistics(str string) StatisticsTable {
	S := &StatisticsTable{Symbol: str}

	// Steps to Populate table
	// Scrape Website
	url := S.setURL()
	htmlStr, _ := S.getWebPage(url)

	// Find data in web page
	tableNode, _ := S.getStatisticsTable(htmlStr)

	// Set headers in table
	S.ColumnTitles = columnList(*tableNode)

	// Set row values in table
	rows, _ := rowGrabber(*tableNode)

	// insert values
	S.Rows = rows

	return *S
}

// Internal functions required for initialization
// Set URL based on Symbol input
func (S *StatisticsTable) setURL() string {
	// Format url to scrape from
	url := strings.Replace(COMPANYURL, "<COMPANY TICKER>", S.Symbol, -1)

	return url
}

// Get Webpage HTML
func (S *StatisticsTable) getWebPage(path string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
		},
	}

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	responseBody, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	// Convert byte slice to string
	return string(responseBody), nil
}

// finding table in webpage html
func (S *StatisticsTable) getStatisticsTable(html string) (*html.Node, error) {
	doc, err := htmlquery.Parse(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	node := htmlquery.FindOne(doc, XPATH)

	return node, nil
}

func rowGrabber(node html.Node) ([]StatisticRow, error) {
	var rows []StatisticRow
	var outerIndex int = 1 // starting at 1 to avoid labels

	// Find all table rows
	rowsNode := htmlquery.Find(&node, "//tr")

	// need to move across each row on decoupled index.
	// This index is actually the length of columnTitles == 6 for now
	for outerIndex <= 6 {
		var currentRow []string // zero out each run
		for _, row := range rowsNode {
			cells := htmlquery.Find(row, "//td")
			for i, n := range cells {
				if i == outerIndex {
					currentRow = append(currentRow, htmlquery.InnerText(n))
				}
			}
		}

		rows = append(rows, StatisticRow{
			MarketCap:              currentRow[0],
			EnterpriseValue:        currentRow[1],
			TrailingPE:             currentRow[2],
			ForwardPE:              currentRow[3],
			PEGRatio:               currentRow[4],
			PriceSales:             currentRow[5],
			PriceBook:              currentRow[6],
			EnterpriseValueRevenue: currentRow[7],
			EnterpriseValueEBITDA:  currentRow[8],
		})
		outerIndex++
	}

	return rows, nil
}

func columnList(node html.Node) []string {
	var columnTitles []string
	columnList := htmlquery.Find(&node, "//th")

	for _, n := range columnList {
		if htmlquery.InnerText(n) == "" {
			continue
		}
		columnTitles = append(columnTitles, htmlquery.InnerText(n))
	}

	return columnTitles
}
