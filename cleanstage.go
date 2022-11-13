package main

import (
	"io/fs"
	"log"
	"os"
	"os/exec"
)

const (
	AddToFolder = iota
	NewFolder
	Trash
	InspectFile
)

type CleanStageModel struct {
	currentFilePointer int
	avlFolderCursor    int
	chosenFolder       string
	files              []fs.FileInfo
	avlFolders         []fs.FileInfo
	folderPath         string
}

func (m *model) initCleanStage(folder string) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	m.chosenFolder = folder
	m.currentFilePointer = 0
	m.avlFolderCursor = 0
	m.avlFolders = getFilesFromDirectory(folder)
	m.files = getAvailableFolders(folder)
	m.folderPath = homedir + folder
}

func (m model) isCleanStage() bool {
	return m.chosenFolder != ""
}

func (m *model) handleCleanStageDownKey() {
	if !m.isCleanStage() {
		return
	}

	if m.cursor == NewFolder {
		m.textInput.Blur()
	} else {
		m.textInput.Focus()
	}

	if m.cursor < len(cleanerOptions)-1 {
		m.cursor++
	}
}
func (m *model) handleCleanStageUpKey() {
	if !m.isCleanStage() {
		return
	}

	if m.cursor == NewFolder {
		m.textInput.Blur()
	} else {
		m.textInput.Focus()
	}

	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *model) handleCleanStageEnter() {
	if !m.isCleanStage() {
		return
	}

	currentFile := m.files[m.currentFilePointer]
	if m.cursor == 0 {
		newFolderPath := m.folderPath + "/" + m.avlFolders[m.avlFolderCursor].Name()

		m.avlFolders = getAvailableFolders(m.currentChoice)
		err := os.Rename(m.folderPath+"/"+currentFile.Name(), newFolderPath+"/"+currentFile.Name())
		if err != nil {
			log.Fatal("bad directory, print actual error")
		}

	}

	if m.cursor == 1 {
		newFolderPath := m.folderPath + "/" + m.textInput.Value()

		err := os.MkdirAll(newFolderPath, 0700)
		if err != nil {
			log.Fatal("bad directory, print actual error")
		}
		m.avlFolders = getAvailableFolders(m.currentChoice)
		err = os.Rename(m.folderPath+"/"+currentFile.Name(), newFolderPath+"/"+currentFile.Name())
		if err != nil {
			log.Fatal("bad directory, print actual error")
		}

	}

	if m.cursor == 2 {
		cmd := exec.Command("trash", m.folderPath+"/"+currentFile.Name())
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	if m.cursor == 3 {
		cmd := exec.Command("open", m.folderPath+"/"+currentFile.Name())
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}

	m.currentFilePointer += 1
	m.avlFolderCursor = 0
	m.textInput.Reset()
}
func (m *model) handleCleanStageLeftKey() {
	if !m.isCleanStage() {
		return
	}
	if m.cursor != AddToFolder {
		return
	}

	if m.avlFolderCursor > 0 {
		csm := m
		csm.avlFolderCursor -= 1
		m = csm
	}
}

func (m *model) handleCleanStageRightKey() {
	if !m.isCleanStage() {
		return
	}
	if m.cursor != AddToFolder {
		return
	}

	if m.avlFolderCursor < len(m.avlFolders)-1 {
		csm := m
		csm.avlFolderCursor += 1
		m = csm
	}
}
