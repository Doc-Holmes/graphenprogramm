#schule #code
# Using go
Rust is too complicated to learn in this short amount of time

# Initialize the project
To create the project
- Create a directory
- Go inside
- Run `go mod init projectname`

To run the program
`go run .`

# The graph program

![](media/Anforderungen.jpg)
## Reading the file and parsing the csv
Reading the file is easy
Converting the csv to a 2d slice was easy as well
- But the resulting slice is a string slice
- In order to convert this to integers 
	- Using a nested for loop
	- Iterate over every item
	- Parse it
	- And write it to the same index in an integer slice

## Potency matrix
After some trial and error I figured out that I need a third nested for loop to iterate over the lines to calculate the cell itself
And after a lot of struggle to figure out which loop vars to use for the calculation and wondering why the solution was wrong
- I used a variable that takes all sum of the calculations to assign it to the cell
- The issue was that the variable with the sums was reset to 0 after the wrong loop