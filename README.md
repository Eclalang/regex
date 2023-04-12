## REGEX LIBRARY FOR ECLA

# Candidate functions :

|  Func Name   |                         Prototype                          |                                  Description                                   | Comments |
|:------------:|:----------------------------------------------------------:|:------------------------------------------------------------------------------:|:--------:|
|     Find     |          Find(regex string, str string) string {}          |               Returns the first match of the regex in the string               |   N/A    |
|   FindAll    |       FindAll(regex string, str string) []string {}        |                 Returns all matches of the regex in the string                 |   N/A    |
| FindAllIndex |      FindAllIndex(regex string, str string) []int {}       |         Returns the indexes of all matches of the regex in the string          |   N/A    |
|  FindIndex   |        FindIndex(regex string, str string) []int {}        | Returns the first and last index of the first match of the regex in the string |   N/A    |
|    Match     |          Match(regex string, str string) bool {}           |                  Returns true if the regex matches the string                  |   N/A    |
|  ReplaceAll  | ReplaceAll(regex string, str string, new string) string {} |      Replaces all matches of the regex in the string with the new string       |   N/A    |