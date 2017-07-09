package main

import (
	"bufio"
	"os"

	"github.com/mtchavez/cuckoo"
	"github.com/prometheus/common/log"
)

// PasswordFilters for the various password filters
type PasswordFilters struct {
	top196  *cuckoo.Filter
	top3575 *cuckoo.Filter
	top95k  *cuckoo.Filter
	top32m  *cuckoo.Filter
}

var filters *PasswordFilters

func createPasswordFilters() *PasswordFilters {
	if filters == nil {
		filters = &PasswordFilters{
			top196:  top196Filter(),
			top3575: top3575Filter(),
			top95k:  top95kFilter(),
			top32m:  top32mFilter(),
		}
	}
	return filters
}

func top196Filter() *cuckoo.Filter {
	filter := cuckoo.New(
		cuckoo.BucketEntries(10),
		cuckoo.BucketTotal(1<<10),
		cuckoo.FingerprintLength(2),
	)
	fd, err := os.Open("./data/Top196-probable.txt")
	defer fd.Close()
	if err != nil {
		log.Errorf(err.Error())
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
		log.Errorf(err.Error())
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
		log.Errorf(err.Error())
		return filter
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := []byte(scanner.Text())
		filter.Insert(word)
	}
	return filter
}

func top32mFilter() *cuckoo.Filter {
	filter := cuckoo.New(
		cuckoo.BucketEntries(36),
		cuckoo.BucketTotal(1<<27),
		cuckoo.FingerprintLength(12),
	)
	fd, err := os.Open("./data/Top32Million-probable.txt")
	defer fd.Close()
	if err != nil {
		log.Errorf(err.Error())
		return filter
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := []byte(scanner.Text())
		filter.Insert(word)
	}
	return filter
}
