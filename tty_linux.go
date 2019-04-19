// +build linux

package tty

import (
	"syscall"
)

const (
	ioctlReadTermios  = syscall.TCGETS   // 0x5401
	ioctlWriteTermios = syscall.TCSETS   // 0x5402
)
