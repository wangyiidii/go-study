package split

import (
	"fmt"
	"reflect"
	"testing"
)

// 测试

func TestSplit(t *testing.T) {
	//got := Split("1,2,3", ",")
	//want := []string{"1", "2", "3"}
	got := Split("我爱你", "爱")
	want := []string{"我", "你"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("want: %v, got:%v", want, got)
	} else {
		fmt.Println("done")
	}
}

func TestSplitBatch(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"chinese": {input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"english": {input: "a,b,c", sep: ",", want: []string{"a", "b", "c"}},
		"digital": {input: "1 2 3", sep: " ", want: []string{"1", "2", "3"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name : %v, want: %v, got:%v", name, tc.want, got)
			} else {
				fmt.Println("done")
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	// b.N不是固定的数

	// BenchmarkSplit-10    	10668562	       102.1 ns/op
	// 核心数					次数					每次操作耗费102.1纳秒

	// BenchmarkSplit-10       12515946                97.69 ns/op            112 B/op             3 allocs/op
	// 核心数				   次数					   每次操作耗费102.1纳秒	  每次操作消耗的内存		每次操作会做三次内存申请

	for i := 0; i < b.N; i++ {
		Split("a,b,c", ",")
	}
}
