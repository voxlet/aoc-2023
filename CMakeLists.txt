cmake_minimum_required(VERSION 3.14)

include(cmake/prelude.cmake)

project(
    aoc-2023
    VERSION 0.1.0
    DESCRIPTION "AoC 2023"
    HOMEPAGE_URL "https://adventofcode.com/2023"
    LANGUAGES CXX
)

include(cmake/project-is-top-level.cmake)
include(cmake/variables.cmake)

add_subdirectory(source/day1)
add_subdirectory(source/day2)

# ---- Declare library ----

# add_library(
#     aoc-2023_lib OBJECT
#     source/lib.cpp
# )

# target_include_directories(
#     aoc-2023_lib ${warning_guard}
#     PUBLIC
#     "$<BUILD_INTERFACE:${PROJECT_SOURCE_DIR}/source>"
# )

# target_compile_features(aoc-2023_lib PUBLIC cxx_std_20)

# find_package(fmt REQUIRED)
# target_link_libraries(aoc-2023_lib PRIVATE fmt::fmt)

# ---- Declare executable ----

# add_executable(aoc-2023_exe source/main.cpp)
# add_executable(aoc-2023::exe ALIAS aoc-2023_exe)

# set_property(TARGET aoc-2023_exe PROPERTY OUTPUT_NAME aoc-2023)

# target_compile_features(aoc-2023_exe PRIVATE cxx_std_20)

# target_link_libraries(aoc-2023_exe PRIVATE aoc-2023_lib)

# ---- Install rules ----

# if(NOT CMAKE_SKIP_INSTALL_RULES)
#   include(cmake/install-rules.cmake)
# endif()

# ---- Developer mode ----

if(NOT aoc-2023_DEVELOPER_MODE)
  return()
elseif(NOT PROJECT_IS_TOP_LEVEL)
  message(
      AUTHOR_WARNING
      "Developer mode is intended for developers of aoc-2023"
  )
endif()

include(cmake/dev-mode.cmake)
