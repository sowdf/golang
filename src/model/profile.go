package model

import "encoding/json"

type Profile struct {
	Name       string
	Gender     string //性别
	Age        int
	Height     int
	Weight     int
	Income     string // 收入
	Marriage   string //婚姻状况
	Education  string //教育情况
	Occupation string //职业
	Hokou      string //籍贯
	Xingzuo    string //星座
	Car        string //车
	House      string //房子
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, e := json.Marshal(o)
	if e != nil {
		return profile, e
	}

	e = json.Unmarshal(s, &profile)

	return profile, e
}
