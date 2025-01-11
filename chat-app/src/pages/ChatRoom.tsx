import React, { useEffect, useState } from "react";
import { useSocket } from "../context/SocketContext";
import { LOCAL_STORAGE_KEYS, DEFAULTS } from "../utils/constants";

interface Message {
  userName: string;
  text: string;
}

const ChatRoom: React.FC = () => {
  const { socket, connectSocket, disconnectSocket } = useSocket();
  const [messages, setMessages] = useState<Message[]>([]);
  const [newMessage, setNewMessage] = useState("");

  const userName =
    localStorage.getItem(LOCAL_STORAGE_KEYS.USER_NAME) || DEFAULTS.USER_NAME;
  const roomName =
    localStorage.getItem(LOCAL_STORAGE_KEYS.ROOM_NAME) || DEFAULTS.ROOM_NAME;

  useEffect(() => {
    connectSocket(userName, roomName);

    return () => {
      disconnectSocket();
    };
  }, []);

  useEffect(() => {
    if (socket) {
      socket.onmessage = (event) => {
        const message: Message = JSON.parse(event.data);
        setMessages((prev) => [...prev, message]);
      };

      socket.onerror = (error) => {
        console.error("WebSocket error:", error);
      };
    }
  }, [socket]);

  const sendMessage = () => {
    if (socket && newMessage) {
      // Only send the text value to the backend
      socket.send(newMessage);
      setNewMessage("");
    }
  };

  console.log(messages);

  return (
    <div className="chat-container">
      <h1>Room: {roomName}</h1>
      <div className="messages">
        {messages.map((msg, index) => (
          <div key={index}>
            <strong>{msg.userName}:</strong> {msg.text}
          </div>
        ))}
      </div>
      <div className="input-container">
        <input
          type="text"
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}
          placeholder="Type a message..."
        />
        <button onClick={sendMessage}>Send</button>
      </div>
    </div>
  );
};

export default ChatRoom;
