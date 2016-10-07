# goparsejson
A simple, but useful JSON parser for Go that is a little more intuitive when handling nested JSON structures and int64/float64 values.

Simply parses JSON data from a file or byte array and allows you to get at any item in the JSON object however many levels deep you want to go. Also handles type safety around numbers as for some reason Go wants to type everything as a Float64 by default (probably a good reason for this, but I prefer actual types). If you want to use this then just add it either under /json directory in your project directory and include using dot notation ./json like I did for the test or include it in your packages directory (library/etc).
