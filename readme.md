# Signaling Server for Omegle-like Site

This is a Go-based signaling server designed to enable real-time peer-to-peer communication for an Omegle-like site. The server facilitates WebRTC signaling for establishing connections between clients, enabling video chat in a web-based application.

## Features
- **WebSocket-Based Signaling**: Clients communicate via WebSocket to exchange signaling messages for establishing a peer-to-peer WebRTC connection.
- **Support for Multiple Clients**: Handles multiple simultaneous connections.
- **Session Management**: Allows users to initiate and join sessions in a matchmaking-like manner.
- **Pair Skipping**: Just like the real Omegle, the server can automatically skip matched pairs, so users can get a new random connection if they don't want to interact with the current one.
