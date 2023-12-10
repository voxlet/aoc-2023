#include <fstream>
#include <iostream>
#include <ranges>
#include <algorithm>
#include <unordered_map>
#include <cassert>

const auto digits = std::unordered_map<std::string, uint>{
    {"1", 1},
    {"2", 2},
    {"3", 3},
    {"4", 4},
    {"5", 5},
    {"6", 6},
    {"7", 7},
    {"8", 8},
    {"9", 9},
    {"one", 1},
    {"two", 2},
    {"three", 3},
    {"four", 4},
    {"five", 5},
    {"six", 6},
    {"seven", 7},
    {"eight", 8},
    {"nine", 9},
};

const auto max_digit_length = std::ranges::max(
    digits
    | std::views::keys
    | std::views::transform(&std::string::length)
);

auto find_digit(std::string_view line, size_t start, size_t end) {
    for (const auto& digit : digits | std::views::keys) {
        if (line.substr(start, end - start).find(digit) != std::string::npos) {
            return digits.find(std::string(digit));
        }
    }
    return digits.cend();
}

uint find_left(std::string_view line) {
    size_t start = 0;
    for (const auto end : std::views::iota(size_t{1}, line.length() + 1)) {
        if (end - start > max_digit_length) {
            ++start;
        }
        assert(end - start <= max_digit_length);

        if (const auto it = find_digit(line, start, end); it != digits.cend()) {
            return it->second;
        }
    }
    throw std::runtime_error(std::string("No left digit found: ") + std::string(line));
}

uint find_right(std::string_view line) {
    size_t end = line.length();
    for (const auto start : std::views::iota(size_t{0}, line.length())
                            | std::views::reverse) {
        if (end - start > max_digit_length) {
            --end;
        }
        assert(end - start <= max_digit_length);

        if (const auto it = find_digit(line, start, end); it != digits.cend()) {
            return it->second;
        }
    }
    throw std::runtime_error("No right digit found");
}

int main() {
    uint sum = 0;

    std::ifstream input("../../../../source/day1/input.txt");

    for (std::string line; std::getline(input, line);) {
        auto val = find_left(line) * 10 + find_right(line);

        sum += val;

        std::cout << sum << ":" << val << " <- " << line << '\n';
    }

    std::cout << "Sum: " << sum << '\n';

    return 0;
}