package main

import (
	"time"
)

type repository struct {
	Repository_name    string      `json:"-"`
	Ch_Repository_name string      `bson:"ch_repname" json:"ch_repname"`
	Create_user        string      `json:"create_user,omitempty"`
	Repaccesstype      string      `json:"repaccesstype,omitempty"`
	Comment            string      `json:"comment"`
	Optime             string      `json:"optime,omitempty"`
	Items              int         `json:"items"`
	CooperateItems     int         `json:"cooperateitems"`
	Label              interface{} `json:"label"`
	Ct                 time.Time   `json:"-"`
	St                 time.Time   `bson:"st,omitempty", json:"-"`
	Cooperate          interface{} `bson:"cooperators", json:"-"`
	Rank               float32     `json:"-"`
}
