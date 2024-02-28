# tournament_dzw

A simple program to request current german rating (DWZ) by german player id (PKZ) and to use [http://www.isewase.de/dwz/](http://www.isewase.de/dwz/) to obtain an estimation of the change of rating.

## Build

```
git clone https://github.com/niclashofmann/tournament_dwz.git
cd tournament_dwz
go build .
```

## Usage

```
./tournament_dwz TOURNAMENT.CSV DWZ YEAR_OF_BIRTH
```

The `TOURNAMENT.CSV` file has to be comma separated and has to contain the pkz of the opponent in the first and the result (1 for win, 0.5 for draw, 0 for loss) in the second column.

`DWZ` may contain the index (number of previously evaluated tournaments).
