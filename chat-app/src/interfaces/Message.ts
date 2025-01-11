interface Message {
  type: string;
  payload: {
    userName: string;
    text: string;
  };
}
