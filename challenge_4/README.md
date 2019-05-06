# Challenge 4 - Candy patterns
We had a massive event recently and decided to give free candy to everyone who attended. We were optimistic and assumed that bringing one piece of candy for every attendee would be enough. But there were some people who got more than one piece of candy. What's more, every attendee wanted at least one piece of candy! Consequently, we didn’t have enough candy to go around.

We suffered a major defeat when we ran out of candy that day, so we want to be better prepared for the next time. The employee in charge of the candies had written down the total number of candies each person got, for each candy. For example, if an attendee got 4 candies, the employee wrote down the number 4 four times. Unfortunately, for compression purposes, the list of numbers got shuffled and cut before you got it. Thankfully, you know that if you repeat your list X times, each number from your list will appear exactly the same number of times as in the original list.

Given the numbers on your list, we want you to find out the exact average number of candies each attendee wanted.

#### Input
The first line has an integer **C**, which is the number of cases for the problem. After that, two lines follow for each case. The first has an integer **N**, indicating the number of entries on the list. The second line has **N** integers **M0 .. M<sub>N-1</sub>** (separated by whie space), that represent the numbers of the list.

#### Output
For each case, there should be a line starting with "Case #x: " followed by the average number of pieces of candy per person as an irreducible fraction (**numerator/denominator**).

#### Limits
> 1 ≤ C ≤ 100 \
> 1 ≤ N ≤ 100 \
> 1 ≤ M ≤ 30

#### Examples
##### Example 1:

> 1 3 3 3

In this example, we can suppose that the list repeats only once (X=1) and that two attendees got candies (one got one candy, while the other got three). That would mean that we have four candies for two attendees, so the average would be 2/1. If we repeat the list any number of times (X>1), the answer would still be the same.

##### Example 2:

> 3 1 3 3

The list is the same as in the previous example but in different order, so the answer would also be 2/1.

#### Example 3:

> 1 2 4

We know that X>=4 since the number 4 needs to appear at least 4 times. If we suppose X=4, we would have 12 candies for 7 attendees (one attendee got 4 candies, two got 2 and four got 1 candy each). So the answer is 12/7 (this also applies for each valid value of X, like 8, 12, etc.).

#### Sample Input
> 3 \
> 4 \
> 1 3 3 3 \
> 4 \
> 3 1 3 3 \
> 3 \
> 1 2 4

#### Sample Output
> Case #1: 2/1 \
> Case #2: 2/1 \
> Case #3: 12/7