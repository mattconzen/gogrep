# gogrep

## Synopsis 

Simple tool to pass a comma-separated list into Git Grep to find lists of strings in repositories.

Used in a migration project that required lots of searching for many things across many repos.

## Usage

Command:
`./gogrep -input input.csv -output output.csv -repo C:\users\mconzen\repos\examplerepo\`

Output:
```
  prefs.csv:192.168,ARLINGTON,GASPSQL,aspexfiles,FT1
 - ARLINGTON

  prefs.csv:192.168,GASPSQL,aspexfiles,FT1
 - GASPSQL

  prefs.csv:192.168,aspexfiles,FT1
 - aspexfiles
 
   
  'SQL4' not found in repository.
  
  'SQL5' not found in repository.
  
  'SQL6' not found in repository.
  
  'SQL7' not found in repository.


 ```
