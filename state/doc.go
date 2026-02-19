// Defines the deployment state specification.
//
// State records what resources have been deployed and their runtime
// identifiers. It is persisted after each successful deployment and read
// back during incremental planning so the system can determine what has
// changed since the last deployment.
//
// State is encoded as JSON. Use [Encode] and [Decode] to convert between
// a [State] value and its JSON representation. Both functions validate
// the state automatically — [Encode] before marshaling, [Decode] after
// unmarshaling.
//
// Decoding a state:
//
//	s, err := state.Decode(data)
//	if err != nil {
//		log.Fatal(err) // malformed JSON or validation failure
//	}
//	fmt.Println("deployed at", s.Deployment.DeployedAt)
//	for _, svc := range s.Services {
//		fmt.Println(svc.ID, svc.ResourceID)
//	}
//
// Encoding a state:
//
//	s := &state.State{
//		Deployment: state.Deployment{DeployedAt: time.Now()},
//		Services: []state.Service{{
//			ID:         "hub",
//			Reference:  "crucible/hub@1.0.0",
//			ResourceID: "hub-abc123",
//		}},
//	}
//	data, err := state.Encode(s)
package state
