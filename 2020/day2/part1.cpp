#include <fstream>
#include <iostream>
#include <tuple>
#include <string>
#include <algorithm>

std::tuple<int, int, char> parseConstraint(auto constraint)
{
    auto char_count_delimmiter = constraint.find_first_of("-");
    return {
        std::stoi(constraint.substr(0, char_count_delimmiter)),
        std::stoi(constraint.substr(char_count_delimmiter + 1, constraint.size() - 4)),
        constraint[constraint.size() - 2]
    };
}

int main(int argc, char* argv[])
{
    std::fstream passwords("data/passwords.txt", std::ios_base::in);

    std::string line, constraint, password;
    std::size_t delimmiter_pos;
    std::tuple<int, int, char> parsed_constraint;
    int char_count;
    auto valid_password_count = 0;
    while (std::getline(passwords, line)) {
        delimmiter_pos = line.find_last_of(" ");
        constraint = line.substr(0, delimmiter_pos);
        password = line.substr(delimmiter_pos + 1);
        parsed_constraint = parseConstraint(constraint);
        char_count = std::count(password.cbegin(), password.cend(), std::get<char>(parsed_constraint));
        if (char_count >= std::get<0>(parsed_constraint) && char_count <= std::get<1>(parsed_constraint)) {
            ++valid_password_count;
        }
    }

    std::cout << valid_password_count << std::endl;
    return 0;
}

