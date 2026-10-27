#!/usr/bin/env bash
# One-time GitHub setup: labels, sprint milestones and the full backlog as issues.
# Needs the GitHub CLI (https://cli.github.com), logged in:  gh auth login
# Run from the repo root after pushing:  bash scripts/setup-github.sh
# Works in Linux, WSL and Git Bash on Windows. On macOS install GNU date first:
#   brew install coreutils && export PATH="$(brew --prefix)/opt/coreutils/libexec/gnubin:$PATH"
set -euo pipefail

REPO="$(gh repo view --json nameWithOwner -q .nameWithOwner)"
START_DATE="${START_DATE:-2026-11-01}"   # first day of Sprint 0 (a Sunday)
echo "Setting up $REPO, sprints starting $START_DATE"

label() { gh label create "$1" --color "$2" --description "${3:-}" --force >/dev/null; echo "label $1"; }

for t in story:1d76db bug:d73a4a spike:fbca04 chore:c5def5 docs:0e8a16; do label "type:${t%%:*}" "${t##*:}"; done
for a in api ai infra web mobile docs; do label "area:$a" "bfdadc"; done
label size:S c2e0c6 "half a day or less"
label size:M fef2c0 "1-2 days"
label size:L f9d0c4 "3+ days - split before Ready"
label priority:must b60205
label priority:should d93f0b
label priority:could fbca04

# Milestones: one per sprint, due on the Thursday of that week.
declare -A MS
for n in $(seq 0 13); do
  due="$(date -u -d "$START_DATE + $((n*7 + 4)) days" +%Y-%m-%dT18:00:00Z)"
  title="Sprint $n"
  num="$(gh api "repos/$REPO/milestones?state=all&per_page=100" -q ".[] | select(.title==\"$title\") | .number")"
  if [[ -z "$num" ]]; then
    num="$(gh api "repos/$REPO/milestones" -f title="$title" -f due_on="$due" -q .number)"
  fi
  MS[$n]="$title"
  echo "milestone $title (due $due)"
done

# Issues from the backlog file.
while IFS='|' read -r sprint labels title body; do
  [[ -z "$sprint" || "$sprint" == \#* ]] && continue
  gh issue create --title "$title" --body "${body:-See sprint plan.}" \
    --label "$labels" --milestone "${MS[$sprint]}" >/dev/null
  echo "issue [S$sprint] $title"
done < "$(dirname "$0")/backlog.psv"

echo
echo "Done. Now create the board: GitHub -> Projects -> New project -> Board,"
echo "add columns Backlog / Ready / In progress / In review / Done, and add these issues."
