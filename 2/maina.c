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
	int doubles = 0, triples = 0, length = 0;
	char str[1000];	
	FILE *input;
	input =  fopen("input", "r");

	if(NULL == input){
		printf("Can't open file!");
		exit(1);
	}
	for(int i = 0; i < 20; i ++){
		fgets(str, 1000, input);
		bool first_dub = true;
		bool first_trip = true;
		length = strlen(str) -1 ; /* -1 to remove \n character */
		for(int j = 0; j < length; j++ ){
			bool nothing = false;
			bool found_double = false;
			bool found_triple = false;

			for(int k = j; k < strlen(str);){
				if(str[j] == str[k] && k != j){
					if(found_double && !found_triple && !nothing) {
						found_double = false;
						found_triple = true;
					}
					else if(!found_double && !found_triple && !nothing){
						found_double = true;
					}
					else if(!found_double && found_triple && !nothing){
						found_triple = false;
						nothing = true;
					}
					memmove(&str[k], &str[k+1], strlen(str) - k);
				}
				else{
					k++;
				}
				
			}
			if(found_double && first_dub){
				printf("Double, %i", i);
				doubles++;
				first_dub = false;
			}
			if(found_triple && first_trip){
				printf("Triple, %i", i);
				triples++;
				first_trip = false;
			}
		}
		
	}
	//	while(fgets(str, 1000, input) != '\n'){
//	}
	printf("Doubles(%d) * triples(%d) = %d\n", doubles, triples, doubles * triples);
	fclose(input);
	return 0;
}
