package main

func main() {

	file1 := NewFile("resume.pdf", 120)
	file2 := NewFile("tasks.txt", 20)
	file3 := NewFile("project.zip", 820)

	documentFolder := NewFolder("Documents")
	documentFolder.Add(file1)
	documentFolder.Add(file2)
	documentFolder.Add(file3)

	documentFolder.Display()
}
