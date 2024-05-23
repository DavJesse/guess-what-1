# math-skills

This program is designed to perform statistical calculations on a set of data provided in a text file. It can calculate the average, median, variance, and standard deviation of the data set.
## Usage
To run the program, execute the following command in your terminal:

```
bash
$ go run . text-file.txt
```

Replace text-file.txt with the path to the text file containing the data you wish parse.

The data in the file should be presented as follows:
```
189
113
121
114
145
110
...
```

Each line in the file represents one value of the statistical population. While the program can read signs like '+' and '-', it only handles whole numbers.

After reading the file, the program will execute each of the calculations mentioned above and print the results to the terminal in the following format (rounded to the nearest whole number):
```
bash
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

## Requirements
- Go programming language
- Text file containing the statistical data
### Instructions
1. Ensure you have Go installed on your system.
2. Prepare a text file containing the statistical data.
3. Run the program using the provided command format, replacing test-file.txt with the name of your data file.
4. View the calculated results in the terminal.
## Code Explanation
The program reads the data from the specified file, calculates the average, median, variance, and standard deviation of the data set, and then prints the results to the terminal.
## Technical Details
- The program is written in Go.
- It uses built-in functions to read files and perform mathematical calculations.
- Error handling is included to handle cases where the file cannot be read or the data is invalid.
## Contributing
If you encounter any issues or have suggestions for improvements, feel free to contribute to the project. You can submit bug reports or pull requests on the GitHub repository.
## License
This project is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License). Feel free to use, modify, and distribute the code for personal or commercial purposes.