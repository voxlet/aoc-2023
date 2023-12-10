#include <fstream>
#include <iostream>
#include <ranges>
#include <unordered_set>

const auto digits = std::unordered_set{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'};

uint to_int(char ascii) {
    return static_cast<uint>(ascii - '0');
}

auto find_first_digit(const std::ranges::range auto& line) {
    return std::ranges::find_if(line, [](const auto c) {
        return digits.contains(c);
    });
}

int main() {
    size_t sum = 0;

    std::ifstream input("../../../../source/day1/input.txt");

    for (std::string line; std::getline(input, line);) {
        const auto first_digit = find_first_digit(line);
        const auto last_digit = find_first_digit(std::ranges::reverse_view(line));

        auto val = to_int(*first_digit) * 10 + to_int(*last_digit);
        sum += val;

        std::cout << sum << ":" << val << " <- " << line << '\n';
    }

    std::cout << "Sum: " << sum << '\n';

    return 0;
}
