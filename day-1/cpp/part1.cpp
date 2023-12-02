// To run do: c++ part1.cpp -Wall -std=c++17 -o output1.out && ./output1.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include "../../utils/utils.cpp"

using namespace std;

int main()
{
    // Initialize a vector to store the file input and set the input file name.
    std::string fileName = "../input.txt";
    std::vector<std::string> lines = read_file(fileName);

    int total = 0;

    // Iterate over the lines vector and print each line.
    for (std::string line : lines)
    {
        std::vector<int> numbers_in_line;

        // Iterate over each character in the line
        for (int i = 0; i < line.length(); i++)
        {
            char symbol = line[i];

            // Check if the symbol is a number using ASCII
            if (symbol >= 48 && symbol <= 57)
            {
                numbers_in_line.push_back(symbol - '0');
            }
        }
        total += numbers_in_line[0] * 10 + numbers_in_line[numbers_in_line.size() - 1];
    }

    // Print the result.
    std::cout << "Result: " << total << std::endl;
}
