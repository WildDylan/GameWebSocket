package request

import (
	. "game"
)

type RoomRequest struct {

	Play 		Play        `json:"play"`
	Owner		Player		`json:"owner"`

	Id 			string		`json:"id"`
	Lock 		bool		`json:"lock"`
	Password	string		`json:"password"`

}