#!/bin/sh

set -e

golint -set_exit_status
exit $?
