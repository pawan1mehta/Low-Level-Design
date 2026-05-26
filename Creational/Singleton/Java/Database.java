

public class Database {

    private static Database database;
    public String connection;

    private Database(String connection) {
        this.connection = connection;
    }

    public static Database createDatabase(String connection) {
        if (database == null) {
            database = new Database(connection);
        }
        return database;
    }
}