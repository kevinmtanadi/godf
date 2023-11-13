package godf

import (
	"fmt"
	"testing"
)

func InitiateDummyDF() *dataframe {
	df := DataFrame(map[string]interface{}{
		"x": []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5},
		"y": []float64{1.5, 1.4, 1.3, 1.2, 1.1, 1.0, 0.9, 0.8, 0.7, 0.6, 0.5, 0.4, 0.3, 0.2, 0.1},
		"z": []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35},
		"a": []int{25, 24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11},
	})

	return df
}

func TestDataframe(t *testing.T) {
	df := InitiateDummyDF()

	df.Head()
}

func TestReadCSV(t *testing.T) {
	df := ReadCSV("https://raw.githubusercontent.com/mk-gurucharan/Regression/master/Startups_Data.csv")

	df.SetOption(DataframeOption{
		StringLimit: 10,
	})

	df.Head()

	dfLocal := ReadCSV("test_data/news.csv")
	dfLocal.OneHotEncode("label")
	dfLocal.Head()
}

func TestMerge(t *testing.T) {
	df1 := DataFrame(map[string]interface{}{
		"x": []float64{0.1, 0.2, 0.3, 0.4, 0.5},
		"y": []float64{1.1, 1.2, 1.3, 1.4, 1.5},
	})

	df2 := DataFrame(map[string]interface{}{
		"x": []float64{0.6, 0.7, 0.8, 0.9, 1.0},
		"y": []float64{1.6, 1.7, 1.8, 1.9, 2.0},
	})

	merged := df1.Merge(df2)

	merged.Show()
}

func TestJoin(t *testing.T) {
	df1 := DataFrame(map[string]interface{}{
		"x": []float64{0.1, 0.2, 0.3, 0.4, 0.5},
		"y": []float64{1.1, 1.2, 1.3, 1.4, 1.5},
	})

	df2 := DataFrame(map[string]interface{}{
		"z": []float64{21, 22, 23, 24, 25},
	})

	joined := df1.Join(df2)
	joined.Show()
}

func TestAppend(t *testing.T) {
	df := DataFrame(map[string]interface{}{
		"x": []float64{0.1, 0.2, 0.3, 0.4, 0.5},
		"y": []float64{1.1, 1.2, 1.3, 1.4, 1.5},
	})

	df.Append([]interface{}{0.6, "test"})
	df.Head()
}

func TestDropRow(t *testing.T) {
	df := InitiateDummyDF()

	df.DropRow(3)
	df.Head()
}

func TestDropCol(t *testing.T) {
	df := InitiateDummyDF()

	df.DropCol("z")

	fmt.Println(df)
	df.Head()
}

func TestGetRow(t *testing.T) {
	df := InitiateDummyDF()

	row := df.GetRow(2, 3)
	row.Show()
}

func TestGetCol(t *testing.T) {
	df := InitiateDummyDF()

	col := df.GetCol("x", "y")
	col.Show()

}

func TestFilterEq(t *testing.T) {
	df := InitiateDummyDF()

	filtered := df.Where(Eq{"z": 21})
	filtered.Show()
}

func TestFilterNotEq(t *testing.T) {
	df := InitiateDummyDF()

	filtered := df.Where(NotEq{"x": 0.1})
	filtered.Show()
}

func TestFilterGT(t *testing.T) {
	df := InitiateDummyDF()

	filtered := df.Where(GT{"x": 0.3})
	filtered.Show()
}

func TestFilterGTE(t *testing.T) {
	df := InitiateDummyDF()

	filtered := df.Where(Eq{"x": 0.3})
	filtered.Show()
}

func TestFilterLT(t *testing.T) {
	df := InitiateDummyDF()

	filtered := df.Where(LT{"x": 0.3})
	filtered.Show()
}

func TestFilterLTE(t *testing.T) {
	df := InitiateDummyDF()

	filtered := df.Where(LTE{"x": 0.3})
	filtered.Show()
}

func TestFilterIn(t *testing.T) {
	df := InitiateDummyDF()

	filtered := df.Where(In{"x": []float64{0.1, 0.2, 0.3}})
	filtered.Show()
}

func TestLimit(t *testing.T) {
	df := InitiateDummyDF()

	limited := df.Limit(8)
	limited.Show()
}

func TestExtractData(t *testing.T) {
	df := InitiateDummyDF()

	col := df.GetCol("x")

	data := col.ExtractData()
	fmt.Println(data)

	arrData := df.ExtractData()
	fmt.Println(arrData)
}

func TestCorr(t *testing.T) {
	df := InitiateDummyDF()

	df.Corr()
}
