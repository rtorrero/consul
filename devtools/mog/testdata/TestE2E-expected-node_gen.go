// Code generated by mog. DO NOT EDIT.

package sourcepkg

import "github.com/hashicorp/mog/internal/e2e/core"

func NodeToCore(s *Node, t *core.ClusterNode) {
	if s == nil {
		return
	}
	t.ID = s.ID
	WorkloadToCore(&s.F1, &t.F1)
	{
		if s.F2 != nil {
			var x core.Workload
			WorkloadToCore(s.F2, &x)
			t.F2 = &x
		}
	}
	if s.F3 != nil {
		WorkloadToCore(s.F3, &t.F3)
	}
	{
		var x core.Workload
		WorkloadToCore(&s.F4, &x)
		t.F4 = &x
	}
}
func NodeFromCore(t *core.ClusterNode, s *Node) {
	if s == nil {
		return
	}
	s.ID = t.ID
	WorkloadFromCore(&t.F1, &s.F1)
	{
		if t.F2 != nil {
			var x Workload
			WorkloadFromCore(t.F2, &x)
			s.F2 = &x
		}
	}
	{
		var x Workload
		WorkloadFromCore(&t.F3, &x)
		s.F3 = &x
	}
	if t.F4 != nil {
		WorkloadFromCore(t.F4, &s.F4)
	}
}
func WorkloadToCore(s *Workload, t *core.Workload) {
	if s == nil {
		return
	}
	t.ID = s.ID
}
func WorkloadFromCore(t *core.Workload, s *Workload) {
	if s == nil {
		return
	}
	s.ID = t.ID
}