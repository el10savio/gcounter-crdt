package gcounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	// testNode is the node name used
	// for single node GCounter tests
	testNode = "test-node"
)

var (
	gcounter GCounter
)

func init() {
	// Initialize the test GCounter
	gcounter = Initialize(testNode)
}

// TestGetCount checks the basic functionality of GCounter GetCount()
// GetCount() should return all count map entries added to the GCounter
func TestGetCount(t *testing.T) {
	gcounter.Count = gcounter.Increment(testNode)

	expectedCount := map[string]int{testNode: 1}
	actualCount := gcounter.GetCount()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestGetCount_UpdatedValue checks the functionality of GCounter GetCount()
// when values are incremented in the GCounter it should return
// all the count map entries added to the GCounter
func TestGetCount_UpdatedValue(t *testing.T) {
	gcounter.Count = gcounter.Increment(testNode)
	gcounter.Count = gcounter.Increment(testNode)
	gcounter.Count = gcounter.Increment(testNode)

	expectedCount := map[string]int{testNode: 3}
	actualCount := gcounter.GetCount()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestGetCount_NoValue checks the functionality of GCounter GetCount()
// when no values are added to GCounter, it should return an
// initialized empty count map
func TestGetCount_NoValue(t *testing.T) {
	expectedCount := map[string]int{testNode: 0}
	actualCount := gcounter.GetCount()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestSetCount checks the basic functionality of GCounter SetCount()
// SetCount(testNode,value) should return all count map entries updated to the GCounter
func TestSetCount(t *testing.T) {
	gcounter.Count = gcounter.SetCount(testNode, 5)
	gcounter.Count = gcounter.SetCount("testNode2", 7)

	expectedCount := map[string]int{testNode: 5, "testNode2": 7}
	actualCount := gcounter.GetCount()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestSetCount_UpdatedValue checks the functionality of GCounter SetCount(testNode, value)
// when multiple SetCounts are done to GCounter it should return
// the updated count map entries added to the GCounter
func TestSetCount_UpdatedValue(t *testing.T) {
	gcounter.Count = gcounter.SetCount(testNode, 5)
	gcounter.Count = gcounter.SetCount(testNode, 7)
	gcounter.Count = gcounter.SetCount(testNode, 9)

	expectedCount := map[string]int{testNode: 9}
	actualCount := gcounter.GetCount()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestSetCount_NoValue checks the functionality of GCounter SetCount(testNode, value)
// when no values are added into the GCounter, it should return an empty count map
func TestSetCount_NoValue(t *testing.T) {
	gcounter.Count = gcounter.SetCount(testNode, 0)

	expectedCount := map[string]int{testNode: 0}
	actualCount := gcounter.GetCount()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestIncrement checks the basic functionality of GCounter Increment(testNode)
// it should return the GCounter node count incremented by 1
func TestIncrement(t *testing.T) {
	expectedCount := map[string]int{testNode: 1}
	actualCount := gcounter.Increment(testNode)

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestClear checks the basic functionality of GCounter Clear(testNode)
// this utility function it clears all the values in a GCounter
func TestClear(t *testing.T) {
	gcounter.Count = gcounter.Increment(testNode)
	gcounter.Count = gcounter.Increment(testNode)
	gcounter.Count = gcounter.Clear(testNode)

	expectedCount := map[string]int{testNode: 0}
	actualCount := gcounter.GetCount()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestClear_EmptyStore checks the functionality of GCounter Clear(testNode)
// utility function when no values are in it
func TestClear_EmptyStore(t *testing.T) {
	gcounter.Count = gcounter.Clear(testNode)

	expectedCount := map[string]int{testNode: 0}
	actualCount := gcounter.GetCount()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestGetTotal checks the basic functionality of GCounter GetTotal()
// this function should return the total of all nodes in count
func TestGetTotal(t *testing.T) {
	gcounter.Count = gcounter.SetCount("testNode1", 1)
	gcounter.Count = gcounter.SetCount("testNode2", 3)
	gcounter.Count = gcounter.SetCount("testNode3", 5)

	expectedCount := 9
	actualCount := gcounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestGetTotal checks GetTotal function when
// no node counts are present, it should then return 0
func TestGetTotal_EmptyCount(t *testing.T) {
	expectedCount := 0
	actualCount := gcounter.GetTotal()

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestMerge checks the basic functionality of the Merge() function on multiple GCounters
// it returns all the GCounters merged together with unique elements as one single GCounter
func TestMerge(t *testing.T) {
	gcounter1 := GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}}
	gcounter2 := GCounter{map[string]int{"node1": 4, "node2": 6, "node3": 8}}
	gcounter3 := GCounter{map[string]int{"node1": 2, "node2": 4, "node3": 9}}

	gcounterExpected := GCounter{map[string]int{"node1": 4, "node2": 6, "node3": 9}}
	actualCount := Merge(gcounter1, gcounter2, gcounter3)

	assert.Equal(t, gcounterExpected, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestMerge_Empty checks the functionality of the Merge() function on multiple GCounters
// when one GCounters are empty, it returns an empty GCounter followed by an error
func TestMerge_Empty(t *testing.T) {
	gcounter1 := GCounter{}
	gcounter2 := GCounter{}
	gcounter3 := GCounter{}

	expectedCount := GCounter{}
	actualCount := Merge(gcounter1, gcounter2, gcounter3)

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}

// TestMerge_Duplicate checks the functionality of the Merge() function on multiple GCounters
// when duplicate values are passed with the GCounter it returns all the GCounters
// merged together with maximum counts as one single GCounter
func TestMerge_Duplicate(t *testing.T) {
	gcounter1 := GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}}
	gcounter2 := GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}}
	gcounter3 := GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}}

	expectedCount := GCounter{map[string]int{"node1": 3, "node2": 5, "node3": 7}}
	actualCount := Merge(gcounter1, gcounter2, gcounter3)

	assert.Equal(t, expectedCount, actualCount)

	gcounter.Count = gcounter.Clear(testNode)
}
