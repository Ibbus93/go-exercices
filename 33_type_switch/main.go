package main

import "fmt"

type cloud interface {
	launch() string
}

type aws struct {
	computeSvcName string
}

type azure struct {
	computeSvcName string
}

func (cloud aws) launch() string {
	return fmt.Sprintf("Launching instance %s...", cloud.computeSvcName)
}

func (cloud azure) launch() string {
	return fmt.Sprintf("Launching virutal machine %s", cloud.computeSvcName)
}

func compute (cloud interface{}) {
	switch cloudPlatform := cloud.(type) {
	case aws:
		aws := cloud.(aws) // casting
		fmt.Printf("AWS: %s -> %s\n", aws, cloudPlatform.launch()) // polymorphism
	case azure: 
		azure := cloud.(azure)
		fmt.Printf("Azure: %s -> %s\n", azure, cloudPlatform.launch())
	}
}

func main() {
	var clouds []cloud = []cloud{
		aws{"ec2"}, azure{"vm"},
	}

	for _, cloud := range clouds {
		compute(cloud)
	}
}