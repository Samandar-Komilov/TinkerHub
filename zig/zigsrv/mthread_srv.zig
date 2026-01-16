const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    var stream_server = try std.net.StreamServer.listen(.{ .port = 8081 });
    defer stream_server.close();

    std.log.info("Server listening on port 8081", .{});

    while (true) {
        const connection = try stream_server.accept();
        _ = try std.Thread.spawn(.{}, handleConnection, .{ connection, allocator });
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
        // Just read and ignore the request for this simple server
        if (std.mem.indexOf(u8, &buf, "\r\n\r\n")) |end_of_headers| {
            _ = end_of_headers;
            break;
        }
    }

    const response = "HTTP/1.1 200 OK\nContent-Type: text/plain\nContent-Length: 23\n\nHello from mthread_srv!";

    _ = try writer.writeAll(response);
}
