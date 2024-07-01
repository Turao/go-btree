package btree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockBTree = btree{
	maxDegree: 3,
	root: &node{
		items: []*item{
			{
				key:   500,
				value: 500,
			},
			{
				key:   700,
				value: 700,
			},
		},
		children: []*node{
			{
				items: []*item{
					{
						key:   100,
						value: 100,
					},
				},
				children: []*node{},
			},
			{
				items: []*item{
					{
						key:   600,
						value: 600,
					},
				},
				children: []*node{},
			},
			{
				items: []*item{
					{
						key:   800,
						value: 800,
					},
				},
				children: []*node{},
			},
		},
	},
}

func TestString(t *testing.T) {
	str := mockBTree.String()
	assert.NotNil(t, str)
}

func TestGet(t *testing.T) {
	tests := map[string]struct {
		Key           int
		BTree         BTree
		ExpectedValue interface{}
		ExpectedError error
	}{
		"not found": {
			Key:           1,
			BTree:         &btree{},
			ExpectedValue: nil,
			ExpectedError: ErrKeyNotFound,
		},
		"found at root node": {
			Key:           500,
			BTree:         &mockBTree,
			ExpectedValue: 500,
			ExpectedError: nil,
		},
		"found at leaf node (left)": {
			Key:           100,
			BTree:         &mockBTree,
			ExpectedValue: 100,
			ExpectedError: nil,
		},
		"found at leaf node (middle)": {
			Key:           600,
			BTree:         &mockBTree,
			ExpectedValue: 600,
			ExpectedError: nil,
		},
		"found at leaf node (right)": {
			Key:           800,
			BTree:         &mockBTree,
			ExpectedValue: 800,
			ExpectedError: nil,
		},
	}
	for name, tx := range tests {
		test := tx
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			value, err := test.BTree.Get(test.Key)
			assert.Equal(t, test.ExpectedError, err)
			assert.Equal(t, test.ExpectedValue, value)
		})
	}
}
