package model

import "time"

type marketEvent struct{
	Type	string 'json:"type"'	//event type
	Data	Stock	'json:"data"'	//payload of data
	Time	time.Time 'json:"time"'	//time of event
}
