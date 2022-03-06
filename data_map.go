package payload

// DataMap maps string path elements to its values.
type DataMap map[string]interface{}

// Update copies all keys from anotherDataMap into original DataMap instance.
func (dataMap DataMap) Update(anotherDataMap DataMap) {
	for k, v := range anotherDataMap {
		dataMap[k] = v
	}
}
