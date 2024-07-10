package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackOperations(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// Test if push worked correctly.
	peekedValue, err := stack.Peek()
	assert.NoError(t, err, "Unexpected error while peeking.")
	assert.Equal(t, 3, peekedValue, "Push did not work correctly. Peeked value Mismatch.")

	// Test Peek.
	poppedValue, err := stack.Pop()
	assert.NoError(t, err, "Unexpected error while popping.")
	assert.Equal(t, 3, poppedValue, "Popped value mismatch.")

	/*---------------- Tests for an empty stack.-----------------------*/
	for i := 0; i < 2; i++ {
		_, _ = stack.Pop()
	}

	// Test popping off an empty stack.
	_, err = stack.Pop()
	assert.Error(t, err, "Expected a pop Error as the stack is empty.")

	// Test peeking an empty stack
	_, err = stack.Peek()
	assert.Error(t, err, "Expected a peek Error as the stack is empty.")
}
