package gomind

import (
	"fmt"
	"math"
)

type Neuron struct {
	weights []float64
}

func NewNeuron(weights []float64) (*Neuron, error) {
	if len(weights) == 0 {
		return nil, fmt.Errorf("unable to create neuron without any weights")
	}
	return &Neuron{
		weights: weights,
	}, nil
}

func (neuron *Neuron) GetOutput(input []float64) float64 {
	return neuron.squash(neuron.getNetInput(input))
}

func (neuron *Neuron) getNetInput(input []float64) float64 {
	netInput := float64(0)
	for i := range input {
		netInput += input[i] * neuron.weights[i]
	}
	return netInput
}

func (neuron *Neuron) squash(input float64) float64 {
	return 1.0 / (1.0 + math.Exp(-input))
}
