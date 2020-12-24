#include <fstream>
#include <iostream>
#include <string>
#include <sstream>
#include <vector>

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
    std::istringstream stream;
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
            key = data.substr(0, data.find_first_of(":"));
            if (key == "byr") {
                passport.byr = true;
            }
            if (key == "iyr") {
                passport.iyr = true;
            }
            if (key == "eyr") {
                passport.eyr = true;
            }
            if (key == "hgt") {
                passport.hgt = true;
            }
            if (key == "hcl") {
                passport.hcl = true;
            }
            if (key == "ecl") {
                passport.ecl = true;
            }
            if (key == "pid") {
                passport.pid = true;
            }
        }
    }

    if (passport.byr && passport.iyr && passport.eyr && passport.hgt && passport.hcl && passport.ecl && passport.pid) {
        ++valid_passports;
    }

    std::cout << valid_passports << std::endl;
    return 0;
}

