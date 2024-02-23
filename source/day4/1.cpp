#include <cstdint>
#include <exception>
#include <iostream>

#include "parse.hpp"

auto uintpow2(uint64_t exp) -> uint64_t {
  auto res = 1u;
  for (uint64_t i = 0; i < exp; ++i) {
    res *= 2;
  }
  return res;
}

auto main() -> int {
  try {
    auto cards = parse();

    uint64_t sum = 0;

    for (const auto& card : cards) {
      auto matches = matchingNumbers(card);
      if (!matches.empty()) {
        sum += uintpow2(matches.size() - 1);
      }
    }

    std::cout << "Sum: " << sum << '\n';

    return 0;

  } catch (const std::exception& e) {
    std::cerr << "Error: " << e.what() << '\n';
    return 1;
  }
}
