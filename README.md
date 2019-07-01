# goidioms
idioms and samples of GO program

hello       - a simple hello world sample

echo        - print out whatever typed in

cmdrepeater - repeat what commander entered, with a Yes Sir or Sure Yes

dupline     - count how many repeats of each words you entered until you enter-enter

keywordInfiles - search the keyword in the files your supplied. print their line number and text if exist

tailcommand  - a GOLANG implementation of the tailcommand, followed the C++ implementation at my <a href=https://github.com/cndpost/cppidioms>cppidioms </a>

topk  - print the largest K integers from an input file which contains the integers. Followed the C++ implementation at my <a href=https://github.com/cndpost/cppidioms>cppidioms </a>

producerconsumer - a GOLANG implementation of the producer consumer. Current version does not see interleaving between producer and consumer.


====================

linking golang code to C language code. It is straight forward using the cgo tool to compile the c file and link to go program

linking golang code to C++ langauge code, need write a set of C fucntions to wrap all the methods of a C++ class, then use above cgo way to compile and link to
the c wrapper function, which in turn make calls to the C++ class inside.

