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
	fmt.Println(fmt.Sprintf("numberOfInputs: %d", numberOfInputs))

	fmt.Println(fmt.Sprintf("numberOfHiddenNeurons: %d", numberOfHiddenNeurons))
	hiddenLayer, err := NewLayer(numberOfHiddenNeurons, numberOfInputs)
	if err != nil {
		return nil, fmt.Errorf("error creating a hidden layer: %v", err)
	}

	fmt.Println(fmt.Sprintf("numberOfOutputs: %d", numberOfOutputs))
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
	nn.UpdateOutputLayerWeight()
	nn.UpdateHiddenLayerWeight()
}

func (nn *NeuralNetwork) GetOutput(input []float64) []float64 {
	hiddenOutput := nn.hiddenLayer.GetOutput(input)
	return nn.outputLayer.GetOutput(hiddenOutput)
}

func (nn *NeuralNetwork) UpdateOutputLayerWeight() {
	fmt.Println("inside UpdateOutputLayerWeight()")
	for _, neuron := range nn.outputLayer.neurons {
		fmt.Println(neuron)
		for _, weight := range neuron.weights {
			fmt.Println(weight)
			pdErrorWrtWeight := nn.getPartialDifferentialWithRespectToWeight()
			fmt.Println(pdErrorWrtWeight)
		}
	}
}

func (nn *NeuralNetwork) UpdateHiddenLayerWeight() {
}

func (nn *NeuralNetwork) getPartialDifferentialWithRespectToWeight() float64 {
	return 0
}
