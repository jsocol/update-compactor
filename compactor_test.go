package compactor_test

import (
	"fmt"
	"testing"

	compactor "github.com/jsocol/update-compactor"
	"github.com/jsocol/update-compactor/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestCompactor(t *testing.T) {
	entity := &proto.Person{
		Id:   int32(423),
		Name: "Miles Morales",
	}
	j, err := protojson.Marshal(entity)
	fmt.Printf("before: %s\n", string(j))

	update1 := &proto.UpdatePerson{
		Person: &proto.Person{
			Address: &proto.Address{
				Street1: "543 New St",
			},
		},
	}

	compactor.Update(entity, update1.Person, update1.UpdateMask)

	j, err = protojson.Marshal(entity)
	assert.NoError(t, err)
	fmt.Printf("result: %s\n", string(j))
}
