#!/bin/sh

FILE=$(readlink -f "$0")
VYDON_ROOT="$(dirname "$FILE")/../../"

cluster_destroy()
{
  if ! command -v ctlptl > /dev/null ; then
    echo "requires ctptl to run, see https://github.com/tilt-dev/ctlptl"
    exit 1
  fi


  VYDON_DEV_HOSTPATH="${VYDON_ROOT}/.data"
  mkdir -p "$VYDON_DEV_HOSTPATH"
  chmod 777 "$VYDON_DEV_HOSTPATH"
  sed 's|{VYDON_DEV_HOSTPATH}|'"$VYDON_DEV_HOSTPATH"'|' < "$VYDON_ROOT/tilt/kind/cluster.yaml" | ctlptl delete -f -
}

cluster_destroy
