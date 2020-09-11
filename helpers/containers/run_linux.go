// +build linux

package containers

import (
	"github.com/zerodayz/rcli/vars"
	"golang.org/x/sys/unix"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func SetUpNS(rootfs string) {
	_, err := os.Stat(rootfs)
	if os.IsNotExist(err) {
		log.Fatal("ERROR: rootfs directory doesn't exist.")
	}
	if err := unix.Sethostname([]byte("container")); err != nil {
		log.Println("ERROR: failed to set hostname.")
		os.Exit(1)
	}
	if err := unix.Mount("proc", rootfs + "/proc", "proc", 0, ""); err != nil {
		log.Println("ERROR: failed to mount proc.")
		os.Exit(1)
	}
	if err := unix.Mount(rootfs, rootfs, "", unix.MS_BIND|unix.MS_REC, ""); err != nil {
		log.Println("ERROR: failed to mount rootfs.")
		os.Exit(1)
	}
	if err := os.MkdirAll(rootfs + "/.old_root", 0700); err != nil {
		log.Println("ERROR: failed to create .old_root directory.")
		os.Exit(1)
	}
	if err := unix.PivotRoot(rootfs, rootfs + "/.old_root"); err != nil {
		log.Println("ERROR: failed to pivot to new root.")
		os.Exit(1)
	}
	if err := unix.Chdir("/"); err != nil {
		log.Println("ERROR: failed to change dir to /.")
		os.Exit(1)
	}
	if err := unix.Unmount("/.old_root", unix.MNT_DETACH); err != nil {
		log.Println("ERROR: failed to unmount .old_root.")
		os.Exit(1)
	}
	if err := os.RemoveAll("/.old_root"); err != nil {
		log.Println("ERROR: failed to remove .old_root")
		os.Exit(1)
	}
}

func Child(command, rootfs string) {
	SetUpNS(rootfs)
	var cmd *exec.Cmd

	commandArgs := strings.Split(command, " ")
	if vars.Debug == true {
		log.Printf("DEBUG: Executing command %v in container.\n", commandArgs)
	}
	if len(commandArgs) == 1 {
		cmd = exec.Command(commandArgs[0])
	} else {
		cmd = exec.Command(commandArgs[0], commandArgs[1:]...)
	}
	cmd.Env = []string{"PATH=/bin:/sbin",`PS1=[\u@\h]\$ `}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	unix.Unmount("/proc", 0)
}

func Run(command, image string) {
	var cmd *exec.Cmd
	if vars.Debug == true {
		cmd = exec.Command("/proc/self/exe",
			append([]string{"container", "-d", "run", "fork",
				"-i", image,
				"-c", command})...)
	} else {
		cmd = exec.Command("/proc/self/exe",
			append([]string{"container", "run", "fork",
				"-i", image,
				"-c", command})...)
	}
	cmd.SysProcAttr = &unix.SysProcAttr{
		Cloneflags: unix.CLONE_NEWUTS |
					unix.CLONE_NEWPID |
					unix.CLONE_NEWNS  |
					unix.CLONE_NEWIPC |
					// Make sure you have enabled user_namespaces:
					// sudo su -c 'echo "user.max_user_namespaces=15064" > /etc/sysctl.d/00-namespaces.conf'
					// sudo sysctl --system
					unix.CLONE_NEWUSER |
					unix.CLONE_NEWNET,
		Unshareflags: unix.CLONE_NEWNS,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}