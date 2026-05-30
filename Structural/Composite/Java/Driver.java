
public class Driver {

    public static void main(String[] args) {
        File file1 = new File("resume.pdf", 120);
        File file2 = new File("project.zip", 120);
        File file3 = new File("notes.txt", 120);

        Folder documentFolder = new Folder("Documents");
        documentFolder.add(file1);
        documentFolder.add(file2);
        documentFolder.add(file3);

        documentFolder.display();
    }
}