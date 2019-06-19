package main

// Uses the Utsname struct and Uname function defined in:
// https://godoc.org/golang.org/x/sys/unix
import (
	"flag"
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"strings"
)

const (
	version = `Version 1`
	help    = `Print certain system information.  With no OPTION, same as -s.

  -a, --all                print all information, in the following order,
                             except omit -p and -i if unknown:
  -s, --kernel-name        print the kernel name
  -n, --nodename           print the network node hostname
  -r, --kernel-release     print the kernel release
  -v, --kernel-version     print the kernel version
  -m, --machine            print the machine hardware name
  -p, --processor          print the processor type (non-portable)
  -i, --hardware-platform  print the hardware platform (non-portable)
  -o, --operating-system   print the operating system
      --help     display this help and exit
      --version  output version information and exit`
)

// The Utsname struct is embeded as to output 'all' values, we need to
// write a method for Utsname, but we are not allowed to add methods to
// non-local types.
// Using 'type Uname unix.Utsname' would require extra casting back and
// forth to use unix.Uname
type Uname struct {
	utsname unix.Utsname
}

type Flags struct {
	help             *bool
	version          *bool
	all              *bool
	sysname          *bool
	nodename         *bool
	kernelrelease    *bool
	kernelversion    *bool
	machine          *bool
	processor        *bool
	hardwareplatform *bool
	operatingsystem  *bool
}

func main() {
	fs := setupFlags()
	flag.Parse()
	// TODO: Flags ignored for now as '--' is not supported. See TODO in
	// root of repo for explaination.

	uname := new(Uname)
	unix.Uname(&(uname.utsname))

	// TODO: Fail when arguments passed
	hasFlags := flag.NFlag() == 0
	switch {
	case *fs.help:
		flag.Usage()
	case *fs.version:
		fmt.Println(version)
	case *fs.all:
		fmt.Println(uname)
	case hasFlags:
		fmt.Printf("%s\n", uname.utsname.Sysname)
	}
	if *fs.version || *fs.all || hasFlags {
		os.Exit(0)
	}
	if flag.NArg() > 0 {
		name := os.Args[0]
		arg := flag.Arg(0)
		fmt.Printf("%s: extra operand ‘%s’\nTry '%s -help' for more information.", name, arg, name)
		os.Exit(1)
	}
	fields := buildOutput(fs, uname)
	fmt.Println(fields)
}

func (un *Uname) String() string {
	u := un.utsname
	return fmt.Sprintf("%s %s %s %s %s %s %s %s", u.Sysname, u.Nodename, u.Release, u.Version, u.Machine, u.Machine, u.Machine, "GNU/Linux")
}

func setupFlags() *Flags {
	fs := new(Flags)
	fs.all = flag.Bool("all", false, "print all information, in the following order, except omit -p and -i if unknown:")
	fs.sysname = flag.Bool("kernel-name", false, "print the kernel name")
	fs.nodename = flag.Bool("nodename", false, "print the network node hostname")
	fs.kernelrelease = flag.Bool("kernel-release", false, "print the kernel release")
	fs.kernelversion = flag.Bool("kernel-version", false, "print the kernel version")
	fs.machine = flag.Bool("machine", false, "print the machine hardware name")
	fs.processor = flag.Bool("processor", false, "print the processor type (non-portable)")
	fs.hardwareplatform = flag.Bool("hardware-platform", false, "print the hardware platform (non-portable)")
	fs.operatingsystem = flag.Bool("operating-system", false, "print the operating system")
	fs.help = flag.Bool("help", false, "display this help and exit")
	fs.version = flag.Bool("version", false, "output version information and exit")
	return fs
}

func buildOutput(fs *Flags, un *Uname) string {
	u := un.utsname
	var sb strings.Builder
	if *fs.sysname {
		sb.Write(u.Sysname[:])
		sb.WriteByte(' ')
	}
	if *fs.nodename {
		sb.Write(u.Nodename[:])
		sb.WriteByte(' ')
	}
	if *fs.kernelrelease {
		sb.Write(u.Release[:])
		sb.WriteByte(' ')
	}
	if *fs.kernelversion {
		sb.Write(u.Version[:])
		sb.WriteByte(' ')
	}
	if *fs.machine {
		sb.Write(u.Machine[:])
		sb.WriteByte(' ')
	}
	// TODO: (non-portable) processor - same as machine
	if *fs.processor {
		sb.Write(u.Machine[:])
		sb.WriteByte(' ')
	}
	// TODO: (non-portable) hadware-platform - same as machine
	if *fs.hardwareplatform {
		sb.Write(u.Machine[:])
		sb.WriteByte(' ')
	}
	// TODO: hardwareplatform
	if *fs.operatingsystem {
		sb.WriteString("GNU/Linux")
	}
	return sb.String()
}
