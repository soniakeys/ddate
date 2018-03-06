/* Command ddate

Usage:

ddate [+format] [day month year]

Options:

+format specifies a custom format, described below.  If not specified,
default format is "Pungenday, Discord 5, 3131 YOLD".

day month year, if given must be integer day, month and year of the Gregorian
calendar.  If not given, the current date is formatted.

Custom format:

Custom formats follow the prototype scheme of the time package of the Go
standard library.  Available prototype elements are

    Long season:   Discord
    Short season:  Dsc
    Long day:      Pungenday
    Short day:     PD
    Ordinal day:   5
    Cardinal day:  5th
    Holy day:      Mojoday
    Year:          3131

So for example,

    $ ddate 2 1 2006
    Boomtime, Chaos 2, 3172 YOLD

    $ ddate '+Pungenday, Discord 5, 3131 YOLD' 2 1 2006
    Boomtime, Chaos 2, 3172 YOLD

    $ ddate '+5 Dsc 3131' 2 1 2006
    2 Chs 3172

Kinda different:

    $ ddate 29 2 2008
    St. Tib's Day, 3174 YOLD
*/
package main
