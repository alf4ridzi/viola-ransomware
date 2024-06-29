// for read file 

package readwrite

import (
	"io/ioutil"
)

func ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err 
	}

	return string(data[:]), err

}