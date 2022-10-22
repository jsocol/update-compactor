package compactor_test

import (
	"fmt"
	"testing"

	compactor "github.com/jsocol/update-compactor"
	"github.com/jsocol/update-compactor/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func TestCompactor(t *testing.T) {
	initial := []byte(`{"id": 423, "name": "Miles Morales"}`)
	entity := &proto.Person{}
	protojson.Unmarshal(initial, entity)
	// entity = &proto.Person{
	// 	Id:   int32(423),
	// 	Name: "Miles Morales",
	// }
	j, err := protojson.Marshal(entity)
	fmt.Printf("before: %s\n", string(j))

	update1 := &proto.UpdatePerson{
		Person: &proto.Person{
			Address: &proto.Address{
				Street1: "543 New St",
				Street2: "Not this",
			},
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{"address.street1"},
		},
	}

	compactor.Update(entity, update1.Person, update1.UpdateMask)

	j, err = protojson.Marshal(entity)
	assert.NoError(t, err)
	fmt.Printf("result: %s\n", string(j))

	expected := &proto.Person{
		Id:   int32(423),
		Name: "Miles Morales",
		Address: &proto.Address{
			Street1: "543 New St",
		},
	}
	ej, err := protojson.Marshal(expected)
	assert.NoError(t, err)

	assert.Equal(t, string(ej), string(j))
}
