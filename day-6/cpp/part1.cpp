// To run do: c++ part1.cpp -Wall -std=c++17 -o output1.out && ./output1.out
#include <string>
#include <vector>
#include "../../utils/utils.cpp"

using namespace std;

// grabs the whole numbers from a string
vector<int> parse_numbers(string line)
{
    vector<int> numbers;
    string number = "";
    for (int i = 0; i < line.size(); i++)
    {
        if (isdigit(line[i]))
        {
            number += line[i];
        }
        else
        {
            if (number != "")
            {
                numbers.push_back(stoi(number));
                number = "";
            }
        }
    }
    return numbers;
}

int main()
{
    // Initialize a vector to store the file input and set the input file name.
    std::string fileName = "../test_part1.txt";
    std::vector<std::string> lines = read_file(fileName);

    int scratch_total = 0;
    vector<int> times = parse_numbers(lines[0]);
    vector<int> distances = parse_numbers(lines[1]);

    // print the times and distances
    for (int i = 0; i < times.size(); i++)
    {
        std::cout << times[i] << " " << distances[i] << std::endl;
    }

    // Print the result.
    std::cout << "Result: " << scratch_total << std::endl;
}
