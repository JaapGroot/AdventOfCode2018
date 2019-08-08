#include <stdio.h>
#include <stdlib.h>


int main() 
{
	FILE *fptr;
	int num = 0, sum = 0;

	fptr = fopen("input", "r");

	if(NULL == fptr)
	{
		printf("Can't open file!");
		exit(1);
	}

	while(fscanf(fptr, "%d\n", &num) > 0) {
		sum += num;
	}


	printf("Value is: %d\n", sum);

	fclose(fptr);
}
