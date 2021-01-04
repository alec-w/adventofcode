#include <fstream>
#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <algorithm>
#include <map>

int main(int argc, char* argv[])
{
    std::fstream answers("data/answers.txt", std::ios_base::in);

    auto total_yeses = 0u;
    auto all_yeses_in_group = 0u;
    auto number_in_group = 0u;
    std::map<char, unsigned long> answers_given_by_group;
    std::string line;
    std::string data;
    std::istringstream stream;
    while (std::getline(answers, line)) {
        if (line.empty()) {
            std::for_each(answers_given_by_group.begin(), answers_given_by_group.end(), [&all_yeses_in_group, number_in_group](auto pair){ all_yeses_in_group += pair.second == number_in_group; });
            total_yeses += all_yeses_in_group;
            answers_given_by_group.clear();
            all_yeses_in_group = 0;
            number_in_group = 0;
            continue;
        }
        stream.str(line);
        stream.clear();
        while (stream >> data) {
            for (const auto character : data) {
                answers_given_by_group[character] = answers_given_by_group.count(character) ? answers_given_by_group[character] + 1 : 1;
            }
            ++number_in_group;
        }
    }

    std::for_each(answers_given_by_group.begin(), answers_given_by_group.end(), [&all_yeses_in_group, number_in_group](auto pair){ all_yeses_in_group += pair.second == number_in_group; });
    total_yeses += all_yeses_in_group;

    std::cout << total_yeses << std::endl;
    return 0;
}

