package processor

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/docker/docker/api/types/mount"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// definition encompasses all of the jobs that need to be run for a
// single submission.
type definition struct {
	// internal definition identifier.
	id string

	Jobs       []*Job
	Submission *Submission

	hostTmpDir         string
	hostInputFile      string
	hostSourceCodeFile string
	mounts             []mount.Mount

	ctx context.Context
	wg  *sync.WaitGroup
}

// NewDefinition takes a job definition and writes its files
// to a tempdir on the host. It also generates all of the jobs that
// need to be run.
func NewDefinition(ctx context.Context, sub *Submission) (*definition, error) {
	var d = &definition{
		id:         uuid.New().String(),
		Submission: sub,
		ctx:        ctx,
		wg:         &sync.WaitGroup{},
	}

	if err := d.WriteFiles(); err != nil {
		return nil, err
	}
	d.GenerateMounts()
	if err := d.GenerateJobs(); err != nil {
		return nil, err
	}
	d.wg.Add(len(d.Jobs))

	return d, nil
}

func (d *definition) WriteFiles() error {
	var err error
	d.hostTmpDir, err = os.MkdirTemp("", fmt.Sprintf("bg-%v_", d.id))
	if err != nil {
		return err
	}

	if d.Submission.Extension == "" {
		d.Submission.Extension = "ext"
	}

	if d.Submission.Input != nil {
		d.hostInputFile = filepath.Join(d.hostTmpDir, "input.txt")
		f, err := os.Create(d.hostInputFile)
		if err != nil {
			return errors.Wrap(err, "Failed to create input text file")
		}

		if _, err := f.Write([]byte(*d.Submission.Input)); err != nil {
			return errors.Wrap(err, "Failed to write input text file")
		}

		f.Close()
	}

	d.hostSourceCodeFile = filepath.Join(d.hostTmpDir, fmt.Sprintf("main.%s", d.Submission.Extension))
	f, err := os.Create(d.hostSourceCodeFile)
	if err != nil {
		return errors.Wrap(err, "Failed to create source code file")
	}
	if _, err := f.Write([]byte(d.Submission.Script)); err != nil {
		return errors.Wrap(err, "Failed to write source code file")
	}

	return nil
}

// CleanTmpDir removes the temp directory that was created. This should be called
// once all jobs for the definition have been run.
func (d *definition) CleanTmpDir() error {
	if err := os.RemoveAll(d.hostTmpDir); err != nil {
		return errors.Wrap(err, "Failed to remove temp dir")
	}

	return nil
}

func (d *definition) GenerateMounts() {
	d.mounts = []mount.Mount{
		{
			Type:     mount.TypeBind,
			Source:   d.hostSourceCodeFile,
			Target:   fmt.Sprintf("/bg/%s", filepath.Base(d.hostSourceCodeFile)),
			ReadOnly: true,
		},
	}

	if d.hostInputFile != "" {
		d.mounts = append(d.mounts, mount.Mount{
			Type:     mount.TypeBind,
			Source:   d.hostInputFile,
			Target:   "/bg/input.txt",
			ReadOnly: true,
		})
	}
}

// GenerateJobs generates all of the jobs that need to be run for this definition.
func (d *definition) GenerateJobs() error {
	for i := 0; i < int(d.Submission.Count); i++ {
		job := d.NewJob(fmt.Sprintf("%v", i))
		if err := job.CreateContainer(d.ctx); err != nil {
			return err
		}
		d.Jobs = append(d.Jobs, job)
	}
	return nil
}

func (d *definition) NewJob(id string) *Job {
	return &Job{
		ID:  id,
		Def: d,

		logger: logrus.WithFields(logrus.Fields{
			"Job":        id,
			"Definition": d.id,
		}),
		done: make(chan struct{}, 1),
	}
}

// Run starts all of the jobs for the definition and waits for them all to
// complete or timeout. This function should NOT be called in a goroutine.
func (d *definition) Run() {
	for _, job := range d.Jobs {
		JobChan <- job
	}

	d.wg.Wait()
}
