package characters

type Character struct {
	ToonName       string  `json:"toon_name"`
	SpecName       string  `json:"spec_name"`
	Class          string  `json:"class_"`
	Race           string  `json:"race"`
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
