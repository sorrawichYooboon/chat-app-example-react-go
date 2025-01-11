import React, { createContext, useContext, useState } from "react";

interface SocketContextValue {
  socket: WebSocket | null;
  connectSocket: (userName: string, roomName: string) => void;
  disconnectSocket: () => void;
}

export const SocketContext = createContext<SocketContextValue>({
  socket: null,
  connectSocket: () => {},
  disconnectSocket: () => {},
});

export const SocketProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [socket, setSocket] = useState<WebSocket | null>(null);

  const connectSocket = (userName: string, roomName: string) => {
    const wsUrl = `${
      import.meta.env.VITE_CHAT_APP_SERVER_URL || "ws://localhost:3000"
    }/ws?userName=${userName}&roomName=${roomName}`;

    const newSocket = new WebSocket(wsUrl);

    newSocket.onopen = () => {
      console.log("WebSocket connection established.");
    };

    newSocket.onclose = () => {
      console.log("WebSocket connection closed.");
    };

    newSocket.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    setSocket(newSocket);
  };

  const disconnectSocket = () => {
    if (socket) {
      socket.close();
      setSocket(null);
    }
  };

  return (
    <SocketContext.Provider value={{ socket, connectSocket, disconnectSocket }}>
      {children}
    </SocketContext.Provider>
  );
};

export const useSocket = (): SocketContextValue => {
  const context = useContext(SocketContext);

  if (!context) {
    throw new Error("useSocket must be used within a SocketProvider");
  }

  return context;
};
