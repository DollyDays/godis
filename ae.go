package main

import "golang.org/x/sys/unix"

type FeType int

const (
	AE_READABLE FeType = 1
	AE_WRITABLE FeType = 2
)

type TeType int

const (
	AE_NORMAL TeType = 1
	AE_ONCE   TeType = 2
)

type FileProc func(loop *AeLoop, fd int, extra interface{})
type TimeProc func(loop *AeLoop, fd int, extra interface{})

type AeFileEvent struct {
	fd    int
	mask  FeType
	proc  FileProc
	extra interface{}
}

type AeTimeEvent struct {
	id       int
	mask     TeType
	when     int64 // ms
	interval int64 // ms
	proc     TimeProc
	extra    interface{}
	next     *AeTimeEvent
}

type AeLoop struct {
	FileEvents      map[int]*AeFileEvent
	AimeEvents      *AeTimeEvent
	fileEventFd     int
	timeEventNextId int
	stop            bool
}

var fe2ep [3]uint32 = [3]uint32{0, unix.POLLIN, unix.POLLOUT}

func getFeKey(fd int, mask FeType) int {
	if mask == AE_READABLE {
		return fd
	} else {
		return fd * -1
	}
}

func (loop *AeLoop) AddFileEvent(fd int, mask FeType, proc FileProc, extra interface{}) {
	// epoll ctl
	op := 
}
