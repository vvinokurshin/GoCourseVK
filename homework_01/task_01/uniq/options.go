package uniq

import "flag"

type Options struct {
	count       bool
	duplicated  bool
	unique      bool
	field       int
	chars       int
	insensitive bool
}

func Init(options *Options) {
	flag.BoolVar(&options.count, "c", false, "count the number of occurrences of a string in the input data")
	flag.BoolVar(&options.duplicated, "d", false, "output only those lines that are repeated in the input data")
	flag.BoolVar(&options.unique, "u", false, "output only those lines that are not repeated in the input data")
	flag.IntVar(&options.field, "f", 0, "ignore the first num_fields of fields in a row")
	flag.IntVar(&options.chars, "s", 0, "ignore the first num_chars characters in the string")
	flag.BoolVar(&options.insensitive, "i", false, "ignore case of letters")

}
