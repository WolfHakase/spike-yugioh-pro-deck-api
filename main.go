package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const allCardsEndpoint = "https://db.ygoprodeck.com/api/v7/cardinfo.php"
const allSpellEndpoint = allCardsEndpoint + "?type=spell"
const allTrapEndpoint = allCardsEndpoint + "?type=trap"
const allNormalMonstersEndpoint = allCardsEndpoint + "?type=normal%20monster"
const allEffectMonstersEndpoint = allCardsEndpoint + "?type=effect%20monster"
const allSynchroMonstersEndpoint = allCardsEndpoint + "?type=synchro%20monster"
const allPendulumMonstersEndpoint = allCardsEndpoint + "?type=pendulum%20monster"
const allFusionMonstersEndpoint = allCardsEndpoint + "?type=fusion%20monster"
var cardTypes = []string {"Effect%20Monster", "Flip%20Effect%20Monster", "Flip%20Tuner%20Effect%20Monster", "Gemini%20Monster", "Normal%20Monster", "Normal%20Tuner%20Monster", "Pendulum%20Effect%20Monster", "Pendulum%20Flip%20Effect%20Monster", "Pendulum%20Normal%20Monster", "Pendulum%20Tuner%20Effect%20Monster", "Ritual%20Effect%20Monster", "Ritual%20Monster", "Skill%20Card", "Spell%20Card", "Spirit%20Monster", "Toon%20Monster", "Trap%20Card", "Tuner%20Monster", "Union%20Effect%20Monster", "Fusion%20Monster", "Link%20Monster", "Pendulum%20Effect%20Fusion%20Monster", "Synchro%20Monster", "Synchro%20Pendulum%20Effect%20Monster", "Synchro%20Tuner%20Monster", "XYZ%20Monster", "XYZ%20Pendulum%20Effect%20Monster"}

type NormalMonsters struct {
	Data []struct {
		Archetype  string `json:"archetype"`
		Atk        int64  `json:"atk"`
		Attribute  string `json:"attribute"`
		CardImages []struct {
			ID            int64  `json:"id"`
			ImageURL      string `json:"image_url"`
			ImageURLSmall string `json:"image_url_small"`
		} `json:"card_images"`
		CardPrices []struct {
			AmazonPrice       string `json:"amazon_price"`
			CardmarketPrice   string `json:"cardmarket_price"`
			CoolstuffincPrice string `json:"coolstuffinc_price"`
			EbayPrice         string `json:"ebay_price"`
			TcgplayerPrice    string `json:"tcgplayer_price"`
		} `json:"card_prices"`
		CardSets []struct {
			SetCode       string `json:"set_code"`
			SetName       string `json:"set_name"`
			SetPrice      string `json:"set_price"`
			SetRarity     string `json:"set_rarity"`
			SetRarityCode string `json:"set_rarity_code"`
		} `json:"card_sets"`
		Def   int64  `json:"def"`
		Desc  string `json:"desc"`
		ID    int64  `json:"id"`
		Level int64  `json:"level"`
		Name  string `json:"name"`
		Race  string `json:"race"`
		Type  string `json:"type"`
	} `json:"data"`
}

func main() {
	getAllNormalMonsters()
}

func getAllNormalMonsters() {
	resp, err := http.Get(allNormalMonstersEndpoint)
	if err != nil {
		panic(err)
	}

	var cards NormalMonsters
	err = json.NewDecoder(resp.Body).Decode(&cards)
	if err != nil {
		panic(err)
	}

	for _, card := range cards.Data {
		fmt.Printf("Retrieved %v \n", card.Name)
	}
}
