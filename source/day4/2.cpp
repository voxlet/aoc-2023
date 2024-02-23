#include <algorithm>
#include <cstddef>
#include <exception>
#include <iostream>
#include <numeric>
#include <ranges>
#include <span>
#include <vector>

#include "parse.hpp"

auto toMatchCounts(std::span<const Card> cards) -> std::vector<size_t> {
  auto matchCounts = std::vector<size_t>(cards.size());

  std::ranges::transform(cards, matchCounts.begin(), [](const Card& card) {
    return matchingNumbers(card).size();
  });

  return matchCounts;
}

auto wonCardCounts(std::span<const size_t> cardCounts,
                   std::span<const size_t> matchCounts) -> std::vector<size_t> {
  auto wonCardCounts = std::vector<size_t>(cardCounts.size());

  for (const auto cardIndex : std::views::iota(0uz, cardCounts.size())) {
    const auto matchCount = matchCounts[cardIndex];

    if (matchCount == 0) {
      continue;
    }

    const auto start = cardIndex + 1;
    const auto end = start + matchCount;

    for (auto wonIndex : std::views::iota(start, end)) {
      wonCardCounts[wonIndex] += cardCounts[cardIndex];
    }
  }

  return wonCardCounts;
}

auto addWonCards(std::span<const size_t> wonCounts,
                 std::span<size_t> cardCounts) -> void {
  for (const auto cardIndex : std::views::iota(0uz, cardCounts.size())) {
    cardCounts[cardIndex] += wonCounts[cardIndex];
  }
}

auto any_postive(std::span<const size_t> counts) -> bool {
  return std::ranges::any_of(counts,
                             [](const auto count) { return count > 0; });
}

auto totalWonCardCount(std::span<const Card> cards) -> size_t {
  auto cardCounts = std::vector<size_t>(cards.size(), 1uz);
  auto matchCounts = toMatchCounts(cards);

  for (auto wonCounts = wonCardCounts(cardCounts, matchCounts);
       any_postive(wonCounts);
       wonCounts = wonCardCounts(wonCounts, matchCounts))
  {
    addWonCards(wonCounts, cardCounts);
  }

  return std::accumulate(cardCounts.begin(), cardCounts.end(), 0uz);
}

auto main() -> int {
  try {
    const auto cards = parse();

    std::cout << "Total Count: " << totalWonCardCount(cards) << '\n';

    return 0;

  } catch (const std::exception& e) {
    std::cerr << "Error: " << e.what() << '\n';
    return 1;
  }
}
