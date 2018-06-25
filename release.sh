#!/bin/bash
set -e
set -u
shopt -s nullglob

function usage {
  cat <<"USAGE"
usage: release.sh version-string
USAGE

  exit 1
}

GHR=$(which ghr)
GOX=$(which gox)
ZIP=$(which zip)
SHASUM=$(which shasum)
RM=$(which rm)

GH_ORG="asicsdigital"
APP="dudewheresmy"
PKGDIR="./pkg"
SHASUM_FILE="SHASUMS"

GOX_OSARCH="darwin/amd64 linux/386 linux/amd64"
GOX_OUTPUT="${PKGDIR}/{{.OS}}/{{.Arch}}/{{.Dir}}"

VERSION=${1:-UNSET}
if [[ "$VERSION" == "UNSET" ]]; then
  usage
else
  echo "$(date -R) Releasing version ${VERSION}"
fi

# BUILD
$GOX -osarch="${GOX_OSARCH}" -output="${GOX_OUTPUT}"

pushd $PKGDIR >/dev/null

# TIDY
rm -f $SHASUM_FILE *.zip

# PACKAGE
for os in *; do # os
  pushd $os >/dev/null
  for arch in *; do #arch
    $ZIP -m -j "../${os}_${arch}.zip" "${arch}/${APP}"
    $RM -rf $arch
  done
  popd >/dev/null
  $RM -rf $os
done

# CHECKSUM
$SHASUM -U *.zip > $SHASUM_FILE
popd >/dev/null

# RELEASE
$GHR -u $GH_ORG --replace $VERSION $PKGDIR
echo "$(date -R) Release published"
