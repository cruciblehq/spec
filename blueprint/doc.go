// Defines the blueprint specification for system composition.
//
// A blueprint declares which resources should be deployed together and how
// they are composed and exposed. Blueprints are the primary input to the
// deployment process: the CLI reads one, resolves resource references against
// the registry, and produces a [plan.Plan] that describes the deployment
// structure and infrastructure.
//
// Blueprints are encoded as JSON. Use [Encode] and [Decode] to convert
// between a [Blueprint] value and its JSON representation. Both functions
// validate the blueprint automatically.
//
// Encoding a blueprint:
//
//	bp := &blueprint.Blueprint{
//		Services: []blueprint.Service{{
//			ID:        "hub",
//			Reference: "crucible/hub ^1.0.0",
//			Prefix:    "/api",
//		}},
//	}
//	data, err := blueprint.Encode(bp)
//
// Decoding a blueprint:
//
//	bp, err := blueprint.Decode(data)
//	if err != nil {
//		log.Fatal(err) // malformed JSON or validation failure
//	}
//	for _, svc := range bp.Services {
//		fmt.Println(svc.ID, svc.Reference)
//	}
package blueprint
