#include <cstdint>
#include <fstream>
#include <iostream>
#include <ranges>
#include <sstream>
#include <string>
#include <string_view>
#include <unordered_set>
#include <vector>

#include <concurrencpp/concurrencpp.h>  // NOLINT(*-include-cleaner)
#include <concurrencpp/forward_declarations.h>

struct Card {
  std::unordered_set<uint32_t> winning_numbers;
  std::vector<uint32_t> numbers;
};

inline auto operator<<(std::ostream& os, const Card& card) -> std::ostream& {
  os << "Winning numbers: ";
  for (auto number : card.winning_numbers) {
    os << number << ' ';
  }
  os << '\n';

  os << "Numbers: ";
  for (auto number : card.numbers) {
    os << number << ' ';
  }
  os << '\n';

  return os;
}

inline auto lines(std::string_view path)
  -> concurrencpp::generator<const std::string&> {
  std::ifstream input(path.data());

  for (std::string line; std::getline(input, line);) {
    co_yield line;
  }
}

namespace details {

using std::literals::operator""sv;

auto to_string_view(auto range) -> std::string_view {
  return {range.begin(), range.end()};
}

auto to_istringstream(auto range) -> std::istringstream {
  return std::istringstream{std::string{to_string_view(range)}};
}

inline auto parse_card(std::string_view line) -> Card {
  auto all_numbers_str_range =
    *(line | std::views::split(": "sv) | std::views::drop(1)).begin();
  auto all_numbers_str = to_string_view(all_numbers_str_range);
  auto all_numbers = all_numbers_str | std::views::split(" | "sv);

  auto card = Card{};

  auto numbers_str_it = all_numbers.begin();

  auto winning_numbers_strstream = to_istringstream(*numbers_str_it);

  for (auto n : std::views::istream<uint32_t>(winning_numbers_strstream)) {
    card.winning_numbers.insert(n);
  }

  ++numbers_str_it;

  auto numbers_strstream = to_istringstream(*numbers_str_it);

  for (auto n : std::views::istream<uint32_t>(numbers_strstream)) {
    card.numbers.push_back(n);
  }

  return card;
}

}  // namespace details

inline auto matchingNumbers(const Card& card) -> std::vector<uint32_t> {
  auto matches = std::vector<uint32_t>{};
  matches.reserve(card.numbers.size());

  for (auto number : card.numbers) {
    if (card.winning_numbers.contains(number)) {
      matches.push_back(number);
    }
  }

  return matches;
}

inline auto parse() -> std::vector<Card> {
  constexpr auto path = "../../../../source/day4/input.txt";

  auto cards = std::vector<Card>{};

  for (const auto& line : lines(path)) {
    cards.emplace_back(details::parse_card(line));
  }

  return cards;
}
