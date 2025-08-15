Net-Cat 🖥️💬

Net-Cat is a Go-based TCP chat project inspired by Netcat, supporting a multi-client server, user name customization, member listing, and chat logging.

🚀 Features

Multi-client TCP server support.

Change your username using /name.

List all connected members with /members.

Send and receive messages with timestamps and sender names.

Automatic logging of all messages in logs.txt.

Notify all users when someone joins or leaves the chat.

Prevent sending empty messages.

📦 Requirements

Go >= 1.20

OS: Linux / Windows / macOS

⚡ Installation
git clone https://github.com/mohamedhill/net-cat.git 
cd net-cat


💻 Usage
Start the Server

go run . [port]

Starts the TCP server on the default port (e.g., 8080) and waits for clients.

Start a Client
./net-cat client


After connecting, you can use the following commands:

Change your name
/name Mohammed


Updates your username in the chat.

List connected members
/members


Shows all currently connected users.

Send a normal message
Hello everyone!


Broadcasts your message to all connected clients with timestamp and name.

Example Chat
[2025-08-15 14:30:12][Mohammed]: Hello everyone!
[2025-08-15 14:30:15][Alice]: Hi Mohammed!

📄 Logs (logs.txt)

Every message is saved automatically in logs.txt.

Useful for reviewing chat history later.
