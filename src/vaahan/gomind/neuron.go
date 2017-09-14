package gomind

import (
	"fmt"
	"math"
)

type Neuron struct {
	inputs                  []float64
	weights                 []float64
	netInput                float64
	output                  float64
	pdErrorWrtTotalNetInput float64
}

func (neuron *Neuron) String() string {
	return fmt.Sprintf("Neuron{\n\tinputs: %v, \n\tweights: %v, \n\tnetInput: %v, \n\toutput: %v}", neuron.inputs, neuron.weights, neuron.netInput, neuron.output)
}

func NewNeuron(weights []float64) (*Neuron, error) {
	if len(weights) == 0 {
		return nil, fmt.Errorf("unable to create neuron without any weights")
	}
	return &Neuron{
		weights: weights,
	}, nil
}

func (neuron *Neuron) CalculateOutput(inputs []float64) float64 {
	neuron.inputs = inputs
	neuron.netInput = neuron.calculateNetInput(neuron.inputs)
	neuron.output = neuron.squash(neuron.netInput)
	return neuron.output
}

func (neuron *Neuron) calculateNetInput(input []float64) float64 {
	netInput := float64(0)
	for i := range input {
		netInput += input[i] * neuron.weights[i]
	}
	return netInput
}

func (neuron *Neuron) squash(input float64) float64 {
	return 1.0 / (1.0 + math.Exp(-input))
}

func (neuron *Neuron) calculatePdErrorWrtTotalNetInput(targetOutput float64) float64 {
	pdErrorWrtOutput := neuron.calculatePdErrorWrtOutput(targetOutput)
	pdTotalNetInputWrtInput := neuron.calculatePdTotalNetInputWrtInput()
	neuron.pdErrorWrtTotalNetInput = pdErrorWrtOutput * pdTotalNetInputWrtInput
	return neuron.pdErrorWrtTotalNetInput
}

func (neuron *Neuron) calculatePdErrorWrtOutput(targetOutput float64) float64 {
	return -(targetOutput - neuron.output)
}

func (neuron *Neuron) calculatePdTotalNetInputWrtInput() float64 {
	return neuron.output * (1 - neuron.output)
}

func (neuron *Neuron) calculatePdTotalNetInputWrtWeight(index int) float64 {
	return neuron.inputs[index]
}

func (neuron *Neuron) CalculateError(targetOutput float64) float64 {
	return 0.5 * math.Pow(targetOutput-neuron.output, 2)
}
