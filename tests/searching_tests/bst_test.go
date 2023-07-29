package tests

import (
	"testing"

	bst "github.com/AlbertRossJoh/itualgs_go/searching"
	utils "github.com/AlbertRossJoh/itualgs_go/sharedfunctions"
)

func TestNewBST(t *testing.T) {

	b := bst.NewBST[int32, string](22, "Albert")
	val, _ := b.Get(22)
	if val != "Albert" {
		t.Errorf("New tree should only have a root")
	}
}

func TestPut(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	b.Put(23, "Ole")
	b.Put(21, "Ole2")
	val, _ := b.Get(23)
	val2, _ := b.Get(21)
	if val != "Ole" && val2 != "Ole2" {
		t.Errorf("Value should be in tree")
	}
}

func TestGet(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	b.Put(24, "Poul")
	b.Put(23, "Ole")
	b.Put(21, "Birgit")
	val, err := b.Get(23)

	if err != nil {
		t.Errorf("Value is not found")
	}

	if val != "Ole" {
		t.Errorf("Value is not correct")
	}
}

func TestDelete(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	b.Put(24, "Poul")
	b.Put(23, "Ole")
	b.Put(25, "Pernille")
	b.Put(20, "Birgit")
	b.Put(21, "Peter")
	val, _ := b.Delete(24)
	del, _ := b.Delete(22)

	if val != "Poul" {
		t.Errorf("Value is not correct")
	}

	val, _ = b.Get(25)

	if val != "Pernille" {
		t.Errorf("Value is not correct")
	}
	if del != "Albert" {
		t.Errorf("Value is not correct")
	}

	val, _ = b.Get(23)

	if val != "Ole" {
		t.Errorf("Value is not correct")
	}

}

func TestDeleteMin(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	b.Put(24, "Poul")
	b.Put(23, "Ole")
	b.Put(25, "Pernille")
	b.Put(20, "Birgit")
	b.Put(21, "Peter")
	b.DeleteMin()

	val, err := b.Get(20)

	if err == nil {
		t.Errorf("minvalue was not deleted correctly, value recieved: %v", val)
	}

}

func TestDeleteMax(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	b.Put(24, "Poul")
	b.Put(23, "Ole")
	b.Put(25, "Pernille")
	b.Put(20, "Birgit")
	b.Put(21, "Peter")
	b.DeleteMax()

	val, err := b.Get(25)

	if err == nil {
		t.Errorf("maxvalue was not deleted correctly, val recieved: %v", val)
	}

}

func TestGetAllKeys(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	b.Put(24, "Poul")
	b.Put(23, "Ole")
	b.Put(25, "Pernille")
	b.Put(20, "Birgit")
	b.Put(21, "Peter")
	arr, err1 := b.GetAllKeys()

	if err1 != nil {
		t.Errorf("Something went wrong")
	}

	expected := []int32{20, 21, 22, 23, 24, 25}
	_, err := b.Get(25)

	if err != nil {
		t.Errorf("The tree got fucked up")
	}

	if !utils.CompareArrays(&arr, &expected) {
		t.Errorf("The array from the bst is not returned correctly, expected: %v, got: %v", expected, arr)
	}

}

func TestGetAllValues(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	b.Put(24, "Poul")
	b.Put(23, "Ole")
	b.Put(25, "Pernille")
	b.Put(20, "Birgit")
	b.Put(21, "Peter")
	arr, err1 := b.GetAllValues()

	if err1 != nil {
		t.Errorf("Something went wrong")
	}

	expected := []string{"Birgit", "Peter", "Albert", "Ole", "Poul", "Pernille"}
	_, err := b.Get(25)

	if err != nil {
		t.Errorf("The tree got fucked up")
	}

	if !utils.CompareArrays(&arr, &expected) {
		t.Errorf("The array from the bst is not returned correctly, expected: %v, got: %v", expected, arr)
	}

}

func TestGetAllKeysFromSingleTree(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	arr, err1 := b.GetAllKeys()

	if err1 != nil {
		t.Errorf("Something went wrong")
	}

	expected := []int32{22}
	_, err := b.Get(22)

	if err != nil {
		t.Errorf("The tree got fucked up")
	}

	if !utils.CompareArrays(&arr, &expected) {
		t.Errorf("The array from the bst is not returned correctly, expected: %v, got: %v", expected, arr)
	}

}

func TestDeleteFromEmptyTree(t *testing.T) {
	b := bst.BST[int, int]{}

	defer func() {
		if err := recover(); err == nil {
			t.Errorf("The function should cause the program to panic")
		}
	}()
	b.DeleteMax()
}

func TestDeleteLastKey(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	str, err1 := b.Delete(22)

	if err1 != nil || str != "Albert" {
		t.Errorf("Something went wrong")
	}

	if !b.IsEmpty() {
		t.Errorf("The tree should be empty")
	}
}

func TestMapify(t *testing.T) {
	b := bst.NewBST[int32, string](22, "Albert")
	b.Put(24, "Poul")
	b.Put(23, "Ole")
	b.Put(25, "Pernille")
	b.Put(20, "Birgit")
	b.Put(21, "Peter")
	keys, err1 := b.GetAllKeys()
	vals, err2 := b.GetAllValues()
	m, err3 := b.Mapify()
	if err1 != nil {
		t.Errorf("Something went wrong")
	}
	if err2 != nil {
		t.Errorf("Something went wrong")
	}
	if err3 != nil {
		t.Errorf("Something went wrong")
	}

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val := vals[i]
		if m[key] != val {
			t.Errorf("Mapify failed, expected: %v, got: %v", val, m[key])
		}
	}
}
