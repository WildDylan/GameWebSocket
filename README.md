# GameWebSocket
** CODE WAS IN DEVELOPE **



#### Step 1: CREATE ROOM

[POST] http://localhost:8181/create with json body:

```json
{
	"play": {
		"id": "play-id",
		"name": "play-name",
		"introduce": "play-introduce",
		"main_image": "play-main_image",
		"player_num": 5,
		"steps": [
			{
				"id": "steps-id",
				"name": "steps-name"
			},
			{
				"id": "steps-id2",
				"name": "steps-name2"
			}
			]
	},
	"owner": {
		"id": "user-id",
		"name": "user-name",
		"owner": true
	},
	"id": "123456",
	"lock": true,
	"password": "123456666"
}

```

#### Step2: CONNECT WEBSOCKET

**connect** localhost:8282/echo?token=xxx&roomNumber=123456

