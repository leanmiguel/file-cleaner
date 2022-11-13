package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices       []string
	cursor        int
	currentChoice string
	textInput     textinput.Model
	CleanStageModel
}

func (m model) isBaseStage() bool {
	return m.chosenFolder == ""
}

func (m *model) handleEnter() {
	if !m.isBaseStage() {
		return
	}
	if m.cursor == Desktop {
		m.initCleanStage("/desktop")
	}
	if m.cursor == Downloads {
		m.initCleanStage("/downloads")
	}
	m.cursor = 0
}

func (m *model) handleDownKey() {
	if !m.isBaseStage() {
		return
	}

	if m.cursor < len(m.choices)-1 {
		m.cursor++
	}
}
func (m *model) handleUpKey() {
	if !m.isBaseStage() {
		return
	}

	if m.cursor > 0 {
		m.cursor--
	}
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "New Folder"

	ti.CharLimit = 156
	ti.Width = 20
	return model{
		choices:   []string{"Desktop", "Downloads"},
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func getAvailableFolders(directory string) []fs.FileInfo {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(homedir + directory)
	if err != nil {
		log.Fatal(err)
	}

	var avlFolders []fs.FileInfo
	for _, v := range files {
		if v.IsDir() {
			avlFolders = append(avlFolders, v)
		}
	}
	return avlFolders

}
func getFilesFromDirectory(directory string) []fs.FileInfo {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(homedir + directory)
	if err != nil {

		log.Fatal(err)
	}

	var removedHiddenFiles []fs.FileInfo
	for _, v := range files {

		if !strings.HasPrefix(v.Name(), ".") && !v.IsDir() {
			removedHiddenFiles = append(removedHiddenFiles, v)
		}
	}
	return removedHiddenFiles
}
