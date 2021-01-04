#include <fstream>
#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <algorithm>

void make_unique(auto &list)
{
    std::sort(list.begin(), list.end());
    auto last = std::unique(list.begin(), list.end());
    list.erase(last, list.end());
}

int main(int argc, char* argv[])
{
    std::fstream answers("data/answers.txt", std::ios_base::in);

    auto total_yeses = 0u;
    std::vector<char> answers_given_by_group;
    std::string line;
    std::string data;
    std::istringstream stream;
    while (std::getline(answers, line)) {
        if (line.empty()) {
            make_unique(answers_given_by_group);
            total_yeses += answers_given_by_group.size();
            answers_given_by_group.clear();
            continue;
        }
        stream.str(line);
        stream.clear();
        while (stream >> data) {
            for (const auto character : data) {
                answers_given_by_group.push_back(character);
            }
        }
    }

    make_unique(answers_given_by_group);
    total_yeses += answers_given_by_group.size();

    std::cout << total_yeses << std::endl;
    return 0;
}

