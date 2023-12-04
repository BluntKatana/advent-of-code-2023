// To run do: c++ part2.cpp -Wall -std=c++17 -o output2.out && ./output2.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include <map>
#include "../../utils/utils.cpp"

using namespace std;

// filter empty strings from vector
void filter_empty_strings(std::vector<std::string> *str_vector)
{
    for (int i = 0; i < str_vector->size(); i++)
    {
        if ((*str_vector)[i] == "")
        {
            str_vector->erase(str_vector->begin() + i);
        }
    }
}

int main()
{
    // Initialize a vector to store the file input and set the input file name.
    std::string fileName = "../input.txt";
    std::vector<std::string> lines = read_file(fileName);

    // initialize map
    map<int, int> scratch_map;

    for (int scratch_num = 0; scratch_num < lines.size(); scratch_num++)
    {
        // parse out the winning and own numbers
        std::vector<std::string> words = splitter(lines[scratch_num], ": ");
        std::vector<std::string> split_numbers = splitter(words[1], " | ");
        std::vector<std::string> winning_numbers = splitter(trim(split_numbers[0]), " ");
        std::vector<std::string> own_numbers = splitter(trim(split_numbers[1]), " ");

        // filter out empty strings
        filter_empty_strings(&winning_numbers);
        filter_empty_strings(&own_numbers);

        // count the total number of winners
        int count_winners = 0;
        for (int j = 0; j < winning_numbers.size(); j++)
        {
            for (int k = 0; k < own_numbers.size(); k++)
            {
                if (winning_numbers[j] == own_numbers[k])
                {
                    count_winners++;
                }
            }
        }

        // check if we won
        if (count_winners > 0)
        {
            // if we won, add the count to the next count_winners scratch_maps
            for (int j = 1; j <= count_winners; j++)
            {
                scratch_map[scratch_num + j] += 1 + scratch_map[scratch_num];
            }
        }

        // add the scratch to self
        scratch_map[scratch_num] += 1;
    }

    // loop through map and add up all the values
    int scratch_total = 0;
    for (auto scratch : scratch_map)
    {
        scratch_total += scratch.second;
    }

    // Print the result.
    std::cout << "Result: " << scratch_total << std::endl;
}
