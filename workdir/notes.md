Unique Locations: 413

example:
   Maputo;27.5
   Sapporo;-3.9
   Bissau;22.7
   Philadelphia;20.0
   Odesa;2.5
   Belgrade;12.4
   Gabès;27.3
   Vienna;6.2
   Brisbane;17.5
   Chicago;11.0

output: <city> = <min><mean><max>

| Solution Type | Threading | Time | Tag |
|---------------|-----------|------|-----|
| simple unoptimized | single threaded | 134.44 s | [1.0](https://github.com/absenth/1brc/releases/tag/1.0) |

## 100 Days of Code log:

#### Day 1: Golang Adventure Begins!

Today, I started tackling the [1 billion rows challenge](https://1brc.dev) in Golang. Initially, I used `os.ReadFile` to read the input file, but soon realized that it wasn't splitting each line correctly when trying to extract temperature values. This was because the next city's name was being appended to the temperature value due to not properly handling newlines.

To overcome this issue, I switched to using `os.Open` and `bufio.NewScanner` to read the file line by line. This approach allowed me to process each row individually and avoid any unwanted concatenations.

Next, I encountered a challenge when trying to convert the `temp` string into a float. After some consideration, I decided to use `strings.Replace` to simply drop the decimal place. While this means that later on I'll need to divide all values by 10, it significantly reduces the amount of floating-point math in the program overall. I believe this will ultimately be a big win for performance and simplicity.

#### Day 2: 1 Billion Row Challenge - Making Progress

Today, I picked up where I left off yesterday on the [1 billion rows
challenge](https://1brc.dev). I started out by creating a map to hold the data as the
program parses the measurements.txt file. The map uses the city name as its key and
uses a slice to hold the temperature data as values.

I then created a small loop to generate an alphabetical list of cities, which will be
necessary to meet the requirements of the challenge. While working on this, I noticed
that the record splitting mechanism I wrote yesterday doesn't actually work. The city
"St. Louis" shows up in the map as "St." and the city "Washington, DC" shows up as
"Washington,", so I'll need to revisit how I'm splitting records in the `for
scanner.Scan` loop.

I manually entered a city name into a `key` variable and an integer value into an
`avg` for testing purposes. This allowed me to print the Minimum and Maximum
temperatures for that city, so I think I'm getting close to having a "working"
solution.

#### Day 3: 1 Billion Row Challenge - First Completion

I fixed the parsing problem I found on Day 2 by changing`splitrecord := strings.Split(scanner.Text(), ";")` to `splitrecord := strings.SplitN(scanner.Text(), ";", 2)`.  However, on further reflection, I think the actual issue was the subsequent two lines where I was using `fmt.Sscanf`.  Instead, I set `city` and `stemp` values from `splitrecord[0,1]`.

Next, I tackled was getting the min/mean/max for each city in the datamap.  While my current implementation works, I'm confident it can be optimized to run faster.  This is an area I will focus on as I further optimize my code.

My initial goal was to have my solution process all 1 billion rows of data in under 15 minutes.  I'm thrilled to report that, on my first attempt, I achieved this in an avearge of *2 minutes and 14.4 seconds*!  This will serve as the baseline for future optimizations and improvement.

