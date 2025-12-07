# 💰 Go Investment Calculator

This is a simple terminal-based investment calculator written in Go. It calculates the future value of an investment, both nominal and inflation-adjusted, based on user inputs.

## 🧠 Go Concepts Practiced

- `import` usage (`fmt`, `math`)
- Constant and variable declaration (`const`, `var`, `:=`)
- Reading user input using `fmt.Scan()`
- Using `math.Pow()` for exponential calculations
- Outputting results with `fmt.Println()`

## 📋 How It Works

1. Prompts user for:
   - Investment amount
   - Expected annual return rate
   - Number of years
2. Calculates:
   - **Future Value**: Using compound interest formula
   - **Real Value**: Adjusted for a constant inflation rate (2.5%)
3. Displays both results in the terminal.

## 🛠 How to Run

```bash
go run main.go
```

Follow the prompts in your terminal to enter the investment amount, return rate, and investment duration.

---

Use this calculator to plan smarter investments with inflation in mind! 📈📉
