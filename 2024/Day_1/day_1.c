#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int
insertion_sort(int* list, int len, int number);

int
binary_search(int* list, int size, int number);

int
main()
{
  FILE* input_file = fopen("./input.txt", "r");
  char* line = malloc(20 * sizeof(char));
  bool is_first = true;

  int* first_list = calloc(1000, sizeof(int));
  int size_first = 0;
  int* second_list = calloc(1000, sizeof(int));
  int size_second = 0;

  while (fgets(line, 8, input_file) != NULL) {
    int num = atoi(line);
    if (is_first) {
      size_first = insertion_sort(first_list, size_first, num);
    } else {
      size_second = insertion_sort(second_list, size_second, num);
    }
    is_first = !is_first;
  }
  int sum = 0;

  for (int i = 0; i < size_first; i++) {
    int idx = binary_search(second_list, size_second, first_list[i]);
    if (idx == -1)
      continue;

    int count = 0;
    // look behind
    if (idx > 0) {
      for (; second_list[idx] == first_list[i]; idx--) {
      }
      idx++;
    }

    // count occurences of first_list[i]
    for (; second_list[idx] == first_list[i]; idx++)
      count++;

    sum += first_list[i] * count;
  }

  free(first_list);
  free(second_list);
  free(line);
  fclose(input_file);

  printf("%d\n", sum);
}

int
insertion_sort(int* list, int len, int number)
{
  if (len <= 0) {
    list[0] = number;
    return len + 1;
  }
  int i = len - 1;

  for (; i >= 0 && list[i] > number; i--)
    list[i + 1] = list[i];

  list[i + 1] = number;
  return ++len;
}

int
binary_search(int* list, int size, int number)
{
  int low = 0;
  int high = size;
  int mid;
  while (low <= high) {
    mid = (low + high) / 2;
    if (list[mid] == number) {
      return mid;
    } else if (list[mid] < number) {
      low = mid + 1;
    } else if (list[mid] > number) {
      high = mid - 1;
    }
  }
  return -1;
}
