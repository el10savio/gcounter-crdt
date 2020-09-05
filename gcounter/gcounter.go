package gcounter

// package gcounter implements the GCounter CRDT data type along with the functionality to
// GetCount & Increment the count in the GCounter. It also provides the functionality to
// merge multiple GCounters together and a utility function to clear a GCounter used in tests

// GCounter is the GCounter CRDT data type
type GCounter struct {
	// Count stores the values of count
	// for each node in the cluster
	Count map[string]int `json:"count"`
}

const (
	// singleNodeName is the node name assigned to
	// the node when it is the only node
	// present in the cluster
	singleNodeName = "node"
)

// Initialize returns a new empty GCounter
func Initialize(node string) GCounter {
	// Set the node name as singleNodeName
	// in case of single node
	if node == "" {
		node = singleNodeName
	}

	// Initialize count map and initialize
	// the count for our node as 0
	countMap := make(map[string]int)
	countMap[node] = 0

	// Return the initialized GCounter
	return GCounter{Count: countMap}
}

// Increment increases the count for
// a given node in the GCounter by 1
func (gcounter GCounter) Increment(node string) map[string]int {
	// Set the node name as singleNodeName
	// in case of single node
	if node == "" {
		node = singleNodeName
	}

	// Increment the GCounter Count
	gcounter.Count[node]++

	// Return the updated GCounter
	// Count value
	return gcounter.Count
}

// GetCount returns the GCounter's Count value
func (gcounter GCounter) GetCount() map[string]int {
	return gcounter.Count
}

// GetTotal returns the sum of the counts
// of all the nodes in GCounter.Count
func (gcounter GCounter) GetTotal() int {
	// Initialize the total count as 0
	total := 0

	// Iterative over teh count map
	// and sum each node's count
	for _, count := range gcounter.Count {
		total += count
	}

	// Return the total count
	return total
}

// Merge combines multiple GCounters into a single GCounter
// for the same node in multiple GCounters the merged
// node's count is the max count obtained
func Merge(GCounters ...GCounter) GCounter {
	// Initialize the merged GCounter
	gcounterMerged := GCounters[0]

	// gcounterMerged = Max(gcounterMerged, gcounterToMergeWith)
	for _, gcounter := range GCounters {
		for node, value := range gcounter.Count {
			gcounterMerged.Count[node] = Max(gcounterMerged.Count[node], value)
		}
	}

	// Return the merged GCounter
	return gcounterMerged
}

// SetCount is a utility function  used in tests
// that assigns the GCounter's Count
// value to a specified value
func (gcounter GCounter) SetCount(node string, value int) map[string]int {
	// Set the node name as singleNodeName
	// in case of single node
	if node == "" {
		node = singleNodeName
	}

	// Set the count of the node
	gcounter.Count[node] = value

	// Return the updated GCounter Count
	return gcounter.Count
}

// Clear is a utility function used for tests that clears the GCounter
func (gcounter GCounter) Clear(node string) map[string]int {
	// Iterate over the GCounter Count
	// map and delete each entry
	for node := range gcounter.Count {
		delete(gcounter.Count, node)
	}

	// Re-Initialize the GCounter
	gcounter = Initialize(node)

	// Return the Cleared GCounter
	return gcounter.Count
}

// Max computes the maximum of
// 2 int values passed to it
func Max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	}
	return value2
}
