#include "parse.hpp"

constexpr size_t red_count = 12;
constexpr size_t green_count = 13;
constexpr size_t blue_count = 14;

bool is_viable_game(const Game& game) {
  return std::ranges::all_of(game.cube_counts, [](const auto& cube_count) {
    return cube_count.red <= red_count
           && cube_count.green <= green_count
           && cube_count.blue <= blue_count;
  });
}

int main() {
  const auto games = parse();

  size_t id_sum = 0;

  for (const auto& game : games) {
    if (is_viable_game(game)) {
      std::cout << "viable: ";
      id_sum += game.id;
    } else {
      std::cout << "not viable: ";
    }
    std::cout << game << '\n';
  }

  std::cout << id_sum << '\n';
}
