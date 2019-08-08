#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

/* - Get the lenght of the string. 
 * - Keep track of two variables: int doubles and int triples.
 * - Find doubles and triples in string.
 * - Do this for every row.
 * - Multiply doubles and triples
 */

int main() {
	int length = 0;
	char str[255];
	char str2[255];
	FILE *input;

	for(int i = 0; i < 255; i ++){
		input =  fopen("input", "r");
		for(int d = 0; d <= i; d++){
			fgets(str, 1000, input);
		}
		length = strlen(str) -1 ; /* -1 to remove \n character */
		for(int j = i; j < 255; j++){
			int faults = 0;
			fgets(str2, 1000, input);
			for(int k  = 0; k < length; k++){
				if(str2[k] != str[k]){
					faults++;
				}
			}
			if(faults <= 1){
				printf("FOUND! %s, %s", str2, str);
			}
		}
		fclose(input);
	}
	return 0;
}
