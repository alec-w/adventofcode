#include <iostream>
#include <fstream>
#include <vector>

int main(int argc, char* argv[])
{
    std::fstream report("data/report.txt", std::ios_base::in);
    
    std::vector<int> values = {};
    int value;
    auto iter = values.cbegin();
    
    while (report >> value) {
        values.push_back(value);
        if (values.size() > 1) {
          iter = values.cbegin();
          for (iter = values.cbegin(); iter != std::prev(values.cend()); ++iter) {
            if (*iter + value == 2020) {
                std::cout << *iter * value << std::endl;
                return 0;
            }
          }
        }
    }
    
    std::cout << "None found" << std::endl;
    
    return 1;
}

