#include <fstream>
#include <vector>
#include <ranges>
#include <string_view>
#include <iostream>
#include <string>
#include <algorithm>
#include <stdexcept>
#include <ranges>

struct CubeCount {
  size_t red = 0;
  size_t green = 0;
  size_t blue = 0;
};

std::ostream& operator<<(std::ostream& os, const CubeCount& cube_count) {
  os << cube_count.red << " red, " << cube_count.green << " green, " << cube_count.blue << " blue" << '\n';
  return os;
}

struct Game {
  uint id;
  std::vector<CubeCount> cube_counts;
};

std::ostream& operator<<(std::ostream& os, const Game& game) {
  os << "Game " << game.id << ": " << '\n';
  for (const auto& cube_count : game.cube_counts) {
    os << cube_count;
  }
  return os;
}

using std::literals::operator""sv;

constexpr auto game_token_length = "Game "sv.length();
constexpr auto game_delim = ": "sv;
constexpr auto grab_delim = "; "sv;
constexpr auto cube_delim = ", "sv;
constexpr auto color_delim = " "sv;

std::string_view to_string_view(std::ranges::subrange<const char*> range) {
  return std::string_view(
    range.begin(),
    static_cast<size_t>(std::distance(range.begin(), range.end()))
  );
}

void assign_cube_count(CubeCount& cube_count, std::string_view color_str, size_t count) {
  if (color_str == "red") {
    cube_count.red = count;
  } else if (color_str == "green") {
    cube_count.green = count;
  } else if (color_str == "blue") {
    cube_count.blue = count;
  } else {
    throw std::runtime_error(std::string("Unknown color: ") + std::string(color_str));
  }
}

std::vector<CubeCount> parse_cube_counts(std::string_view grabs_str) {
  std::vector<CubeCount> cube_counts;

  for (const auto& grab : grabs_str | std::views::split(grab_delim)) {
    CubeCount cube_count;
    for (const auto& cube : grab | std::views::split(cube_delim)) {
      const auto cube_str = to_string_view(cube);

      const auto color_pos = cube_str.find(' ');
      const auto color_str = cube_str.substr(color_pos + 1);
      const auto count = std::stoi(std::string(cube_str.substr(0, color_pos)));

      assign_cube_count(cube_count, color_str, static_cast<size_t>(count));
    }
    cube_counts.push_back(cube_count);
  }

  return cube_counts;
}

Game parse_game(std::string_view line) {
  const auto game_pos = line.find(game_delim);
  const auto game_id_str = line.substr(game_token_length, game_pos - game_token_length);
  const auto id = std::stoi(std::string(game_id_str));

  line.remove_prefix(game_pos + game_delim.length());

  return Game{
    .id = static_cast<uint>(id),
    .cube_counts = parse_cube_counts(line),
  };
}

std::vector<Game> parse() {
  std::ifstream input("../../../../source/day2/input.txt");
  std::vector<Game> games;

  for (std::string line; std::getline(input, line);) {
    games.push_back(parse_game(line));
  }

  return games;
}
