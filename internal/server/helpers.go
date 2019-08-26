package server

func removeChanFromSlice(slice []chan []byte, target chan []byte) []chan []byte {
	var retSlice []chan []byte
	removeIndex := -1
	for k, v := range slice {
		if v == target {
			removeIndex = k
		}
	}

	if removeIndex == -1 {
		return slice
	}
	if len(slice) == 1 && removeIndex == 0 {
		return retSlice
	}
	retSlice = append(slice[:removeIndex], slice[removeIndex+1:]...)
	return retSlice
}
