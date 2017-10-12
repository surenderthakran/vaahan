package gomind

import (
	"fmt"
	"math/rand"
)

type Layer struct {
	neurons []*Neuron
}

func newLayer(numberOfNeurons, numberOfNeuronsInPreviousLayer int) (*Layer, error) {
	if numberOfNeurons <= 0 {
		return nil, fmt.Errorf("%d is not a valid number of neurons", numberOfNeurons)
	}

	var neurons []*Neuron
	for i := 0; i < numberOfNeurons; i++ {
		var weights []float64
		for i := 0; i < numberOfNeuronsInPreviousLayer; i++ {
			weights = append(weights, rand.Float64())
		}
		fmt.Println(fmt.Sprintf("weights: %v", weights))
		neuron, err := newNeuron(weights)
		if err != nil {
			return nil, fmt.Errorf("error creating a neuron: %v", err)
		}
		neurons = append(neurons, neuron)
	}
	return &Layer{
		neurons: neurons,
	}, nil
}

// calculateOutput function returns the output array from a layer of neurons for an
// array of input for the current set of weights of its neurons.
func (layer *Layer) calculateOutput(input []float64) []float64 {
	var output []float64
	for _, neuron := range layer.neurons {
		output = append(output, neuron.calculateOutput(input))
	}
	return output
}

func (layer *Layer) describe() {
	fmt.Println("Neurons:")
	for _, neuron := range layer.neurons {
		fmt.Println(neuron)
	}
}
