package main

import (
    "fmt"
    "sort"
    "os"
    "strings"
    "log"
    "github.com/Myriad-Dreamin/go-ves/types"
    "github.com/Myriad-Dreamin/minimum-lib/sugar"
    "flag"
)

type X struct {
    c types.Code
    d string
}

type Xs []X

func (a Xs) Len() int           { return len(a) }
func (a Xs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Xs) Less(i, j int) bool { return a[i].c < a[j].c }


func getSlice() Xs {
    xs := make(Xs, 0, len(types.CodeDesc))
    for k, v := range types.CodeDesc {
        xs = append(xs, X{c: k,d: v})
    }
    sort.Sort(xs)
    return xs
}

var (
    tl = flag.String("target-language", "python", "specify target language")
    o = flag.String("o", "out", "specify name")
)

func init() {
    flag.Parse()
}

func main() {
    switch *tl {
    case "python", "python2", "python3", "py", "py2", "py3":
        toPython(getSlice())
    default:
        log.Fatal("unknown target-language", *tl)
    }
}

func toPython(xs Xs) {
    op := *o
    if !strings.HasSuffix(op, ".py") {
        op += ".py"
    }
    sugar.WithWriteFile(func (f *os.File) {
    f.WriteString("import enum\n\n\nclass Code(enum.Enum):\n")
    for _, x := range xs {
        f.WriteString(fmt.Sprintf(`    %v = %v
`, x.d[4:], x.c))
    }

    f.WriteString("\n\ncode_desc = {\n")
    for _, x := range xs {
        f.WriteString(fmt.Sprintf(`    %v: '%v',
`, x.c, x.d))
    }

    f.WriteString("}\n\ncode_revert_desc = {\n")
    for _, x := range xs {
        f.WriteString(fmt.Sprintf(`    '%v': %v,
`, x.d, x.c))
    }
    f.WriteString(`}


class VESError(Exception):
    pass


`)

    for _, x := range xs {
        f.WriteString(fmt.Sprintf(`class %v(VESError):
    def __init__(self, msg):
        self.msg = msg
        self.code = Code.%v
    
    def __str__(self):
        return f'<code:{self.code}, error:{self.msg}>'


`, x.d[4:], x.d[4:]))
}
	f.WriteString(`def response_to_error(r):
    return eval(code_desc[r.code][4:])(r.get_error())

`)

    }, op)
}
