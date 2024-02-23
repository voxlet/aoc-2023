#include <cctype>
#include <cstddef>
#include <cstdint>
#include <fstream>
#include <iostream>
#include <ranges>
#include <string>
#include <string_view>
#include <variant>
#include <vector>

struct Number {
  struct Range {
    size_t start;
    size_t size;
  };

  uint64_t value;
  Range range;
};

inline auto operator<<(std::ostream& os, const Number& number)
  -> std::ostream& {
  os << "Number: " << number.value << " @" << number.range.start << ":"
     << number.range.size;
  return os;
}

struct Symbol {
  char value;
  size_t pos;
};

inline auto operator<<(std::ostream& os, const Symbol& symbol)
  -> std::ostream& {
  os << "Symbol: " << symbol.value << " @" << symbol.pos;
  return os;
}

struct Row {
  std::vector<Number> numbers;
  std::vector<Symbol> symbols;
};

inline auto operator<<(std::ostream& os, const Row& row) -> std::ostream& {
  os << "Row: " << '\n';
  for (const auto& number : row.numbers) {
    os << number << ", ";
  }
  os << '\n';
  for (const auto& symbol : row.symbols) {
    os << symbol << ", ";
  }
  os << '\n';
  return os;
}

struct Schematic {
  size_t width = 0;
  std::vector<Row> rows;
};

inline auto operator<<(std::ostream& os, const Schematic& schematic)
  -> std::ostream& {
  os << "Schematic: " << '\n';
  for (const auto& row : schematic.rows) {
    os << row;
  }
  return os;
}

namespace state {

struct SingleChar {};

struct NumberSequence {
  size_t start;
  std::string sequence;
};

}  // namespace state

inline auto to_number(const state::NumberSequence& numberSequence) -> Number {
  return Number{
    .value = std::stoul(numberSequence.sequence),
    .range =
      Number::Range{
        .start = numberSequence.start,
        .size = numberSequence.sequence.length(),
      },
  };
}

using State = std::variant<state::SingleChar, state::NumberSequence>;

template<class... Ts>
struct overload : Ts... {  // NOLINT(*-multiple-inheritance)
  using Ts::operator()...;
};

inline auto parse_row(std::string_view line) -> Row {
  Row row;
  State state = state::SingleChar{};

  for (const auto x : std::views::iota(size_t{0}, line.length())) {
    const auto c = line[x];

    auto visitor = overload{
      [&](state::SingleChar&) {
        if (std::isdigit(c) != 0) {
          state = state::NumberSequence{.start = x, .sequence = {c}};
        } else if (c != '.') {
          row.symbols.emplace_back(Symbol{.value = c, .pos = x});
        }
      },
      [&](state::NumberSequence& numberSequence) {
        if (std::isdigit(c) != 0) {
          numberSequence.sequence.append(1, c);
        } else {
          if (c != '.') {
            row.symbols.emplace_back(Symbol{.value = c, .pos = x});
          }
          row.numbers.emplace_back(to_number(numberSequence));
          state = state::SingleChar{};
        }
      },
    };

    std::visit(visitor, state);
  }

  if (std::holds_alternative<state::NumberSequence>(state)) {
    const auto& numberSequence = std::get<state::NumberSequence>(state);
    row.numbers.emplace_back(to_number(numberSequence));
  }

  return row;
}

inline auto parse() -> Schematic {
  Schematic schematic;

  std::ifstream input("../../../../source/day3/input.txt");

  for (std::string line; std::getline(input, line);) {
    schematic.rows.emplace_back(parse_row(line));
    schematic.width = line.size();
  }

  return schematic;
}

inline auto adjacent_rows(std::span<const Row> rows, size_t y)
  -> std::span<const Row> {
  const size_t offset = std::max(y, 1uz) - 1;
  const size_t last = std::min(y, rows.size() - 2) + 1;
  return rows.subspan(offset, last - offset + 1);
}

inline auto adjacent_segment(Number::Range r, size_t width)
  -> std::pair<size_t, size_t> {
  return {
    std::max(r.start, 1uz) - 1,
    std::min(r.start + r.size - 1, width - 2) + 1,
  };
}
