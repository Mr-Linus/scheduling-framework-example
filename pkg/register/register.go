package register

import (
	"github.com/spf13/cobra"

	"k8s.io/kubernetes/cmd/kube-scheduler/app"

	"github.com/NJUPT-ISL/scheduling-framework-example/pkg/scheduler"

)

func Register() *cobra.Command {
	return app.NewSchedulerCommand(
		app.WithPlugin(scheduler.Name, scheduler.New),
	)
}