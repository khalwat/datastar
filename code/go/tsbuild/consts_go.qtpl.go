// Code generated by qtc from "consts_go.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line code/go/tsbuild/consts_go.qtpl:1
package tsbuild

//line code/go/tsbuild/consts_go.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line code/go/tsbuild/consts_go.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line code/go/tsbuild/consts_go.qtpl:1
func streamgoConsts(qw422016 *qt422016.Writer, data *ConstTemplateData) {
//line code/go/tsbuild/consts_go.qtpl:1
	qw422016.N().S(`package datastar

import "time"

const (
    Version                   = "`)
//line code/go/tsbuild/consts_go.qtpl:7
	qw422016.E().S(data.Version)
//line code/go/tsbuild/consts_go.qtpl:7
	qw422016.N().S(`"
    VersionClientByteSize     = `)
//line code/go/tsbuild/consts_go.qtpl:8
	qw422016.N().D(data.VersionClientByteSize)
//line code/go/tsbuild/consts_go.qtpl:8
	qw422016.N().S(`
    VersionClientByteSizeGzip = `)
//line code/go/tsbuild/consts_go.qtpl:9
	qw422016.N().D(data.VersionClientByteSizeGzip)
//line code/go/tsbuild/consts_go.qtpl:9
	qw422016.N().S(`
    DatastarKey               = "`)
//line code/go/tsbuild/consts_go.qtpl:10
	qw422016.E().S(data.DatastarKey)
//line code/go/tsbuild/consts_go.qtpl:10
	qw422016.N().S(`"
    DefaultSettleDuration         = `)
//line code/go/tsbuild/consts_go.qtpl:11
	qw422016.N().D(durationToMs(data.DefaultSettleDuration))
//line code/go/tsbuild/consts_go.qtpl:11
	qw422016.N().S(` * time.Millisecond
    DefaultSSERetryDuration   = `)
//line code/go/tsbuild/consts_go.qtpl:12
	qw422016.N().D(durationToMs(data.DefaultSSERetryDuration))
//line code/go/tsbuild/consts_go.qtpl:12
	qw422016.N().S(` * time.Millisecond
    DefaultUseViewTransitions = `)
//line code/go/tsbuild/consts_go.qtpl:13
	qw422016.E().V(data.DefaultUseViewTransitions)
//line code/go/tsbuild/consts_go.qtpl:13
	qw422016.N().S(`
)

`)
//line code/go/tsbuild/consts_go.qtpl:16
	for _, enum := range data.Enums {
//line code/go/tsbuild/consts_go.qtpl:16
		qw422016.N().S(`type `)
//line code/go/tsbuild/consts_go.qtpl:17
		qw422016.E().S(enum.Name.Pascal)
//line code/go/tsbuild/consts_go.qtpl:17
		qw422016.N().S(` string

const (
`)
//line code/go/tsbuild/consts_go.qtpl:20
		if enum.Default != nil {
//line code/go/tsbuild/consts_go.qtpl:20
			qw422016.N().S(`    // Default value for `)
//line code/go/tsbuild/consts_go.qtpl:21
			qw422016.E().S(enum.Name.Pascal)
//line code/go/tsbuild/consts_go.qtpl:21
			qw422016.N().S(`
    Default`)
//line code/go/tsbuild/consts_go.qtpl:22
			qw422016.E().S(enum.Name.Pascal)
//line code/go/tsbuild/consts_go.qtpl:22
			qw422016.N().S(` = `)
//line code/go/tsbuild/consts_go.qtpl:22
			qw422016.E().S(enum.Name.Pascal)
//line code/go/tsbuild/consts_go.qtpl:22
			qw422016.E().S(enum.Default.Name.Pascal)
//line code/go/tsbuild/consts_go.qtpl:22
			qw422016.N().S(`

`)
//line code/go/tsbuild/consts_go.qtpl:24
		}
//line code/go/tsbuild/consts_go.qtpl:25
		for _, entry := range enum.Values {
//line code/go/tsbuild/consts_go.qtpl:25
			qw422016.N().S(`    // `)
//line code/go/tsbuild/consts_go.qtpl:26
			qw422016.E().S(entry.Description)
//line code/go/tsbuild/consts_go.qtpl:26
			qw422016.N().S(`
    `)
//line code/go/tsbuild/consts_go.qtpl:27
			qw422016.E().S(enum.Name.Pascal)
//line code/go/tsbuild/consts_go.qtpl:27
			qw422016.E().S(entry.Name.Pascal)
//line code/go/tsbuild/consts_go.qtpl:27
			qw422016.N().S(` `)
//line code/go/tsbuild/consts_go.qtpl:27
			qw422016.E().S(enum.Name.Pascal)
//line code/go/tsbuild/consts_go.qtpl:27
			qw422016.N().S(` = "`)
//line code/go/tsbuild/consts_go.qtpl:27
			qw422016.E().S(entry.Value)
//line code/go/tsbuild/consts_go.qtpl:27
			qw422016.N().S(`"

`)
//line code/go/tsbuild/consts_go.qtpl:29
		}
//line code/go/tsbuild/consts_go.qtpl:29
		qw422016.N().S(`)

`)
//line code/go/tsbuild/consts_go.qtpl:32
	}
//line code/go/tsbuild/consts_go.qtpl:33
}

//line code/go/tsbuild/consts_go.qtpl:33
func writegoConsts(qq422016 qtio422016.Writer, data *ConstTemplateData) {
//line code/go/tsbuild/consts_go.qtpl:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line code/go/tsbuild/consts_go.qtpl:33
	streamgoConsts(qw422016, data)
//line code/go/tsbuild/consts_go.qtpl:33
	qt422016.ReleaseWriter(qw422016)
//line code/go/tsbuild/consts_go.qtpl:33
}

//line code/go/tsbuild/consts_go.qtpl:33
func goConsts(data *ConstTemplateData) string {
//line code/go/tsbuild/consts_go.qtpl:33
	qb422016 := qt422016.AcquireByteBuffer()
//line code/go/tsbuild/consts_go.qtpl:33
	writegoConsts(qb422016, data)
//line code/go/tsbuild/consts_go.qtpl:33
	qs422016 := string(qb422016.B)
//line code/go/tsbuild/consts_go.qtpl:33
	qt422016.ReleaseByteBuffer(qb422016)
//line code/go/tsbuild/consts_go.qtpl:33
	return qs422016
//line code/go/tsbuild/consts_go.qtpl:33
}
