package table

import (
	"reflect"
	"testing"
)

type RecordContext struct {
	Record   []string
	Expected []string
}

func TestTrimRecord(t *testing.T) {
	context := []RecordContext{
		{
			[]string{
				"<root_account>",
				"dont care",
				"dont care",
				"not_supported",
				"dont care",
				"dont care",
				"dont care",
				"true",
				"dont care",
			},
			[]string{
				"<root_account>",
				"not_supported",
				"true",
			},
		},
		{
			[]string{
				"user1",
				"dont care",
				"dont care",
				"false",
				"dont care",
				"dont care",
				"dont care",
				"false",
				"dont care",
			},
			[]string{
				"user1",
				"false",
				"false",
			},
		},
		{
			[]string{
				"user2",
				"dont care",
				"dont care",
				"true",
				"dont care",
				"dont care",
				"dont care",
				"true",
				"dont care",
			},
			[]string{
				"user2",
				"true",
				"true",
			},
		},
	}

	for _, c := range context {
		v := TrimRecord(c.Record)
		if reflect.DeepEqual(v, c.Expected) == false {
			t.Errorf("expected %s got %s", c.Expected, v)
		}
	}

}
