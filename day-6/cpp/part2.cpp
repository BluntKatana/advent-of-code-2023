// To run do: c++ part2.cpp -Wall -std=c++17 -o output2.out && ./output2.out
#include <string>
#include <vector>
#include "../../utils/utils.cpp"

using namespace std;

// grabs the whole numbers from a string and combines them
long parse_numbers_and_combine(string line)
{
    long total = 0;
    while (line.size() > 0)
    {
        char c = line[0];
        line.erase(0, 1);
        if (c >= '0' && c <= '9')
        {
            total = total * 10 + (c - '0');
        }
    }

    return total;
}

long solve_quadratic_equation(long time, long dist)
{
    long ds = sqrt(time * time - 4 * dist);
    long from = floor((time - ds) / 2);
    long to = ceil((time + ds) / 2);

    return to - from;
}

int main()
{
    // Initialize a vector to store the file input and set the input file name.
    string fileName = "../test_part1.txt";
    vector<string> lines = read_file(fileName);

    long time = parse_numbers_and_combine(lines[0]);
    long distance = parse_numbers_and_combine(lines[1]);

    cout << time << "|" << distance << endl;

    long total_combo = solve_quadratic_equation(time, distance);

    // Print the result.
    std::cout << "Result: " << total_combo << std::endl;
}