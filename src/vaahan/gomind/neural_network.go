package gomind

import (
	"fmt"
)

type NeuralNetwork struct {
	numberOfInputs int
	hiddenLayer    *Layer
	outputLayer    *Layer
}

func NewNeuralNetwork(numberOfInputs, numberOfHiddenNeurons, numberOfOutputs int) (*NeuralNetwork, error) {
	hiddenLayer, err := NewLayer(numberOfHiddenNeurons, numberOfInputs)
	if err != nil {
		return nil, fmt.Errorf("error creating a hidden layer: %v", err)
	}

	outputLayer, err := NewLayer(numberOfOutputs, numberOfHiddenNeurons)
	if err != nil {
		return nil, fmt.Errorf("error creating output layer: %v", err)
	}

	return &NeuralNetwork{
		numberOfInputs: numberOfInputs,
		hiddenLayer:    hiddenLayer,
		outputLayer:    outputLayer,
	}, nil
}

func (nn *NeuralNetwork) Train(trainingInput, trainingOutput []float64) {
	fmt.Println("trainingInput: %v", trainingInput)
	fmt.Println("trainingOutput: %v", trainingOutput)
	output := nn.GetOutput(trainingInput)
	fmt.Println("output: %v", output)
}

func (nn *NeuralNetwork) GetOutput(input []float64) []float64 {
	hiddenOutput := nn.hiddenLayer.GetOutput(input)
	return nn.outputLayer.GetOutput(hiddenOutput)
}
