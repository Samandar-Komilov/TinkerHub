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