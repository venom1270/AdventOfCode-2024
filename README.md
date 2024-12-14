# AdventOfCode-2024

## In short...

This year it's GO time 😎 ...

... ~~and VIM time~~ for the VIM 😎 ... *courtesy of @nsk4*

... and stream time 😎 (although we will see how that goes)!


That's about it. Simple, right?

*Also a side note: this chapter was written in VIM!*

Oh, yeah, trying to use a split keyboard (Sofle) with EN key layout! 😎

Phew, now that really is everything! 😅

## Day 1

It was actually really easy, but had some problems with keyboard (physical and EN layout). Otherwise fine, but streaming this is a bit of a challenge.

## Day 2

That was a rollercoaster. Part 1 was easy as hell, but I wanted to do a fancy solution for part 2... didn't work out 😅. Spent *wayyyy* to much time here and in the end opted for the straightforward "bruteforce" solution. Not my best day.

## Day 3

THis day was straightforward and easy. I was a bit scared part 2 was gonna be more complex (nesting) and that I would have to rewrite my loop in recursion form, but that was not neccessary thankfully! 😎

## Day 4

Pretty easy, got it (almost) first try (stupid GO slices!!). Part 2 easier than part 1, or at least less verbose 😄. Onwards!!! 😎

## Day 5

Parsing was a bit more involved (nothing like last year though, thankfully). Part 1 was easy, part 2 aswell after some thought. This year is going pretty well so far! 😎 *Famous last words...* 😅

## Day 6

Pretty easy overall. I liked it a lot. Took a bit more time than previous days. Part 2 took a while to compute (5+ seconds), but the answer was correct. Both part 1 and part 2 worked first try (not counting syntax error and typos). Maybe I could optimeze part 2 a bit by starting the test just 1 move before encountering the new (placed) obstacle, but I would need to save the Position struct when doing Part 1.

## Day 7

That was easier than expected! Dusting off recursion after a while, feels good! An additional challenge for me was talking (IRL, not on stream) while doing this day. But I managed to complete it in ~30min. 😎🥳

## Day 8

Ramping up... this took almost an hour (actually more if we count input parsing). Had some trouble with GOs references (some array elements were changed because some reference pointed to them), but otherwise the solution was relatively straightforward. Used a map to gather all connected antennas, and then just calulated the distance and extrapoled until valid for part 2. Yeah, we are slowly ramping up, exited for the next days... 😎

## Day 9

Disk fragmentation-like task. Pretty nice, enjoyed it. Had some bugs which took me some time to solve (especiall part 2), but otherwise not that difficult algorithm-wise. Still going strong!! 😎😎

## Day 10

Finally some recursion!! That was quite easy I think I set a personal record 😄. I completed both parts even faster than Day 1 😅😅. Accidentaly did part 2 first and then was bamboozled why part 1 didn't work - note to myself: LEARN TO READ!! 🚨😅🚨

## Day 11

First hard one - managed to do only day 11. Different methods for day 2 that I tried are too slow 🤔. I think my method of DFS and memoization should work but I'm clearly missing something. A mystery for another day! 😄

EDIT 12.12.2024: I did it! Day 12 gave me the motivation to try solving part 2 - but I can't believe what a stupid mistake I made. The challenge requires only a simple memoization. Yet I thought it was not enough because I forgot to change recursion call from naive() to memoization() on one of the branches!!!!!! 😅😅 A VERY EASY day if you're not a complete dum-dum (which I was yesterday apparently - I blame my late coming home and expending all my brain power on board games that evening 😄😅). And I kept trying to discover hot water with cycles, some complex memoization, at the end even a trace which would update the rolling sum (kind of dynamic programming basically). 🤪 Anyway, glad that I made it! Got all 24 stars so far. Half of AoC done! 😎🥳

## Day 12

I feel so proud for this one!! Managed to figure out part 2 on my own without any noteworthy prior knowledge on how to approach those kinds of problems 🥳😄😎 I really liked todays challenge - totally unbiased statement! 😄 Flood fill (I guess?) with some form of edge detection (part 1 simple, part 2 quite a lot more complex).

## Day 13

Part 1 easy, I just complicated it because the memoization condition was wrong (missing =). For part 2 I am completely out of ideas though. Tried some thing (division, A*, etc.) but none seem to be appropriate. 🤔 Will have to give this one some more thought. I blame it on my low quality sleep as I watched The Game Awards last night and went to sleep some time after 5 AM (and woke up at 8:30AM) 😎😅 So far I have had an excuse for every situation 😉 (coincidence 🤞)

EDIT: After briefly looking on the AoC subreddit, I noticed a helpful keyword: *linear equasions*. Of course! In hindsight it's obvious, this problem is nothing more than a series of linear equasions (2 equasions, 2 variables). How to solve them programatically? I went digging into my brain and on the web and the obvious answer presents itself - *Gauss elimination*, of course! But it looked a bit tiresome to write in code, so I went digging some more - there's another method that I've long forgotten even exists (long time since uni 😅) - *Cramer's Law*. It just calculates 3 easy matrix determinants and you get the answer. The only trick is that you have to check if the solutions actually solves the equasions (sometimes it doesn't - found this by trial&error). So yeah, used online (non-AI) resources to implement actual __MATH__, how cool is that? 😎😎 Though I have to say this one bamboozled me hard, with going deep into DFS, A* and other useless methods. 🤪

*Side note: even though this commit was made after midnight, I still count this under the same day since the answer was submitted the same day*

## Day 14

Part 1 suspiciously easy, while for part 2 the instructions are not really clear... what constitutes a "christmas tree" exactly? This one came a bit out of the blue 🤔

EDIT: After returning from the barber, I solved it in 5 mins. 😎 All I had to do is make sure the height is more than 4 - I wasn't sure if smaller trees count so that condition was just a wild guess on my part, but it provided the right answer. 😎😎 While this part 2 twist was pretty cool, the instructions do not give any direction on how the tree should look like at all (min size, trunk size, do fields around it have to be empty etc.) so it was a rather poor (but, again, cool) challenge. 😉