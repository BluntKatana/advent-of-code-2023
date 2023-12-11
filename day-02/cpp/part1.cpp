// To run do: c++ part1.cpp -Wall -std=c++17 -o outpu1.out && ./output1.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include <map>
#include "../../utils/utils.cpp"

using namespace std;


int main()
{
    // Initialize a vector to store the file input and set the input file name.
    std::string fileName = "../input.txt";
    std::vector<std::string> lines = read_file(fileName);

    map<string, int> MAP_MAX_PER_COLOR = {
        {"red", 12}, {"green", 13}, {"blue", 14}};

    int game_id_total = 0;

    // Iterate over the lines vector and print each line.
    for (std::string line : lines)
    {
        // Parse the game_id and sets of cubes
        std::vector<std::string> game_id_and_cubes = splitter(line, ": ");
        int game_id = stoi(splitter(game_id_and_cubes[0], " ")[1]);
        std::vector<std::string> sets_of_cubes = splitter(game_id_and_cubes[1], "; ");

        bool is_valid_set = true;
        for (std::string set : sets_of_cubes)
        {
            std::vector<std::string> cubes_in_set = splitter(set, ", ");
            for (std::string cube : cubes_in_set)
            {
                std::vector<std::string> split_cube = splitter(cube, " ");
                int num = stoi(split_cube[0]);
                std::string color = split_cube[1];

                // Check if the cube amount if larger than the maximum allowed
                if (num > MAP_MAX_PER_COLOR[color])
                {
                    is_valid_set = false;
                }
            }
        }

        if (is_valid_set)
        {
            game_id_total += game_id;
        }
    }

    // Print the result.
    std::cout << "Result: " << game_id_total << std::endl;
}
