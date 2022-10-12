package msd

import "unsafe"

func (s *Series) DataI64() []int64 {
	if s.Dtype == FieldKind_Int64 {
		ptr := (*int64)(unsafe.Pointer(&s.Datas[0]))
		return unsafe.Slice(ptr, len(s.Datas)/8)
	}
	return nil
}

func (s *Series) DataU64() []uint64 {
	if s.Dtype == FieldKind_UInt64 {
		ptr := (*uint64)(unsafe.Pointer(&s.Datas[0]))
		return unsafe.Slice(ptr, len(s.Datas)/8)
	}
	return nil
}

func (s *Series) DataF64() []float64 {
	if s.Dtype == FieldKind_Float64 {
		ptr := (*float64)(unsafe.Pointer(&s.Datas[0]))
		return unsafe.Slice(ptr, len(s.Datas)/8)
	}
	return nil
}

func (s *Series) DataD64() []D64 {
	if s.Dtype == FieldKind_Int64 {
		ptr := (*D64)(unsafe.Pointer(&s.Datas[0]))
		return unsafe.Slice(ptr, len(s.Datas)/8)
	}
	return nil
}
