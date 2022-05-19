package main

import "errors"

type HashMap struct {
	m map[int]string
}

// O(1)
func NewHashMap() *HashMap {
	return &HashMap{
		m: make(map[int]string),
	}
}

func (h *HashMap) Get(key int) (string, error) {
	val, ok := h.m[key]
	if !ok {
		return "", errors.New("not found")
	}

	return val, nil
}

func (h *HashMap) Put(key int, value string) error {
	if _, ok := h.m[key]; ok {
		return errors.New("key exists in the hashmap")
	}
	h.m[key] = value

	return nil
}

func (h *HashMap) GetRange(from, to int) ([]string, error) {
	var result []string

	// 1 : john
	// 2 : jane

	// 11 : joko
	// 12 : john 2
	// dst

	// 1 sampai 10

	// key >= 1 dan key <= 10
	// tambahkan datanya

	for key := range h.m {
		if key >= from && key <= to {
			result = append(result, h.m[key])
		}
	}

	return result, nil
}
