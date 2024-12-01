# Advent of Code 2024
## Description
Another year another try at keeping up with the Advent of Code. Once again in the beautiful Go language.  

## Days
|Day|Name|Stars|
|---|----|-----|
|01|Historian Hysteria|⭐⭐|


##  Running the code
At the time of AoC 2024 I am working with Go version 1.23.3. I don't believe I'm using any recent additions to the language so using older versions should be fine.

I wrote a little program to run the code for each day, it can be run by using
```bash
go run main.go -d <day> -p <puzzle>
```

If that's not explanatory enough with `-d` you can specify the day you want to run and with `-p` the specific puzzle from that day. Running it with just the day shows the solutions to all puzzles for the given day. Not specifying either day or puzzle shows an overview of the days and puzzles.

In addition to the final calculated solution the code of most puzzles will also print the one or two last sources for that answer. For example an array or slice with the values to be added up together in the next step. This is mostly for fun.

## License
The inputs and any puzzle material that may have made it into this repository was created and is owned by Advent of Code. Any code written by me for solving the puzzles is to be considered public domain.