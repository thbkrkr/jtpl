#!/bin/bash -eu

test_jtpl() {
  declare name="$1"

  result=$(JTPL_JSON_DATA=$(jq -c . test/$name.json) ./jtpl -f test/$name.tpl)
  expected=$(cat test/$name.expected)

  if [[ "$result" != "$expected" ]]; then
    echo "Test KO"
    echo -e "\nresult:\n\n$result"
    echo -e "\nexpected:\n\n$expected"
    exit 1
  fi

  echo "Test OK"
}

main() {
  test_jtpl example
}

main