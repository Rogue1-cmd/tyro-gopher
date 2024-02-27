//The dataset used for this code has the following columns
//Book,Author(s),Original language,First published,Approximate sales in millions,Genre

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {
	//>>Open the csv File
	csvfile, err := os.Open("best-selling-books.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	csvfile2, err := os.Open("books.csv")

	if err != nil {
		log.Fatal(err)
	}
	//>> Close the file when the function returns
	defer csvfile2.Close()

	//>>CSV<<
	df := dataframe.ReadCSV(csvfile)
	df2 := dataframe.ReadCSV(csvfile2)
	//fmt.Println(df)

	//>>json<<
	/*df := dataframe.ReadJSON()
	fmt.Println(df)*/

	//>>ANALYZING DATAFRAME<<
	//fmt.Println(df.Dims())//rows and columns
	//fmt.Println(df.Ncol())//columns
	//fmt.Println(df.Nrow())//rows
	//fmt.Println(df.Names())//titles
	//fmt.Println(df.Types())// data types

	//fmt.Println(df.Describe())

	//>>QUERING/SORTING AND FILTERING<<
	//col1 := df.Select("Author(s)") //Dataframe object
	//fmt.Println(col1)

	//row1 := df.Subset(12) //Dataframe Object
	//fmt.Println(row1)

	//ds := df.Col("Author(s)")
	//fmt.Println(ds.IsNaN()) //IsNaN returns an array that identifies which of the elements are NaN.

	/*df = df.Arrange(
		dataframe.Sort("Approximate sales in millions"),
	)
	fmt.Println(df)*/

	/*df = df.Arrange(
		dataframe.RevSort("Approximate sales in millions"),
	)
	fmt.Println(df)*/

	//>>FILTER<<
	/*
		df = df.Filter(
			dataframe.F{5, "Approximate sales in millions", ">", "100"},
		)
		fmt.Println(df)

		df = df.Filter(
			dataframe.F{5, "Genre", "==", "Fantasy"},
		)
		fmt.Println(df)*/

	//>>EXPANSION<<
	df3 := df.Concat(df2)
	//After Concat
	fmt.Println(df3)

	fmt.Println(df)
	fmt.Println(df2)

	fmt.Println("\n\n\n\t Inner Join Happens Here!!!")
	df4 := df.InnerJoin(df3, "Genre")
	fmt.Println(df4)

}
