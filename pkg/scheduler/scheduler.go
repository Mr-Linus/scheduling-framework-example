package scheduler

import (
	"context"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const (
	Name = "Scheduling-framework-example"
)

var (
	_ framework.FilterPlugin = &Scheduler{}
	_ framework.ScorePlugin  = &Scheduler{}
)

type Scheduler struct {
	handle framework.Handle
}

func (s *Scheduler) Name() string {
	return Name
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &Scheduler{
		handle: h,
	}, nil
}

func (s *Scheduler) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	klog.Infof("Filter Node: %v while Scheduling Pod: %v/%v. ", nodeInfo.Node().GetName(), pod.GetNamespace(), pod.GetName())
	// TODO: Write Your Filter Policy here.
	// ..
	return nil
}

func (s *Scheduler) Score(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {
	klog.Infof("Scoring Node: %v while scheduling Pod: %v/%v", nodeName, pod.GetNamespace(), pod.GetName())
	// TODO: Write Your Score Policy here.
	// ...
	return 0, nil
}

func (s *Scheduler) NormalizeScore(ctx context.Context, cycleState *framework.CycleState, pod *v1.Pod, scores framework.NodeScoreList) *framework.Status {
	var (
		highest int64 = 0
		lowest        = scores[0].Score
	)

	for _, nodeScore := range scores {
		if nodeScore.Score < lowest {
			lowest = nodeScore.Score
		}
		if nodeScore.Score > highest {
			highest = nodeScore.Score
		}
	}

	if highest == lowest {
		lowest--
	}

	// Set Range to [0-100]
	for i, nodeScore := range scores {
		scores[i].Score = (nodeScore.Score - lowest) * framework.MaxNodeScore / (highest - lowest)
		klog.Infof("Node: %v, Score: %v in Plugin: Mandalorian When scheduling Pod: %v/%v", scores[i].Name, scores[i].Score, pod.GetNamespace(), pod.GetName())
	}
	return nil
}

func (s *Scheduler) ScoreExtensions() framework.ScoreExtensions {
	return s
}
