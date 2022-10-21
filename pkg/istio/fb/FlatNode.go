// Code generated by the FlatBuffers compiler. DO NOT EDIT.

// Copyright (c) 2019, 2020, and 2021 Cisco and/or its affiliates. All rights reserved.

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlatNode struct {
	_tab flatbuffers.Table
}

func GetRootAsFlatNode(buf []byte, offset flatbuffers.UOffsetT) *FlatNode {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FlatNode{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *FlatNode) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FlatNode) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FlatNode) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatNode) Namespace() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatNode) Labels(obj *KeyVal, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *FlatNode) LabelsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FlatNode) Owner() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatNode) WorkloadName() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatNode) PlatformMetadata(obj *KeyVal, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *FlatNode) PlatformMetadataLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FlatNode) IstioVersion() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatNode) MeshId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *FlatNode) AppContainers(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *FlatNode) AppContainersLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *FlatNode) ClusterId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func FlatNodeStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func FlatNodeAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0)
}
func FlatNodeAddNamespace(builder *flatbuffers.Builder, namespace flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(namespace), 0)
}
func FlatNodeAddLabels(builder *flatbuffers.Builder, labels flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(labels), 0)
}
func FlatNodeStartLabelsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FlatNodeAddOwner(builder *flatbuffers.Builder, owner flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(owner), 0)
}
func FlatNodeAddWorkloadName(builder *flatbuffers.Builder, workloadName flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(workloadName), 0)
}
func FlatNodeAddPlatformMetadata(builder *flatbuffers.Builder, platformMetadata flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(platformMetadata), 0)
}
func FlatNodeStartPlatformMetadataVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FlatNodeAddIstioVersion(builder *flatbuffers.Builder, istioVersion flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(istioVersion), 0)
}
func FlatNodeAddMeshId(builder *flatbuffers.Builder, meshId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(meshId), 0)
}
func FlatNodeAddAppContainers(builder *flatbuffers.Builder, appContainers flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(appContainers), 0)
}
func FlatNodeStartAppContainersVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FlatNodeAddClusterId(builder *flatbuffers.Builder, clusterId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(clusterId), 0)
}
func FlatNodeEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
