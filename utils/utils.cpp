#include <string>
#include <vector>
#include <fstream>
#include <iostream>

using namespace std;

// yoinked from friend sandeyez
std::vector<std::string> read_file(std::string &fileName)
{
    std::vector<std::string> lines;
    // Attempt to open the file.
    std::ifstream file(fileName);

    // If the file could not be opened, print an error message and return 1.
    if (!file)
    {
        printf("Error opening file %s\n", fileName.c_str());
        return lines;
    }

    // Read each line in the file and store it in the lines vector.
    std::string line;
    while (std::getline(file, line))
    {
        lines.push_back(line);
    }

    // Close the file
    file.close();

    // Return 0 to indicate success.
    return lines;
}

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