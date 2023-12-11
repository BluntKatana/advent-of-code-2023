// To run do: c++ part2.cpp -Wall -std=c++17 -o output2.out && ./output2.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include "../../utils/utils.cpp"

using namespace std;

int main()
{
    // Read the file and store the result in the lines vector.
    std::string fileName = "../input.txt";
    std::vector<std::string> lines = read_file(fileName);

    int game_id_total = 0;

    // Iterate over the lines vector and print each line.
    for (std::string line : lines)
    {
        // Parse the game_id and sets of cubes
        std::vector<std::string> game_id_and_cubes = splitter(line, ": ");
        std::vector<std::string> sets_of_cubes = splitter(game_id_and_cubes[1], "; ");

        int l_blue = 1;
        int l_red = 1;
        int l_green = 1;

        for (std::string set : sets_of_cubes)
        {
            std::vector<std::string> cubes_in_set = splitter(set, ", ");
            for (std::string cube : cubes_in_set)
            {
                std::vector<std::string> split_cube = splitter(cube, " ");
                int num = stoi(split_cube[0]);
                std::string color = split_cube[1];

                // Check for each color if the current num is largest
                if (color == "blue")
                {
                    if (num > l_blue)
                    {
                        l_blue = num;
                    }
                }
                else if (color == "red")
                {
                    if (num > l_red)
                    {
                        l_red = num;
                    }
                }
                else if (color == "green")
                {
                    if (num > l_green)
                    {
                        l_green = num;
                    }
                }
            }
        }

        game_id_total += l_blue * l_red * l_green;
    }

    // Print the result.
    std::cout << "Result: " << game_id_total << std::endl;
}
