package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cuelang.org/go/cue"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/tools/flow"
)

func main() {
	ctx := cuecontext.New()
	v := ctx.CompileString(`
	a: {
		input: "world"
		output: string
	}
	b: {
		input: a.output
		output: string
	}
	`)
	if err := v.Err(); err != nil {
		log.Fatal(err)
	}
	controller := flow.New(nil, v, ioTaskFunc)
	logGraph(controller)
	if err := controller.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
	logGraph(controller)
	// Output:
	// setting a.output to "hello world"
	// setting b.output to "hello hello world"
}

func logGraph(controller *flow.Controller) {
	graph := mermaidGraph(controller)
	log.Printf("Graph: \n%v", graph)
}

func ioTaskFunc(v cue.Value) (flow.Runner, error) {
	inputPath := cue.ParsePath("input")

	input := v.LookupPath(inputPath)
	if !input.Exists() {
		return nil, nil
	}

	return flow.RunnerFunc(func(t *flow.Task) error {
		inputVal, err := t.Value().LookupPath(inputPath).String()
		if err != nil {
			return fmt.Errorf("input not of type string")
		}

		outputVal := fmt.Sprintf("hello %s", inputVal)
		fmt.Printf("setting %s.output to %q\n", t.Path(), outputVal)

		return t.Fill(map[string]string{
			"output": outputVal,
		})
	}), nil
}

// We need to escape quotes in the path, per
// https://mermaid-js.github.io/mermaid/#/flowchart?id=entity-codes-to-escape-characters
// This also requires that we escape the quoting character #.
var mermaidQuote = strings.NewReplacer("#", "#35;", `"`, "#quot;")

// mermaidGraph generates a mermaid graph of the current state. This can be
// pasted into https://mermaid-js.github.io/mermaid-live-editor/ for
// visualization.
func mermaidGraph(c *flow.Controller) string {
	w := &strings.Builder{}
	fmt.Fprintln(w, "graph TD")
	for i, t := range c.Tasks() {
		path := mermaidQuote.Replace(t.Path().String())
		fmt.Fprintf(w, "  t%d(\"%s [%s]\")\n", i, path, t.State())
		for _, t := range t.Dependencies() {
			fmt.Fprintf(w, "  t%d-->t%d\n", i, t.Index())
		}
	}
	return w.String()
}
