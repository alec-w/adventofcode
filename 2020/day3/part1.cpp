#include <fstream>
#include <iostream>
#include <string>
#include <vector>

int main(int argc, char* argv[])
{
    std::fstream topology_data("data/topology.txt", std::ios_base::in);

    std::vector<std::vector<char>> topology;
    std::vector<char> row;
    std::string line;
    while (std::getline(topology_data, line)) {
        row = {};
        for(auto character : line) {
            row.push_back(character);
        }
        topology.push_back(row);
    }

    auto trees = 0;
    auto x_pos = 0;
    for (const auto& active_row : topology) {
        x_pos = x_pos % active_row.size();
        if (active_row[x_pos] == '#') {
            ++trees;
        }
        x_pos += 3;
    }

    std::cout << trees << std::endl;
    return 0;
}

