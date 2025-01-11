# Chat App Example with React and Go

A real-time chat application built with React (frontend) and Go (backend). This project demonstrates a WebSocket-based chat system where users can join specific rooms and exchange messages in real time.

## **Project Features**

- **Real-Time Communication**: Powered by WebSockets for seamless message exchange.
- **Room-Based Chat**: Users can join specific rooms with unique names.
- **Frontend**: React-based UI with WebSocket connection management.
- **Backend**: Go server using Echo framework and Gorilla WebSocket library for WebSocket handling.

## **Technologies Used**

### Frontend

- React (with TypeScript)
- Vite
- WebSocket API

### Backend

- Go
- Echo Framework
- Gorilla WebSocket

## **Setup Instructions**

### Prerequisites

1. **Node.js** (v16 or later)
2. **Yarn** (package manager)
3. **Go** (v1.18 or later)

### Backend Setup

1. Clone the repository and navigate to the backend folder

   ```bash
   git clone https://github.com/yourusername/chat-app.git
   cd chat-app/chat-app-server
   ```

2. Install dependencies

   ```bash
   go mod tidy
   ```

3. Run the backend server

   ```bash
   go run main.go
   ```

### Frontend Setup

1. Navigate to the frontend folder

   ```bash
   cd chat-app-example-react-go/chat-app
   ```

2. Install dependencies

   ```bash
   yarn install
   ```

3. Create a .env file in the chat-app directory and add the following

   ```bash
   VITE_CHAT_APP_SERVER_URL=ws://localhost:3000
   ```

4. Run the frontend development server

   ```bash
   yarn dev
   ```

## **Code Overview**

### Backend (chat-app-server)

- **main.go**: Entry point of the server. Configures routes and middleware (e.g., CORS) and starts the WebSocket server.

- **handlers/websocket_handler.go**: Handles WebSocket upgrades and manages client connections to specific chat rooms.

- **models/room.go**: Manages chat rooms, including adding/removing clients and broadcasting messages.

- **models/client.go**: Represents a single client connection and handles message reading and writing.

### Frontend (chat-app)

- **context/SocketContext.tsx**: Manages the WebSocket connection lifecycle and provides context to components.

- **pages/JoinRoom.tsx**: Allows users to enter their username and room name to join a chat room.

- **pages/ChatRoom.tsx**: Displays messages in a chat room and allows users to send messages.

- **WebSocket Integration**: The frontend communicates with the backend using the native WebSocket API.

## **How the Code Works**

1. **Joining a Room**:

   - The user enters their username and room name in JoinRoom.tsx.
   - The frontend saves this information in localStorage and navigates to the chat room.

2. **Connecting to WebSocket**:

   - When the user lands in ChatRoom.tsx, the WebSocket connection is established with the server.
   - The backend creates or retrieves the specified room and adds the user to it.

3. **Message Exchange**:

   - The frontend sends plain text messages to the backend.
   - The backend wraps the message with the sender’s username and broadcasts it to all clients in the same room.

4. **Leaving a Room**:
   - When the user leaves the chat room or closes the tab, the WebSocket connection is closed, and the backend removes the user from the room.

## **How to Test**

1. **Start Both Servers**:

   - Start the backend server: `go run main.go`
   - Start the frontend development server: `yarn dev`

2. **Join a Room**:

   - Open your browser and navigate to `http://localhost:5173`.
   - Enter a unique username and a room name to join a chat room.
   - Open another browser tab and repeat the same process with a different username but the same room name.

3. **Send Messages**:

   - In one tab, type and send a message.
   - Verify that the message appears in real-time in both tabs with the correct sender’s username.

4. **Leave a Room**:
   - Close one of the tabs.
   - Check the backend logs to ensure the user is removed from the room.
