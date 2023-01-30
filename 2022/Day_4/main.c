#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int split_term(char *);

int main() {
  FILE *f_ptr;
  char *line_bfr = malloc(10);
  size_t __n = 10;
  ssize_t read;
  int count;

  f_ptr = fopen("./input.txt", "r");

  if (f_ptr == NULL)
    return 1;

  while ((read = getline(&line_bfr, &__n, f_ptr)) != -1) {
    count += split_term(line_bfr);
    ;
  }

  printf("%d\n", count);
  free(line_bfr);
  free(f_ptr);
}

int split_term(char *term) {
  char *sub_inp;
  int *digits = malloc(sizeof(int) * 4);
  char *delim;
  char *comma;
  long int delim_idx;
  long int comma_idx;
  int out;

  sub_inp = malloc(2);

  comma = strstr(term, ",");
  comma_idx = comma - term;

  delim = strstr(term, "-");
  delim_idx = delim - term;

  strncpy(sub_inp, term, delim_idx);
  delim_idx = delim - term;
  digits[0] = atoi(sub_inp);

  strncpy(sub_inp, term + delim_idx + 1, comma - delim - 1);
  digits[1] = atoi(sub_inp);

  delim = strstr(term + comma_idx, "-");
  delim_idx = delim - term;
  sub_inp = malloc(2);

  strncpy(sub_inp, term + comma_idx + 1, delim - comma - 1);
  digits[2] = atoi(sub_inp);

  strncpy(sub_inp, term + delim_idx + 1, strlen(delim + 1));
  digits[3] = atoi(sub_inp);

  free(sub_inp);
  out = (digits[0] >= digits[2] && digits[0] <= digits[3]) ||
                (digits[1] >= digits[2] && digits[1] <= digits[3]) ||
                (digits[2] >= digits[0] && digits[2] <= digits[1]) ||
                (digits[2] >= digits[1] && digits[3] <= digits[1])
            ? 1
            : 0;
  return out;
}
