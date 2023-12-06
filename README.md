Introduction:
This project is an instant messaging system implemented in the Go programming language. The system utilizes WebSocket for basic user communication and employs WebRTC technology for audio and video communication, structured as a Selective Forwarding Unit (SFU) architecture. The server distributes streams to enable multi-party communication.

Running the Project:

```bash
go run *.go
```

Feature List:

- User login and registration (completed)
- Pushing historical messages (completed)
- Public chat (completed)
- Private chat (completed)
- Text messages (completed)
- Image messages (completed)
- File transfer (in progress)
- Real-time audio and video communication (completed)
- Online user statistics (completed)
- Historical user statistics (completed)
- Displaying a list of online users (completed)
- Editing user profile (avatar, username, password) (in progress)
- Message persistence (in progress)

Technology Stack:
- Golang
- Gin
- Gorilla/websocket
- Pion/webrtc
- ion/SFU
- JSON
