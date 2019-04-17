package main

import "fmt"

/*
The Revised Julian Calendar is a calendar system very similar to the familiar
Gregorian Calendar, but slightly more accurate in terms of average year length.
The Revised Julian Calendar has a leap day on Feb 29th of leap years as follows:
	1. Years that are evenly divisible by 4 are leap years.
	2. Exception: Years that are evenly divisible by 100 are not leap years.
	3. Exception to the exception: Years for which the remainder when divided by 900
	is either 200 or 600 are leap years.

	For instance, 2000 is an exception to the exception: the remainder when dividing
	2000 by 900 is 200. So 2000 is a leap year in the Revised Julian Calendar.
	Challenge

Given two positive year numbers (with the second one greater than or equal to the first),
find out how many leap days (Feb 29ths) appear between Jan 1 of the first year, and Jan 1
of the second year in the Revised Julian Calendar. This is equivalent to asking how many
leap years there are in the interval between the two years, including the first but excluding
the second.

leaps(2016, 2017) => 1
leaps(2019, 2020) => 0
leaps(1900, 1901) => 0
leaps(2000, 2001) => 1
leaps(2800, 2801) => 0
leaps(123456, 123456) => 0
leaps(1234, 5678) => 1077
leaps(123456, 7891011) => 1881475
For this challenge, you must handle very large years efficiently, much faster than checking each year in the range.
leaps(123456789101112, 1314151617181920) => 288412747246240

CHALLENGE: NOT DONE(not optimized)
*/
func main() {

	leaps := func(start, end int) int {
		diff := end - start
		count := 0
		for i := 0; i < diff; i++ {
			year := start + i
			div := year / 4
			if div*4 == year { // potential leap year
				hunDiv := year / 100
				nineDiv := year / 900
				if hunDiv*100 != year { // leap year; exception #1
					count++
				} else if year-(nineDiv*900) == 200 || year-(nineDiv*900) == 600 { // leap year; exception #2
					count++
				}
			}
		}
		return count
	}
	fmt.Println(leaps(123456, 7891011))

}
