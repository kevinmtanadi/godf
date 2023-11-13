# godf

[![Go Reference](https://pkg.go.dev/badge/github.com/kevinmtanadi/godf?utm_source=godoc)](https://pkg.go.dev/github.com/kevinmtanadi/godf)

A simple dataframe handler for golang, inspired by pandas from python. I used [@jedib0t](https://github.com/jedib0t)'s go-pretty to render the table.
Here's the library : [go-pretty](https://github.com/jedib0t/go-pretty/tree/main).

It can handle simple data manipulation such as :
- One hot encoding
- Sorting
- Shuffling data
- Filtering

## Install

```
go get -u github.com/kevinmtanadi/godf
```

## Usage
### Basic Usage
```golang
x := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
y := []float64{1.1, 1.2, 1.3, 1.4, 1.5}
z := []float64{21, 22, 23, 24, 25}

df := godf.DataFrame(map[string]interface{}{
  "x": x,
  "y": y,
  "z": z,
})

df.Show()
```

```
┌───┬─────┬─────┬────┐
│ # │   X │   Y │  Z │
├───┼─────┼─────┼────┤
│ 1 │ 0.1 │ 1.1 │ 21 │
│ 2 │ 0.2 │ 1.2 │ 22 │
│ 3 │ 0.3 │ 1.3 │ 23 │
│ 4 │ 0.4 │ 1.4 │ 24 │
│ 5 │ 0.5 │ 1.5 │ 25 │
└───┴─────┴─────┴────┘
```

### Read CSV File
```golang
df := godf.ReadCSV("data.csv")

df.Show()
```

### License
MIT
