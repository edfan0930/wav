package playwav

// PutWAVExtension
func PutWAVExtension(fileName []string) {

	for i, v := range fileName {
		fileName[i] = v + ".wav"
	}
}
