package onnxtest

// this file is auto-generated... DO NOT EDIT

import (
	"github.com/owulveryck/onnx-go/backend/testbackend"
	"gorgonia.org/tensor"
)

func init() {
	testbackend.Register("Hardmax", "TestHardmaxOneHot", NewTestHardmaxOneHot)
}

// NewTestHardmaxOneHot version: 3.
func NewTestHardmaxOneHot() *testbackend.TestCase {
	return &testbackend.TestCase{
		OpType: "Hardmax",
		Title:  "TestHardmaxOneHot",
		ModelB: []byte{0x8, 0x3, 0x12, 0xc, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2d, 0x74, 0x65, 0x73, 0x74, 0x3a, 0x51, 0xa, 0xf, 0xa, 0x1, 0x78, 0x12, 0x1, 0x79, 0x22, 0x7, 0x48, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x78, 0x12, 0x14, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x72, 0x64, 0x6d, 0x61, 0x78, 0x5f, 0x6f, 0x6e, 0x65, 0x5f, 0x68, 0x6f, 0x74, 0x5a, 0x13, 0xa, 0x1, 0x78, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x4, 0x62, 0x13, 0xa, 0x1, 0x79, 0x12, 0xe, 0xa, 0xc, 0x8, 0x1, 0x12, 0x8, 0xa, 0x2, 0x8, 0x1, 0xa, 0x2, 0x8, 0x4, 0x42, 0x2, 0x10, 0x9},

		/*

		   &pb.NodeProto{
		     Input:     []string{"x"},
		     Output:    []string{"y"},
		     Name:      "",
		     OpType:    "Hardmax",
		     Attributes: ([]*pb.AttributeProto) <nil>
		   ,
		   },


		*/

		Input: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 4),
				tensor.WithBacking([]float32{3, 3, 3, 1}),
			),
		},
		ExpectedOutput: []tensor.Tensor{

			tensor.New(
				tensor.WithShape(1, 4),
				tensor.WithBacking([]float32{1, 0, 0, 0}),
			),
		},
	}
}