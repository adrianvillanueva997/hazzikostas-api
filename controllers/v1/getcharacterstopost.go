package v1

import (
	"hazzikostas-api/core/characters"
	"hazzikostas-api/pkg/db"
	"log"
)

// Public function that returns the characters that the bot has to post on discord.
func GetCharactersToPost() (*[]characters.Character, error) {
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
	var characterList []characters.Character
	for results.Next() {
		var tmp characters.Character
		err := results.Scan(&tmp.ToonName, &tmp.Region, &tmp.Realm, &tmp.All, &tmp.Dps, &tmp.Healer,
			&tmp.Tank, &tmp.Spec0, &tmp.Spec1, &tmp.Spec2, &tmp.Spec3, &tmp.RankOverall, &tmp.RankClass,
			&tmp.RankFaction, &tmp.AllDif, &tmp.DpsDif, &tmp.HealerDif, &tmp.TankDif, &tmp.Spec0Dif, &tmp.Spec1Dif,
			&tmp.Spec2Dif, &tmp.Spec3Dif, &tmp.RankClassDif, &tmp.RankFactionDif)
		if err != nil {
			return nil, err
		}
		characterList = append(characterList, tmp)
	}
	return &characterList, nil
}
