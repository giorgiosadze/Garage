#include <stdio.h>

int main()
{
    FILE* fptr = fopen("input", "r"); 
    int i = 0;
    int c;
    while( (c = fgetc(fptr)) != EOF )
    {
        i += c == '(' ? 1 : 0;
        i += c == ')' ? -1 : 0;
    }
    printf("ans: %d\n",i);
}
