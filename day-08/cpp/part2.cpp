// To run do: c++ part2.cpp -Wall -std=c++17 -o output2.out && ./output2.out
#include <string>
#include <vector>
#include <fstream>
#include <iostream>
#include <map>
#include "../../utils/utils.cpp"

using namespace std;

// greatest common divisor (GCD) via Euclidian algorithm
long GCD(long a, long b)
{
    while (b != 0)
    {
        long t = b;
        b = a % b;
        a = t;
    }

    return a;
}

// least common divisor (LCM) via GCD
long LCM(vector<long> integers)
{
    long result = integers[0] * integers[1] / GCD(integers[0], integers[1]);

    for (long i = 2; i < integers.size(); i++)
    {
        result = result * integers[i] / GCD(result, integers[i]);
    }

    return result;
}

int main()
{
    // Initialize a vector to store the file input and set the input file name.
    std::string fileName = "../input.txt";
    std::vector<std::string> lines = read_file(fileName);

    map<string, tuple<string, string>> network;
    string instructions;
    vector<string> starting_elements;

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

        if (parts[0][2] == 'A')
        {
            starting_elements.push_back(parts[0]);
        }
    }

    vector<long> steps_to_z;

    for (long i = 0; i < starting_elements.size(); i++)
    {
        string curr_element = starting_elements[i];
        long instruction_idx = 0;
        long steps = 0;
        while (curr_element[2] != 'Z')
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

        steps_to_z.push_back(steps);
    }

    // Print the result.
    std::cout << "Result: " << LCM(steps_to_z) << std::endl;
}