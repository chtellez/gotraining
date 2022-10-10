package queue

import "testing"

func TestAddQueue(t *testing.T){
	q := New(3)
	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			t.Errorf("queue length should be %d, but got %d", i, len(q.items))
		}
		if !q.Append(i){
			t.Errorf("failed to append item %d to queue", i)
		}
	}
	if q.Append(4){
		t.Errorf("Should not be able to append 4 items to queue")
	}
}

func TestNext(t *testing.T){
	q := New(3)
    for i := 0; i < 3; i++ {
        q.Append(i)
    }

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
            t.Errorf("failed to get next item from queue")
		}
		if item != i {
			t.Errorf("got item in wrong order %v, want %v", i, item)
		}
	}
	// Queue is empty at this point
	item, ok := q.Next()
	if ok {
		t.Errorf("should not be anymore itmes in queue, got %v", item)
	}
}




