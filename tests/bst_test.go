package tests

import (
	"testing"

	bst "github.com/AlbertRossJoh/itualgs_go/searching"
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

	_, err := b.Get(20)

	if err == nil {
		t.Errorf("Value was not deleted correctly")
	}

}
