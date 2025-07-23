#!/bin/sh

#generate pdf files based on Polygon problem directory

source="$1" # Polygon package path
dest="$2" # problem path

workdir="$source/statement-sections/english"

cp scripts/gen_statement/gen_pdf.sh $workdir || exit -1
cp scripts/gen_statement/main.tex $workdir || exit -1
cp scripts/gen_statement/styles.sty $workdir || exit -1
cp scripts/gen_statement/logo.png $workdir || exit -1
cp scripts/gen_statement/replace_deprecated.sh $workdir || exit -1

cd $workdir

# ICPC type doesnt have scoring section, so this is psuedo-file
[ -f yourfile.txt ] && printf "%smyheading{Chấm điểm}\n\n" "\\" | cat - scoring.tex 2>/dev/null > tmp && mv tmp scoring.tex
touch scoring.tex

./gen_pdf.sh || exit -1

cp ./statement.pdf "$dest" || exit -1

exit 0