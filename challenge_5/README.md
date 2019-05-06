# Challenge 5 - Forbidden Love

![](https://contest.tuenti.net/resources/img/25-letter.png)

This is the last non-encoded message Gordon Bowsher sent to his secret lover, WWII soldier Gilbert Bradley. Fearing of being discovered as a couple in a very intolerant 1939, Gordon decides to begin encrypting their messages.

We have in our possession multiple messages they interchanged during their years together but as you may know, they are encrypted.

Now film director Andrew Vallentine wants to take this amazing story to the big screen but he's asking us to decrypt all the messages so he knows exactly how this story unfolds... And honestly, we also want to know all the details.

Can you help us building an algorithm to decrypt these messages?

Our investigators have been working very hard to gather a lot of useful information for you to come up with a solution.

Our investigators' findings:

* The typewriter is an Underwood model-20 which has the following layout:

![](https://contest.tuenti.net/resources/img/25-layout.png)

* They always finished their messages with a single letter, G (Gordon) or B (Bradley).
* They followed a shifting method to encrypt their messages. This means that they displaced the keys by the same offset on each message; for example if they replaced the T for a U, then the G is replaced by a J
* The displacement can also be for letters above or below; or even mixed (two rows above plus three letters to the right)
* This displacement changes between messages

#### Input
The input begins with an integer **N** representing the number of messages to be decrypted.
Each case will be composed of two lines, the first line will have a character ‘**B**’ or ‘**G**’ representing who’s sending the message, **Bradley** or **Gordon**.
The next line is the message to be decrypted.

#### Output
For each case, the output should be one line containing the decrypted message.

#### Sample input
> 2 \
> G \
> P .PFF IQOZ J \
> B \
> J 6JZZ GKH FKK8 4

#### Sample output
> Case #1: I MISS YOU. G \
> Case #2: I MISS YOU TOO. B

_Based on a true story_