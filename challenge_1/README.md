# Challenge 1 - Onion wars
You are celebrating your 20 months anniversary at Tuenti today! As part of the celebration, you have decided to bring some good _tortilla española_ for your coworkers, but you realize that each person only eats _tortillas_ with or without onion.

![](https://contest.tuenti.net/resources/img/tortillas.jpg)

Given that each person eats half _tortilla_, what is the minimum number of _tortillas_ you need to ensure that everyone gets their share? Keep in mind that you can only bring whole _tortillas_, with or without onion.

Note: although it would be correct to bring only _tortilla_ with onion since a _tortilla_ without onion is horrible, we don't want to start an internal war over this, so we want you to bring enough food for both groups.

#### Input
The first line has an integer C, which is the number of cases for the problem. Then C lines follow and each one has two integers N and M, which are the number of people that want _tortillas_ with and without onion, respectively.

#### Output
For each case, there should be a line starting with "Case #x: " followed by the minimum number of _tortilla_ you must bring.

#### Limits
> 1 ≤ C ≤ 100 \
> 0 ≤ N, M ≤ 100000 \
> 1 ≤ N + M

#### Sample Input
> 3 \
> 2 2 \
> 3 2 \
> 0 1

#### Sample Output
> Case #1: 2 \
> Case #2: 3 \
> Case #3: 1

In the first case, we need one _tortilla_ per group.
In the second case, we need two _tortillas_ (since it is not possible to only bring one and a half) with onions and one without onions.
In the third case, we only need half of a _tortilla_ with onion, but we bring one since they are indivisible.