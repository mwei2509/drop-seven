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

The map is a 7 by 7 grid.  

## notes
7x7 grid

0️⃣⃣0️⃣1️⃣1️⃣1️⃣1️⃣
0️⃣⃣0️⃣1️⃣1️⃣1️⃣1️⃣
0️⃣⃣0️⃣1️⃣1️⃣1️⃣1️⃣
0️⃣⃣0️⃣1️⃣1️⃣1️⃣1️⃣
0️⃣⃣0️⃣1️⃣1️⃣1️⃣1️⃣1️⃣
0️⃣⃣0️⃣1️⃣1️⃣1️⃣⏺⏺
0️⃣⃣0️⃣1️⃣1️⃣1️⃣1️⃣1️⃣

2 <- will "explode" bc of 2 vertical
******* <- unknown status

when a number "explodes", the unknown cells adjacent lose a life
each cell has 2 lives