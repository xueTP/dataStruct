package array

import (
	"errors"
	"fmt"
)

type MyArray struct {
	size int
	data []interface{}
}

func NewMyArray(cap int) *MyArray {
	if cap < 2 {
		cap = 2
	}
	data := make([]interface{}, cap)
	return &MyArray{data: data}
}

func (ma MyArray) IsEmpty() bool {
	if ma.size == 0 {
		return true
	}
	return false
}

func (ma MyArray) GetCap() int {
	return cap(ma.data)
}

func (ma MyArray) GetSize() int {
	return ma.size
}

func (ma *MyArray) AddByIndex(index int, val interface{}) error {
	err := ma.checkIndex(index)
	if err != nil {
		return err
	}
	for i := ma.GetSize(); i > index; i-- {
		ma.data[i] = ma.data[i-1]
	}
	ma.data[index] = val
	ma.size++
	return nil
}

func (ma *MyArray) AddAtLast(val interface{}) error {
	return ma.AddByIndex(ma.GetSize(), val)
}

func (ma *MyArray) AddAtHead(val interface{}) error {
	return ma.AddByIndex(0, val)
}

func (ma *MyArray) GetByIndex(index int) (interface{}, error) {
	err := ma.checkIndex(index)
	if err != nil {
		return -1, err
	}
	return ma.data[index], nil
}

func (ma *MyArray) GetIndexByValL(val interface{}) (int, error) {
	for k, v := range ma.data {
		if v == val {
			return k, nil
		}
	}
	return -1, errors.New("this array not have this val")
}

func (ma *MyArray) GetIndexByValR(val interface{}) (int, error) {
	for k := ma.GetSize()-1; k >= 0; k -- {
		if ma.data[k] == val {
			return k, nil
		}
	}
	return -1, errors.New("this array not have this val")
}

func (ma *MyArray) DelByIndex(index int) (interface{}, error) {
	err := ma.checkIndex(index)
	if err != nil {
		return -1, err
	}
	res := ma.data[index]
	for k := index; k < ma.GetSize(); k++ {
		ma.data[k] = ma.data[k+1]
	}
	ma.size --
	ma.deleteAfter()
	return res, nil
}

func (ma *MyArray) DelByValL(val interface{}) (int, error) {
	var err error
	key := -1
	for k, v := range ma.data {
		if v == val {
			_, err = ma.DelByIndex(k)
			key = k
			break
		}
	}
	return key, err
}

func (ma *MyArray) DelByValR(val interface{}) (int, error) {
	var err error
	key := -1
	for k := ma.GetSize(); k >= 0; k-- {
		if ma.data[k] == val {
			_, err = ma.DelByIndex(k)
			key = k
			break
		}
	}
	return key, err
}

func (ma *MyArray) ReplaceByIndex(index int, val interface{}) (error) {
	err := ma.checkIndex(index)
	if err != nil {
		return err
	}
	ma.data[index] = val
	return nil
}

func (ma *MyArray) checkIndex(index int) error {
	if index < 0 || index > ma.GetSize() {
		return errors.New("this index is out of val border")
	}
	if ma.size >= ma.GetCap() {
		ma.redistributionArray(ma.size * 2)
	}
	return nil
}

func (ma *MyArray) deleteAfter() {
	if ma.size <= ma.GetCap()/4 {
		ma.redistributionArray(ma.GetCap()/2)
	}
}

func (ma *MyArray) redistributionArray(newCap int) {
	newArray := make([]interface{}, newCap)
	copy(newArray, ma.data)
	ma.data = newArray
}

func MyArrayDemo() {
	arr := NewMyArray(1)
	err := arr.AddAtLast("1")
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err)
	err = arr.AddAtLast("2")
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err)
	err = arr.AddByIndex(1, "3")
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err)
	err = arr.AddAtHead("4")
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err)
	val, err := arr.GetByIndex(2)
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err, val)
	index, err := arr.GetIndexByValL("3")
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err, index)
	index, err = arr.GetIndexByValR("2")
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err, index)
	err = arr.ReplaceByIndex(1, "2")
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err)
	val, err = arr.DelByIndex(1)
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err, val)
	val, err = arr.DelByValL("3")
	fmt.Println(arr.GetCap(), arr.GetSize(), arr, err, val)
}
