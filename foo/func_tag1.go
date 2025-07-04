//go:build tag1
// +build tag1

package foo

func init() {
	TAG = append(TAG, "tag1")
}
