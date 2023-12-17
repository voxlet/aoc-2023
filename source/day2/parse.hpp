#include <fstream>
#include <iostream>
#include <ranges>
#include <stdexcept>
#include <string>
#include <string_view>
#include <vector>

struct CubeCount {
  size_t red = 0;
  size_t green = 0;
  size_t blue = 0;
};

struct Game {
  uint id;
  std::vector<CubeCount> cube_counts;
};

inline auto operator<<(std::ostream& out, const CubeCount& cube_count)
  -> std::ostream& {
  out << cube_count.red << " red, " << cube_count.green << " green, "
      << cube_count.blue << " blue" << '\n';
  return out;
}

inline auto operator<<(std::ostream& out, const Game& game) -> std::ostream& {
  out << "Game " << game.id << ": " << '\n';
  for (const auto& cube_count : game.cube_counts) {
    out << cube_count;
  }
  return out;
}

namespace details {

using std::literals::operator""sv;

constexpr auto game_token_length = "Game "sv.length();
constexpr auto game_delim = ": "sv;
constexpr auto grab_delim = "; "sv;
constexpr auto cube_delim = ", "sv;
constexpr auto color_delim = " "sv;

inline auto to_string_view(std::ranges::subrange<const char*> range)
  -> std::string_view {
  return {range.begin(),
          static_cast<size_t>(std::distance(range.begin(), range.end()))};
}

inline void assign_cube_count(CubeCount& cube_count,
                              std::string_view color_str,
                              size_t count) {
  if (color_str == "red") {
    cube_count.red = count;
  } else if (color_str == "green") {
    cube_count.green = count;
  } else if (color_str == "blue") {
    cube_count.blue = count;
  } else {
    throw std::runtime_error(std::string("Unknown color: ")
                             + std::string(color_str));
  }
}

inline auto parse_cube_counts(std::string_view grabs_str)
  -> std::vector<CubeCount> {
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

inline auto parse_game(std::string_view line) -> Game {
  const auto game_pos = line.find(game_delim);
  const auto game_id_str =
    line.substr(game_token_length, game_pos - game_token_length);
  const auto game_id = std::stoi(std::string(game_id_str));

  line.remove_prefix(game_pos + game_delim.length());

  return Game{
    .id = static_cast<uint>(game_id),
    .cube_counts = parse_cube_counts(line),
  };
}

}  // namespace details

inline auto parse() -> std::vector<Game> {
  std::ifstream input("../../../../source/day2/input.txt");
  std::vector<Game> games;

  for (std::string line; std::getline(input, line);) {
    games.push_back(details::parse_game(line));
  }

  return games;
}
