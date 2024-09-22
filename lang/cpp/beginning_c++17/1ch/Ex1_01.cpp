// Ex1_01.cpp
// A complete C++ program
#include <iostream>

int main() {
  int answer{42}; // Defines answer with value 42

  std::cout << "The answer to life, the universe, and everything is " << answer
            << std::endl;

  return 0;
}

namespace my_space {
int testVariable{44};
}
