package manifest

// Names of the files a built resource holds next to its manifest.
const (
	ManifestFile   = "crucible.yaml" // Resource manifest.
	PlanFile       = "plan.yaml"     // Deployment plan produced by blueprint builds.
	ImageFile      = "image.tar"     // OCI image archive produced by runtime and service builds.
	WidgetMainFile = "index.js"      // JavaScript bundle produced by widget builds.
)
