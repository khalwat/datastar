// This file will be regenerated by tsbuild
namespace StarFederation.Datastar

open System

type FragmentMergeMode = | Morph | Inner | Outer | Prepend | Append | Before | After | UpsertAttributes

type EventType = | Fragment | Signal | Remove | Redirect | Console

type ConsoleMode = | Assert | Clear | Count | CountReset | Debug | Dir | Dirxml | Error | Group | GroupCollapsed | GroupEnd | Info | Log | Table | Time | TimeEnd | TimeLog | Trace | Warn


module Consts =

    let [<Literal>] DatastarKey = "datastar"
    let [<Literal>] Version                   = "0.20.0"
    let [<Literal>] VersionClientByteSize     = 43803
    let [<Literal>] VersionClientByteSizeGzip = 14866

    let DefaultSettleDuration     = TimeSpan.FromMilliseconds 300
    let DefaultSSERetryDuration   = TimeSpan.FromMilliseconds 1000
    let [<Literal>] DefaultUseViewTransitions = false
    let [<Literal>] DefaultOnlyIfMissing = false

    let [<Literal>] datastarDatalineSelector = "selector "
    let [<Literal>] datastarDatalineMerge = "merge "
    let [<Literal>] datastarDatalineSettleDuration = "settleDuration "
    let [<Literal>] datastarDatalineFragment = "fragment "
    let [<Literal>] datastarDatalineUseViewTransition = "useViewTransition "
    let [<Literal>] datastarDatalineStore = "store "
    let [<Literal>] datastarDatalineOnlyIfMissing = "onlyIfMissing "
    let [<Literal>] datastarDatalineUrl = "url "
    let [<Literal>] datastarDatalinePaths = "paths "

    module FragmentMergeMode =
        let toString (this.FragmentMergeMode) =
            match this with
                | FragmentMergeMode.Morph -> "morph"
                | FragmentMergeMode.Inner -> "inner"
                | FragmentMergeMode.Outer -> "outer"
                | FragmentMergeMode.Prepend -> "prepend"
                | FragmentMergeMode.Append -> "append"
                | FragmentMergeMode.Before -> "before"
                | FragmentMergeMode.After -> "after"
                | FragmentMergeMode.UpsertAttributes -> "upsertAttributes"

    module EventType =
        let toString (this.EventType) =
            match this with
                | EventType.Fragment -> "datastar-fragment"
                | EventType.Signal -> "datastar-signal"
                | EventType.Remove -> "datastar-remove"
                | EventType.Redirect -> "datastar-redirect"
                | EventType.Console -> "datastar-console"

    module ConsoleMode =
        let toString (this.ConsoleMode) =
            match this with
                | ConsoleMode.Assert -> "assert"
                | ConsoleMode.Clear -> "clear"
                | ConsoleMode.Count -> "count"
                | ConsoleMode.CountReset -> "countReset"
                | ConsoleMode.Debug -> "debug"
                | ConsoleMode.Dir -> "dir"
                | ConsoleMode.Dirxml -> "dirxml"
                | ConsoleMode.Error -> "error"
                | ConsoleMode.Group -> "group"
                | ConsoleMode.GroupCollapsed -> "groupCollapsed"
                | ConsoleMode.GroupEnd -> "groupEnd"
                | ConsoleMode.Info -> "info"
                | ConsoleMode.Log -> "log"
                | ConsoleMode.Table -> "table"
                | ConsoleMode.Time -> "time"
                | ConsoleMode.TimeEnd -> "timeEnd"
                | ConsoleMode.TimeLog -> "timeLog"
                | ConsoleMode.Trace -> "trace"
                | ConsoleMode.Warn -> "warn"