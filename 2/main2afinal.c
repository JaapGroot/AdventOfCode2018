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

int get_line_numbers(FILE *file);
int search_multiple(char* str);

int main() {
	int doubles = 0, triples = 0, line_nr = 0;
	char str[1000];	
	FILE *input;
	input =  fopen("input", "r");

	if(NULL == input){
		printf("Can't open file!");
		exit(1);
	}

	line_nr = get_line_numbers(input);
	rewind(input);
	
	for(int i = 0; i < line_nr; i ++){
		fgets(str, 1000, input);

		int result = search_multiple(str);
		
		if(result & 0x01){
			doubles++;
		}
		if(result & 0x02){
			triples++;
		}
	}
	printf("Doubles(%d) * triples(%d) = %d\n", doubles, triples, doubles * triples);
	fclose(input);
	return 0;
}
/* returns 1 if found double and 2 for triple. */
int search_multiple(char* str){
	int counter = 0;
	int retvalue = 0;
	for(int k = 0; k < strlen(str); k++) {
		for(int j = k + 1; j <= strlen(str);){
			if(str[k] == str[j]) {
				counter++;
				memmove(&str[j], &str[j+1], strlen(str) - j);
			}
			else {
				j++;
			}
		}
		if(counter == 1){
			retvalue |= 0x01;
		}
		if(counter == 2){
			retvalue |= 0x02;
		}
		counter = 0;
	}
	return retvalue;
}

int get_line_numbers(FILE *file){
	int count = 0;
	char c;
	// Extract characters from file and store in character c 
    for (c = getc(file); c != EOF; c = getc(file)){  
        if (c == '\n') {//Increment count if this character is newline
            count = count + 1;
		}
	}
	return count;
}
