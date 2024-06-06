using System.IO;

// AOC p1
new StreamReader("./input.txt")
    .ReadLine()
    .Select( e => e == '(' ? 1 : -1 )
    .Sum()

// AOC p2
var x = 1;
new StreamReader("./input.txt")
    .ReadLine()
    .Select(e => e == '(' ? 1 : -1)
    .TakeWhile(e => (x+=e) > -1 ? true : false)
    .Count()
