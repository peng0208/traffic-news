#!/bin/bash

BIN=/app/traffic-news/bin/traffic-collectd
CONF=/app/traffic-news/conf/collectd.ini
LOG=/app/traffic-news/logs/collectd.log

. /etc/init.d/functions

_start() {
  cmd="/usr/bin/nohup $BIN -c $CONF 2>$LOG 1>/dev/null &"
  daemon $cmd
  echo "Starting..."
}

_stop() {
  killproc $BIN
  echo "Stopping..."

}

_status() {
 status $BIN
}

case $1 in
  start)
    _start
  ;;
  stop)
    _stop
  ;;
  restart)
    _stop && _start
  ;;
  status)
    _status
  ;;
  *)
    echo "Usage: $0 (start|stop|restart|status)"
esac
