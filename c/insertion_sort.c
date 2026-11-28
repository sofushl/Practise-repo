
#include <stdio.h>
void insertion_sort(int n, int a[n]) {
  for (int i = 1; i < n; i++) {
    int key = a[i];
    int j = i - 1;
    while (j >= 0 && a[j] > key) {
      a[j + 1] = a[j];
      j--;
    }
    a[j + 1] = key;
  }
}

int main() {
  int A[10] = {1, 2, 3, 412, 2, 45, 211, 12, 23, 3};
  for (int i = 0; i < 10; i++) {
    printf("%i ", A[i]);
  }
  printf("\n");

  insertion_sort(10, A);

  for (int i = 0; i < 10; i++) {
    printf("%i ", A[i]);
  }
  printf("\n");
}
