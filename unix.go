// +build darwin dragonfly freebsd linux netbsd openbsd solaris

package main

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"golang.org/x/sys/unix"
)

func do(name string, args []string) {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, unix.SIGINT, unix.SIGTERM)

	err := cmd.Start()
	if err != nil {
		die(err)
	}

	errc := make(chan error)
	go func() { errc <- cmd.Wait() }()

	select {
	case sig := <-c:
		err := syscall.Kill(-cmd.Process.Pid, sig.(syscall.Signal))
		if err != nil {
			die(err)
		}

	case err := <-errc:
		if err != nil {
			if ee, ok := err.(*exec.ExitError); ok {
				os.Exit(ee.Sys().(syscall.WaitStatus).ExitStatus())
			} else {
				die(err)
			}
		}
	}
}
