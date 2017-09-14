package gomind

import (
	"fmt"
)

type NeuralNetwork struct {
	numberOfInputs int
	hiddenLayer    *Layer
	outputLayer    *Layer
}

const (
	LearningRate = 0.5
)

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
	// fmt.Println("trainingInput: %v", trainingInput)
	// fmt.Println("trainingOutput: %v", trainingOutput)
	outputs := nn.CalculateOutput(trainingInput)
	// fmt.Println("outputs: %v", outputs)
	// fmt.Println("------------------------------------------------------------")
	nn.UpdateOutputLayerWeight(outputs, trainingOutput)
	// fmt.Println("------------------------------------------------------------")
	nn.UpdateHiddenLayerWeight()
}

func (nn *NeuralNetwork) CalculateOutput(input []float64) []float64 {
	hiddenOutput := nn.hiddenLayer.CalculateOutput(input)
	return nn.outputLayer.CalculateOutput(hiddenOutput)
}

func (nn *NeuralNetwork) UpdateOutputLayerWeight(outputs, targetOutputs []float64) {
	// fmt.Println("inside UpdateOutputLayerWeight()")
	for neuronIndex, neuron := range nn.outputLayer.neurons {
		// fmt.Println(neuron)

		pdErrorWrtTotalNetInput := neuron.calculatePdErrorWrtTotalNetInput(targetOutputs[neuronIndex])
		// fmt.Println("pdErrorWrtTotalNetInput: %v", pdErrorWrtTotalNetInput)

		for weightIndex, weight := range neuron.weights {
			// fmt.Println("\nweight: %v", weight)

			pdTotalNetInputWrtWeight := neuron.calculatePdTotalNetInputWrtWeight(weightIndex)
			// fmt.Println("pdTotalNetInputWrtWeight: %v", pdTotalNetInputWrtWeight)

			pdErrorWrtWeight := pdErrorWrtTotalNetInput * pdTotalNetInputWrtWeight
			// fmt.Println("pdErrorWrtWeight: %v", pdErrorWrtWeight)

			weight -= LearningRate * pdErrorWrtWeight
			// fmt.Println("NewWeight: %v", weight)

			neuron.weights[weightIndex] = weight
		}
	}
}

func (nn *NeuralNetwork) UpdateHiddenLayerWeight() {
	// fmt.Println("inside UpdateHiddenLayerWeight()")
	for neuronIndex, neuron := range nn.hiddenLayer.neurons {
		// fmt.Println(neuron)

		dErrorWrtOutput := float64(0)
		for _, outputNeuron := range nn.outputLayer.neurons {
			dErrorWrtOutput += outputNeuron.pdErrorWrtTotalNetInput * outputNeuron.weights[neuronIndex]
		}

		pdTotalNetInputWrtInput := neuron.calculatePdTotalNetInputWrtInput()

		pdErrorWrtTotalNetInput := dErrorWrtOutput * pdTotalNetInputWrtInput
		// fmt.Println("pdErrorWrtTotalNetInput: %v", pdErrorWrtTotalNetInput)

		for weightIndex, weight := range neuron.weights {
			// fmt.Println("\nweight: %v", weight)

			pdTotalNetInputWrtWeight := neuron.calculatePdTotalNetInputWrtWeight(weightIndex)
			// fmt.Println("pdTotalNetInputWrtWeight: %v", pdTotalNetInputWrtWeight)

			pdErrorWrtWeight := pdErrorWrtTotalNetInput * pdTotalNetInputWrtWeight
			// fmt.Println("pdErrorWrtWeight: %v", pdErrorWrtWeight)

			weight -= LearningRate * pdErrorWrtWeight
			// fmt.Println("NewWeight: %v", weight)

			neuron.weights[weightIndex] = weight
		}
	}
}

func (nn *NeuralNetwork) CalculateError(targetOutput []float64) float64 {
	error := float64(0)
	for index, neuron := range nn.outputLayer.neurons {
		error += neuron.CalculateError(targetOutput[index])
	}
	return error
}

func (nn *NeuralNetwork) State() {
	fmt.Println("Hidden Neurons")
	for _, hiddenNeuron := range nn.hiddenLayer.neurons {
		fmt.Println(hiddenNeuron)
	}
	fmt.Println("Output Neurons")
	for _, outputNeuron := range nn.outputLayer.neurons {
		fmt.Println(outputNeuron)
	}
}
