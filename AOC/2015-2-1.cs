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
        var ans = 2 * l * w + 2 * w * h + 2 * h * l + (a[0] * a[1]);
        total += ans;
    } 
}
Console.WriteLine(total);
