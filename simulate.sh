#!/bin/bash
projectDir=$@

osascript -  "$projectDir"  <<EOF
  on run argv -- argv is a list of strings
    tell application "Terminal"
      set bounds of first window to {0, 25, 829, 536}
      do script ("cd " & quoted form of item 1 of argv) in window 1
      do script ("make server") in window 1
  end tell
  end run
EOF

osascript -  "$projectDir"  <<EOF
  on run argv -- argv is a list of strings
    tell application "Terminal"
      set bounds of second window to {830, 25, 1680, 536}
      do script ("cd " & quoted form of item 1 of argv) in window 2
      do script ("make user") in window 2
  end tell
  end run
EOF

osascript -  "$projectDir"  <<EOF
  on run argv -- argv is a list of strings
    tell application "Terminal"
      set bounds of third window to {0, 525, 829, 1050}
      do script ("cd " & quoted form of item 1 of argv) in window 3
      do script ("make user") in window 3
  end tell
  end run
EOF

osascript -  "$projectDir"  <<EOF
  on run argv -- argv is a list of strings
    tell application "Terminal"
      activate
      set bounds of fourth window to {830, 525, 1680, 1050}
      do script ("cd " & quoted form of item 1 of argv) in window 4
      do script ("make user") in window 4
  end tell
  end run
EOF
