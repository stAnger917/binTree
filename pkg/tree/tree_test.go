package tree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNode_FindMax(t1 *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "positive case",
			want: 13,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			oak := Node[map[string]interface{}]{}
			err := oak.Insert(4, map[string]interface{}{
				"some_key": fmt.Sprintf("some_value %v", 4),
			})
			require.NoError(t1, err)
			err = oak.Insert(13, map[string]interface{}{
				"some_key": fmt.Sprintf("some_value %v", 13),
			})
			require.NoError(t1, err)
			if got := oak.FindMax(); got != tt.want {
				t1.Errorf("FindMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_FindMin(t1 *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "positive case",
			want: 4,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			oak := Node[map[string]interface{}]{}
			err := oak.Insert(4, map[string]interface{}{
				"some_key": fmt.Sprintf("some_value %v", 4),
			})
			require.NoError(t1, err)
			err = oak.Insert(13, map[string]interface{}{
				"some_key": fmt.Sprintf("some_value %v", 13),
			})
			require.NoError(t1, err)
			if got := oak.FindMin(); got != tt.want {
				t1.Errorf("FindMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_GetValueByIndex(t1 *testing.T) {
	tests := []struct {
		name    string
		index   int
		want    any
		wantErr bool
	}{
		{
			name:  "positive case - left side",
			index: 4,
			want: map[string]interface{}{
				"some_key": fmt.Sprintf("some_value %v", 4),
			},
		},
		{
			name:  "positive case - right side",
			index: 13,
			want: map[string]interface{}{
				"some_key": fmt.Sprintf("some_value %v", 13),
			},
		},
		{
			name:    "negative case",
			index:   8,
			want:    map[string]interface{}{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t1.Run(tt.name, func(t1 *testing.T) {
				oak := Node[map[string]interface{}]{}
				err := oak.Insert(4, map[string]interface{}{
					"some_key": fmt.Sprintf("some_value %v", 4),
				})
				require.NoError(t1, err)
				err = oak.Insert(13, map[string]interface{}{
					"some_key": fmt.Sprintf("some_value %v", 13),
				})
				require.NoError(t1, err)
				got, err := oak.GetValueByIndex(tt.index)
				if tt.wantErr {
					assert.NotNil(t1, err)
					return
				}
				assert.Equal(t1, got, tt.want)
			})
		})
	}
}

func TestNode_Insert(t1 *testing.T) {
	oak := Node[map[string]interface{}]{}
	type args struct {
		index int
		value map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		preTest func()
	}{
		{
			name: "positive case",
			args: args{
				index: 1,
				value: map[string]interface{}{"some_key": "some_value"},
			},
		},
		{
			name: "negative case - index already exist",
			args: args{
				index: 2,
				value: map[string]interface{}{"some_key": "some_value"},
			},
			wantErr: true,
			preTest: func() {
				err := oak.Insert(2, map[string]interface{}{"some_key": "some_value"})
				require.NoError(t1, err)
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			if tt.preTest != nil {
				tt.preTest()
			}
			err := oak.Insert(tt.args.index, tt.args.value)
			if tt.wantErr {
				assert.NotNil(t1, err)
				return
			}
			data, err := oak.GetValueByIndex(tt.args.index)
			require.NoError(t1, err)
			assert.Equal(t1, data, tt.args.value)
		})
	}
}
