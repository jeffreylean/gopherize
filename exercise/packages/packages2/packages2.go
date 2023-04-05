package main

// Make it compile!
// In Go only exported code elements in a package can be used in the source file which imports the packge.
// An exported code element uses exported identifier as its name, meaning if the first character of the identifier
// is in upper case letter, then the identifer is exported, otherwise it is private hence not being export.
import(
"github.com/jeffreylean/gopherize/exercise/packages/packages2/lib"
)

// I AM NOT DONE

func main(){
    a,b:=1,2

    println("The sum is",lib.sum(a,b))
}


