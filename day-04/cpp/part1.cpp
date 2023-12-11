// To run do: c++ part1.cpp -Wall -std=c++17 -o output1.out && ./output1.out
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

    int scratch_total = 0;

    for (int i = 0; i < lines.size(); i++)
    {
        std::vector<std::string> words = splitter(lines[i], ": ");
        std::vector<std::string> split_numbers = splitter(words[1], " | ");
        std::vector<std::string> winning_numbers = splitter(trim(split_numbers[0]), " ");
        std::vector<std::string> own_numbers = splitter(trim(split_numbers[1]), " ");

        filter_empty_strings(&winning_numbers);
        filter_empty_strings(&own_numbers);

        int count_winners = 0;
        for (int j = 0; j < winning_numbers.size(); j++)
        {
            for (int k = 0; k < own_numbers.size(); k++)
            {
                if (winning_numbers[j] == own_numbers[k])
                {
                    if (count_winners == 0)
                    {
                        count_winners++;
                    }
                    else
                    {
                        count_winners *= 2;
                    }
                }
            }
        }

        scratch_total += count_winners;
    }

    // Print the result.
    std::cout << "Result: " << scratch_total << std::endl;
}
