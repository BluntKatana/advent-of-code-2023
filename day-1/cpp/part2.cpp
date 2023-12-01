// To run do: c++ part2.cpp -Wall -std=c++17 -o output2.out && ./output2.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include <map>

using namespace std;
map<string, int> mappedNumbers = {
    {"one", 1}, {"two", 2}, {"three", 3}, {"four", 4}, {"five", 5}, {"six", 6}, {"seven", 7}, {"eight", 8}, {"nine", 9}};

bool hasSuffix(std::string const &str, std::string const &suffix)
{
    // Check if the string is longer than the suffix
    if (str.length() >= suffix.length())
    {
        // Check if the last characters are the same
        return (0 == str.compare(str.length() - suffix.length(), suffix.length(), suffix));
    }
    return false;
}

int main()
{

    // Initialize a vector to store the file input and set the input file name.
    std::vector<std::string> lines;
    std::string fileName = "../input.txt";

    // Read the file and store the result in the lines vector.
    std::ifstream file(fileName);
    if (file.is_open())
    {
        std::string line;
        while (getline(file, line))
        {
            lines.push_back(line);
        }
        file.close();
    }

    int total = 0;

    // Iterate over the lines vector and print each line.
    for (std::string line : lines)
    {
        std::cout << line << std::endl;
        std::vector<int> numbers_in_line;
        std::string potential_num = "";

        // Iterate over each character in the line
        for (int i = 0; i < line.length(); i++)
        {
            int symbol = line[i];
            potential_num += symbol;

            // Check if the symbol is a number using ASCII
            if (symbol >= 48 && symbol <= 57)
            {
                numbers_in_line.push_back(symbol - '0');
            }

            // Loop through each keyval in the map
            for (auto const &keyval : mappedNumbers)
            {
                // If the potential number ends with the keyval.first,
                // then add the keyval.second to the numbers_in_line vector
                if (hasSuffix(potential_num, keyval.first))
                {
                    numbers_in_line.push_back(keyval.second);
                }
            }
        }

        potential_num = "";

        total += numbers_in_line[0] * 10 + numbers_in_line[numbers_in_line.size() - 1];
    }

    // Print the result.
    std::cout << "Result: " << total << std::endl;
}
