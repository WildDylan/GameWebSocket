package game

type Play struct {

	Id 			    string      `json:"id"`
	Name 		    string      `json:"name"`
	Introduce 	    string      `json:"introduce"`
	MainImage 	    string      `json:"main_image"`

	Steps    		[]Step      `json:"steps"`
	PlayerNum       int         `json:"player_num"`

}