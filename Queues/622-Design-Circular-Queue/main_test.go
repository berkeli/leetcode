package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Example 1:

// Input
// ["MyCircularQueue", "enQueue", "enQueue", "enQueue", "enQueue", "Rear", "isFull", "deQueue", "enQueue", "Rear"]
// [[3], [1], [2], [3], [4], [], [], [], [4], []]
// Output
// [null, true, true, true, false, 3, true, true, true, 4]

// Explanation
// MyCircularQueue myCircularQueue = new MyCircularQueue(3);
// myCircularQueue.enQueue(1); // return True
// myCircularQueue.enQueue(2); // return True
// myCircularQueue.enQueue(3); // return True
// myCircularQueue.enQueue(4); // return False
// myCircularQueue.Rear();     // return 3
// myCircularQueue.isFull();   // return True
// myCircularQueue.deQueue();  // return True
// myCircularQueue.enQueue(4); // return True
// myCircularQueue.Rear();     // return 4

func TestMyCircularQueue(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		myCircularQueue := Constructor(3)
		boolVal := myCircularQueue.EnQueue(1) // return True
		require.True(t, boolVal)

		boolVal = myCircularQueue.EnQueue(2) // return True
		require.True(t, boolVal)

		boolVal = myCircularQueue.EnQueue(3) // return True
		require.True(t, boolVal)

		boolVal = myCircularQueue.EnQueue(4) // return False
		require.False(t, boolVal)

		val := myCircularQueue.Rear() // return 3
		require.Equal(t, 3, val)

		boolVal = myCircularQueue.IsFull() // return True
		require.True(t, boolVal)

		boolVal = myCircularQueue.DeQueue() // return True
		require.True(t, boolVal)

		boolVal = myCircularQueue.EnQueue(4) // return True
		require.True(t, boolVal)

		val = myCircularQueue.Rear() // return 4
		require.Equal(t, 4, val)
	})
}
