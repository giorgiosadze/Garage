var reader = new StreamReader("input");
var total = 0;

var hashMap = new Dictionary<Tuple<int, int>, int>();
int x = 0;
int y = 0;

int x2 = 0;
int y2 = 0;


var turn = 0;
foreach(var d in reader.ReadLine()) {

    switch(d)
    {
        case '>':
            if(turn % 2 == 0) { x++; } else { x2++; } 
            break;
        case '<':
            if(turn % 2 == 0) { x--; } else { x2--; } 
            break;
        case '^':
            if(turn % 2 == 0) { y++; } else { y2++; } 
            break;
        case 'v':
            if(turn % 2 == 0) { y--; } else { y2--; } 
            break;
    }

    Tuple<int,int> key;
    if(turn % 2 == 0) {
        key = Tuple.Create(x, y);
    } else {
        key = Tuple.Create(x2, y2);
    }
    turn++;

    if (hashMap.ContainsKey(key)) {
        hashMap[key] += 1;
    } else {
        hashMap.Add(key, 1);
    }

    
}
Console.WriteLine(hashMap.Count);
