keyword-parser
==============

This is supposed to parse a given code in the form of a string. We are trying to allow developers to write code in their 
own spoken lanugage (eg. marathi, hindi, chinese etc). This should include keywords from that specific spoken 
language. 
Now this parser is supposed to parse given code, for a given programming language, and replace the keywords with 
the original english keywords, to be given to compiler/interpreter.

currently working example is for ruby.
Should use - https://github.com/RainingClouds/rubyvernac-marathi or https://github.com/RainingClouds/rubyvernac-hindi - this has bunch of useful examples and a gem with translated aliases for method calls.

Usage - 
>keyword-parser test/test.rb ruby test/keywords.txt

Installation -

Install Go - https://golang.org/doc/install

> go get github.com/rainingclouds/keyword-parser 

*Very early stage*
