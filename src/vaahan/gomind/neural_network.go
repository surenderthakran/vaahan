// Package gomind for a simple Multi Layer Perceptron (MLP) Feed Forward Artificial Neural Network library.
package gomind

import (
	"fmt"
)

// Type NeuralNetwork to describe a single hidden layer MLP feed forward neural network.
type NeuralNetwork struct {
	numberOfInputs int
	hiddenLayer    *layer
	outputLayer    *layer
}

const (
	learningRate = 0.5
)

// NewNeuralNetwork function returns a new NeuralNetwork object.
func NewNeuralNetwork(numberOfInputs, numberOfHiddenNeurons, numberOfOutputs int) (*NeuralNetwork, error) {
	fmt.Println(fmt.Sprintf("numberOfInputs: %d", numberOfInputs))

	fmt.Println(fmt.Sprintf("numberOfHiddenNeurons: %d", numberOfHiddenNeurons))
	hiddenLayer, err := newLayer(numberOfHiddenNeurons, numberOfInputs)
	if err != nil {
		return nil, fmt.Errorf("error creating a hidden layer: %v", err)
	}

	fmt.Println(fmt.Sprintf("numberOfOutputs: %d", numberOfOutputs))
	outputLayer, err := newLayer(numberOfOutputs, numberOfHiddenNeurons)
	if err != nil {
		return nil, fmt.Errorf("error creating output layer: %v", err)
	}

	return &NeuralNetwork{
		numberOfInputs: numberOfInputs,
		hiddenLayer:    hiddenLayer,
		outputLayer:    outputLayer,
	}, nil
}

// Train function trains the neural network using the given set of inputs and outputs.
func (network *NeuralNetwork) Train(trainingInput, trainingOutput []float64) {
	// fmt.Println("trainingInput: %v", trainingInput)
	// fmt.Println("trainingOutput: %v", trainingOutput)
	outputs := network.CalculateOutput(trainingInput)
	// fmt.Println("outputs: %v", outputs)
	// fmt.Println("------------------------------------------------------------")
	network.updateOutputLayerWeight(outputs, trainingOutput)
	// fmt.Println("------------------------------------------------------------")
	network.updateHiddenLayerWeight()
}

// CalculateOutput function returns the output array from the neural network for the given
// input array based on the current weights.
func (network *NeuralNetwork) CalculateOutput(input []float64) []float64 {
	hiddenOutput := network.hiddenLayer.calculateOutput(input)
	return network.outputLayer.calculateOutput(hiddenOutput)
}

// GetLastOutput function returns the array of last output computed by the network.
func (network *NeuralNetwork) GetLastOutput() []float64 {
	var output []float64
	for _, neuron := range network.outputLayer.neurons {
		output = append(output, neuron.output)
	}
	return output
}

// updateOutputLayerWeight function updates the weights from the hidden layer to the output layer,
// after calculating how much each weight affects the total error in the final output of the network.
// i.e. the partial differential of error with respect to the weight. ∂Error/∂Weight.
//
// By applying the chain rule, https://en.wikipedia.org/wiki/Chain_rule
// ∂TotalError/∂OutputNeuronWeight = ∂TotalError/∂TotalNetInputToOutputNeuron * ∂TotalNetInputToOutputNeuron/∂OutputNeuronWeight
func (network *NeuralNetwork) updateOutputLayerWeight(outputs, targetOutputs []float64) {
	// fmt.Println("inside updateOutputLayerWeight()")
	for neuronIndex, neuron := range network.outputLayer.neurons {
		// fmt.Println(neuronIndex)
		// fmt.Println(neuron)

		// Since a neuron has only one total net input and one output, we need to calculate
		// the partial derivative of error with respect to the total net input (∂TotalError/∂TotalNetInputToOutputNeuron) only once.
		//
		// The total error of the network is a sum of errors in all the output neurons.
		// ex: Total Error = error1 + erro2 + error3 + ...
		// But when calculating the partial derivative of the total error with respect to the total net input
		// of only one output neuron, we need to find partial derivative of only the corresponding neuron's error because
		// the errors due to other neurons would be constant for it and their derivative wouldn't matter.
		pdErrorWrtTotalNetInputOfOutputNeuron := neuron.calculatePdErrorWrtTotalNetInputOfOutputNeuron(targetOutputs[neuronIndex])
		// fmt.Println("pdErrorWrtTotalNetInputOfOutputNeuron: %v", pdErrorWrtTotalNetInputOfOutputNeuron)

		for weightIndex, weight := range neuron.weights {
			// fmt.Println("\nweight: %v", weight)

			// For each weight of the neuron we calculate the partial derivative of
			// total net input with respect to the weight i.e. ∂TotalNetInputToOutputNeuron/∂OutputNeuronWeight.
			pdTotalNetInputWrtWeight := neuron.calculatePdTotalNetInputWrtWeight(weightIndex)
			// fmt.Println("pdTotalNetInputWrtWeight: %v", pdTotalNetInputWrtWeight)

			// Finally, the partial derivative of error with respect to the output neuron weight is:
			// ∂TotalError/∂OutputNeuronWeight = ∂TotalError/∂TotalNetInputToOutputNeuron * ∂TotalNetInputToOutputNeuron/∂OutputNeuronWeight
			pdErrorWrtWeight := pdErrorWrtTotalNetInputOfOutputNeuron * pdTotalNetInputWrtWeight
			// fmt.Println("pdErrorWrtWeight: %v", pdErrorWrtWeight)

			// Now that we know how much the output neuron's weight affects the error in the output, we adjust the weight
			// by subtracting the affect from the current weight after multiplying it with the learning rate.
			// The learning rate is a constant value chosen at fro a network to control the correction in
			// a networks weight based on a sample.
			weight -= learningRate * pdErrorWrtWeight
			// fmt.Println("NewWeight: %v", weight)

			neuron.weights[weightIndex] = weight
		}
	}
}

// updateHiddenLayerWeight function updates the weights from the input layer to the hidden layer,
// after calculating how much each weight affects the error in the final output of the network.
// i.e. the partial differential of error with respect to the weight. ∂Error/∂Weight.
//
// By applying the chain rule, https://en.wikipedia.org/wiki/Chain_rule
// ∂TotalError/∂HiddenNeuronWeight = ∂TotalError/∂HiddenNeuronOutput * ∂HiddenNeuronOutput/∂TotalNetInputToHiddenNeuron * ∂TotalNetInputToHiddenNeuron/∂HiddenNeuronWeight
func (network *NeuralNetwork) updateHiddenLayerWeight() {
	// fmt.Println("inside updateHiddenLayerWeight()")

	// First we calculate the derivative of total error with respect to the output of each hidden neuron.
	// i.e. ∂TotalError/∂HiddenNeuronOutput.
	for neuronIndex, neuron := range network.hiddenLayer.neurons {
		// Since the total error is a summation of errors in each output neuron's output, we need to calculate the
		// derivative of error in each output neuron with respect to the output of the hidden neuron and add them.
		// i.e. ∂TotalError/∂HiddenNeuronOutput = ∂Error1/∂HiddenNeuronOutput + ∂Error2/∂HiddenNeuronOutput + ...
		dErrorWrtOutputOfHiddenNeuron := float64(0)
		for _, outputNeuron := range network.outputLayer.neurons {
			// The partial derivative of an output neuron's output's error with respect to the output of the hidden neuron can be expressed as:
			// ∂Error/∂HiddenNeuronOutput = ∂Error/∂TotalNetInputToOutputNeuron * ∂TotalNetInputToOutputNeuron/∂HiddenNeuronOutput
			//
			// We already have partial derivative of output neuron's error with respect to its total net input for each neuron from previous calculations.
			// The partial derivative of total net input to output neuron with respect to the current hidden neuron (∂TotalNetInputToOutputNeuron/∂HiddenNeuronOutput),
			// is the weight from the current hidden neuron to the current output neuron.
			dErrorWrtOutputOfHiddenNeuron += outputNeuron.pdErrorWrtTotalNetInputOfOutputNeuron * outputNeuron.weights[neuronIndex]
		}

		// We calculate the derivative of hidden neuron outout with respect to total net input to hidden neuron,
		// dHiddenNeuronOutput/dTotalNetInputToHiddenNeuron
		dHiddenNeuronOutputWrtTotalNetInputToHiddenNeuron := neuron.calculateDerivativeOutputWrtTotalNetInput()

		// Next the partial derivative of error with respect to the total net input of the hidden neuron is:
		// ∂TotalError/∂TotalNetInputToHiddenNeuron = ∂TotalError/∂HiddenNeuronOutput * dHiddenNeuronOutput/dTotalNetInputToHiddenNeuron
		pdErrorWrtTotalNetInput := dErrorWrtOutputOfHiddenNeuron * dHiddenNeuronOutputWrtTotalNetInputToHiddenNeuron
		// fmt.Println("pdErrorWrtTotalNetInput: %v", pdErrorWrtTotalNetInput)

		for weightIndex, weight := range neuron.weights {
			// fmt.Println("\nweight: %v", weight)

			// For each weight of the neuron we calculate the partial derivative of
			// total net input with respect to the weight i.e. ∂TotalNetInputToHiddenNeuron/∂HiddenNeuronWeight
			pdTotalNetInputWrtWeight := neuron.calculatePdTotalNetInputWrtWeight(weightIndex)
			// fmt.Println("pdTotalNetInputWrtWeight: %v", pdTotalNetInputWrtWeight)

			// Finally, the partial derivative of total error with respect to the hidden neuron weight is:
			// ∂TotalError/∂HiddenNeuronWeight = ∂TotalError/∂TotalNetInputToHiddenNeuron * ∂TotalNetInputToHiddenNeuron/∂HiddenNeuronWeight
			pdErrorWrtWeight := pdErrorWrtTotalNetInput * pdTotalNetInputWrtWeight
			// fmt.Println("pdErrorWrtWeight: %v", pdErrorWrtWeight)

			// Now that we know how much the hidden neuron's weight affects the error in the output, we adjust the weight
			// by subtracting the affect from the current weight after multiplying it with the learning rate.
			// The learning rate is a constant value chosen at fro a network to control the correction in
			// a networks weight based on a sample.
			weight -= learningRate * pdErrorWrtWeight
			// fmt.Println("NewWeight: %v", weight)

			neuron.weights[weightIndex] = weight
		}
	}
}

// CalculateTotalError computes and returns the total error for the given training set.
func (network *NeuralNetwork) CalculateTotalError(trainingSet [][][]float64) float64 {
	totalError := float64(0)
	for _, set := range trainingSet {
		output := network.CalculateOutput(set[0])
		_ = output // we don't need output here.
		totalError += network.CalculateError(set[1])
	}
	return totalError
}

// CalculateError function generates the error value for the given target output against the network's last output.
func (network *NeuralNetwork) CalculateError(targetOutput []float64) float64 {
	error := float64(0)
	for index, neuron := range network.outputLayer.neurons {
		error += neuron.calculateError(targetOutput[index])
	}
	return error
}

// Describe function prints the current state of the neural network and its components.
func (network *NeuralNetwork) Describe() {
	fmt.Println("Hidden Layer:")
	network.hiddenLayer.describe()
	fmt.Println("\nOutput Layer:")
	network.outputLayer.describe()
}
