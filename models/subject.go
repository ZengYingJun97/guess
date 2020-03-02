package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

type Subject struct {
	Id        int
	Option    string
	AnswerKey string
	Status    int8
	Img       string
}

func init() {
	orm.RegisterModel(new(Subject))
}

func GetSubject(Id int) (s Subject, err error) {
	o := orm.NewOrm()
	o.Using("guess")
	s = Subject{Id: Id}
	err = o.Read(&s)

	if err != nil {
		return s, err
	}
	return
}

func Answer(Id int, answerKey string) bool {
	s, err := GetSubject(Id)

	if err != nil {
		return false
	}
	return strings.Compare(s.AnswerKey, strings.ToUpper(answerKey)) == 0
}
