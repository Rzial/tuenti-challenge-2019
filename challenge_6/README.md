# Challenge 6 - Alphabet from outer space

![](https://contest.tuenti.net/resources/img/alien_decoder.png)

Some strange documents have been found in a Salt mine in the area of Turda, Romania. The documents are full of words that seems to be written in some kind of characters. Scientists think they're part of a dictionary so the words are ordered alphabetically. We want to infer the alphabetic ordering of the characters.

To simplify things, we'll use traditional characters to represent the alien characters.

For example, let’s suppose we’ve found an alphabet with four characters - a,b,c,d - and the following list of ordered \`words\`:

dad < bad < cab < cda

We know that the order of the characters is **a < d < b < c** , because:

* from ‘cab < cda’ we know that ‘a < d’
* from ‘dad < bad’ we know that ‘d < b’
* from ‘bad < cab’ we know that ‘b < c’

#### Input
The first line is an int **N**. **N** test cases follow. Each scenario is represented by a line with an integer **M**. A ordered list of **M** strings follows with one per line. Input is consistent so you won’t be able to infer a < b and b < a in the same test case.

#### Output
**N** lines with the ordered list of characters separated by whitespace for each test case if there is a unique possible order. Otherwise, the string “AMBIGUOUS” if there is more than one possible order.

#### Limits
> 1 < N <= 100 \
> 1 < M <= 1000 \
> 1 ≤ len(strings) <= 80

#### Sample Input
> 2 \
> 4 \
> dad \
> bad \
> cab \
> cda \
> 2 \
> ab \
> bc

#### Sample Output
> Case #1: a d b c \
> Case #2: AMBIGUOUS 