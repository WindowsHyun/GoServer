package structure

type ReqRegist struct {
	Email      string `json:"email" bson:"email"`
	NickName   string `json:"nickName" bson:"nickName"`
	BirthYear  string `json:"birthYear" bson:"birthYear"`
	BirthMonth string `json:"birthMonth" bson:"birthMonth"`
}

type ResRegist struct {
	Message string `json:"msg" bson:"msg"`
}

type ReqLogin struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type ResLogin struct {
	Message string `json:"msg" bson:"msg"`
	Token   string `json:"token" bson:"token"`
}

type ResDefaultMessage struct {
	Message string `json:"msg" bson:"msg"`
}
