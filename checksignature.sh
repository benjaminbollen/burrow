#!/bin/sh
# Thanks to https://mikegerwitz.com/papers/git-horror-story#merge-1
# Validate signatures on each and every commit within the given range
##

# if a ref is provided, append range spec to include all children
chkafter="${1+$1..}"

# parse list of public keys for retrieval
keys=$(git log --show-signature "${chkafter:-HEAD}" \
  | grep 'key ID' \
  | grep -o '[A-Z0-9]\+$' \
  | sort \
  | uniq)

# retrieve public keys from following locations
gpg --keyserver http://keyserver.pgp.com --recv-keys $keys
gpg --keyserver pgp.mit.edu --recv-keys $keys

# make sure all keys are trusted
# The "-E" makes this work with both GNU sed and OS X sed
gpg --list-keys --fingerprint --with-colons |
  sed -E -n -e 's/^fpr:::::::::([0-9A-F]+):$/\1:6:/p' |
  gpg --import-ownertrust

# note: bash users may instead use $'\t'; the echo statement below is a more
# portable option
t=$( echo '\t' )

# Check every commit after chkafter (or all commits if chkafter was not
# provided) for a trusted signature, listing invalid commits. %G? will output
# "G" if the signature is trusted.
git log --pretty="format:%H$t%aN$t%s$t%G?" "${chkafter:-HEAD}" \
  | grep -v "${t}G$"

# grep will exit with a non-zero status if no matches are found, which we
# consider a success, so invert it
[ $? -gt 0 ]
