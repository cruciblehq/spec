// Defines the deployment plan specification.
//
// A plan describes the concrete infrastructure and routing configuration
// needed to deploy a set of resources. Plans are produced by executing a
// [blueprint.Blueprint] against the registry: service references are
// resolved, compute resources allocated, bindings established, and
// gateway routes configured.
//
// Plans are encoded as JSON. Use [Encode] and [Decode] to convert between
// a [Plan] value and its JSON representation. Both functions validate the
// plan automatically.
//
// Decoding a plan:
//
//	p, err := plan.Decode(data)
//	if err != nil {
//		log.Fatal(err) // malformed JSON or validation failure
//	}
//	for _, svc := range p.Services {
//		fmt.Println(svc.ID, svc.Reference)
//	}
//
// Encoding a plan:
//
//	p := &plan.Plan{
//		Services: []plan.Service{{ID: "hub", Reference: "crucible/hub@1.0.0"}},
//		Compute:  []plan.Compute{{ID: "main", Provider: "local"}},
//		Bindings: []plan.Binding{{Service: "hub", Compute: "main"}},
//		Gateway:  plan.Gateway{Routes: []plan.Route{{Pattern: "/api", Service: "hub"}}},
//	}
//	data, err := plan.Encode(p)
package plan
