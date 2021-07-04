package models

import "time"

type MarketwideOI struct {
	position_date time.Time

	client ParticipantOI
	dii    ParticipantOI
	fii    ParticipantOI
	pro    ParticipantOI
}

type ParticipantOI struct {
	future_index_long  int
	future_index_short int

	future_stock_long  int
	future_stock_short int

	option_index_call_long int
	option_index_put_long  int

	option_index_call_short int
	option_index_put_short  int

	option_stock_call_long int
	option_stock_put_long  int

	option_stock_call_short int
	option_stock_put_short  int
}
