/*
 * @Author: damao
 * @Date: 2021-11-11 10:13:09
 */

package ch5

import (
	"testing"
)

func Test_square(t *testing.T) {
	tests := []struct {
		name string
		want func() int
	}{
		// TODO: Add test cases.
	}

	bigSlowOperation()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/* if got := square(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("square() = %v, want %v", got, tt.want)
			} */
		})
	}
}
