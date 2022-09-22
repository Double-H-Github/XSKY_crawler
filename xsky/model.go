package xsky

type GetTokenResp struct {
	Code int `json:"code"`
	Data struct{
		Token string `json:"token"`
	} `json:"data"`
	Message string `json:"message"`
}

type Category struct {
	Name string `json:"name"`
	EnName string `json:"en_name"`
	I18nName string `json:"i18n_name"`
}

type RecruitType struct {
	Name string `json:"name"`
}

type City struct {
	Name string `json:"name"`
}

type JobInfo struct {
	Title string `json:"title"`
	Desc string `json:"description"`
	Requirement string `json:"requirement"`
	JobCategory *Category `json:"job_category"`
	RecruitType *RecruitType `json:"recruit_type"`
	CityList []*City
}

type Data struct {
	JobPostList []*JobInfo `json:"job_post_list"`
	Count int `json:"count"`
}

type GetJobResp struct {
	Code int `json:"code"`
	Data Data `json:"data"`
	Message string `json:"message"`
}
