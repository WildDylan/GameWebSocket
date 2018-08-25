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

// Options
func (room *Room) AddPlayer(player *Player)  {
	room.Players = append(room.Players, player)
	room.Owner = room.Players[0]
}

// Scheduler
func (room *Room) Execute() {
	// Every seconds, call this room
	// TODO: But, it's not best, every step should have this own count down timer.
}

// Listeners
func (room *Room) OnMessage(bytes []byte)  {
	room.BroadCast(bytes)
}

// Message functions
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