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
	"io/ioutil"
	"path/filepath"
)

func SetUpNS(rootPath string) {
	if vars.Debug == true {
		log.Printf("DEBUG: checking rootfs directory: %s.\n", rootPath)
	}
	_, err := os.Stat(rootPath)
	if os.IsNotExist(err) {
		log.Fatalf("ERROR: root directory doesn't exist: %v\n", err)
	}
	pivotRoot, err := ioutil.TempDir(rootPath, ".pivot_root")
	if err != nil {
		log.Printf("ERROR: setting up pivot dir: %v\n", err)
		os.Exit(1)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: setting up pivot_root directory: %s.\n", pivotRoot)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: setting hostname in the container.\n")
	}
	if err := unix.Sethostname([]byte("container")); err != nil {
		log.Printf("ERROR: failed to set hostname: %v\n", err)
		os.Exit(1)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: mounting proc onto %sproc.\n", rootPath)
	}
	if err := unix.Mount("proc", filepath.Join(rootPath, "/proc"), "proc", 0, ""); err != nil {
		log.Printf("ERROR: failed to mount proc: %v\n", err)
		os.Exit(1)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: bind mounting root onto itself (workaround for pivot_root).\n")
	}
	if err := unix.Mount(rootPath, rootPath, "", unix.MS_BIND|unix.MS_REC, ""); err != nil {
		log.Printf("ERROR: failed to mount root: %v\n", err)
		os.Exit(1)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: creating new pivot directory: %s.\n", pivotRoot)
	}
	if err := os.MkdirAll(pivotRoot, 0700); err != nil {
		log.Printf("ERROR: failed to create pivot_root: %v\n", err)
		os.Exit(1)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: executing pivot_root %s %s.\n", rootPath, pivotRoot)
	}
	if err := unix.PivotRoot(rootPath, pivotRoot); err != nil {
		log.Printf("ERROR: failed to pivot to new root: %v\n", err)
		os.Exit(1)
	}
	pivotRoot = filepath.Join("/", filepath.Base(pivotRoot))
	if vars.Debug == true {
		log.Printf("DEBUG: changing directory to /.\n")
	}
	if err := unix.Chdir("/"); err != nil {
		log.Printf("ERROR: failed to change dir to /: %v\n", err)
		os.Exit(1)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: unmounting pivot_root directory: %s.\n", pivotRoot)
	}
	if err := unix.Unmount(pivotRoot, unix.MNT_DETACH); err != nil {
		log.Printf("ERROR: failed to unmount %s: %v\n", pivotRoot, err)
		os.Exit(1)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: removing pivot_root directory: %s.\n", pivotRoot)
	}
	if err := os.RemoveAll(pivotRoot); err != nil {
		log.Printf("ERROR: failed to remove %s: %v\n", pivotRoot, err)
		os.Exit(1)
	}
}

func Child(command, rootfs string) {
	SetUpNS(rootfs)
	var cmd *exec.Cmd

	commandArgs := strings.Split(command, " ")
	if vars.Debug == true {
		log.Printf("DEBUG: executing command %v in container.\n", commandArgs)
	}
	if len(commandArgs) == 1 {
		cmd = exec.Command(commandArgs[0])
	} else {
		cmd = exec.Command(commandArgs[0], commandArgs[1:]...)
	}
	if vars.Debug == true {
		log.Printf("DEBUG: setting up extra PATH=/bin:/sbin inside new namespace.\n")
	}
	cmd.Env = []string{"PATH=/bin:/sbin",`PS1=[\u@\h]\$ `}
	if vars.Debug == true {
		log.Printf("DEBUG: mapping stdin, stdout and stderr.\n")
	}
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
		log.Printf("DEBUG: executing command: /proc/self/exe container -d run fork -i %s -c %s\n", image, command)
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