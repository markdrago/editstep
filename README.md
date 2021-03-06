This problem is called "Edit Step Ladder".  I found it in the book
["Programming Challenges" by Steven S. Skiena & Miguel A. Revilla](http://www.acmsolver.org/books/Programming_Challenges_Miguel_Skiena.pdf).
The problem is numbered 9.6.5 and it is on page 210 of the book.  I have copied
the problem description here:

-----------------------

An edit step is a transformation from one word x to another word y such that x
and y are words in the dictionary, and x can be transformed to y by adding,
deleting, or changing one letter. The transformations from dig to dog and from
dog to do are both edit steps. An edit step ladder is a lexicographically
ordered sequence of words w1, w2, . . . , wn such that the transformation from
wi to wi+1 is an edit step for all i from 1 to n − 1.

For a given dictionary, you are to compute the length of the longest edit step
ladder.

Input
The input to your program consists of the dictionary: a set of lowercase words
in lexicographic order at one word per line. No word exceeds 16 letters and
there are at most 25,000 words in the dictionary.

Output
The output consists of a single integer, the number of words in the longest
edit step ladder.

Sample Input
cat
dig
dog
fig
fin
fine
fog
log
wine

Sample Output
5

-----------------------

The only modification I have made to the description is the word length is
between 3 and 8 characters inclusive and there are ~35000 words in the list.

The program you write will read the input from stdin and output a single number
to stdout.  You should be able to run your program by running
"./yourproghere < wordlisthere.txt" and the only output should be a number
followed by a newline.

