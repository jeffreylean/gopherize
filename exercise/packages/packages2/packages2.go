package packages2

// Make it compile!
// In Go, a code package consist of several source files. There source files has to locate in the same folder.
// The source files in a folder (not including subfolders) must belong to the same package.
// So, a folder corresponds to a code package, and vice versa. The folder containing the source files of a code package is called the folder of the package.
// The package name typically are the path to the folder of the package.
import (
	"lib"
)

func main() {
	var a, b int32 = 2, 2
	println("The sum of a and b is", lib.Sum(a, b))
}
