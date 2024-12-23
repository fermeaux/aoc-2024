printf -v day "%02d" $1
mkdir -p ./inputs/
mkdir -p ./inputs/aoc-$day
curl --cookie "session=$AOC_SESSION_ID" https://adventofcode.com/2024/day/$1/input --output ./inputs/aoc-$day/data.txt