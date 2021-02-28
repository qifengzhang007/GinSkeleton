package data_type

type StuInfo struct {
	Code int    `form:"code" json:"code"`
	Msg  string `form:"msg" json:"msg"`
	Data struct {
		UserName  string `form:"user_name" json:"user_name"`
		Age       int    `form:"age" json:"age"`
		ScoreList []struct {
			Score   float64 `form:"score"  json:"score"`
			Subject string  `form:"subject" json:"subject"`
		} `form:"score_list" json:"score_list"`
	} `form:"data" json:"data"`
}
