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

/* Get line numbers for for loop. */
int get_line_numbers(FILE *file);
/* Retrieve string from specified line. */
void get_string_from_line(FILE *file, int nr, char* buf);
/* Compare two string and return the number of faults. */
int compare_strings(char* str1, char* str2);

int main() {
	int line_nr = 0;
	bool found = false;
	char str1[255], str2[255];
	FILE *input;
	
	input =  fopen("input", "r");

    line_nr = get_line_numbers(input);
	/* For loop for the lines. */
	for(int i = 0; i < line_nr; i ++){
		get_string_from_line(input, i, str1);
		/* For loop to compare one line with all other lines. */
		for(int j = i + 1; j < line_nr; j++){
			get_string_from_line(input, j, str2);
			/* Check values on the go. */
			if(compare_strings(str1, str2) <= 1){
				found = true;
				break;
			}
		}
		if(found){
			/* Found the two strings, print them. */
			printf("%s%s", str1, str2);
			for(int i = 0; i < strlen(str1); i++) {
				if(str1[i] == str2[i]) {
					printf("%c", str1[i]);
				}
			}
			break;
		}
	}


	fclose(input);
	return 0;
}

int compare_strings(char* str1, char* str2){
	int faults = 0;
	for(int i  = 0; i < strlen(str1); i++){
		if(str1[i] != str2[i]){
			faults++;
		}
		/* More than one fault, no match. */
		if(faults > 1){
			break;
		}
	}
	return faults;
}

void get_string_from_line(FILE *file, int nr, char* buf){

	/* Go to first line */
	rewind(file);
	for(int d = 0; d <= nr; d++){
		/* Start at specific line */
		fgets(buf, 1000, file);
	}
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

