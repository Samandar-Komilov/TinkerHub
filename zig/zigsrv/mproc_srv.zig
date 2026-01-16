const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var stream_server = try std.net.StreamServer.listen(.{ .port = 8080 });
    defer stream_server.close();

    std.log.info("Server listening on port 8080", .{});

    while (true) {
        var connection = try stream_server.accept();
        const pid = try std.os.fork();

        if (pid == 0) {
            // Child process
            defer connection.close();
            handleConnection(connection, allocator) catch |err| {
                std.log.err("Error handling connection: {s}", .{@errorName(err)});
            };
            std.os.exit(0);
        } else {
            // Parent process
            connection.close(); // The child has the connection
        }
    }
}

fn handleConnection(connection: std.net.StreamServer.Connection) !void {
    var reader = connection.reader();
    var writer = connection.writer();

    var buf: [1024]u8 = undefined;
    while (true) {
        const bytes_read = try reader.read(&buf);
        if (bytes_read == 0) break; // Connection closed
        // Just read and ignore the request for this simple server
        if (std.mem.indexOf(u8, &buf, "\r\n\r\n")) |end_of_headers| {
            _ = end_of_headers;
            break;
        }
    }

    const response =
        \\ HTTP/1.1 200 OK
        \\ Content-Type: text/plain
        \\ Content-Length: 21
        \\ Hello from mproc_srv!
    ;

    _ = try writer.writeAll(response);
}
