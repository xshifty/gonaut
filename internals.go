package gonaut

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
	"strings"
)

type routingPreload struct {
	a *RoutingAnnotationData
	p string
	c string
	f string
}

func parseControllerDir(d string) map[string]*ast.Package {
	var fset = token.NewFileSet()
	pack, err := parser.ParseDir(fset, d, nil, parser.ParseComments)
	if err != nil {
		LogFatal("Cannot parse controller directory: %s", err)
	}

	return pack
}

func buildParsingMapFromControllersDir(d string) []interface{} {
	var pmap = make(map[int]interface{})
	var pack = parseControllerDir(d)

	for _, p := range pack {
		for _, f := range p.Files {
			for _, d := range f.Decls {
				if fn, isFn := d.(*ast.FuncDecl); isFn {
					rt := ""
					for _, r := range fn.Recv.List {
						rt = fmt.Sprintf("%s", r.Type)
					}
					if len(rt) > 0 {
						pmap[int(fn.Pos())] = fmt.Sprintf("CONTROLLER,%s.%s:%s", p.Name, rt, fn.Name.String())
					}
				}
			}
			for _, c := range f.Comments {
				for _, l := range c.List {
					if a := parseAnnotation(l.Text); a != nil {
						pmap[int(l.Pos())] = a
					}
				}
			}
		}
	}

	var keys []int
	for k := range pmap {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	var values []interface{}
	for _, k := range keys {
		values = append(values, pmap[k])
	}

	return values
}

func buildRoutingPreload(parsingMap []interface{}) []routingPreload {
	rp := []routingPreload{}
	annotations := []Annotation{}

	for _, v := range parsingMap {
		switch x := v.(type) {
		case *Annotation:
			annotations = append(annotations, Annotation{x.kind, x.data})
			break
		case string:
			sym := ""
			if strings.HasPrefix(x, "CONTROLLER,") {
				sym = strings.Split(x, ",")[1]
				tsym := strings.Split(sym, ".")
				p := tsym[0]
				sign := strings.Split(tsym[1], ":")
				c, f := sign[0], sign[1]
				if len(annotations) > 0 {
					for _, a := range annotations {
						rp = append(rp, routingPreload{a.data.(*RoutingAnnotationData), p, c, f})
					}
				}
				annotations = []Annotation{}
			}
		}
	}

	return rp
}
