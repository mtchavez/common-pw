package filters

import (
	"bufio"
	"log"
	"os"
	"time"

	"github.com/mtchavez/cuckoo"
)

// PasswordFilters for the various password filters
type PasswordFilters struct {
	Top196  *cuckoo.Filter
	Top3575 *cuckoo.Filter
	Top95k  *cuckoo.Filter
	Top32m  *cuckoo.Filter
}

// PWFilters is a package global for the created filters
var PWFilters *PasswordFilters

// BuildFilters will create all the password filters
func BuildFilters() {
	log.Println("Starting to build filters")
	startTime := time.Now()
	createPasswordFilters()
	endTime := time.Now()
	log.Printf("Done building filters. Took %v to build.\n", endTime.Sub(startTime))
}

func createPasswordFilters() *PasswordFilters {
	if PWFilters == nil {
		PWFilters = &PasswordFilters{
			Top196:  top196Filter(),
			Top3575: top3575Filter(),
			Top95k:  top95kFilter(),
			Top32m:  top32mFilter(),
		}
	}
	return PWFilters
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
		log.Println(err.Error())
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
		log.Println(err.Error())
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
		log.Println(err.Error())
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
		log.Println(err.Error())
		return filter
	}

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		word := []byte(scanner.Text())
		filter.Insert(word)
	}
	return filter
}
