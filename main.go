package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"
)

type ContributorCount struct {
	Contributor string
	Count       int
}

type ByCount []ContributorCount

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

type Pair struct {
	Contributor1 string
	Contributor2 string
	Count        int
}

type ByPairCount []Pair

func (a ByPairCount) Len() int           { return len(a) }
func (a ByPairCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPairCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

func main() {
	repoPath, _ := os.Getwd()
	log.Default().Printf("Repository path: %s", repoPath)

	cmd := exec.Command("git", "ls-files")
	cmd.Dir = repoPath

	output, err := cmd.Output()
	if err != nil {
		log.Fatal("Something went wrong! Try running the command inside a git repository and make sure you have git installed on your machine.")
		return
	}

	files := strings.Fields(string(output))

	contributorCounts := make(map[string]map[string]int)
	pairCounts := make(map[string]int)

	for _, file := range files {
		cmd := exec.Command("git", "log", "--format=%aN", file)
		cmd.Dir = repoPath
		output, err := cmd.Output()
		if err != nil {
			log.Fatal("Something went wrong! It seems like there are no commits on the current branch.")
			return
		}

		contributors := strings.Split(strings.TrimSpace(string(output)), "\n")

		fileContributorCounts := make(map[string]int)

		for _, contributor := range contributors {
			contributor = strings.TrimSpace(contributor)
			if contributor != "" {
				fileContributorCounts[contributor]++
			}
		}

		contributorCounts[file] = fileContributorCounts

		var contributorCountsSlice []ContributorCount
		for contributor, count := range fileContributorCounts {
			contributorCountsSlice = append(contributorCountsSlice, ContributorCount{Contributor: contributor, Count: count})
		}

		sort.Sort(ByCount(contributorCountsSlice))

		if len(contributorCountsSlice) >= 2 {
			pair := sortPair(contributorCountsSlice[0].Contributor, contributorCountsSlice[1].Contributor)
			pairCounts[pair]++
		}
	}

	var pairs []Pair
	for pair, count := range pairCounts {
		pairSplit := strings.Split(pair, "_")
		if len(pairSplit) >= 2 {
			pairs = append(pairs, Pair{
				Contributor1: pairSplit[0],
				Contributor2: pairSplit[1],
				Count:        count,
			})
		}
	}

	sort.Sort(ByPairCount(pairs))

	fmt.Println("Top contributor pairs frequently pushing to the same file:")
	for _, pair := range pairs {
		fmt.Printf("%s and %s: %d\n", pair.Contributor1, pair.Contributor2, pair.Count)
	}
}

func sortPair(a, b string) string {
	if a < b {
		return fmt.Sprintf("%s_%s", a, b)
	}
	return fmt.Sprintf("%s_%s", b, a)
}
