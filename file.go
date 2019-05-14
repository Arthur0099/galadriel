package pgc

import (
	"io/ioutil"
)

// WriteToFile write the data to file. if the file doesn't exist,
// it will create a new file. if the file exists, it will cover
// the data.
func WriteToFile(data []byte, path string) {
	if err := ioutil.WriteFile(path, data, 0644); err != nil {
		panic(err)
	}
}
