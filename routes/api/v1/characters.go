package v1

import (
	"hazzikostas-api/pkg/db"
	"log"
)

type Character struct {
	ToonName       string  `json:"toon_name"`
	Region         string  `json:"region"`
	Realm          string  `json:"realm"`
	All            float32 `json:"all"`
	Dps            float32 `json:"dps"`
	Healer         float32 `json:"healer"`
	Tank           float32 `json:"tank"`
	Spec0          float32 `json:"spec_0"`
	Spec1          float32 `json:"spec_1"`
	Spec2          float32 `json:"spec_2"`
	Spec3          float32 `json:"spec_3"`
	RankOverall    float32 `json:"rank_overall"`
	RankClass      float32 `json:"rank_class"`
	RankFaction    float32 `json:"rank_faction"`
	AllDif         float32 `json:"all_dif"`
	DpsDif         float32 `json:"dps_dif"`
	HealerDif      float32 `json:"healer_dif"`
	TankDif        float32 `json:"tank_dif"`
	Spec0Dif       float32 `json:"spec_0_dif"`
	Spec1Dif       float32 `json:"spec_1_dif"`
	Spec2Dif       float32 `json:"spec_2_dif"`
	Spec3Dif       float32 `json:"spec_3_dif"`
	RankOverallDif float32 `json:"rank_overall_dif"`
	RankClassDif   float32 `json:"rank_class_dif"`
	RankFactionDif float32 `json:"rank_faction_dif"`
}

// Public function that returns all the characters stored in the database
func GetCharacters() (*[]Character, error) {
	cursor, err := db.SetConnection()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer func() {
		err = cursor.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	results, err := cursor.Query("SELECT db.Toon_Name, db.region, db.realm, db._all, db.dps," +
		"db.healer, db.tank, db.spec_0, db.spec_1, db.spec_2, db.spec_3, db.rank_overall," +
		"db.rank_class, db.rank_faction, db._all_diff, db.dps_diff, db.healer_diff," +
		"db.tank_diff, db.spec_0_diff, db.spec_1_diff, db.spec_2_diff, db.spec_3_diff," +
		"db.rank_class_diff, db.rank_faction_diff FROM HK_Toons_1 db")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var characters []Character
	for results.Next() {
		var tmp Character
		err := results.Scan(&tmp.ToonName, &tmp.Region, &tmp.Realm, &tmp.All, &tmp.Dps, &tmp.Healer,
			&tmp.Tank, &tmp.Spec0, &tmp.Spec1, &tmp.Spec2, &tmp.Spec3, &tmp.RankOverall, &tmp.RankClass,
			&tmp.RankFaction, &tmp.AllDif, &tmp.DpsDif, &tmp.HealerDif, &tmp.TankDif, &tmp.Spec0Dif, &tmp.Spec1Dif,
			&tmp.Spec2Dif, &tmp.Spec3Dif, &tmp.RankClassDif, &tmp.RankFactionDif)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		characters = append(characters, tmp)
	}
	return &characters, nil
}

// Public function that returns the characters that the bot has to post on discord.
func GetCharactersToPost() (*[]Character, error) {
	cursor, err := db.SetConnection()
	if err != nil {
		return nil, err
	}
	defer func() {
		err := cursor.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	results, err := cursor.Query("SELECT db.Toon_Name, db.region, db.realm, db._all, db.dps," +
		"db.healer, db.tank, db.spec_0, db.spec_1, db.spec_2, db.spec_3, db.rank_overall," +
		"db.rank_class, db.rank_faction, db._all_diff, db.dps_diff, db.healer_diff," +
		"db.tank_diff, db.spec_0_diff, db.spec_1_diff, db.spec_2_diff, db.spec_3_diff," +
		"db.rank_class_diff, db.rank_faction_diff FROM HK_Toons_1 db WHERE Post = 1")
	if err != nil {
		return nil, err
	}
	var characters []Character
	for results.Next() {
		var tmp Character
		err := results.Scan(&tmp.ToonName, &tmp.Region, &tmp.Realm, &tmp.All, &tmp.Dps, &tmp.Healer,
			&tmp.Tank, &tmp.Spec0, &tmp.Spec1, &tmp.Spec2, &tmp.Spec3, &tmp.RankOverall, &tmp.RankClass,
			&tmp.RankFaction, &tmp.AllDif, &tmp.DpsDif, &tmp.HealerDif, &tmp.TankDif, &tmp.Spec0Dif, &tmp.Spec1Dif,
			&tmp.Spec2Dif, &tmp.Spec3Dif, &tmp.RankClassDif, &tmp.RankFactionDif)
		if err != nil {
			return nil, err
		}
		characters = append(characters, tmp)
	}
	return &characters, nil
}

func UpdatePostCharacterStatus(characterName string) error {

	return nil
}
