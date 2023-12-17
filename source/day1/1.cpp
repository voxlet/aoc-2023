#include <fstream>
#include <iostream>
#include <ranges>
#include <unordered_set>

auto to_int(char ascii) -> uint {
  return static_cast<uint>(ascii - '0');
}

auto find_first_digit(const std::ranges::range auto& line) {
  static const auto digits =
    std::unordered_set{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};

  auto found = std::ranges::find_if(
    line, [](const auto chr) { return digits.contains(chr); });

  if (found == std::ranges::end(line)) {
    throw std::runtime_error("No digit found in line");
  }

  return to_int(*found);
}

auto find_last_digit(const std::ranges::range auto& line) {
  return find_first_digit(std::ranges::reverse_view(line));
}

auto main() -> int {
  try {
    size_t sum = 0;

    std::ifstream input("../../../../source/day1/input.txt");

    for (std::string line; std::getline(input, line);) {
      const auto first_digit = find_first_digit(line);
      const auto last_digit = find_last_digit(line);

      constexpr auto tens = 10;
      auto val = tens * first_digit + last_digit;
      sum += val;

      std::cout << sum << ":" << val << " <- " << line << '\n';
    }

    std::cout << "Sum: " << sum << '\n';

    return 0;

  } catch (const std::exception& e) {
    std::cerr << "Error: " << e.what() << '\n';
    return 1;
  }
}
