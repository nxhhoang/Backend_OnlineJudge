package impl_test

import (
	"github.com/bibimoni/Online-judge/submission-judge/src/service/problem/impl"
	"testing"
)

func TestGetCheckerAddr(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		problemId string
		want      string
		wantErr   bool
	}{
		// TODO: Add test cases.
		{
			name:      "test1",
			problemId: "445985",
			want:      "/problems_dir/445985/checker",
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := impl.GetCheckerAddr(tt.problemId)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("GetCheckerAddr() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("GetCheckerAddr() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetCheckerAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
