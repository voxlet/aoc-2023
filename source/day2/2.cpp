#include "parse.hpp"

CubeCount max_cube_count(const Game& game) {
  CubeCount max;

  for (const auto& cube_count : game.cube_counts) {
    max.red = std::max(max.red, cube_count.red);
    max.green = std::max(max.green, cube_count.green);
    max.blue = std::max(max.blue, cube_count.blue);
  }

  return max;
}

size_t power(const CubeCount& cube_count) {
  return cube_count.red * cube_count.green * cube_count.blue;
}

int main() {
  const auto games = parse();

  size_t power_sum = 0;

  for (const auto& game : games) {
    std::cout << "game: " << game << '\n';

    CubeCount max = max_cube_count(game);
    std::cout << "max: " << max << '\n';

    power_sum += power(max);
  }

  std::cout << power_sum << '\n';
}
