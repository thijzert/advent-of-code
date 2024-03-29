package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/thijzert/go-resemble"
)

type job func(ctx context.Context) error

type compileConfig struct {
	Development    bool
	Quick          bool
	GOOS           string
	GOARCH         string
	PackageVersion string
}

func main() {
	var conf compileConfig
	watch := false
	run := false
	flag.BoolVar(&conf.Development, "development", false, "Create a development build")
	flag.BoolVar(&conf.Quick, "quick", false, "Create a development build")
	flag.StringVar(&conf.GOARCH, "GOARCH", "", "Cross-compile for architecture")
	flag.StringVar(&conf.GOOS, "GOOS", "", "Cross-compile for operating system")
	flag.StringVar(&conf.PackageVersion, "version", "", "Override embedded version number")
	flag.BoolVar(&watch, "watch", false, "Watch source tree for changes")
	flag.BoolVar(&run, "run", false, "Run the program upon successful compilation")
	flag.Parse()

	var theJob job

	if run {
		theJob = func(ctx context.Context) error {
			err := compile(ctx, conf)
			if err != nil {
				return err
			}
			runArgs := append([]string{"build/aoc"}, flag.Args()...)
			return passthru(ctx, runArgs...)
		}
	} else {
		theJob = func(ctx context.Context) error {
			return compile(ctx, conf)
		}
	}

	if watch {
		theJob = watchSourceTree([]string{"."}, []string{"*.go"}, theJob)
	}

	err := theJob(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func compile(ctx context.Context, conf compileConfig) error {
	// Embed static assets
	if err := os.Chdir("data"); err != nil {
		return errors.Errorf("Error: cannot find data directory. (error: %s)", err)
	}
	var emb resemble.Resemble
	emb.OutputFile = "assets.go"
	emb.PackageName = "data"
	emb.Debug = conf.Development
	emb.AssetPaths = []string{
		"inputs",
		"results.txt",
	}
	if err := emb.Run(); err != nil {
		return errors.WithMessage(err, "error running 'resemble'")
	}

	os.Chdir("../")

	// Build main executable
	execOutput := "build/aoc"
	if runtime.GOOS == "windows" || conf.GOOS == "windows" {
		execOutput += ".exe"
	}

	gofiles, err := filepath.Glob("cmd/aoc/*.go")
	if err != nil || len(gofiles) == 0 {
		return errors.WithMessage(err, "error: cannot find any go files to compile.")
	}
	compileArgs := append([]string{
		"build",
		"-o", execOutput,
	}, gofiles...)

	compileCmd := exec.CommandContext(ctx, "go", compileArgs...)

	compileCmd.Env = append(compileCmd.Env, os.Environ()...)
	if conf.GOOS != "" {
		compileCmd.Env = append(compileCmd.Env, "GOOS="+conf.GOOS)
	}
	if conf.GOARCH != "" {
		compileCmd.Env = append(compileCmd.Env, "GOARCH="+conf.GOARCH)
	}

	err = passthruCmd(compileCmd)
	if err != nil {
		return errors.WithMessage(err, "compilation failed")
	}

	if conf.Development && !conf.Quick {
		log.Printf("Development build finished.")
	} else {
		log.Printf("Compilation finished.")
	}

	return nil
}

func passthru(ctx context.Context, argv ...string) error {
	c := exec.CommandContext(ctx, argv[0], argv[1:]...)
	return passthruCmd(c)
}

func passthruCmd(c *exec.Cmd) error {
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	return c.Run()
}

func watchSourceTree(paths []string, fileFilter []string, childJob job) job {
	return func(ctx context.Context) error {
		var mu sync.Mutex
		for {
			lastHash := sourceTreeHash(paths, fileFilter)
			current := lastHash
			cctx, cancel := context.WithCancel(ctx)
			go func() {
				mu.Lock()
				err := childJob(cctx)
				if err != nil {
					log.Printf("child process: %s", err)
				}
				mu.Unlock()
			}()

			for lastHash == current {
				time.Sleep(250 * time.Millisecond)
				current = sourceTreeHash(paths, fileFilter)
			}

			log.Printf("Source change detected - rebuilding")
			cancel()
		}
	}
}

func sourceTreeHash(paths []string, fileFilter []string) string {
	h := sha1.New()
	for _, d := range paths {
		h.Write(directoryHash(0, d, fileFilter))
	}
	return hex.EncodeToString(h.Sum(nil))
}

func directoryHash(level int, filePath string, fileFilter []string) []byte {
	h := sha1.New()
	h.Write([]byte(filePath))

	fi, err := os.Stat(filePath)
	if err != nil {
		return h.Sum(nil)
	}
	if fi.IsDir() {
		base := filepath.Base(filePath)
		if level > 0 {
			if base == ".git" || base == ".." || base == "node_modules" {
				return []byte{}
			}
		}
		// recurse
		var names []string
		f, err := os.Open(filePath)
		if err == nil {
			names, err = f.Readdirnames(-1)
		}
		if err == nil {
			for _, name := range names {
				if name == "" || name[0] == '.' {
					continue
				}
				h.Write(directoryHash(level+1, path.Join(filePath, name), fileFilter))
			}
		}
	} else {
		if fileFilter != nil {
			found := false
			for _, pattern := range fileFilter {
				if ok, _ := filepath.Match(pattern, filePath); ok {
					found = true
				} else if ok, _ := filepath.Match(pattern, filepath.Base(filePath)); ok {
					found = true
				}
			}
			if !found {
				return []byte{}
			}
		}
		f, err := os.Open(filePath)
		if err == nil {
			io.Copy(h, f)
			f.Close()
		}
	}
	return h.Sum(nil)
}
