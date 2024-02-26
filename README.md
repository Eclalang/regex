# Regex library

## Candidate functions :

|  Func Name   |                          Prototype                           |                                   Description                                    | Comments |
|:------------:|:------------------------------------------------------------:|:--------------------------------------------------------------------------------:|:--------:|
|     Find     |          Find(pattern string, str string) string {}          |               Returns the first match of the pattern in the string               |   N/A    |
|   FindAll    |       FindAll(pattern string, str string) []string {}        |                 Returns all matches of the pattern in the string                 |   N/A    |
| FindAllIndex |      FindAllIndex(pattern string, str string) []int {}       |         Returns the indexes of all matches of the pattern in the string          |   N/A    |
|  FindIndex   |        FindIndex(pattern string, str string) []int {}        | Returns the first and last index of the first match of the pattern in the string |   N/A    |
|    Match     |          Match(pattern string, str string) bool {}           |                  Returns true if the pattern matches the string                  |   N/A    |
|   Replace    |     Replace(pattern, str, new string, nb int) string {}      |  Replaces the nb first matches of the pattern in the string with the new string  |   N/A    |
|  ReplaceAll  | ReplaceAll(pattern string, str string, new string) string {} |      Replaces all matches of the pattern in the string with the new string       |   N/A    |