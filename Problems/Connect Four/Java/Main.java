public class Main {
    public static void main(String[] args) {
        Player player1 = new Player("pawan1", DiscColor.RED);
        Player player2 = new Player("pawan1", DiscColor.YELLOW);

        Game game = new Game(player1, player2);

        game.makeMove(0, player1);
        game.makeMove(0, player2);

        game.getWiner();

    }
}