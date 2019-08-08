#include <stdio.h>
#include <stdlib.h>


int main() 
{
	FILE *fptr, *wptr;
	int num = 0, sum = 0, i = 0;
	int result[1011];

	int test[4] = {1, 1, -2, 3};

	fptr = fopen("input", "r");
	wptr = fopen("output", "w");
	if(NULL == fptr)
	{
		printf("Can't open file!");
		exit(1);
	}

	while(fscanf(fptr, "%d\n", &num) > 0) {
//	for(i = 0; i < 4; i ++){
	//	num = test[i];
		result[i] = sum;
		sum += num;
		fprintf(wptr, "%d\n", sum);
		for(int j = 0; j < i; j++) {
			if(sum == result[j]) {
				printf(" found duplicate!sum: %d\n", sum);
			}
//			printf("sum=%d, result[%d] = %d\n", sum, j, result[j]);

		}
		i++;
	}


//			printf("sum=%d, result[%d] = %d\n", sum, i, result[i]);

	fclose(fptr);
	fclose(wptr);
}
