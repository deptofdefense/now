#!/bin/bash

# =================================================================
#
# Work of the U.S. Department of Defense, Defense Digital Service.
# Released as open source under the MIT License.  See LICENSE file.
#
# =================================================================

set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

testEpochSeconds() {
  "${DIR}/../bin/now" -e -p s
}

testEpochMilliseconds() {
  "${DIR}/../bin/now" -e -p ms
}

testEpochMicroseconds() {
  "${DIR}/../bin/now" -e -ps us
}

testEpochNanoseconds() {
  "${DIR}/../bin/now" -e -p ns
}

testKitchen() {
  "${DIR}/../bin/now" -f Kitchen
}

testRFC3339Nano() {
  "${DIR}/../bin/now" -f RFC3339Nano
}

testRFC3339() {
  "${DIR}/../bin/now" -f RFC3339
}

testYearMonthDay() {
  "${DIR}/../bin/now" -f 2006-01-02
}

testTimeZoneFixed() {
  "${DIR}/../bin/now" -z UTC+09:30
}

testTimeZoneNamed() {
  "${DIR}/../bin/now" -z America/New_York
}

oneTimeSetUp() {
  echo "Using temporary directory at ${SHUNIT_TMPDIR}"
}

oneTimeTearDown() {
  echo "Tearing Down"
}

# Load shUnit2.
. "${DIR}/shunit2"
