#!/bin/sh
set -e

OS_NAME=$(uname -s)
ARCH_NAME=$(uname -m)
OS=""
ARCH=""
INSTALL_DIR=${INSTALL_DIR:-/usr/local/bin}

if [ "$OS_NAME" = "Linux" ]; then
	OS="linux"
elif [ "$OS_NAME" = "Darwin" ]; then
	OS="darwin"
elif [ "$OS_NAME" = "FreeBSD" ]; then
	OS="freebsd"
elif [ "$OS_NAME" = "OpenBSD" ]; then
	OS="openbsd"
elif [ "$OS_NAME" = "NetBSD" ]; then
	OS="netbsd"
else
	echo "Unsupported OS: $OS_NAME"
	exit 1
fi

if [ "$ARCH_NAME" = "x86_64" ]; then
	ARCH="amd64"
elif [ "$ARCH_NAME" = "i386" ]; then
	ARCH="386"
else
	echo "Unsupported architecture: $ARCH_NAME"
	exit 1
fi

sudo wget -O "$INSTALL_DIR/git-lazy-commit" "https://github.com/spenserblack/git-lazy-commit/releases/latest/download/git-lazy-commit-$OS-$ARCH"
sudo chmod +x "$INSTALL_DIR/git-lazy-commit"
