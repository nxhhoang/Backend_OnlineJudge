#!/bin/sh

printf "" > examples.tex

printf "%sbegin{exampletable}\n" "\\" >> examples.tex

for f in example.*; do
  case "$f" in
    *.a) continue ;;
  esac
  printf "  %sexample{%s}{%s.a}\n" "\\" "$f" "$f"
done >> examples.tex

printf "%send{exampletable}" "\\" >> examples.tex

./replace_deprecated.sh

latexmk -xelatex -pdfxe -interaction=nonstopmode main.tex -jobname=statement
latexmk -xelatex -pdfxe -interaction=nonstopmode main.tex -jobname=statement

exit 0