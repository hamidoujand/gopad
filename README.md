# Gopad - A Minimalist Text Editor in Golang
Gopad is a lightweight, terminal-based text editor written in Golang. It provides basic editing features with a clean, simple interface. Ideal for quick edits and minimalistic use, it operates with intuitive key bindings for navigation and editing.

# Features
 - modes (VIEW/EDIT)
 - display buffer & status bar
 - inserting characters
 - deleting characters
 - inserting lines
 - deleting lines
 - navigation
 - copy/paste
 - undo/redo

# Key bindigns
       ESC: enter the 'VIEW' mode
         e: enter the 'EDIT' mode
         q: quit from the text editor
         w: write file to disk
         x: cut current line to copy buffer
         c: copy current line to copy buffer
         p: paste line from copy buffer
         s: push text buffer to undo buffer
         l: pull text buffer from undo buffer
    Arrows: move cursor
    PgDown: scroll 1/4 of the screen downwards
      PgUp: scroll 1/4 of the screen upwards
      HOME: move cursor to the begining of the current line
       END: move cursor to the end of the current line


# Usage
    $ gopad                 # opens editor with 'out.txt' source file name
    $ gopad my_file.txt     # opens editor with 'my_file.txt' if it exists,
                            # otherwise sets source filename to 'my_file.txt'

# Build from sources
```bash
make build