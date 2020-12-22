#include <iostream>
#include <fstream>
#include <vector>

int main(int argc, char* argv[])
{
    std::fstream report("data/report.txt", std::ios_base::in);
    
    std::vector<int> values = {};
    int value;
    auto iter = values.cbegin();
    auto iter2 = values.cbegin();
    
    while (report >> value) {
        values.push_back(value);
        if (values.size() > 2) {
          iter = values.cbegin();
          for (iter = values.cbegin(); iter != std::prev(values.cend(), 2); ++iter) {
            for (iter2 = std::next(iter); iter2 != std::prev(values.cend()); ++iter2) {
              if (*iter + *iter2 + value == 2020) {
                std::cout << *iter * *iter2 * value << std::endl;
                return 0;
              }
            }
          }
        }
    }
    
    std::cout << "None found" << std::endl;
    
    return 1;
}

