const std = @import("std");

pub fn main() !void {
    var stream_server = try std.net.StreamServer.listen(.{ .port = 8082 });
    defer stream_server.close();

    std.log.info("Server listening on port 8082", .{});

    while (true) {
        var connection = try stream_server.accept();
        async handleConnection(connection) catch |err| {
            std.log.err("Error handling connection: {s}", .{@errorName(err)});
        };
    }
}

fn handleConnection(connection: std.net.StreamServer.Connection) !void {
    defer connection.close();
    var reader = connection.reader();
    var writer = connection.writer();

    var buf: [1024]u8 = undefined;
    while (true) {
        const bytes_read = try reader.read(&buf);
        if (bytes_read == 0) break; // Connection closed
        if (std.mem.indexOf(u8, &buf, "\r\n\r\n")) |end_of_headers| {
            _ = end_of_headers;
            break;
        }
    }

    const response = "HTTP/1.1 200 OK\nContent-Type: text/plain\nContent-Length: 21\n\nHello from async_srv!\n";

    _ = try writer.writeAll(response);
}
