//go:build tag2
// +build tag2

package foo

func init() {
	TAG = append(TAG, "tag2")
}
