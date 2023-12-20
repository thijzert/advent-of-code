package aoc23

import (
	"fmt"
	"strings"

	"github.com/thijzert/advent-of-code/ch"
)

func Dec20a(ctx ch.AOContext) (interface{}, error) {
	mm, err := dec20read(ctx)
	if err != nil {
		return nil, err
	}
	ctx.Print(len(mm.modules))

	for i := 0; i < 1000; i++ {
		err = mm.Press(LOW)
		if err != nil {
			return nil, err
		}
		if i < 5 || (i < 50 && i%10 == 9) || (i%100 == 99) {
			ctx.Printf("  after %d presses: low: %d, high: %d", i+1, mm.bus.lowCount, mm.bus.highCount)
		}
	}

	ctx.Printf("Low pulses: %d, high pulses: %d", mm.bus.lowCount, mm.bus.highCount)
	return mm.bus.lowCount * mm.bus.highCount, nil
}

var Dec20b ch.AdventFunc = nil

// func Dec20b(ctx ch.AOContext) (interface{}, error) {
// 	return nil, errNotImplemented
// }

type pulse bool

const LOW pulse = false
const HIGH pulse = true

type pulseBus struct {
	highCount, lowCount int

	queue []struct {
		Source      string
		Destination string
		Value       pulse
	}
	offset int
}

func (p *pulseBus) Yeet(source, destination string, value pulse) {
	if value == HIGH {
		p.highCount++
	} else if value == LOW {
		p.lowCount++
	}

	if p.queue == nil {
		p.queue = make([]struct {
			Source      string
			Destination string
			Value       pulse
		}, 0, 100)
	}
	if p.offset > cap(p.queue)/2 {
		copy(p.queue, p.queue[p.offset:])
		p.queue = p.queue[:len(p.queue)-p.offset]
		p.offset = 0
	}
	p.queue = append(p.queue, struct {
		Source      string
		Destination string
		Value       pulse
	}{source, destination, value})
}

func (p *pulseBus) Yoink() (string, string, pulse) {
	rv := p.queue[p.offset]
	p.offset++
	return rv.Source, rv.Destination, rv.Value
}

func (p *pulseBus) Len() int {
	if p.queue == nil {
		return 0
	}
	return len(p.queue[p.offset:])
}

type module interface {
	handlePulse(string, pulse) []pulse
}

type flipflopModule struct {
	state bool
}

func (ffm *flipflopModule) handlePulse(source string, value pulse) []pulse {
	if value == HIGH {
		return nil
	}
	ffm.state = !ffm.state
	if ffm.state {
		return []pulse{HIGH}
	} else {
		return []pulse{LOW}
	}
}

type conjunctionModule struct {
	inputState map[string]pulse
}

func (cjm *conjunctionModule) handlePulse(source string, value pulse) []pulse {
	if _, ok := cjm.inputState[source]; !ok {
		return nil
	}
	cjm.inputState[source] = value
	rv := LOW
	for _, v := range cjm.inputState {
		if v == LOW {
			rv = HIGH
		}
	}
	return []pulse{rv}
}

type broadcasterModule struct {
}

func (cjm *broadcasterModule) handlePulse(source string, value pulse) []pulse {
	return []pulse{value}
}

type debugModule struct {
}

func (dbm *debugModule) handlePulse(source string, value pulse) []pulse {
	return nil
}

type moduleNetwork struct {
	modules map[string]struct {
		module  module
		outputs []string
	}
	bus *pulseBus
}

func (n *moduleNetwork) Press(value pulse) error {
	n.bus.Yeet("button", "roadcaster", LOW)
	for n.bus.Len() > 0 {
		source, dest, value := n.bus.Yoink()
		mod, ok := n.modules[dest]
		if !ok {
			return fmt.Errorf("module '%s' not found", dest)
		}
		reply := mod.module.handlePulse(source, value)
		for _, v := range reply {
			for _, d := range mod.outputs {
				n.bus.Yeet(dest, d, v)
			}
		}
	}
	return nil
}

func dec20read(ctx ch.AOContext) (*moduleNetwork, error) {
	lines, err := ctx.DataLines("inputs/2023/dec20.txt")
	if err != nil {
		return nil, err
	}

	inputs := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		name := parts[0][1:]
		outputs := strings.Split(parts[1], ", ")

		for _, out := range outputs {
			a := inputs[out]
			a = append(a, name)
			inputs[out] = a
		}
	}

	rv := moduleNetwork{
		modules: make(map[string]struct {
			module  module
			outputs []string
		}),
		bus: &pulseBus{},
	}
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		name := parts[0][1:]
		outputs := strings.Split(parts[1], ", ")
		var mod module
		if parts[0][0] == 'b' {
			mod = &broadcasterModule{}
		} else if parts[0][0] == '%' {
			mod = &flipflopModule{}
		} else if parts[0][0] == '&' {
			m := make(map[string]pulse)
			for _, in := range inputs[name] {
				m[in] = LOW
			}
			mod = &conjunctionModule{inputState: m}
		} else {
			return nil, fmt.Errorf("Unknown module type '%c': '%s'", parts[0][0], line)
		}
		rv.modules[name] = struct {
			module  module
			outputs []string
		}{mod, outputs}
	}

	// Sanity check: make sure all outputs are connected to... _something_
	for name := range inputs {
		if _, ok := rv.modules[name]; !ok {
			rv.modules[name] = struct {
				module  module
				outputs []string
			}{&debugModule{}, nil}
		}
	}

	return &rv, nil
}
