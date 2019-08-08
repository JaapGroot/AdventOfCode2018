#include <stdio.h>
#include <stdlib.h>


int main() 
{
	FILE *fptr, *output;
	int num = 0, num2 = 0, sum = 0;

	fptr = fopen("input", "r");
	output = fopen("output", "r");
	if(NULL == fptr)
	{
		printf("Can't open file!");
		exit(1);
	}

	while(fscanf(fptr, "%d\n", &num) > 0) {
		sum += num;
		while(fscanf(output, "%d\n", &num2) > 0){
			if(sum == num2){
				printf(" Found dup!");
			}
		}
	}


	printf("Value is: %d\n", sum);
	fclose(output);
	fclose(fptr);
}
