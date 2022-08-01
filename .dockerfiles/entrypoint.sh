#!/bin/sh

[ -n "${PUID}" ] && usermod -u "${PUID}" tube
[ -n "${PGID}" ] && groupmod -g "${PGID}" tube

printf "Configuring tube ..."
[ -z "${DATA}" ] && DATA="/data"

export DATA

printf "Switching UID=%s and GID=%s\n" "${PUID}" "${PGID}"
exec su-exec tube:ytube "$@"
