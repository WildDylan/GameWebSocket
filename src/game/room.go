package game

type Room struct {

	Play 		*Play
	Players 	[]*Player
	Owner		*Player

	Id 			string
	Lock 		bool
	Password	string

}

var Rooms = make(map[string]*Room)

func (room *Room) BroadCast(bytes []byte) {
	player := room.Owner
	player.Connection.To(room.Id).EmitMessage(bytes)
}

func (room *Room) BroadCastExclude(bytes []byte, uid string) {
	for index := range room.Players {
		currentPlayer := room.Players[index]
		if currentPlayer.Id == uid {
			continue
		}
		currentPlayer.Connection.EmitMessage(bytes)
	}
}

func (room *Room) Execute() {

}

func (room *Room) OnMessage(bytes []byte)  {
	room.BroadCast(bytes)
}