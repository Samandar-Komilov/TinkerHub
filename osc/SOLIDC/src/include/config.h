#ifndef CONFIG_H
#define CONFIG_H

#define MAX_BUFFER_SIZE 1024
#define MAX_TRANSACTION_AMOUNT_TO_LOG 1000000.0
#define LOG_FILE "transactions.log"
#define LOG_PORT 8087
#define LOG_HOST "127.0.0.1"

enum FormatTypeEnum {
    TEXT,
    JSON,
};

enum TransportTypeEnum { DISK, TCP, UDP };

#endif // CONFIG_H
