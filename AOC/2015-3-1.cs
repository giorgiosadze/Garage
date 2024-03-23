using System;
using System.IO;

var reader = new StreamReader("input");
var hashMap = new Dictionary<Tuple<int, int>, int>();
int x = 0;
int y = 0;

foreach(var d in reader.ReadLine()) {
    var key = Tuple.Create(x, y);

    if (hashMap.ContainsKey(key)) {
        hashMap[key] += 1;
    } else {
        hashMap.Add(key, 1);
    }
    
    switch(d)
    {
        case '>':
            x++;
            break;
        case '<':
            x--;
            break;
        case '^':
            y++;
            break;
        case 'v':
            y--;
            break;
    }
}
Console.WriteLine(hashMap.Count);
