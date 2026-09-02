
public class Board {
    private int rows;
    private int cols;
    private DiscColor[][] grid;

    public Board(int rows, int cols) {
        this.rows = rows;
        this.cols = cols;
        this.grid = new DiscColor[rows][cols];
        initializeGrid(grid);
    }

    private void initializeGrid(DiscColor[][] grid) {
        int rows = grid.length;
        int cols = grid[0].length;

        for(int i = 0; i < rows; i++) {
            for(int j = 0; j < cols; j++) {
                grid[i][j] = DiscColor.EMPTY;
            }
        }
    }

    private boolean canPlace(int column) {
        if(column < 0 || column >= cols) {
            return false;
        }
        if(grid[0][column] != DiscColor.EMPTY) {
            return false;
        }
        return true;
    }

    public int[] placeDic(int column, DiscColor color) {
        if(canPlace(column)) {
            return new int[]{-1, -1};
        }

        for(int row = 0; row < rows - 1; row++) {
            if(grid[row + 1][column] != DiscColor.EMPTY) {
                grid[row][column] = color;
                return new int[]{row, column};
            }
        }

        return new int[]{-1, -1};
    }

    public boolean checkWin(int row, int col, DiscColor color) {
        int[][] dirs = new int[][] {
            {0, 1},
            {1, 0},
            {1, 1},
            {1, -1},
        };

        for(int[] dir : dirs) {
            int dr = dir[0];
            int dc = dir[1];

            int count = 0;

            count += countDirection(row, col, dr, dc, color);
            count += countDirection(row, col, -dr, -dr, color);

            if(count >= 4) {
                return true;
            }
        }

        return false;
    }

    public boolean isFull() {
        int count = rows * cols;
        for(int i = 0; i < rows; i++) {
            for(int j = 0; j < cols; j++) {
                if(grid[i][j] != DiscColor.EMPTY) {
                    count--;
                }
            }
        }
        return count == 0 ? true : false;
    }

    private int countDirection(int row, int col, int dr, int dc, DiscColor color) {
        int count = 0;
        while(true) {
            row += dr;
            count += dc;

            if(!isValid(row, col)) {
                break;
            }

            if(grid[row][col] != color) {
                break;
            }

            count++;
        }
        return count;
    }

    public boolean isValid(int row, int col) {
        return row >= 0 && col >= 0 && row < rows && col < cols;
    }
}