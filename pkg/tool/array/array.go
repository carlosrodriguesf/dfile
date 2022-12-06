package array

import (
	"context"
	"math"
)

func Each[V any](ctx context.Context, arr []V, fn func(item V)) {
	for _, item := range arr {
		select {
		case <-ctx.Done():
			return
		default:
			fn(item)
		}
	}
}

func ToBoolMap[V comparable](arr []V) map[V]bool {
	outputMap := make(map[V]bool, len(arr))
	for _, current := range arr {
		outputMap[current] = true
	}
	return outputMap
}

func ToBoolMapFunc[V any, C comparable](arr []V, fn func(V) C) map[C]bool {
	outputMap := make(map[C]bool, len(arr))
	for _, current := range arr {
		outputMap[fn(current)] = true
	}
	return outputMap
}

func Chunk[V any](arr []V, size int) [][]V {
	arrLen := len(arr)
	chunkQuantity := int(math.Ceil(float64(arrLen) / float64(size)))
	chunks := make([][]V, 0, chunkQuantity)
	for i := 0; i < chunkQuantity; i++ {
		start, end := i*size, (i+1)*size
		if end > arrLen {
			end = arrLen
		}
		chunks = append(chunks, arr[start:end])
	}
	return chunks
}

func Fill[V any](arr []V, value V) []V {
	arr[0] = value
	for j := 1; j < len(arr); j *= 2 {
		copy(arr[j:], arr[:j])
	}
	return arr
}
