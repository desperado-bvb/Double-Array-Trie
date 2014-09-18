package antism

import (
	"fmt"
	"sort"
	"errors"
	//"os"
)

var Base 	= make([]int, 1)
var Check	= make([]int, 1)
var Used	= make([]int, 1)
var Max		= 65535
var Next int
var maxCode int

func (wordarr WordSlice)Build()(err error) {
	err = nil
	var done  [] *ElementNode
	size := len(wordarr)
	if size == 0 {
		return nil
	}
	sort.Sort(wordarr)
	Base[0] = 1
	Next 	= 0
	root    := &ElementNode{depth:0, left:0, right:size}
	err, done = antiExtract(root, wordarr, done)
	if err != nil {
		return
	}
	err, _= antiConsturct(wordarr, done)
	if err != nil {
		return
	}
	for k, ele := range Base {
		if ele > 0 {
			fmt.Println(k, ",", ele)
		}
	}
	return
}

func antiConsturct(wordarr WordSlice, done []*ElementNode)(error, int) {
	var begin 	= 0
	var nonZero	= 0
	var first	= false
	var err error
	err = nil

	if len(done) == 0 {
		return nil, begin
	}

	var pos int
	if done[0].text+1 > Next {
		pos  = done[0].text
	} else {
		pos = Next -1
	}

	if pos >= len(Used) {
		antialloct(pos+1)
	}
	for {
		pos++

		for pos >= len(Used) {
			antialloct(pos+Max)
		}

		if Check[pos] != 0 || pos <  done[0].text {
			nonZero++;
			continue
		} else if !first {
			Next  = pos
			first = true
		}

		begin = pos - done[0].text
		t := begin + done[len(done) - 1].text
		if t > len(Used) {
			antialloct(t+Max)
		}

		if Used[begin] == 1 {
			continue
		}

		flag := false
		for i := 1; i< len(done); i++ {
			if Check[begin + done[i].text] != 0 {
				flag = true
				break
			}
		}
		if !flag {
			break
		}
	}

	if float32(1.0 * nonZero / (pos-Next + 1)) >= 0.95 {
		Next = pos
	}

	Used[begin] = 1
	for _, element := range done {
		Check[begin + element.text] = begin
	}

	for _, element := range done {
		var newdone [] *ElementNode
		err, newdone = antiExtract(element, wordarr, newdone)
		if err != nil {
			return err, 0
		}
		if len(newdone) == 0 {
			Base[begin + element.text] = -element.text -1
		} else {
			err, ins := antiConsturct(wordarr, newdone)
			if err != nil {
				return err, 0
			}
			Base[begin+element.text] = ins
		}

		if begin + element.text >  maxCode {
			maxCode = begin + element.text
		}
	}
	return nil, begin
}

func antiExtract(root *ElementNode, wordarr WordSlice, done []*ElementNode)(error,  []*ElementNode) {
	var err error
	err = nil
	preNode  := new(ElementNode)
	var prev = 0
	var chars []byte
	var charLen int
	var cur int
	for i := root.left; i< root.right; i++ {
		tmpNode  := new(ElementNode)
		chars = wordarr[i].charlist
		charLen = len(chars)
		if charLen == 0 || charLen < root.depth {
			continue
		}
		cur = 0
		if charLen != root.depth {
			cur = int(chars[root.depth]) + 1
		}

		if(prev > cur) {
			err = errors.New("get the char from the word is error")
			return err, nil
		}

		if cur != prev || len(done) == 0 {
			tmpNode.depth = root.depth + 1;
			tmpNode.left  = i
			tmpNode.text  = int(cur)
			if preNode != nil {
				preNode.right = i
			}
			preNode = tmpNode
			done = Append(done, tmpNode)

		}
		prev = cur
	}

	if preNode != nil {
		preNode.right = root.right
	} 
	return err, done
}


func Append(slice []*ElementNode , data *ElementNode) []*ElementNode{
    l := len(slice)
    if l + 1 > cap(slice) {
        newSlice := make([]*ElementNode, (l+1)*2)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+1]
    slice[l] = data
    return slice
}

func antialloct(size int) {
	if size >  cap(Base) {
		newSlice := make([]int ,size)
		copy(newSlice, Base)
		Base = newSlice
	}

	if size >  cap(Used) {
		newSlice := make([]int ,size)
		copy(newSlice, Used)
		Used = newSlice
	}

	if size >  cap(Check) {
		newSlice := make([]int ,size)
		copy(newSlice, Check)
		Check = newSlice
	}
}

