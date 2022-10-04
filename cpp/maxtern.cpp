#include <iostream>
using std::cout;
using std::endl;

int
main ()
{
  int i = 30, j = 10, k = 20;
  int max = i > j ? (i > k ? i : k) : (j > k ? j : k);
  cout << max << endl;
  return 0;
}
