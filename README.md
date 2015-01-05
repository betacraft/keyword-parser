keyword-parser
==============

This is supposed to parse a given code in form of a string. We are trying to allow developers to write code in their ]
own spoken lanugage (eg. spanish, marathi, hindi, chinese etc). This should include keywords from that specific spoken 
language. 
Now this parser is supposed to parse given code, for a given programming language, and replace this keywords with 
the original original english keywords, to be given to compiler/interpreter.

currently working example is for ruby.
Should use - https://github.com/RainingClouds/rubyvernac-marathi - this has bunch of useful examples and a gem with translated aliases for method calls

usage - parser test/test.rb ruby test/keywords.txt


*Very early stage*
