## Hands-On Exercises (20 min)

Create `day02/main.go` in your learn-golang repo.

### Exercise 1: Type Declaration Practice

**Requirements:**

1. Declare a variable `temperature` as `float64` with value `98.6`
2. Declare a variable `isActive` using short declaration, set to `true`
3. Declare a variable `errorCount` with zero value (just type, no value)
4. Print all three with labels

**Expected output:**

```
Temperature: 98.6
Is Active: true
Error Count: 0
```

---

### Exercise 2: Zero Values Explorer

**Requirements:**

1. Declare variables of each type with NO initial value: `int`, `float64`, `bool`, `string`
2. Print each with a format that makes empty string visible (use quotes)

**Expected output:**

```
int zero: 0
float64 zero: 0
bool zero: false
string zero: ""
```

**Hint:** Use `fmt.Printf` with `%q` verb for quoted strings.

---

### Exercise 3: Type Conversion (Explicit Only!)

In Python: `x = int("42")` or `y = str(123)` just works.  
In Go: Conversions must be explicit, and `string(123)` does NOT do what you think!

**Requirements:**

1. Declare `intVal := 42`
2. Convert to `float64` and store in `floatVal`
3. Declare `floatPi := 3.14159`
4. Convert to `int` and store in `intPi` (note: truncates, doesn't round!)
5. Print both conversions

**Expected output:**

```
int 42 -> float64: 42
float64 3.14159 -> int: 3
```

**Syntax reminder:** `float64(intVal)` and `int(floatPi)`

---

### Exercise 4: Constants and iota

**Requirements:**

1. Create a `const` block defining task statuses: `Pending`, `Running`, `Complete`, `Failed` using `iota`
2. Create a constant `MaxRetries = 3`
3. Print the status values and max retries

**Expected output:**

```
Status values - Pending: 0, Running: 1, Complete: 2, Failed: 3
Max retries: 3
```

---

### Exercise 5: Type Inference Detective

**Requirements:**

1. Use short declaration for these values and use `fmt.Printf` with `%T` to print their inferred types:
   - `42`
   - `3.14`
   - `"hello"`
   - `true`
   - `'A'` (single character — what type is this?)

**Expected output:**

```
42 has type: int
3.14 has type: float64
hello has type: string
true has type: bool
A has type: int32
```
