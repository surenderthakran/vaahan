package gomind

import (
	"fmt"
	"math/rand"
)

type Layer struct {
	neurons []*Neuron
}

func NewLayer(numberOfNeurons, numberOfNeuronsInPreviousLayer int) (*Layer, error) {
	if numberOfNeurons <= 0 {
		return nil, fmt.Errorf("%d is not a valid number of neurons", numberOfNeurons)
	}

	var neurons []*Neuron
	for i := 0; i < numberOfNeurons; i++ {
		var weights []float64
		for i := 0; i < numberOfNeuronsInPreviousLayer; i++ {
			weights = append(weights, rand.Float64())
		}
		neuron, err := NewNeuron(weights)
		if err != nil {
			return nil, fmt.Errorf("error creating a neuron: %v", err)
		}
		neurons = append(neurons, neuron)
	}
	return &Layer{
		neurons: neurons,
	}, nil
}

func (layer *Layer) GetOutput(input []float64) []float64 {
	var output []float64
	for _, neuron := range layer.neurons {
		output = append(output, neuron.GetOutput(input))
	}
	return output
}
