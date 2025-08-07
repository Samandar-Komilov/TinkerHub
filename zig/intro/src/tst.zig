const std = @import("std");

pub fn main() !void {
    var age: u8 = 11;
    age = 12;

    std.debug.print("The age is {}.\n", .{age});
}
