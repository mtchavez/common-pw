package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mtchavez/cuckoo"
)

type PasswordFilters struct {
	top196  *cuckoo.Filter
	top3575 *cuckoo.Filter
	top95k  *cuckoo.Filter
}

var filters *PasswordFilters

func createPasswordFilters() *PasswordFilters {
	if filters == nil {
		filters = &PasswordFilters{
			top196:  top196Filter(),
			top3575: top3575Filter(),
			top95k:  top95kFilter(),
		}
	}
	return filters
}

func top196Filter() *cuckoo.Filter {
	filter := cuckoo.New(
		cuckoo.BucketEntries(1<<10),
		cuckoo.BucketTotal(10),
		cuckoo.FingerprintLength(2),
	)
	fd, err := os.Open("./data/Top196-probable.txt")
	defer fd.Close()
	if err != nil {
		fmt.Errorf(err.Error())
		return filter
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := []byte(scanner.Text())
		filter.Insert(word)
	}
	return filter
}

func top3575Filter() *cuckoo.Filter {
	filter := cuckoo.New()
	fd, err := os.Open("./data/Top3575-probable.txt")
	defer fd.Close()
	if err != nil {
		fmt.Errorf(err.Error())
		return filter
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := []byte(scanner.Text())
		filter.Insert(word)
	}
	return filter
}

func top95kFilter() *cuckoo.Filter {
	filter := cuckoo.New()
	fd, err := os.Open("./data/Top95Thousand-probable.txt")
	defer fd.Close()
	if err != nil {
		fmt.Errorf(err.Error())
		return filter
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := []byte(scanner.Text())
		filter.Insert(word)
	}
	return filter
}
