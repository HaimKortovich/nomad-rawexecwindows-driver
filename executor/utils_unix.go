// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package executor

import (
	"os/exec"
	"syscall"
)

// isolateCommand sets the setsid flag in exec.Cmd to true so that the process
// becomes the process leader in a new session and doesn't receive signals that
// are sent to the parent process.
func isolateCommand(cmd *exec.Cmd) {
	if cmd.SysProcAttr == nil {
		cmd.SysProcAttr = &syscall.SysProcAttr{}
	}
	cmd.SysProcAttr.Setsid = true
}
