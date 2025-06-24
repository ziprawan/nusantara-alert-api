package utils

import (
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const MAX_PAGE = 5

func findFirstElement(node *html.Node, dataAtom atom.Atom) *html.Node {
	if node == nil {
		return nil
	}

	for child := range node.ChildNodes() {
		if child == nil {
			continue
		}
		if child.Type == html.ElementNode && child.DataAtom == dataAtom {
			return child
		}
	}

	return nil
}

func findFirstText(node *html.Node) string {
	if node == nil {
		return ""
	}

	for child := range node.ChildNodes() {
		if child == nil {
			continue
		}
		if child.Type == html.TextNode {
			return strings.TrimSpace(child.Data)
		}
	}

	return ""
}

func FetchAdinetPage(page int64) ([]DisasterData, error) {
	res := make([]DisasterData, 0)
	filterDisaster := []string{"flood", "tornadoes"}
	layout := "2006-01-02 15:04:05"

	u, _ := url.Parse("https://adinet.ahacentre.org/report/list?keywords=Indonesia&sort=new") // Harusnya gabisa error kan?
	q := u.Query()
	q.Set("page", strconv.FormatInt(page, 10))
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Println("adinet: Failed to fetch URL:", err)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("adinet: Failed to parse HTML response")
		return nil, err
	}

	rows := doc.Find("table tbody tr")
	if rows == nil {
		fmt.Printf("adinet: Unexpected nil query result while finding \"table tbody tr\"\n")
		return nil, err
	}

	for i, row := range rows.Nodes {
		var disaster DisasterData
		if row == nil {
			fmt.Printf("adinet: Unexpected nil row at index %d\n", i)
			continue
		}

		invalid := false
		tdIdx := 0
		for child := range row.ChildNodes() {
			if child == nil {
				fmt.Printf("adinet: Unexpected nil row child\n")
				continue
			}
			if child.Type != html.ElementNode || child.DataAtom != atom.Td {
				// fmt.Printf("adinet: child is not a td, skipping\n")
				continue
			}

			if tdIdx == 1 {
				aElem := findFirstElement(child, atom.A)
				bElem := findFirstElement(aElem, atom.B)
				text := findFirstText(bElem)

				if text == "" {
					fmt.Printf("adinet: Empty string, skipping\n")
					invalid = true
					break
				}

				disaster.LocationName = strings.Split(text, " in ")[1]
			} else if tdIdx == 2 {
				text := strings.ToLower(findFirstText(child))
				if !slices.Contains(filterDisaster, text) {
					fmt.Printf("adinet: Not the type, skipping\n")
					invalid = true
					break
				}

				disaster.Type = text
			} else if tdIdx == 4 {
				text := findFirstText(child)
				if text == "" {
					fmt.Printf("adinet: Empty string, skipping\n")
					invalid = true
					break
				}

				parsed, err := time.Parse(layout, text)
				if err != nil {
					fmt.Printf("adinet: Invalid time: %s\n", err.Error())
					invalid = true
					break
				}

				disaster.IncidentTime = parsed.UnixMilli()
			}

			tdIdx++
		}

		if invalid {
			fmt.Printf("adinet: Invalid row data at index %d and tdIdx %d\n", i, tdIdx)
			continue
		}

		res = append(res, disaster)
	}

	return res, nil
}

func FetchAdinet() ([]DisasterData, error) {
	res := []DisasterData{}

	for i := range MAX_PAGE {
		dis, err := FetchAdinetPage(int64(i))
		if err != nil {
			return res, err
		}
		res = append(res, dis...)
	}

	return res, nil
}
