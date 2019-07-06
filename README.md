# DROP SEVEN

Trying to recreate a simple phone game as a CLI game using go.

## run
```
go run .
```

## exit game
Hit `esc` to quit game

## how to play
This game is based on the Drop7 app available on android and iphone devices.

The map is a 7 by 7 grid.  The objective of the game is to not allow all 7 rows to be filled up in any column.  The game starts with one row filled with `*`.  A `*` is a cell with an unknown value until it's "shattered".  The player controls a random number at the top of the grid and can choose which row to drop the number.  If the rows or columns adjacent to the dropped number equal the number, the number explodes and shatters its adjacent cells.  A `*` cell needs to be shattered twice before revealing its value, which will then allow it to explode.  When a cell explodes, it disappears.  As time passes, more rows will appear at the bottom of the grid.