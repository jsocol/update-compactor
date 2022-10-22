package compactor

import (
	"strings"

	"github.com/pkg/errors"
	"go.einride.tech/aip/fieldbehavior"
	"go.einride.tech/aip/fieldmask"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

func pop(s []string) ([]string, string) {
	l := len(s)
	return s[0 : l-1], s[l-1]
}

func Update(dst, src proto.Message, fm *fieldmaskpb.FieldMask) error {
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

func UpdateJSON(dst, src map[string]interface{}, paths []string) error {
	for _, path := range paths {
		currDst := dst
		currSrc := src

		steps := strings.Split(path, ".")
		steps, leaf := pop(steps)

		for _, step := range steps {
			var ok bool

			if _, ok = currDst[step]; !ok {
				currDst[step] = make(map[string]interface{}, 0)
			}
			if currDst, ok = currDst[step].(map[string]interface{}); !ok {
				return errors.Errorf("unexpected dst leaf at %s of %s", step, path)
			}

			if _, ok = currSrc[step]; !ok {
				return errors.Errorf("source is missing %s", path)
			}
			if currSrc, ok = currSrc[step].(map[string]interface{}); !ok {
				return errors.Errorf("unexpected src leaf at %s of %s", step, path)
			}
		}

		currDst[leaf] = currSrc[leaf]
	}

	return nil
}

type option struct{}

func SkipUnwritable() option {
	return option{}
}

func GRPCPathToJSON(s string, m proto.Message, opts ...option) (string, error) {
	skipUnwritable := len(opts) > 0
	parts := []string{}

	md0 := m.ProtoReflect().Descriptor()
	md := md0
	for _, field := range strings.Split(s, ".") {
		if md == nil {
			break
		}
		fd := md.Fields().ByName(protoreflect.Name(field))
		if fd == nil {
			gd := md.Fields().ByName(protoreflect.Name(strings.ToLower(field)))
			if gd != nil && gd.Kind() == protoreflect.GroupKind && string(gd.Message().Name()) == field {
				fd = gd
			}
		} else if fd.Kind() == protoreflect.GroupKind && string(fd.Message().Name()) != field {
			fd = nil
		}
		if fd == nil {
			return "", errors.Errorf("message does not have field %s", field)
		}
		if skipUnwritable {
			if fieldbehavior.Has(fd, annotations.FieldBehavior_IMMUTABLE) ||
				fieldbehavior.Has(fd, annotations.FieldBehavior_OUTPUT_ONLY) {
				return "", errors.Errorf("field cannot be updated %s", field)
			}
		}
		md = fd.Message()
		if fd.IsMap() {
			md = fd.MapValue().Message()
		}
		parts = append(parts, fd.JSONName())
	}

	return strings.Join(parts, "."), nil
}

func FieldMaskToJSONPaths(fm *fieldmaskpb.FieldMask, m proto.Message, opts ...option) ([]string, error) {
	paths := []string{}
	for _, p := range fm.Paths {
		np, err := GRPCPathToJSON(p, m, opts...)
		if err != nil {
			return nil, err
		}
		paths = append(paths, np)
	}
	return paths, nil
}
