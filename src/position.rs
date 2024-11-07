pub struct Position {
    index: usize,
    row: usize,
    col: usize,
}

impl Position {
    pub fn new() -> Position {
        Position {
            index: 0,
            row: 0,
            col: 0,
        }
    }

    pub fn coords(&self) -> (usize, usize) {
        (self.row, self.col)
    }

    pub fn index(&self) -> usize {
        self.index
    }

    pub fn col_advance(&mut self) {
        self.index += 1;
        self.col += 1;
    }

    pub fn row_advance(&mut self) {
        self.index += 1;
        self.row += 1;
        self.col = 0;
    }
}

impl Clone for Position {
    fn clone(&self) -> Position {
        return Position {
            index: self.index,
            row: self.row,
            col: self.col,
        };
    }
}

impl Copy for Position {}
