# Parent project does not export its library target, so this CML implicitly
# depends on being added from it, i.e. the testing is done only from the build
# tree and is not feasible from an install location

project(aoc-2023Tests LANGUAGES CXX)

# ---- Dependencies ----

find_package(Catch2 REQUIRED)
include(Catch)

# ---- Tests ----

add_executable(aoc-2023_test source/aoc-2023_test.cpp)
target_link_libraries(
    aoc-2023_test PRIVATE
    aoc-2023_lib
    Catch2::Catch2WithMain
)
target_compile_features(aoc-2023_test PRIVATE cxx_std_20)

catch_discover_tests(aoc-2023_test)

# ---- End-of-file commands ----

add_folders(Test)
