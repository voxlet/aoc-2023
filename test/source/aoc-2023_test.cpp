#include <catch2/catch_test_macros.hpp>

#include "lib.hpp"

TEST_CASE("Name is aoc-2023", "[library]")
{
  auto const lib = library {};
  REQUIRE(lib.name == "aoc-2023");
}
