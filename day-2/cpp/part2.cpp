// To run do: c++ part2.cpp -Wall -std=c++17 -o output2.out && ./output2.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>

using namespace std;

int print_string_vector(std::vector<std::string> str_vector)
{
    std::cout << str_vector.size() << " [";
    for (std::string str : str_vector)
    {
        std::cout << "'" << str << "'"
                  << ", ";
    }

    std::cout << "]" << std::endl;

    return 0;
}

std::vector<std::string> splitter(std::string s, std::string seperator)
{
    std::vector<std::string> res;

    // initialize index for seperator
    // size_t as the result of s.find is size_t??
    size_t pos_of_seperator;
    pos_of_seperator = s.find(seperator);

    while (pos_of_seperator != std::string::npos)
    {
        // Grab the substring
        std::string substring = s.substr(0, pos_of_seperator);
        // std::cout << s << "|" << substring << std::endl;

        // Add substring to array and erase the substring from original string
        res.push_back(substring);
        s.erase(0, pos_of_seperator + seperator.length());

        // find new seperator index
        pos_of_seperator = s.find(seperator);
        // std::cout << s << "|" << substring << "|" << 0 << ":" << pos_of_seperator << std::endl;
    }

    res.push_back(s);

    return res;
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
