#include <fstream>
#include <iostream>
#include <string>
#include <vector>
#include <array>

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
    auto total_trees = -1ll;
    auto x_pos = 0;
    std::vector<char> active_row;
    std::array<std::pair<int, int>, 5> slopes = { { {1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2} } };
    for (const auto& slope : slopes) {
        trees = 0;
        x_pos = 0;
        std::cout << "slope " << slope.first << " " << slope.second << std::endl;
        for (auto y = 0u; y < topology.size(); y += slope.second) {
            active_row = topology[y];
            x_pos = x_pos % active_row.size();
            if (active_row[x_pos] == '#') {
                ++trees;
            }
            x_pos += slope.first;
        }
        if (-1 == total_trees) {
            total_trees = trees;
        } else {
            total_trees *= trees;
        }
        std::cout << trees << std::endl;
    }

    std::cout << total_trees << std::endl;
    return 0;
}

