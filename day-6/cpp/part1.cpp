// To run do: c++ part1.cpp -Wall -std=c++17 -o output1.out && ./output1.out
#include <string>
#include <vector>
#include "../../utils/utils.cpp"

using namespace std;

// grabs the whole numbers from a string
vector<int> parse_numbers(string line)
{
    vector<int> numbers;
    int number = 0;
    while (line.size() > 0)
    {
        char c = line[0];
        line.erase(0, 1);
        if (c >= '0' && c <= '9')
        {
            number = number * 10 + (c - '0');
        }
        else
        {
            if (number > 0)
            {
                numbers.push_back(number);
            }
            number = 0;
        }
    }

    if (number > 0)
    {
        numbers.push_back(number);
    }

    return numbers;
}

int solve_quadratic_equation(float time, float dist)
{
    float ds = sqrt(time * time - 4 * dist);
    int from = floor((time - ds) / 2);
    int to = ceil((time + ds) / 2);

    return to - from - 1;
}

int main()
{
    // Initialize a vector to store the file input and set the input file name.
    std::string fileName = "../input.txt";
    std::vector<std::string> lines = read_file(fileName);

    vector<int> times = parse_numbers(lines[0]);
    vector<int> distances = parse_numbers(lines[1]);

    int total_combos = 1;
    for (int i = 0; i < times.size(); i++)
    {
        total_combos *= solve_quadratic_equation(times[i], distances[i]);
    }

    // Print the result.
    std::cout << "Result: " << total_combos << std::endl;
}