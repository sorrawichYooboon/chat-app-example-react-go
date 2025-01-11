import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { LOCAL_STORAGE_KEYS } from "../utils/constants";

const JoinRoom: React.FC = () => {
  const [name, setName] = useState("");
  const [room, setRoom] = useState("");
  const navigate = useNavigate();

  const handleJoin = () => {
    if (name && room) {
      localStorage.setItem(LOCAL_STORAGE_KEYS.USER_NAME, name);
      localStorage.setItem(LOCAL_STORAGE_KEYS.ROOM_NAME, room);
      navigate("/chat");
    }
  };

  return (
    <div className="join-container">
      <h1>Join Chat Room</h1>
      <input
        type="text"
        placeholder="Enter your name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
      <input
        type="text"
        placeholder="Enter room name"
        value={room}
        onChange={(e) => setRoom(e.target.value)}
      />
      <button onClick={handleJoin}>Join</button>
    </div>
  );
};

export default JoinRoom;
