package sequence

import "sync/atomic"

type Int32Sequence func() int32
type IntSequence func() int
type Uint32Sequence func() uint32

func NewIntSequence(startId int) IntSequence {

	seq := int64(startId)

	return func() int {
		pint64Id := atomic.AddInt64(&seq, 1)
		int64Id := *&pint64Id
		return int(int64Id)
	}

}

func (s Int32Sequence) Next() int32 {
	return s()
}

func NewInt32Sequence(startId int32) Int32Sequence {

	seq := startId

	return func() int32 {
		return atomic.AddInt32(&seq, 1)
	}

}

func (s IntSequence) Next() int {
	return s()
}

func NewUint32Sequence(startId uint32) Uint32Sequence {

	seq := startId

	return func() uint32 {
		return atomic.AddUint32(&seq, 1)
	}

}

func (s Uint32Sequence) Next() uint32 {
	return s()
}
