package models

import //I'm not exactly sure what will be needed here. the tutorial I am following is using mongo and I would like to use postgres, I believe it needs a driver

type MakeupList struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`

	Task string `json:"task,omitempty"`
	Status bool `json:"status,omitempty`
}