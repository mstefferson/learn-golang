// Declaration patterns
var x int = 10 // explicit type and value
var y = 10 // inferred type
z := 10 // short declaration (functions only)
var w int // zero value

// Type conversions
float64(intVar)
int(floatVar)
strconv.Itoa(intVar) // int to string
strconv.Atoi(stringVar) // string to int (returns error too!)

// Printf verbs
%v // default format
%T // type
%d // integer
%f // float
%s // string
%q // quoted string
%t // boolean
