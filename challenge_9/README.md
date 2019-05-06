# Challenge 9 - Helping Nobita
Nobita is not good at maths. He's spent all weekend studying maths and he's finished all his homework problems and all the answers are correct! Unfortunately, Tsuneo and Giant have played a trick on Nobita. They stole the answer sheet and changed questions and answers.

Nobita needs help! Doraemon will pull the program you write out of his pocket to resolve this desperate situation!

![](https://contest.tuenti.net/resources/2019/img/helping_nobita.jpeg)

Helping Nobita
The problems are arithmetic operations: addition, subtraction and multiplication:

**A OPERATOR B = C** where **1 ≤ A,B,C ≤ 99999 and OPERATOR={+,-,\*}**

#### Examples

Original operations

\#1 2017 + 223 = 2240

\#2 250 * 65 = 16250

\#3 8102 - 1747 = 6355

Tsuneo and Giant:

They first converted the western numbers to Japanese using Chinese Kanji characters:

\#1 二千十七 + 二百二十三 = 二千二百四十

\#2 二百五十 * 六十五 = 一万六千二百五十

\#3 八千百二 - 千七百四十七 = 六千三百五十五

And then they swapped Kanji and hid the operators:

\#1 七十二千 OPERATOR 二三百十二 = 四二千百十二

\#2 五十百二 OPERATOR 五十六 = 千十万二百一五六

\#3 千八百二 OPERATOR 七百四千十七 = 千五百三六十五

**The program you have to write must get the original operations.**

#### Input
The first line will contain an integer **N**, which is the number of cases for the problem. Each case is an altered and hidden arithmetic operation.

#### Output
For each case, there should be a line starting with "Case #x: " followed by the correct original arithmetic operation. 
Each case has an unique solution.

#### Limits
* Western numbers: 1 ≤ A,B,C ≤ 99999   (western numbers)
* Kanji numbers: 一 ≤ A,B,C ≤ 九万九千九百九十九
* OPERATOR = [+,-,*]

#### Sample input
> 3 \
> 七十二千 OPERATOR 二三百十二 = 四二千百十二 \
> 五十百二 OPERATOR 五十六 = 千十万二百一五六 \
> 千八百二 OPERATOR 七百四千十七 = 千五百三六十五

#### Sample Output
> Case #1: 2017 + 223 = 2240 \
> Case #2: 250 * 65 = 16250 \
> Case #3: 8102 - 1747 = 6355

##### Some words about Japanese numerals
In order to be able to solve this challenge, you’d better to study a bit about Japanese numerals. It’s easy and funny!

[Wikipedia Japanese Numerals](https://en.wikipedia.org/wiki/Japanese_numerals) \
[Convert western numbers to Japanese](https://www.sljfaq.org/cgi/numbers.cgi)

The system of Japanese numerals is the system of number names used in the Japanese language. There are Japanese characters (Kanjis) representing basic numbers:


| **Kanji** | 一 | 二 | 三 | 四 | 五 | 六 | 七 | 八 | 九 | 十 | 百 | 千 | 万 |
|--------|----|----|----|----|----|----|----|----|----|----|-----|------|-------|
| **Number** | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 100 | 1000 | 10000 |

Intermediate numbers are made by combining these elements:

* Tens from 20 to 90 are "(digit)十" as in 二十 (20) to 九十 (90).
* Hundreds from 200 to 900 are "(digit)百".
* Thousands from 2000 to 9000 are "(digit)千".

**NOTE**: Starting at 万, numbers begin with 一 if no digit would otherwise precede. That is, 100 is just 百, and 1000 is just 千, but 10000 is 一万, not just 万.

Some examples:

> 十一 : 11  made as concatenation of 十 (10) and 一 (1) \
> 四十八 : 48 made as concatenation of 四 (4) and 十 (10) and 八 (8) \
> 百五十一 : 151 made as concatenation of 百 (100) and 五十 (50) and 一 (1) \
> 三百二 : 302 \
> 四百六十九 : 469 \
> 二千二十五 : 2025 \
> 千百十 : 1110 \
> 二千三百四十五 : 2345 = 2x1000(二千) + 300 (三百) + 40(四十) + 5 (五) \
> 一万百十四 : 10114   Notice that 10000 is 一万, not just 万 \
五万四千七百三十二 : 54732 = 5x10000 + 4x1000 + 7x100 + 3x10 + 2