package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Player struct {
	Name string
}

type Team struct {
	Name    string
	Stats   TeamStats
	Players []Player
}

type TeamStats struct {
	AverageMarketValue int
	TotalMarketValue   int
}

func main() {
	resp, err := http.Get("https://www.transfermarkt.de/regionalliga-nordost/startseite/wettbewerb/RLN4")
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	teams := []Team{}

	doc.Find("#yw1 > table > tbody > tr > td").Each(func(i int, s *goquery.Selection) {

		if s.HasClass("rechts") {
			if teams[len(teams)-1].Stats.AverageMarketValue != 0 {
				teams[len(teams)-1].Stats.TotalMarketValue = parseValueString(s.Text())
			} else {
				teams[len(teams)-1].Stats.AverageMarketValue = parseValueString(s.Text())
			}
		}

		if s.HasClass("hauptlink") && s.HasClass("no-border-links") {
			teamLink, _ := s.Find("a").Attr("href")
			fmt.Printf("Team link for %s: %s\n", s.Text(), teamLink)
			players := scrapePlayers(teamLink)
			teams = append(teams, Team{Name: strings.TrimSpace(s.Text()), Players: players})
		}
	})

	fmt.Printf("Got %d teams", len(teams))

	j, _ := json.MarshalIndent(teams, "", "\t")
	// fmt.Println(string(j))

	os.WriteFile("data.json", j, 0644)

}

func scrapePlayers(teamLink string) []Player {
	resp, err := http.Get("https://www.transfermarkt.de" + teamLink)
	if err != nil {
		panic(err)
	}

	d, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	players := []Player{}

	d.Find("#yw1 > table > tbody > tr > td").Each(func(i int, s *goquery.Selection) {
		// is there a better way to identify attributes in the table?
		if s.HasClass("hide") {
			println(s.Text())
			players = append(players, Player{Name: strings.TrimSpace(s.Text())})
		}
	})
	return players
}

func parseValueString(valueString string) int {
	// Example values
	// 101 Tsd. €
	// 2,13 Mio. €
	parts := strings.Split(valueString, " ")
	valueStringWithDecimalPoint := strings.Replace(parts[0], ",", ".", -1)
	value, err := strconv.ParseFloat(valueStringWithDecimalPoint, 64)
	if err != nil {
		panic(err)
	}
	switch parts[1] {
	case "Tsd.":
		return int(value * 1000)
	case "Mio.":
		return int(value * 1000000)
	}
	panic("could not parse value " + valueString)
}
