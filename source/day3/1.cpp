#include <cstddef>
#include <cstdint>
#include <iostream>
#include <ranges>
#include <vector>

#include "parse.hpp"

auto is_adjacent_to_symbol(const Number& number,
                           size_t number_y,
                           const Schematic& schematic) -> bool {
  const auto [start, last] = adjacent_segment(number.range, schematic.width);

  for (const auto& row : adjacent_rows(schematic.rows, number_y)) {
    for (const auto& symbol : row.symbols) {
      if (symbol.pos >= start && symbol.pos <= last) {
        return true;
      }
      if (symbol.pos > last) {
        break;
      }
    }
  }

  return false;
}

auto main() -> int {
  auto schematic = parse();

  uint64_t sum = 0;

  for (const auto y : std::views::iota(size_t{0}, schematic.rows.size())) {
    const auto& row = schematic.rows[y];
    for (const auto& number : row.numbers) {
      if (is_adjacent_to_symbol(number, y, schematic)) {
        sum += number.value;
      }
    }
  }

  std::cout << "Sum: " << sum << '\n';

  return 0;
}
