// To run do: c++ part2.cpp -Wall -std=c++17 -o output2.out && ./output2.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include <map>
#include "../../utils/utils.cpp"

using namespace std;

struct Point
{
    int x;
    int y;
};

// Check if there are gears in the surrounding of a given position
std::__1::vector<Point> gears_in_surrounding(int row, int col, std::vector<std::vector<char>> grid)
{
    std::vector<Point> gears = {};

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

            if (curr_char == '*')
            {
                gears.push_back({i, j});
            }
        }
    }

    return gears;
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

    map<tuple<int, int>, std::vector<int>> gears = {};

    // Iterate through the engine to find valid words
    for (int row_num = 0; row_num < rows; row_num++)
    {
        // Look for a number in the line
        int curr_num = 0;
        std::vector<Point> gears_of_curr_num = {};
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
                if (curr_num > 0 && gears_of_curr_num.size() > 0)
                {
                    for (Point gear : gears_of_curr_num)
                    {
                        gears[tuple(gear.x, gear.y)].push_back(curr_num);
                    }
                }

                curr_num = 0;
                gears_of_curr_num = {};
            }

            std::vector<Point> surrounding_gears = gears_in_surrounding(row_num, col_num, engine);

            if (is_digit && surrounding_gears.size() > 0)
            {
                // add to gears_of_curr_num only if not already in it
                for (Point gear : surrounding_gears)
                {
                    bool is_in_gears_of_curr_num = false;
                    for (Point gear_of_curr_num : gears_of_curr_num)
                    {
                        if (gear_of_curr_num.x == gear.x && gear_of_curr_num.y == gear.y)
                        {
                            is_in_gears_of_curr_num = true;
                            break;
                        }
                    }

                    if (!is_in_gears_of_curr_num)
                    {
                        gears_of_curr_num.push_back(gear);
                    }
                }
            }
        }

        // check last number
        if (curr_num > 0 && gears_of_curr_num.size() > 0)
        {
            for (Point gear : gears_of_curr_num)
            {
                gears[tuple(gear.x, gear.y)].push_back(curr_num);
            }
        }
    }

    int total_part_numbers = 0;
    for (auto const &gear : gears)
    {
        if (gear.second.size() == 2)
        {
            total_part_numbers += gear.second[0] * gear.second[1];
        }
    }

    // Print the result.
    std::cout << "Result: " << total_part_numbers << std::endl;
}
