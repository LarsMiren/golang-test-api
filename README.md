REST API

1. Create a room: POST to "/room" with json {"number":yourInt, guests:["your", "guests", "names"]}.
2. Get one room: GET to "/room/yourRoomNumber", where yourRoomNumber is int.
3. Get all rooms: GET to "/rooms"
4. Delete room: DELETE to "/room/yourRoomNumber", where yourRoomNumber is int.
5. Change room (replaces old room with new one): PUT to "room/yourRoomNumber", where yourRoomNumber is int with json {"number":yourInt, guests:["your", "guests", "names"]}.

To start server run "go run main.go model.go api.go".