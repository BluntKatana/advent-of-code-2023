// To run do: c++ part1.cpp -Wall -std=c++17 -o outpu1.out && ./output1.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>

using namespace std;

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
