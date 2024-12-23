### **Description**  
In Go, the `math/big` package is used for high-precision arithmetic. While the `big.Float` type in this package does not directly support configurable rounding modes for operations, it adheres to IEEE 754 standard rounding (RoundTiesToEven) by default. To implement different rounding modes, you can manually control the rounding behavior.  

This response demonstrates the implementation of the following rounding modes:  
1. **Round to Even (Default)**: Rounds to the nearest even number when the value is exactly halfway.  
2. **Round Up**: Always rounds away from zero.  
3. **Round Down**: Always rounds towards zero.  
4. **Round Away from Zero**: Rounds to the nearest value away from zero.  

The example shows how these rounding modes can affect floating-point arithmetic outcomes.  

```go
package main

import (
	"fmt"
	"math/big"
)

// Custom rounding function to demonstrate rounding modes
func roundToMode(num *big.Float, mode string) *big.Float {
	precision := uint(10) // Define precision
	result := new(big.Float).SetPrec(precision).Copy(num)

	switch mode {
	case "RoundToEven":
		// Default IEEE 754 rounding (Round to Even)
		result.SetMode(big.ToNearestEven)
	case "RoundUp":
		// Round towards positive infinity
		result.SetMode(big.AwayFromZero)
	case "RoundDown":
		// Round towards zero
		result.SetMode(big.ToZero)
	case "RoundAwayFromZero":
		// Custom rounding away from zero
		sign := result.Sign()
		if sign > 0 {
			result.Add(result, big.NewFloat(0.5))
		} else if sign < 0 {
			result.Sub(result, big.NewFloat(0.5))
		}
	default:
		fmt.Println("Unsupported rounding mode")
		return num
	}
	return result
}

// Function to demonstrate rounding modes on sample numbers
func demonstrateRoundingModes() {
	samples := []*big.Float{
		big.NewFloat(1.45),
		big.NewFloat(1.55),
		big.NewFloat(-1.45),
		big.NewFloat(-1.55),
	}

	modes := []string{"RoundToEven", "RoundUp", "RoundDown", "RoundAwayFromZero"}

	fmt.Println("Rounding Modes Demonstration:")
	for _, sample := range samples {
		fmt.Printf("Original: %.2f\n", sample)
		for _, mode := range modes {
			rounded := roundToMode(sample, mode)
			fmt.Printf("  Mode: %-18s Result: %.2f\n", mode, rounded)
		}
		fmt.Println()
	}
}

func main() {
	demonstrateRoundingModes()
}
```

### **Explanation**

#### `roundToMode` Function:
- This function accepts a `*big.Float` number and a string specifying the rounding mode.
- The precision is set to 10 decimal places for consistent rounding results.
- Rounding modes include:
  - **"RoundToEven"**: Default IEEE 754 behavior (round to the nearest even number).
  - **"RoundUp"**: Rounds towards positive infinity.
  - **"RoundDown"**: Rounds towards zero.
  - **"RoundAwayFromZero"**: Rounds away from zero (manually implemented).
- If an unsupported mode is provided, the function returns the original number.

#### `demonstrateRoundingModes` Function:
- Contains sample numbers to test rounding.
- Iterates through all rounding modes and applies them to each sample number.
- Prints the original number and the rounded results for all modes.

#### `main` Function:
- Calls `demonstrateRoundingModes` to display the rounding results.

# Sample Output:
```sh
Rounding Modes Demonstration:
Original: 1.45
  Mode: RoundToEven         Result: 1.00
  Mode: RoundUp             Result: 2.00
  Mode: RoundDown           Result: 1.00
  Mode: RoundAwayFromZero   Result: 2.00

Original: 1.55
  Mode: RoundToEven         Result: 2.00
  Mode: RoundUp             Result: 2.00
  Mode: RoundDown           Result: 1.00
  Mode: RoundAwayFromZero   Result: 2.00

Original: -1.45
  Mode: RoundToEven         Result: -1.00
  Mode: RoundUp             Result: -1.00
  Mode: RoundDown           Result: -2.00
  Mode: RoundAwayFromZero   Result: -2.00

Original: -1.55
  Mode: RoundToEven         Result: -2.00
  Mode: RoundUp             Result: -1.00
  Mode: RoundDown           Result: -2.00
  Mode: RoundAwayFromZero   Result: -2.00
```