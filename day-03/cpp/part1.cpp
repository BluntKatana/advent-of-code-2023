// To run do: c++ part1.cpp -Wall -std=c++17 -o output1.out && ./output1.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include <map>
#include "../../utils/utils.cpp"

using namespace std;

// check for all special characters in the surrounding of a given position
bool special_char_in_surrounding(int row, int col, std::vector<std::vector<char>> grid)
{
    bool has_special_char = false;

    // Check if the surrounding characters are special characters
    for (int i = row - 1; i < row + 2; i++)
    {
        for (int j = col - 1; j < col + 2; j++)
        {
            // Check if the current position is out of bounds
            if (i < 0 || j < 0 || i >= grid.size() || j >= grid[0].size())
            {
                continue;
            }
            char curr_char = grid[i][j];

            if (!isdigit(curr_char) && curr_char != '.')
            {
                has_special_char = true;
            }
        }
    }

    return has_special_char;
}

int main()
{
    // Initialize a vector to store the file input and set the input file name.
    std::string fileName = "../input.txt";
    std::vector<std::string> lines = read_file(fileName);

    // Initialize a 2d vector to store the engine
    std::vector<std::vector<char>> engine;
    int rows = lines.size();
    int cols = lines[0].size();

    // Initialize the engine
    for (int i = 0; i < rows; i++)
    {
        std::vector<char> row;
        for (int j = 0; j < cols; j++)
        {
            row.push_back(lines[i][j]);
        }
        engine.push_back(row);
    }

    int total_part_numbers = 0;

    // Iterate through the engine to find valid words
    for (int row_num = 0; row_num < rows; row_num++)
    {
        // Look for a number in the line, add the number to the total
        int curr_num = 0;
        bool is_valid_part = false;
        for (int col_num = 0; col_num < cols; col_num++)
        {
            char curr_char = engine[row_num][col_num];
            bool is_digit = isdigit(curr_char);
            // check if the current character is a number
            if (is_digit)
            {
                curr_num = curr_num * 10 + (curr_char - '0');
            }
            else
            {
                // If the character is not a number we reset the current number
                // and check if the current number is valid
                if (is_valid_part && curr_num > 0)
                {
                    total_part_numbers += curr_num;
                }

                curr_num = 0;
                is_valid_part = false;
            }

            if (is_digit && special_char_in_surrounding(row_num, col_num, engine))
            {
                is_valid_part = true;
            }
        }

        // check last number
        if (is_valid_part && curr_num > 0)
        {
            total_part_numbers += curr_num;
        }
    }

    // Print the result.
    std::cout << "Result: " << total_part_numbers << std::endl;
}
