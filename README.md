# kidsmath
a simple program to generate math quizs for my kid.

```
./kidsmath -h
Generate multiplcation, substraction, addition, division quizs within a given range.

Usage:
  kidsmath [command]

Available Commands:
  basic       A brief description of your command
  completion  Generate the autocompletion script for the specified shell
  expr        generate a random expression , eg: (1+2) * 3
  help        Help about any command

Flags:
  -h, --help             help for kidsmath
      --maxdev int8      defines max deviation in percentile of the generated number, 1-99 (default 50)
      --pattern string   defines the boundary of binary-operations (default "100x10")
  -t, --toggle           Help message for toggle

Use "kidsmath [command] --help" for more information about a command.

./kidsmath basic add --maxdev 50 --pattern "500x100"
2023/12/13 11:09:18 N=500, M=100
295 + 85 =
308 + 86 =
446 + 68 =
280 + 50 =
353 + 95 =
409 + 63 =
391 + 70 =
498 + 65 =
306 + 77 =
288 + 95 =
...
```
