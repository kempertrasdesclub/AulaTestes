package dataformat

import (
	"github.com/helmutkemper/util"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
)

type Menu struct {
	Id          string `json:"id" bson:"_id"`
	IdSecondary string `json:"secondaryId" bson:"secondaryId"`
	Text        string `json:"text" bson:"text"`
	Admin       int    `json:"admin" bson:"admin"`
	Icon        string `json:"icon,omitempty" bson:"icon"`
	Url         string `json:"url,omitempty" bson:"url"`
	ItemOrder   int    `json:"itemOrder" bson:"itemOrder"`
	TypeContent int    `json:"typeContent,omitempty" bson:"typeContent"`
	Classroom   int    `json:"classroom,omitempty" bson:"classroom"`
	Menu        []Menu `json:"menu,omitempty" bson:"menu"`
}

func (e *Menu) GetIdAndSecondaryIdAsBSonQuery() (query bson.M) {

	var ok bool
	var secondaryIdMenu reflect.StructField

	secondaryIdMenu, ok = reflect.ValueOf(*e).Type().FieldByName("IdSecondary")
	if ok == false {
		util.TraceToLog()
		return
	}

	tagNameIdSecondary := secondaryIdMenu.Tag.Get("bson")

	query = bson.M{tagNameIdSecondary: e.IdSecondary}
	return
}
