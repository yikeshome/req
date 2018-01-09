package req

import (
	"testing"
)

type testStruct struct {
	ID        int    `req:"id,path,required"`
	AID       int    `req:"aid,path, required"`
	UserName  string `req:"username"`
	Limit     int    `req:"limit,omitempty"`
	Offset    int64  `req:"offset,omitempty"`
	IsDeleted bool   `req:"is_deleted"`
}

type testEmbeddedS struct {
	ID int `req:"id,path,required"`
	subS
}

type subS struct {
	UserName  string `req:"username"`
	Limit     int    `req:"limit,omitempty"`
	Offset    int64  `req:"offset,omitempty"`
	IsDeleted bool   `req:"is_deleted"`
}

var (
	tStructS    = testStruct{11, 1000, "testname", 10, 200, true}
	tStructOmit = testStruct{ID: 12, AID: 1000, UserName: "testname12"}
	tEmbeddedS  = testEmbeddedS{ID: 13}
)

func TestNewEncoder(t *testing.T) {
	resPath, err := Marshal(tStructS)
	if err != nil {
		t.Error(err)
	}
	t.Log(resPath, "First test:")
}

func TestMarshalOmit(t *testing.T) {
	_, err := Marshal(tStructOmit)
	if err != nil {
		t.Error(err)
	}
}

func TestMarshalEmbbedStruct(t *testing.T) {
	tEmbeddedS.Limit = 10
	tEmbeddedS.UserName = "Embedded Name"
	_, err := Marshal(tEmbeddedS)
	if err != nil {
		t.Error(err)
	}
}
