export class Chat {
  constructor() {
    this.messages = [];
    this.messageCount = 0;
    this.socket = new WebSocket('ws://localhost:9000/ws');

    this.socket.onopen = () => {
      console.log('Connected to WebSocket server');
    };

    this.socket.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      this.messages.push(msg);
      this.updateMessageList();
    };
  }

  renderChat() {
    const chatContainer = document.createElement('div');
    chatContainer.classList.add('chat-container');

    const messageList = document.createElement('div');
    messageList.classList.add('message-list');
    chatContainer.appendChild(messageList);

    const messageInputContainer = document.createElement('div');
    messageInputContainer.classList.add('message-input-container');

    const messageInput = document.createElement('input');
    messageInput.setAttribute('placeholder', 'Type a message...');
    messageInput.classList.add('message-input');

    const sendButton = document.createElement('button');
    sendButton.textContent = 'Send';
    sendButton.classList.add('send-button');
    sendButton.addEventListener('click', () => this.sendMessage(messageInput.value));

    messageInputContainer.appendChild(messageInput);
    messageInputContainer.appendChild(sendButton);
    chatContainer.appendChild(messageInputContainer);

    document.getElementById('app').appendChild(chatContainer);
  }

  sendMessage(content) {
    if (content.trim()) {
      const message = { 
        content: content, 
        sender: 'User1', 
        timestamp: new Date().toLocaleTimeString() 
      }; 
      this.socket.send(JSON.stringify(message));
      this.messages.push(message);
      this.updateMessageList();
    }
  }

  updateMessageList() {
    const messageList = document.querySelector('.message-list');
    messageList.innerHTML = '';
    this.messages.forEach((message) => {
      const messageItem = document.createElement('div');
      messageItem.classList.add('message-item');
      messageItem.textContent = `${message.sender}: ${message.content} (${message.timestamp})`;
      messageList.appendChild(messageItem);
    });
  }
}
