# real-time-forum

## Overview

This project is a web-based forum application that allows users to register, login, create posts, comment on posts, and send private messages to each other. The application is built using a combination of technologies including **SQLite** for data storage, **Golang** for backend logic and WebSocket handling, and **JavaScript**, **HTML**, and **CSS** for the frontend. The application is designed as a **single-page application (SPA)**, where all page transitions are handled dynamically using JavaScript.

---

## Features

### **User Authentication**
- **Registration**: Users can register by providing:
  - Nickname
  - Age
  - Gender
  - First Name
  - Last Name
  - Email
  - Password
- **Login**: Users can log in using either their **nickname** or **email** along with their password.
- **Logout**: Users can log out from any page within the forum.

### **Posts and Comments**
- **Create Posts**: Users can create posts and assign them to specific categories.
- **Comment on Posts**: Users can comment on existing posts.
- **Feed Display**: Posts are displayed in a feed, and comments are visible only when a post is clicked.

### **Private Messages**
- **Online/Offline Status**: A section displays users who are online or offline, organized by:
  - The last message sent (like Discord).
  - Alphabetical order if no messages have been sent.
- **Real-Time Messaging**: Users can send and receive private messages in real-time using **WebSockets**.
- **Message History**:
  - Previous messages between users are visible.
  - The chat loads the last 10 messages and loads 10 more when scrolled up (using throttling/debouncing to avoid spamming).

---

## Technologies Used

### **Backend**
- **Golang**: Handles data and WebSocket connections.
- **SQLite**: Stores user data, posts, comments, and messages.
- **Gorilla WebSocket**: Manages real-time WebSocket connections.
- **bcrypt**: Used for password hashing.
- **UUID**: Generates unique identifiers for users and messages.

### **Frontend**
- **JavaScript**: Manages frontend events and WebSocket clients.
- **HTML**: Structures the elements of the page.
- **CSS**: Styles the elements of the page.

---

## Project Structure

- **SQLite Database**: Contains tables for:
  - Users
  - Posts
  - Comments
  - Messages
- **Golang Backend**: Manages:
  - API endpoints
  - WebSocket connections
  - Database interactions
- **JavaScript Frontend**: Handles:
  - Dynamic content loading
  - User interactions
  - Real-time updates
- **Single HTML File**: The entire application is contained within a single HTML file, with JavaScript handling all page transitions.

---

## Installation and Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/johneliud/real-time-forum.git
   cd real-time-forum
   ```

2. **Set Up the Database**:
   - Ensure **SQLite** is installed on your system.
   - Initialize the database by running the provided SQL scripts.

3. **Install Dependencies**:
   - Install Go dependencies:
     ```bash
     go mod tidy
     ```
   - Ensure all required Go packages are installed.

4. **Run the Application**:
   - Start the Go server:
     ```bash
     go run main.go
     ```
   - The default port to access the website is `http://localhost:9000`.

---

## Usage

- **Registration**: Navigate to the registration page and fill out the form to create a new account.
- **Login**: Use your credentials to log in and access the forum features.
- **Creating Posts**: Once logged in, you can create new posts and assign them to categories.
- **Commenting**: Click on any post to view and add comments.
- **Private Messaging**: Use the chat section to send and receive private messages in real-time.

---

## Allowed Packages

### **Backend**
- All standard Go packages.
- **Gorilla WebSocket**
- **sqlite3**
- **bcrypt**
- **gofrs/uuid** or **google/uuid**

### **Frontend**
- No external libraries or frameworks like React, Angular, or Vue are allowed.

---

## Learning Outcomes

- **Web Basics**:
  - HTML
  - HTTP
  - Sessions and cookies
  - CSS
  - Backend and Frontend
  - DOM
- **Go Routines**
- **Go Channels**
- **WebSockets**:
  - Go WebSockets
  - JS WebSockets
- **SQL Language**:
  - Manipulation of databases

---

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request with your changes.

---

## License

This project is licensed under the **MIT License**. See the [LICENSE](https://github.com/johneliud/real-time-forum/blob/main/LICENSE) file for details.

---

## Acknowledgments

- **Gorilla WebSocket**: For providing a robust WebSocket library.
- **SQLite**: For a lightweight and efficient database solution.
- **bcrypt**: For secure password hashing.

---

## Contact

For any questions or suggestions, please contact [me](mailto:johneliud4@gmail.com).
