package antism

type WordObj struct {
	charlist []byte
}

type WordSlice []*WordObj

func Word(content []byte)(word *WordObj) {
	word = &WordObj{content}
	return
}

func (wordarr WordSlice) Len() int {
	return len(wordarr)
}

func (wordarr WordSlice) Swap(i,j int) {
	wordarr[i], wordarr[j] = wordarr[j],  wordarr[i]
}

func (wordarr WordSlice) Less(i,j int) bool {
	var x ,size int
	var size1 int = len(wordarr[i].charlist)
	var size2 int = len(wordarr[j].charlist)
	
	if size1 < size2 {
		size = size1
	} else {
		size = size2
	}
	
	for x = 0; x < size; x++ {
		if wordarr[i].charlist[x] < wordarr[j].charlist[x] {
			return true
		}

		if wordarr[i].charlist[x] > wordarr[j].charlist[x] {
			return false
		}
	}
	return true
}
