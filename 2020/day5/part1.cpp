#include <fstream>
#include <iostream>
#include <bitset>
#include <string>
#include <algorithm>

int main(int argc, char* argv[])
{
    std::fstream seats("data/seats.txt", std::ios_base::in);

    auto max_seat_id = 0u;
    auto seat_id = 0u;
    std::bitset<8> row;
    std::bitset<3> column;
    std::string line;
    while (std::getline(seats, line)) {
        row = std::bitset<8>(line, 0, 7, 'F', 'B');
        column = std::bitset<3>(line, 7, 3, 'L', 'R');
        seat_id = row.to_ulong() * 8 + column.to_ulong();
        max_seat_id = std::max(max_seat_id, seat_id);
    }

    std::cout << max_seat_id << std::endl;
    return 0;
}

