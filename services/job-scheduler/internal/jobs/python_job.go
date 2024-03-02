package jobs

import (
	"context"
	"fmt"
	"job-scheduler/pkg/logging"
	"os"
	"os/exec"
)

type PythonJob struct {
	Dir  string
	Args []string
}

func (job *PythonJob) Execute(ctx context.Context) (err error) {
	logger := logging.FromContext(ctx)

	pythonCmd := exec.Command("python", fmt.Sprintf("%s/job.py", job.Dir))
	pythonCmd.Args = append(pythonCmd.Args, job.Args...)

	pythonCmd.Stdout = os.Stdout
	pythonCmd.Stderr = os.Stderr

	fmt.Printf("\nRunning command %v", pythonCmd)
	err = pythonCmd.Run()

	if err != nil {
		logger.Infow("Could not execute python", "err", err)
		return err
	}

	return

}
