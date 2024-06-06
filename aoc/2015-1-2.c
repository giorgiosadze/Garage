#include <stdio.h>

int main()
{
    FILE* fptr = fopen("input", "r"); 
    int i = 0;
    int pos = 0;
    int c;
    while( (c = fgetc(fptr)) != EOF )
    {
        pos++;
        i += c == '(' ? 1 : 0;
        i += c == ')' ? -1 : 0;
        if(i == -1)
        {
            break;
        }
    }
    printf("ans: %d\n",pos);
}
