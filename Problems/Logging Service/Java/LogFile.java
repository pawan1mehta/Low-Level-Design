import java.io.BufferedOutputStream;
import java.io.IOException;
import java.io.OutputStream;
import java.io.UncheckedIOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardOpenOption;

public class LogFile {

    private final Path path;
    private final OutputStream out;

    public LogFile(String pathStr) {
        try {
            this.path = Paths.get(pathStr);
            if(path.getParent() != null) {
                Files.createDirectories(path.getParent());
            }
            this.out = new BufferedOutputStream(Files.newOutputStream(path, StandardOpenOption.CREATE, StandardOpenOption.APPEND));
        } catch (IOException exception) {
            throw new UncheckedIOException(exception);
        }
    }

    public synchronized void write(byte[] bytes) {
        try {
            out.write(bytes);
        } catch (IOException e) {
            throw new UncheckedIOException(e);
        }
    }

    public synchronized void flush() {
        try {
            out.flush();
        } catch (IOException e) {
            throw new UncheckedIOException(e);
        }
    }

    public synchronized void close() {
        try {
            out.close();
        } catch (IOException e) {
            throw new UncheckedIOException(e);
        }
    }
}
