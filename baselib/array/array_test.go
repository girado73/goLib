package array

import "testing"

func TestEnd(t *testing.T) {
	got := End([]string{"Hello", "World", "Go"})
	want := "Go"

	if got != want {
		t.Errorf("End() = %v, want %v", got, want)
	}
}

func TestCompareSuccess(t *testing.T) {
	got := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 3, 4, 5}

	comp := Compare(got, want)

	if !comp {
		t.Errorf("Compare(%v, %v) = %t, want true", got, want, comp)
	}
}

func TestCompareFail(t *testing.T) {
	got := []int{1, 2, 3, 4, 5}
	want := []int{1, 2, 4, 4, 5}

	comp := Compare(got, want)

	if comp {
		t.Errorf("Compare(%v, %v) = %t, want false", got, want, comp)
	}
}

func TestFilter(t *testing.T) {
	fn := func(n int) bool {
		return n%2 == 0
	}

	got := Filter([]int{1, 2, 3, 4, 5}, fn)
	want := []int{2, 4}

	if !Compare(got, want) {
		t.Errorf("End() = %v, want %v", got, want)
	}
}

func TestMap(t *testing.T) {
	fn := func(n int) int {
		return n + 1
	}

	got := Map([]int{1, 2, 3, 4, 5}, fn)
	want := []int{2, 3, 4, 5, 6}

	if !Compare(got, want) {
		t.Errorf("Map() = %v, want %v", got, want)
	}
}

func TestMapParallel(t *testing.T) {
	fn := func(n int) int {
		return n + 1
	}

	got := MapParallel([]int{1, 2, 3, 4, 5}, fn)
	want := []int{2, 3, 4, 5, 6}

	if !Compare(got, want) {
		t.Errorf("MapParallel() = %v, want %v", got, want)
	}
}

func TestMerge(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}

	got := Merge(arr1, arr2)
	want := []any{1, 2, 3, 4, 5, 6}

	if !Compare(got, want) {
		t.Errorf("Merge() = %v, want %v", got, want)
	}
}

func TestMergeMultiple(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 5, 6}
	arr3 := []int{7, 8, 9}

	got := MergeMultiple(arr1, arr2, arr3)
	want := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	if !Compare(got, want) {
		t.Errorf("MergeMultiple() = %v, want %v", got, want)
	}
}

func TestAll(t *testing.T) {
	got := All([]int{2, 4, 6, 8}, func(n int) bool {
		return n%2 == 0
	})
	want := true

	if got != want {
		t.Errorf("All() = %v, want %v", got, want)
	}
}

func TestAny(t *testing.T) {
	got := Any([]int{1, 2, 3, 4, 5}, func(n int) bool { return n == 3 })
	want := true

	if got != want {
		t.Errorf("Any() = %v, want %v", got, want)
	}
}

func TestSearch(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	value := 3

	got := Search(arr, value)
	want := 2 // index of value 3

	if got != want {
		t.Errorf("Search() = %d, want %d", got, want)
	}
}

func TestPop(t *testing.T) {
	got, gotArr := Pop([]int{1, 2, 3, 4, 5})
	want := 5
	wantArr := []int{1, 2, 3, 4}

	if got != want && !Compare(gotArr, wantArr) {
		t.Errorf("Pop() = %d, %v; want %d, %v", got, gotArr, want, wantArr)
	}
}

func TestPopEmpty(t *testing.T) {
	got, gotArr := Pop([]int{})
	var want int

	if got != want || len(gotArr) != 0 {
		t.Errorf("Pop() on empty array = %d, %v; want %d, %v", got, gotArr, want, []int{})
	}
}

func TestRand(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	got := Rand(arr)

	if got < 1 || got > 5 {
		t.Errorf("Rand() = %d; want value between 1 and 5", got)
	}
}

func TestReduce(t *testing.T) {
	fn := func(acc int, n int) int {
		return acc + n
	}

	got := Reduce([]int{1, 2, 3, 4, 5}, fn, 0)
	want := 15

	if got != want {
		t.Errorf("Reduce() = %d, want %d", got, want)
	}
}

func TestReverse(t *testing.T) {
	got := Reverse([]int{1, 2, 3, 4, 5})
	want := []int{5, 4, 3, 2, 1}

	if !Compare(got, want) {
		t.Errorf("Reverse() = %v, want %v", got, want)
	}
}

func TestProduct(t *testing.T) {
	got := Product([]int{1, 2, 3, 4, 5})
	want := 120 // 1*2*3*4*5

	if got != want {
		t.Errorf("Product() = %d, want %d", got, want)
	}
}

func TestProductParallel(t *testing.T) {
	got := ProductParallel([]int{1, 2, 3, 4, 5})
	want := 120 // 1*2*3*4*5

	if got != want {
		t.Errorf("ProductParallel() = %d, want %d", got, want)
	}
}

func TestDifference(t *testing.T) {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4}

	got := Difference(arr1, arr2)
	want := []int{1, 2, 5}

	if !Compare(got, want) {
		t.Errorf("Difference() = %v, want %v", got, want)
	}
}
