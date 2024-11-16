// This is auto-generated by Datastar. DO NOT EDIT.

namespace StarFederation.Datastar

open System

type FragmentMergeMode =
/// Morphs the fragment into the existing element using idiomorph.
| Morph
/// Replaces the inner HTML of the existing element.
| Inner
/// Replaces the outer HTML of the existing element.
| Outer
/// Prepends the fragment to the existing element.
| Prepend
/// Appends the fragment to the existing element.
| Append
/// Inserts the fragment before the existing element.
| Before
/// Inserts the fragment after the existing element.
| After
/// Upserts the attributes of the existing element.
| UpsertAttributes

type EventType =
/// An event dealing with HTML fragments.
| MergeFragments
/// An event dealing with fine grain signals.
| MergeStore
/// An event dealing with removing elements from the DOM
| RemoveFragments
/// An event dealing with removing signals from the store.
| RemoveFromStore
/// An event dealing with redirecting the browser.
| Redirect
/// An event dealing to execute functions in the browser.
| ExecuteJs
/// An event dealing with custom events.
| DispatchCustomEvent


module Consts =
    let [<Literal>] DatastarKey               = "datastar"
    let [<Literal>] Version                   = "0.20.0"
    let [<Literal>] VersionClientByteSize     = 44905
    let [<Literal>] VersionClientByteSizeGzip = 15065

    /// Default: TimeSpan.FromMilliseconds 300
    let DefaultSettleDuration = TimeSpan.FromMilliseconds 300
    /// Default: TimeSpan.FromMilliseconds 1000
    let DefaultSseRetryDuration = TimeSpan.FromMilliseconds 1000


    /// Default: morph - Morphs the fragment into the existing element using idiomorph.
    let DefaultFragmentMergeMode = Morph

    let [<Literal>] DefaultUseViewTransitions = false
    let [<Literal>] DefaultOnlyIfMissing = false
    let [<Literal>] DefaultAutoRemoveScript = true
    let [<Literal>] DefaultCustomEventCancelable = true
    let [<Literal>] DefaultCustomEventComposed = true
    let [<Literal>] DefaultCustomEventBubbles = true

    let [<Literal>] DefaultCustomEventSelector = "document"
    let [<Literal>] DefaultCustomEventDetailJson = "{}"

    let [<Literal>] DatastarDatalineSelector = "selector"
    let [<Literal>] DatastarDatalineMergeMode = "mergeMode"
    let [<Literal>] DatastarDatalineSettleDuration = "settleDuration"
    let [<Literal>] DatastarDatalineFragment = "fragment"
    let [<Literal>] DatastarDatalineUseViewTransition = "useViewTransition"
    let [<Literal>] DatastarDatalineStore = "store"
    let [<Literal>] DatastarDatalineOnlyIfMissing = "onlyIfMissing"
    let [<Literal>] DatastarDatalineUrl = "url"
    let [<Literal>] DatastarDatalinePaths = "paths"
    let [<Literal>] DatastarDatalineScript = "script"
    let [<Literal>] DatastarDatalineAutoRemoveScript = "autoRemoveScript"
    let [<Literal>] DatastarDatalineEventName = "eventName"
    let [<Literal>] DatastarDatalineCancelable = "cancelable"
    let [<Literal>] DatastarDatalineComposed = "composed"
    let [<Literal>] DatastarDatalineBubbles = "bubbles"
    let [<Literal>] DatastarDatalineDetailJson = "detailJson"

    module FragmentMergeMode =
        let toString this =
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
        let toString this =
            match this with
                | EventType.MergeFragments -> "datastar-merge-fragments"
                | EventType.MergeStore -> "datastar-merge-store"
                | EventType.RemoveFragments -> "datastar-remove-fragments"
                | EventType.RemoveFromStore -> "datastar-remove-from-store"
                | EventType.Redirect -> "datastar-redirect"
                | EventType.ExecuteJs -> "datastar-execute-js"
                | EventType.DispatchCustomEvent -> "datastar-dispatch-custom-event"