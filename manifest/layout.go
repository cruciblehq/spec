package manifest

import "path/filepath"

// File and directory names in a resource project.
const (
	manifestFile   = "crucible.yaml"   // Resource manifest.
	planFile       = "plan.yaml"       // Deployment plan produced by blueprint builds.
	imageFile      = "image.tar"       // OCI image archive produced by runtime and service builds.
	widgetMainFile = "index.js"        // JavaScript bundle produced by widget builds.
	packageFile    = "package.tar.zst" // Packaged resource archive.
	buildDirName   = "build"           // Build output directory.
	distDirName    = "dist"            // Distributable archive directory.
)

// Returns the path to the manifest file in dir.
func Path(dir string) string {
	return filepath.Join(dir, manifestFile)
}

// Returns the path to the deployment plan in dir.
func PlanPath(dir string) string {
	return filepath.Join(dir, planFile)
}

// Returns the path to the OCI image archive in dir.
func ImagePath(dir string) string {
	return filepath.Join(dir, imageFile)
}

// Returns the path to the widget JavaScript bundle in dir.
func WidgetMainPath(dir string) string {
	return filepath.Join(dir, widgetMainFile)
}

// Returns the build output directory of the project in dir.
//
// Its contents depend on the resource type. Widgets produce a JavaScript
// bundle, runtimes and services an OCI image, and blueprints a plan.
func BuildDir(dir string) string {
	return filepath.Join(dir, buildDirName)
}

// Returns the path to the packaged archive of the project in dir.
func PackagePath(dir string) string {
	return filepath.Join(dir, distDirName, packageFile)
}
