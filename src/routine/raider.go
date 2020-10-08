package routine

import (
	"encoding/json"
	"fmt"
	v1 "hazzikostas-api/routes/api/v1"
	"hazzikostas-api/src/db"
	"log"
	"net/http"
	"net/url"
)

type RaiderData struct {
	Name                     string `json:"name"`
	Race                     string `json:"race"`
	Class                    string `json:"class"`
	ActiveSpecName           string `json:"active_spec_name"`
	ActiveSpecRole           string `json:"active_spec_role"`
	Gender                   string `json:"gender"`
	Faction                  string `json:"faction"`
	AchievementPoints        int    `json:"achievement_points"`
	HonorableKills           int    `json:"honorable_kills"`
	Region                   string `json:"region"`
	Realm                    string `json:"realm"`
	MythicPlusScoresBySeason []struct {
		Season string `json:"season"`
		Scores struct {
			All    float32 `json:"all"`
			Dps    float32 `json:"dps"`
			Healer float32 `json:"healer"`
			Tank   float32 `json:"tank"`
			Spec0  float32 `json:"spec_0"`
			Spec1  float32 `json:"spec_1"`
			Spec2  float32 `json:"spec_2"`
			Spec3  float32 `json:"spec_3"`
		} `json:"scores"`
	} `json:"mythic_plus_scores_by_season"`
	MythicPlusRanks struct {
		Overall struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"overall"`
		Class struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"class"`
		FactionOverall struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"faction_overall"`
		FactionClass struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"faction_class"`
		Dps struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"dps"`
		ClassDps struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"class_dps"`
		FactionDps struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"faction_dps"`
		FactionClassDps struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"faction_class_dps"`
		Spec62 struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"spec_62"`
		FactionSpec62 struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"faction_spec_62"`
		Spec63 struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"spec_63"`
		FactionSpec63 struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"faction_spec_63"`
		Spec64 struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"spec_64"`
		FactionSpec64 struct {
			World  float32 `json:"world"`
			Region float32 `json:"region"`
			Realm  float32 `json:"realm"`
		} `json:"faction_spec_64"`
	} `json:"mythic_plus_ranks"`
}

type characterInfo struct {
	region string
	realm  string
	name   string
}

func buildURL(character characterInfo) string {
	return fmt.Sprintf("https://raider.io/api/v1/characters/profile?region=%s&realm=%s&name=%s&fields=mythic_plus_ranks,mythic_plus_scores_by_season:current",
		character.region, character.realm, url.QueryEscape(character.name))
}

func GetRaiderData(region string, realm string, name string) (*RaiderData, error) {
	character := characterInfo{region, realm, name}
	fmt.Println(url.Parse(buildURL(character)))
	var resp, err = http.Get(buildURL(character))
	if err != nil {
		log.Fatal(err)
	}
	data := new(RaiderData)
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

func Routine(characters []v1.Character) {
	for i := 0; i < len(characters); i++ {
		raiderInfo, err := GetRaiderData(characters[i].Region, characters[i].Realm, characters[i].ToonName)
		log.Println("Analyzing: " + characters[i].ToonName)
		if err != nil {
			log.Println(err)
		}
		if raiderInfo.MythicPlusScoresBySeason[0].Scores.All != characters[i].All ||
			raiderInfo.MythicPlusScoresBySeason[0].Scores.Dps != characters[i].Dps ||
			raiderInfo.MythicPlusScoresBySeason[0].Scores.Tank != characters[i].Tank ||
			raiderInfo.MythicPlusScoresBySeason[0].Scores.Healer != characters[i].Healer ||
			raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec0 != characters[i].Spec0 ||
			raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec1 != characters[i].Spec1 ||
			raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec2 != characters[i].Spec2 ||
			raiderInfo.MythicPlusScoresBySeason[0].Scores.Spec3 != characters[i].Spec3{
			log.Println("Updating: " + characters[i].ToonName)
			err = updateCharacter(characters[i], raiderInfo)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func updateCharacterValues(character *v1.Character, raiderInfo *RaiderData) {
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
	character.RankFaction = raiderInfo.MythicPlusRanks.FactionOverall.Realm - character.RankFaction
	character.RankClass = raiderInfo.MythicPlusRanks.Class.Realm - character.RankClass
	character.RankOverall = raiderInfo.MythicPlusRanks.Overall.Realm - character.RankOverall
}

// updateCharacter
func updateCharacter(character v1.Character, raiderInfo *RaiderData) error {
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
		"SET Post=1, `_all`=?, dps=?, healer=?, "+
		"tank=?, spec_0=?, spec_1=?, spec_2=?, "+
		"spec_3=?, rank_overall=?, rank_class=?, "+
		"rank_faction=?, `_all_diff`=?, dps_diff=?, "+
		"healer_diff=?, tank_diff=?, spec_0_diff=?, "+
		"spec_1_diff=?, spec_2_diff=?, spec_3_diff=?, "+
		"rank_overall_diff=?, rank_class_diff=?, rank_faction_diff=? "+
		"WHERE Toon_Name=? ;", character.All, character.Dps, character.Healer, character.Tank, character.Spec0,
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
