package manifest

import (
	"errors"
	"testing"
)

func validPlan() Plan {
	return Plan{
		Version:  PlanVersion,
		Services: map[string]string{"svc": "crucible/foo 1.0.0"},
		Infrastructure: Infrastructure{
			Computes: map[string]Compute{"cmp": {Type: ComputeTypeLocal, Config: &ComputeLocal{Host: "h"}}},
			Networks: map[string]Network{"net": {}},
		},
		Containers:  map[string]Container{"ctr": {}},
		Deployments: []Deployment{{Service: "svc", Container: "ctr", Compute: "cmp", Network: "net"}},
	}
}

func TestPlanValidate(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*Plan)
		want   error
	}{
		{"valid", func(*Plan) {}, nil},
		{"bad version", func(p *Plan) { p.Version = PlanVersion + 1 }, ErrUnsupportedPlanVersion},
		{"empty service target", func(p *Plan) { p.Services["svc"] = "" }, ErrMissingRefTarget},
		{
			"infrastructure error",
			func(p *Plan) { p.Infrastructure.Computes["cmp"] = Compute{Type: ComputeTypeLocal} },
			ErrMissingComputeProvider,
		},
		{
			"container error",
			func(p *Plan) {
				p.Containers["ctr"] = Container{Network: Network{Ingress: []IngressRule{{Protocol: "nope"}}}}
			},
			ErrInvalidPlan,
		},
		{
			"deployment missing field",
			func(p *Plan) { p.Deployments[0].Service = "" },
			ErrMissingDeploymentService,
		},
		{
			"undefined service ref",
			func(p *Plan) { p.Deployments[0].Service = "ghost" },
			ErrInvalidPlan,
		},
		{
			"undefined container ref",
			func(p *Plan) { p.Deployments[0].Container = "ghost" },
			ErrInvalidPlan,
		},
		{
			"undefined compute ref",
			func(p *Plan) { p.Deployments[0].Compute = "ghost" },
			ErrInvalidPlan,
		},
		{
			"undefined network ref",
			func(p *Plan) { p.Deployments[0].Network = "ghost" },
			ErrInvalidPlan,
		},
		{
			"undefined environment ref",
			func(p *Plan) { p.Deployments[0].Environment = "ghost" },
			ErrInvalidPlan,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			p := validPlan()
			c.mutate(&p)
			if got := p.Validate(); !errors.Is(got, c.want) {
				t.Fatalf("Plan.Validate() = %v, want %v", got, c.want)
			}
		})
	}
}
