package routine

import (
	"encoding/json"
	"fmt"
	"hazzikostas-api/core/characters"
	"hazzikostas-api/core/raider"
	"hazzikostas-api/pkg/db"
	"log"
	"net/http"
	"net/url"
)

type characterInfo struct {
	region string
	realm  string
	name   string
}

func buildURL(character characterInfo) string {
	return fmt.Sprintf("https://raider.io/api/v1/characters/profile?region=%s&realm=%s&name=%s&fields=mythic_plus_ranks,mythic_plus_scores_by_season:current",
		character.region, character.realm, url.QueryEscape(character.name))
}

func GetRaiderData(region string, realm string, name string) (*raider.Data, error) {
	character := characterInfo{region, realm, name}
	var resp, err = http.Get(buildURL(character))
	if err != nil {
		log.Fatal(err)
	}
	data := new(raider.Data)
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Routine(characters []characters.Character) {
	for i := 0; i < len(characters); i++ {
		raiderInfo, err := GetRaiderData(characters[i].Region, characters[i].Realm, characters[i].ToonName)
		log.Println("Analyzing: " + characters[i].ToonName)
		if err != nil {
			log.Println(err)
		}
		if len(raiderInfo.MythicPlusScoresBySeason) != 0 {
			if raiderInfo.MythicPlusScoresBySeason[0].Scores.All != characters[i].All ||
				raiderInfo.MythicPlusScoresBySeason[0].Scores.Dps != characters[i].Dps ||
				raiderInfo.MythicPlusScoresBySeason[0].Scores.Tank != characters[i].Tank ||
				raiderInfo.MythicPlusScoresBySeason[0].Scores.Healer != characters[i].Healer ||
				raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec0 != characters[i].Spec0 ||
				raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec1 != characters[i].Spec1 ||
				raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec2 != characters[i].Spec2 ||
				raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec3 != characters[i].Spec3 {
				log.Println("Updating: " + characters[i].ToonName)
				err = updateCharacter(characters[i], raiderInfo)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}

func updateCharacterValues(character *characters.Character, raiderInfo *raider.Data) {
	// differences with old data
	character.TankDif = raiderInfo.MythicPlusScoresBySeason[0].Scores.Tank - character.Tank
	character.HealerDif = raiderInfo.MythicPlusScoresBySeason[0].Scores.Healer - character.Healer
	character.DpsDif = raiderInfo.MythicPlusScoresBySeason[0].Scores.Dps - character.Dps
	character.AllDif = raiderInfo.MythicPlusScoresBySeason[0].Scores.All - character.All
	character.Spec0Dif = raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec0 - character.Spec0
	character.Spec1Dif = raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec1 - character.Spec1
	character.Spec2Dif = raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec2 - character.Spec2
	character.Spec3Dif = raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec3 - character.Spec3
	character.RankFactionDif = raiderInfo.MythicPlusRanks.FactionOverall.Realm - character.RankFaction
	character.RankClassDif = raiderInfo.MythicPlusRanks.Class.Realm - character.RankClass
	character.RankOverallDif = raiderInfo.MythicPlusRanks.Overall.Realm - character.RankOverall
	// assign new data
	character.All = raiderInfo.MythicPlusScoresBySeason[0].Scores.All
	character.Dps = raiderInfo.MythicPlusScoresBySeason[0].Scores.Dps
	character.Tank = raiderInfo.MythicPlusScoresBySeason[0].Scores.Tank
	character.Healer = raiderInfo.MythicPlusScoresBySeason[0].Scores.Healer
	character.Spec0 = raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec0
	character.Spec1 = raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec1
	character.Spec2 = raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec2
	character.Spec3 = raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec3
	character.RankFaction = raiderInfo.MythicPlusRanks.FactionOverall.Realm
	character.RankClass = raiderInfo.MythicPlusRanks.Class.Realm
	character.RankOverall = raiderInfo.MythicPlusRanks.Overall.Realm
	character.SpecName = raiderInfo.ActiveSpecName
	character.Race = raiderInfo.Race
	character.Class = raiderInfo.Class
}

// updateCharacter
func updateCharacter(character characters.Character, raiderInfo *raider.Data) error {
	cursor, err := db.SetConnection()
	if err != nil {
		log.Println(err)
	}
	defer func() {
		err = cursor.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	updateCharacterValues(&character, raiderInfo)
	rows, err := cursor.Query("UPDATE Discord_Bots.HK_Toons_1 "+
		"SET Post=1, activeSpec=?, class=?, race=?, `_all`=?, dps=?, healer=?, "+
		"tank=?, spec_0=?, spec_1=?, spec_2=?, "+
		"spec_3=?, rank_overall=?, rank_class=?, "+
		"rank_faction=?, `_all_diff`=?, dps_diff=?, "+
		"healer_diff=?, tank_diff=?, spec_0_diff=?, "+
		"spec_1_diff=?, spec_2_diff=?, spec_3_diff=?, "+
		"rank_overall_diff=?, rank_class_diff=?, rank_faction_diff=? "+
		"WHERE Toon_Name=? ;", character.SpecName, character.Class, character.Race, character.All, character.Dps,
		character.Healer, character.Tank, character.Spec0,
		character.Spec1, character.Spec2, character.Spec3, character.RankOverall, character.RankClass,
		character.RankFaction, character.AllDif, character.DpsDif, character.HealerDif, character.TankDif,
		character.Spec0Dif, character.Spec1Dif, character.Spec2Dif, character.Spec3Dif, character.RankOverallDif,
		character.RankClassDif, character.RankFactionDif, character.ToonName)
	if rows != nil && err == nil {
		log.Println(err)
		return err
	}
	return nil
}
