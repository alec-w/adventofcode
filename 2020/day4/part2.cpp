#include <fstream>
#include <iostream>
#include <string>
#include <sstream>
#include <vector>
#include <regex>

struct Passport
{
    bool byr = false;
    bool iyr = false;
    bool eyr = false;
    bool hgt = false;
    bool hcl = false;
    bool ecl = false;
    bool pid = false;
};

int main(int argc, char* argv[])
{
    std::fstream passports("data/passports.txt", std::ios_base::in);

    auto valid_passports = 0;
    std::string line;
    std::string data;
    Passport passport = {};
    std::string key;
    std::string value;
    std::size_t delimmiter_pos;
    std::istringstream stream;
    std::regex regex;
    while (std::getline(passports, line)) {
        if (line.empty()) {
            if (passport.byr && passport.iyr && passport.eyr && passport.hgt && passport.hcl && passport.ecl && passport.pid) {
                ++valid_passports;
            }
            passport = {};
            continue;
        }
        stream.str(line);
        stream.clear();
        while (stream >> data) {
            delimmiter_pos = data.find_first_of(":");
            key = data.substr(0, delimmiter_pos);
            value = data.substr(delimmiter_pos + 1);
            if (key == "byr") {
                regex = std::regex("(19[2-9]\\d)|(200[0-2])");
                passport.byr = std::regex_match(value, regex);
            }
            if (key == "iyr") {
                regex = std::regex("20((1\\d)|20)");
                passport.iyr = std::regex_match(value, regex);
            }
            if (key == "eyr") {
                regex = std::regex("20((2\\d)|30)");
                passport.eyr = std::regex_match(value, regex);
            }
            if (key == "hgt") {
                regex = std::regex("(1(([5-8]\\d)|(9[0-3]))cm)|(((59)|(6\\d)|(7[0-6]))in)");
                passport.hgt = std::regex_match(value, regex);
            }
            if (key == "hcl") {
                regex = std::regex("#[0-9a-f]{6}");
                passport.hcl = std::regex_match(value, regex);
            }
            if (key == "ecl") {
                regex = std::regex("(amb)|(blu)|(brn)|(gry)|(grn)|(hzl)|(oth)");
                passport.ecl = std::regex_match(value, regex);
            }
            if (key == "pid") {
                regex = std::regex("\\d{9}");
                passport.pid = std::regex_match(value, regex);
            }
        }
    }

    if (passport.byr && passport.iyr && passport.eyr && passport.hgt && passport.hcl && passport.ecl && passport.pid) {
        ++valid_passports;
    }

    std::cout << valid_passports << std::endl;
    return 0;
}

