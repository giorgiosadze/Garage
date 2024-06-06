#include <fstream>
#include <iostream>
#include <filesystem>

int main() {
    std::filesystem::path filePath = "input";
    std::ifstream file(filePath);

    char ch;
    int count = 0;
    int floor = 0;
    int position = 0;
    while (file.get(ch)) {
        if(ch == '(')  {
            count++;
        } else {
            count--;
        }
        
        if( count == -1 && floor == 0 )
        {
            floor = position;
        }
        position++;
    }
    std::cout << count << " " << floor << "\n";
    file.close();

    return 0;
}
