#!/bin/bash

cd "$(dirname "$0")"

PKG="$1"
YEAR="$2"
if [ -z "$PKG" ]
then
	echo "Usage: $0 pkgname [year]" >&2
	exit 1
fi
if [ -z "$YEAR" ]
then
	YEAR=$(date +%Y)
fi

mkdir "ch/$PKG"

cat >"ch/$PKG/todo.go" <<EOF
package $PKG

import (
	"errors"
)

var errNotImplemented = errors.New("not implemented")
var errFailed = errors.New("failed to find the solution")
EOF

cat >"ch/$PKG/todo.go" <<EOF
package $PKG

import (
	"errors"
)

var errNotImplemented = errors.New("not implemented")
var errFailed = errors.New("failed to find the solution")
EOF

cat >"ch/$PKG/advent.go" <<EOF
package $PKG

import (
	"github.com/thijzert/advent-of-code/ch"
)

var Advent ch.Advent = ch.Advent {
EOF

for i in 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
do

	echo "	Dec${i}a, Dec${i}b," >> "ch/$PKG/advent.go"

	cat >"ch/$PKG/dec${i}.go" <<EOF
package $PKG

import (
	"github.com/thijzert/advent-of-code/ch"
)

var Dec${i}a ch.AdventFunc = nil
// func Dec${i}a(ctx ch.AOContext) error {
// 	sections, err := ctx.DataSections("inputs/${YEAR}/dec${i}a.txt")
// 	if err != nil {
// 		return err
// 	}
//
// 	ctx.Print(len(sections))
// 	return errNotImplemented
// }

var Dec${i}b ch.AdventFunc = nil
// func Dec${i}b(ctx ch.AOContext) error {
// 	return errNotImplemented
// }

EOF

done

echo "}" >> "ch/$PKG/advent.go"
