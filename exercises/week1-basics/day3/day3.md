## Hands-On Exercises (20 min)

Create `day03/main.go` in your learn-golang repo.

### Exercise 1: If Statement Practice

**Requirements:**

1. Declare a variable `score` with value `85`
2. Use an if-else chain to assign a letter grade:
   - 90-100: "A"
   - 80-89: "B"
   - 70-79: "C"
   - Below 70: "F"
3. Print the score and grade

**Expected output:**

```
Score: 85, Grade: B
```

---

### Exercise 2: If with Initialization

**Requirements:**

1. Use an if statement with initialization to check if a random number (0-99) is even or odd
2. Print both the number and whether it's even or odd

**Expected output:**

```
Number 42 is even
```
or
```
Number 37 is odd
```

**Hint:** Use `rand.Intn(100)` and `num%2 == 0` for even check.

---

### Exercise 3: For Loop Variations

Go only has `for` loops, but they can act like different loop types!

**Requirements:**

1. Classic for loop: Print numbers 1 to 5
2. For-as-while: Count down from 5 to 1
3. Infinite loop with break: Print "tick" 3 times then break
4. Continue example: Print odd numbers from 1 to 10 (skip evens)

**Expected output:**

```
Count up: 1 2 3 4 5
Count down: 5 4 3 2 1
Tick: tick tick tick
Odds: 1 3 5 7 9
```

---

### Exercise 4: Range Over Collections

**Requirements:**

1. Create a slice of strings: `[]string{"apple", "banana", "cherry"}`
2. Use range to print each fruit with its index
3. Use range again but ignore the index (use `_`)

**Expected output:**

```
With index: 0:apple 1:banana 2:cherry
Without index: apple banana cherry
```

---

### Exercise 5: Switch Statement Practice

**Requirements:**

1. Create a switch that checks the current day of the week
2. Print different messages for Monday, Friday, weekends, and other days
3. Create a switch without expression to categorize temperature:
   - Below 32: "Freezing"
   - 32-59: "Cold"
   - 60-79: "Comfortable"
   - 80+: "Hot"

**Expected output:**

```
Today is Tuesday: Regular weekday
Temperature 75°F: Comfortable
```

**Hint:** Use `time.Now().Weekday()` and group Saturday/Sunday with commas.

---

### Exercise 6: Defer Demonstration

**Requirements:**

1. Create a function that demonstrates defer execution order
2. Use multiple defer statements to show LIFO (Last In, First Out) behavior
3. Show a common cleanup pattern with defer

**Expected output:**

```
Function starting
Opening resource
Function ending
Closing resource (deferred)
Third defer
Second defer
First defer
```

**Key concept:** Defer statements execute when the function returns, in reverse order of declaration.

---

### Exercise 7: Switch Type Assertion

**Requirements:**

1. Create a variable of type `interface{}` and assign it different values
2. Use a type switch to handle `string`, `int`, and `bool` cases
3. Include a default case for unknown types

**Expected output:**

```
hello is a string with length 5
42 is an integer
true is a boolean
3.14 is unknown type: float64
```

**Syntax reminder:** `switch v := value.(type) { case string: ... }`