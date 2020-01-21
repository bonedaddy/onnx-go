package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Range", "TestRangeInt32TypeNegativeDelta", NewTestRangeInt32TypeNegativeDelta)
}

// NewTestRangeInt32TypeNegativeDelta version: 6.
func NewTestRangeInt32TypeNegativeDelta() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Range",
		Title:  "TestRangeInt32TypeNegativeDelta",
		ModelB: []byte{0x8, 0x6, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x95, 0x1, 0xa, 0x24, 0xa, 0x5, 0x73, 0x74, 0x61, 0x72, 0x74, 0xa, 0x5, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0xa, 0x5, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x5, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x24, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x5a, 0xf, 0xa, 0x5, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x6, 0xa, 0x4, 0x8, 0x6, 0x12, 0x0, 0x5a, 0xf, 0xa, 0x5, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x6, 0xa, 0x4, 0x8, 0x6, 0x12, 0x0, 0x5a, 0xf, 0xa, 0x5, 0x64, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x6, 0xa, 0x4, 0x8, 0x6, 0x12, 0x0, 0x62, 0x14, 0xa, 0x6, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0xa, 0xa, 0x8, 0x8, 0x6, 0x12, 0x4, 0xa, 0x2, 0x8, 0x2, 0x42, 0x2, 0x10, 0xb},

		/*

		   &ir.NodeProto{
		     Input:     []string{"start", "limit", "delta"},
		     Output:    []string{"output"},
		     Name:      "",
		     OpType:    "Range",
		     Attributes: ([]*ir.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{10}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{6}),
			),

			tensor.New(
				tensor.WithShape(1),
				tensor.WithBacking([]float32{-3}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(2),
				tensor.WithBacking([]int32{10, 7}),
			),
		},
	}
}