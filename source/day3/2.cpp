#include <cstddef>
#include <cstdint>
#include <iostream>
#include <optional>
#include <ranges>
#include <span>
#include <vector>

#include "parse.hpp"

auto adjacent_numbers(Symbol symbol,
                      size_t symbol_y,
                      const Schematic& schematic) -> std::vector<Number> {
  std::vector<Number> numbers;

  for (const auto& row : adjacent_rows(schematic.rows, symbol_y)) {
    for (auto number : row.numbers) {
      const auto [start, last] =
        adjacent_segment(number.range, schematic.width);

      if (symbol.pos >= start && symbol.pos <= last) {
        numbers.push_back(number);
      }
    }
  }

  return numbers;
}

auto gear_ratio(const Symbol& symbol,
                size_t symbol_y,
                const Schematic& schematic) -> std::optional<uint64_t> {
  if (symbol.value != '*') {
    return std::nullopt;
  }

  auto numbers = adjacent_numbers(symbol, symbol_y, schematic);

  if (numbers.size() != 2) {
    return std::nullopt;
  }

  return numbers[0].value * numbers[1].value;
}

auto main() -> int {
  auto schematic = parse();

  uint64_t sum = 0;

  for (const auto y : std::views::iota(0uz, schematic.rows.size())) {
    const auto& row = schematic.rows[y];
    for (const auto& symbol : row.symbols) {
      sum += gear_ratio(symbol, y, schematic).value_or(0);
    }
  }
  std::cout << "Sum: " << sum << '\n';

  return 0;
}
