// To run do: c++ part1.cpp -Wall -std=c++17 -o output1.out && ./output1.out
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

    map<string, tuple<string, string>> network;
    string instructions;

    for (int i = 0; i < lines.size(); i++)
    {
        if (i == 1)
        {
            continue;
        }

        if (i == 0)
        {
            instructions = lines[i];
            continue;
        }

        vector<string> parts = splitter(lines[i], " = ");
        vector<string> direction = splitter(parts[1], ", ");
        string left_direction = splitter(direction[0], "(")[1];
        string right_direction = splitter(direction[1], ")")[0];
        network[parts[0]] = tuple(left_direction, right_direction);
    }

    int steps = 0;
    string curr_element = "AAA";
    int instruction_idx = 0;

    while (curr_element != "ZZZ")
    {
        tuple<string, string> directions = network[curr_element];
        char instruction = instructions[instruction_idx];

        if (instruction == 'L')
        {
            curr_element = get<0>(directions);
        }
        else
        {
            curr_element = get<1>(directions);
        }

        if (instruction_idx == instructions.length() - 1)
        {
            instruction_idx = 0;
        }
        else
        {
            instruction_idx += 1;
        }

        steps += 1;
    }

    // Print the result.
    std::cout << "Result: " << steps << std::endl;
}
