package main

import (
	"encoding/json"
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
	Players []Player
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
		if s.HasClass("hauptlink") && s.HasClass("hide-for-small") {
			teamLink, _ := s.Find("a").Attr("href")
			fmt.Printf("Team link for %s: %s\n", s.Text(), teamLink)
			players := []Player{}
			// todo: follow team link, scrape players and fill for each team
			teams = append(teams, Team{Name: strings.TrimSpace(s.Text()), Players: players})
		}
	})

	println(len(teams))

	j, _ := json.MarshalIndent(teams, "", "\t")
	fmt.Println(string(j))

}
