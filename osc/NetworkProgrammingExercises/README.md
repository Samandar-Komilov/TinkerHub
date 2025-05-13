# Network Programming Exercises

References:
    - [Beej's Guide To Network Programming](https://beej.us/guide/bgnet/)
    - 

## ✅ Core Rules
- No copy-paste from AI or tutorials.
- Use man pages and strace as your main tools.
- Stick to IPv4 TCP, unless specified.
- Log everything: cserver/logs/dev.log.

## 🧠 Level 1 – Foundation Builders

### 1. Echo Server with Logging
- Accept one client at a time.
- Echo back every message.
- Log messages to stdout with timestamps.
- Limit buffer to 128 bytes.

Mini Chat Server (2 clients only)

    Allow exactly 2 clients.

    Relay messages between them.

    Detect disconnection and notify the other client.

Port Scanner

    Write a client that scans all ports on 127.0.0.1.

    Mark open/closed ports.

    Bonus: multithreaded scanner.

HTTP HEAD Server

    Accept HTTP requests, but respond only to HEAD method.

    Return static headers like:

        HTTP/1.1 200 OK
        Content-Length: 0

    Simple File Uploader (Client + Server)

        Client sends a file to the server.

        Server saves it under uploads/.

        Max file size = 1MB.

🔥 Level 2 – Real-World Concepts

    Timeout-aware Server

        Set SO_RCVTIMEO and SO_SNDTIMEO.

        Disconnect clients if idle for 5 seconds.

        Print a timeout message.

    DNS Resolver Client

        Use getaddrinfo() to resolve domain names.

        Print all returned IPs for google.com.

    Reverse Proxy Skeleton

        Client → proxy → upstream server.

        Use blocking I/O.

        Proxy just forwards and logs everything.

    Hex Dump Server

        Server accepts binary input.

        Logs a hexdump of each message.

        Bonus: write your own hexdump() logic.

    Telnet-like Terminal

        Connect to your server via netcat.

        Server responds with:

            Timestamp on connect.

            Echo + line number.

            "Bye" on /exit.

🚀 Level 3 – System Integration

    Multi-client Broadcast Server

        Use select() to manage multiple clients.

        Broadcast each client’s message to all others.

        Bonus: implement /nick command.

    Syslog-style Logger Server

        Client sends log lines.

        Server appends them to /var/log/mylog.log.

        Bonus: add log levels (INFO, WARN, ERROR).

    Command Execution Server (Careful)

        Accepts a command string from client.

        Executes it via popen() and returns output.

        Hardcode allowed commands only: ls, date, uptime.

    File Downloader Client

        Connect to a web server.

        Send HTTP GET.

        Parse headers and download the body to a local file.

    Game Server (Guess the Number)

        Server picks a random number.

        Each client guesses.

        Server replies: “Too high”, “Too low”, or “Correct!”.

        Limit guesses per client.

🧩 Bonus Challenges

    Concurrent Port Knocking Daemon

        Server waits for a sequence of connections on ports: 7000 → 8000 → 9000.

        After correct sequence, open port 9999 for a shell.

    Simple SOCKS4 Proxy

        Bare minimum spec.

        Supports TCP CONNECT to any IP+port.

    Mini Load Balancer

        Accept requests from clients.

        Forward to one of two upstream servers in round-robin.

        Log distribution stats.

    Raw Packet Sniffer

        Use AF_PACKET or raw sockets.

        Dump headers for all incoming packets on lo.

    Chatroom with Commands

        JOIN, LEAVE, /msg, /who

        Manage multiple rooms.

        Print a leaderboard of active users.

📓 How to Practice

    Create a folder per challenge.

    Use Makefile, version control, and README.

    Document every bug, every system call, and every confusion.

    If you get stuck, resist asking AI. Instead:

        Run strace.

        Use netcat, tcpdump, and wireshark.

        Read the man page twice.