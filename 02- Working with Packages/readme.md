EACH file must belong to a package.

here for example both bank.go and communication.go belongs to main package.

the function printOptions() can be used directly inside bank.go without importing it from communication.go because both belongs to the same package.

EACH PACKAGE MUST HAS HIS OWN subFOLDER IN THE PROJECT project. and the package must have the same name as the subfolder.

to create another package, first create subfolder in the project folder, give it a name, i.e: 'utilities'.
then create a file in it and give it any name, i.e. 'fileops.go'. 
inside fileops.go the package name must be also 'utilities'.
any function that you need to export it inside fileops.go must start with capitla letter. 'ReadFromFile'.
then when you need to import that package i.e.'utilities' inside the main package do the following:
import (
    "name of the module/name of the package"
    i.e. example.com/bank/utilities
) 
then you can start using the functions inside that package. i.e. utilities.ReadFromFile()