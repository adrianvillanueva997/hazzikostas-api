package raider

type Data struct {
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
