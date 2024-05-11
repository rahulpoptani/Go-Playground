package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount,
		text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// j.entries = append(j.entries[:index], j.entries[index:]...)
}

// Breaking Single Resposibility Priciple (Save, Load, LoadFroWeb), this should be handle by Persistance (added below)

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {

}

func (j *Journal) LoadFromWeb(url *url.URL) {

}

// Separate function for saving file

var lineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, lineSeparator)), 0644)
}

// Separating Read and Writes Operation via Persistance.

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) saveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func main() {
	// the whole idea is to have a single responsibility (and separation of concerns). Like Journal have resposibility related to adding and removing
	// Persistence on the other hand have responsibility for storing or loading the journal
	j := Journal{}
	j.AddEntry("I ate ")
	j.AddEntry("I slept")
	fmt.Println(j.String())

	// separate function
	SaveToFile(&j, "journal.txt")

	// or using persistence
	p := Persistence{"\n"}
	p.saveToFile(&j, "./journal.txt")
}
