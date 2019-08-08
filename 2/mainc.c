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
	int length = 0, count = 0;
	bool found = false;
	char str[255];
	char c;
	char str2[255];
	FILE *input;
	
	input =  fopen("input", "r");

    // Extract characters from file and store in character c 
    for (c = getc(input); c != EOF; c = getc(input)){  
        if (c == '\n') // Increment count if this character is newline
            count = count + 1; 
	}
	
	for(int i = 0; i < count; i ++){
		rewind(input);
		for(int d = 0; d <= i; d++){
			/* Start at specific line */
			fgets(str, 1000, input);
		}
		length = strlen(str) -1 ; /* -1 to remove \n character */
		for(int j = i; j < count; j++){
			int faults = 0;
			fgets(str2, 1000, input);
			for(int k  = 0; k < length; k++){
				if(str2[k] != str[k]){
					faults++;
				}
			}
			if(faults <= 1){
				found = true;
				break;
			}
		}
		if(found){
			printf("%s%s", str, str2);
			break;
		}
	}

	fclose(input);
	return 0;
}
