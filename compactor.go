package compactor

import (
	"github.com/pkg/errors"
	"go.einride.tech/aip/fieldmask"
	gproto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func pop(s []string) ([]string, string) {
	l := len(s)
	return s[0 : l-1], s[l-1]
}

func Update(dst, src gproto.Message, fm *fieldmaskpb.FieldMask) error {
	var err error
	defer func() {
		e := recover()
		if ie, ok := e.(error); ok {
			err = ie
		} else {
			err = errors.Errorf("unknown error %v", e)
		}
	}()
	fieldmask.Update(fm, dst, src)
	return err
}
