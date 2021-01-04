#include <fstream>
#include <iostream>
#include <bitset>
#include <string>
#include <algorithm>
#include <vector>

int main(int argc, char* argv[])
{
    std::fstream seats("data/seats.txt", std::ios_base::in);

    std::vector<unsigned long> seat_numbers;
    std::bitset<8> row;
    std::bitset<3> column;
    std::string line;
    while (std::getline(seats, line)) {
        row = std::bitset<8>(line, 0, 7, 'F', 'B');
        column = std::bitset<3>(line, 7, 3, 'L', 'R');
        seat_numbers.push_back(row.to_ulong() * 8 + column.to_ulong());
    }
    std::sort(seat_numbers.begin(), seat_numbers.end());

    auto my_seat = std::adjacent_find(seat_numbers.begin(), seat_numbers.end(), [](int left, int right){return left+1<right;});

    std::cout << 1 + *my_seat << std::endl;
    return 0;
}

