package gonaut

import (
	"fmt"
	"strings"
)

type AnnotationKind int

const (
	AnnotationKindRouting = iota
)

type Annotation struct {
	kind AnnotationKind
	data interface{}
}

func (a *Annotation) String() string {
	switch a.kind {
	case AnnotationKindRouting:
		return a.data.(*RoutingAnnotationData).String()
	default:
		return ""
	}
}

type RoutingAnnotationData struct {
	method string
	uri    string
}

func (r *RoutingAnnotationData) String() string {
	return fmt.Sprintf("%s %s", r.method, r.uri)
}

func parseAnnotation(a string) *Annotation {
	if !strings.HasPrefix(a, "//gonaut:") {
		return nil
	}

	t := strings.Split(a, ":")
	if len(t) < 2 {
		return nil
	}
	s := t[1]

	if strings.HasPrefix(s, "routing") {
		r := strings.Split(s, " ")
		if len(r) < 3 {
			return nil
		}
		method, uri := strings.ToUpper(strings.TrimSpace(r[1])), strings.TrimSpace(r[2])
		return &Annotation{
			AnnotationKindRouting,
			&RoutingAnnotationData{method, uri},
		}
	}

	return nil
}
