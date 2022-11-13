package main

import "fmt"

func (m model) View() string {

	if m.isCleanStage() {
		s := "How do you want to deal with "
		s += m.files[m.currentFilePointer].Name()
		s += "?\n"

		for i := range cleanerOptions {

			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}
			if i == 0 {
				s += fmt.Sprintf(
					"%s Add to existing folder ",
					cursor,
				)
				if m.cursor == 0 {
					// s += fmt.Sprintf("m.csm.avlFolderCursor %d", m.csm.avlFolderCursor)
					for i, v := range m.avlFolders {
						if i == m.avlFolderCursor {
							s += "> "
						}
						s += v.Name()
						s += " "
					}

				}
				s += "\n"
			}
			if i == 1 {

				if m.cursor == 1 {
					s += fmt.Sprintf(
						"%s Add to new folder: %s\n",
						cursor, m.textInput.View(),
					)
				} else {
					s += fmt.Sprintf(
						"%s Add to new folder\n",
						cursor)
				}
			}

			if i == 2 {
				s += fmt.Sprintf("%s Delete the file\n", cursor)
			}
			if i == 3 {
				s += fmt.Sprintf("%s Inspect file\n", cursor)
			}

		}

		return s
	}
	if m.isBaseStage() {
		s := "Which folder do you want to clean?\n\n"

		// Iterate over our choices
		for i, choice := range m.choices {

			// Is the cursor pointing at this choice?
			cursor := " " // no cursor
			if m.cursor == i {
				cursor = ">" // cursor!
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}

		// The footer
		s += "\nPress ctrl+c to quit.\n"
		// s += fmt.Sprintf("the current chosen is %s\n", m.chosenFolder)
		// s += fmt.Sprintf("the current cursor is %d", m.cursor)
		return s
	}

	return ""
}
