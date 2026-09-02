
public class Game {

    private Player player1;
    private Player player2;
    private Player currentPlayer;
    private Board board;
    private GameState state;
    private Player winer;

    public Game(Player player1, Player player2) {
        this.player1 = player1;
        this.player2 = player2;
        board = new Board(6, 7);
        this.state = null;
        this.winer = null;
    }

    public boolean makeMove(int column, Player player) {
        if (this.state != GameState.IN_PROGRESS) {
            return false;
        }

        if(this.currentPlayer != player) {
            return false;
        }

        int[] position = this.board.placeDic(column, player.getColor());
        if(position[0] == -1) {
            return false;
        }

        if(board.checkWin(position[0], position[1], player.getColor())) {
            this.state = GameState.WON;
            this.winer = player;
        } else if(board.isFull()) {
            this.state = GameState.DRAW;
        } else {
            this.currentPlayer = player == player1 ? player2 : player1;
        }

        return true;
    }

    public GameState getGameState() {
        return this.state;
    }

    public Player getCurrentPlayer() {
        return this.currentPlayer;
    }

    public Player getWiner() {
        return this.winer;
    }
}