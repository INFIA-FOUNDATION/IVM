package main

import (
	"os"
	"runtime/pprof"

	"github.com/INFIA-FOUNDATION/TEST_VM/classpath"
	"github.com/INFIA-FOUNDATION/TEST_VM/cmdline"
	"github.com/INFIA-FOUNDATION/TEST_VM/jutil"
	"github.com/INFIA-FOUNDATION/TEST_VM/interpreter"
	"github.com/INFIA-FOUNDATION/TEST_VM/options"
	"github.com/INFIA-FOUNDATION/TEST_VM/rtda"
	"github.com/INFIA-FOUNDATION/TEST_VM/rtda/heap"
	_ "github.com/INFIA-FOUNDATION/TEST_VM/native"
)

func startJVM(cmd *cmdline.Command) {
	Xcpuprofile := cmd.Options().Xcpuprofile
	if Xcpuprofile != "" {
		f, err := os.Create(Xcpuprofile)
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	options.InitOptions(cmd.Options())

	cp := classpath.Parse(cmd.Options().Classpath())
	heap.InitBootLoader(cp)

	mainClassName := jutil.ReplaceAll(cmd.Class(), ".", "/")
	mainThread := createMainThread(mainClassName, cmd.Args())
	interpreter.Loop(mainThread)
	interpreter.KeepAlive()
}

func createMainThread(className string, args []string) *rtda.Thread {
	mainThread := rtda.NewThread(nil)
	bootMethod := heap.BootstrapMethod()
	bootArgs := []interface{}{className, args}
	mainThread.InvokeMethodWithShim(bootMethod, bootArgs)
	return mainThread
}
