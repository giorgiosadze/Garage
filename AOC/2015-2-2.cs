using System;
using System.IO;

var reader = new StreamReader("input").ReadToEnd().Split("\n");
var total = 0;
foreach(var r in reader.ToList()) {

    if(r.Split("x").Length == 3) {
        var a = r.Split("x")
            .Select(e => int.Parse(e))
            .ToList<int>();
        var l = a[0];
        var w = a[1];
        var h = a[2];
        a.Sort();
        var ans = a[0] * 2 + a[1] * 2 + l * w * h;
        total += ans;
    } 
}
Console.WriteLine(total);
