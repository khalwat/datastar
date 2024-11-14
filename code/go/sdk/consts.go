package datastar

import "time"

const (
    DatastarKey = "datastar"
    Version                   = "0.20.0"
    VersionClientByteSize     = 43803
    VersionClientByteSizeGzip = 14866
    DefaultSettleDuration     = 300 * time.Millisecond
    DefaultSSERetryDuration   = 1000 * time.Millisecond
    DefaultUseViewTransitions = false
    DefaultOnlyIfMissing      = false
    SelectorDatalineLiteral = "selector "
    MergeDatalineLiteral = "merge "
    SettleDurationDatalineLiteral = "settleDuration "
    FragmentDatalineLiteral = "fragment "
    UseViewTransitionDatalineLiteral = "useViewTransition "
    StoreDatalineLiteral = "store "
    OnlyIfMissingDatalineLiteral = "onlyIfMissing "
    UrlDatalineLiteral = "url "
    PathsDatalineLiteral = "paths "
)

type FragmentMergeMode string

const (
    // Default value for FragmentMergeMode
    DefaultFragmentMergeMode = FragmentMergeModeMorph

    // Morphs the fragment into the existing element using idiomorph.
    FragmentMergeModeMorph FragmentMergeMode = "morph"

    // Replaces the inner HTML of the existing element.
    FragmentMergeModeInner FragmentMergeMode = "inner"

    // Replaces the outer HTML of the existing element.
    FragmentMergeModeOuter FragmentMergeMode = "outer"

    // Prepends the fragment to the existing element.
    FragmentMergeModePrepend FragmentMergeMode = "prepend"

    // Appends the fragment to the existing element.
    FragmentMergeModeAppend FragmentMergeMode = "append"

    // Inserts the fragment before the existing element.
    FragmentMergeModeBefore FragmentMergeMode = "before"

    // Inserts the fragment after the existing element.
    FragmentMergeModeAfter FragmentMergeMode = "after"

    // Upserts the attributes of the existing element.
    FragmentMergeModeUpsertAttributes FragmentMergeMode = "upsertAttributes"

)

type EventType string

const (
    // An event dealing with HTML fragments
    EventTypeFragment EventType = "datastar-fragment"

    // An event dealing with fine grain signals
    EventTypeSignal EventType = "datastar-signal"

    // An event dealing with removing elements or signals
    EventTypeRemove EventType = "datastar-remove"

    // An event dealing with redirecting the browser
    EventTypeRedirect EventType = "datastar-redirect"

    // An event dealing with console messages
    EventTypeConsole EventType = "datastar-console"

)

type ConsoleMode string

const (
    // Writes an error message to the console if the assertion is false. If the assertion is true, nothing happens.
    ConsoleModeAssert ConsoleMode = "assert"

    // Clears the console if the console allows it. A graphical console, like those running on browsers, will allow it; a console displaying on the terminal, like the one running on Node, will not support it, and will have no effect (and no error).
    ConsoleModeClear ConsoleMode = "clear"

    // Logs the number of times that this particular call to count() has been called.
    ConsoleModeCount ConsoleMode = "count"

    // Resets counter used with console.count().
    ConsoleModeCountReset ConsoleMode = "countReset"

    // Outputs a message to the console at the &#39;debug&#39; log level. The message is only displayed to the user if the console is configured to display debug output. In most cases, the log level is configured within the console UI. This log level might correspond to the Debug or Verbose log level.
    ConsoleModeDebug ConsoleMode = "debug"

    // Displays a list of the properties of the specified JavaScript object. In browser consoles, the output is presented as a hierarchical listing with disclosure triangles that let you see the contents of child objects. Unlike other logging methods, console.dir() does not attempt to pretty-print the object. For example, if you pass a DOM element to console.dir(), it will not be displayed like in the element inspector, but will instead show a list of properties.
    ConsoleModeDir ConsoleMode = "dir"

    // Displays an interactive tree of the descendant elements of the specified XML/HTML element. If it is not possible to display as an element the JavaScript Object view is shown instead. The output is presented as a hierarchical listing of expandable nodes that let you see the contents of child nodes.
    ConsoleModeDirxml ConsoleMode = "dirxml"

    // Outputs a message to the console at the &#39;error&#39; log level.
    ConsoleModeError ConsoleMode = "error"

    // Creates a new inline group in the Web console log, causing any subsequent console messages to be indented by an additional level, until console.groupEnd() is called.
    ConsoleModeGroup ConsoleMode = "group"

    // Creates a new inline group in the console. Unlike console.group(), however, the new group is created collapsed. The user will need to use the disclosure button next to it to expand it, revealing the entries created in the group. Call console.groupEnd() to back out to the parent group.
    ConsoleModeGroupCollapsed ConsoleMode = "groupCollapsed"

    // Exits the current inline group in the console.
    ConsoleModeGroupEnd ConsoleMode = "groupEnd"

    // Outputs a message to the console at the &#39;info&#39; log level.
    ConsoleModeInfo ConsoleMode = "info"

    // Outputs a message to the console.
    ConsoleModeLog ConsoleMode = "log"

    // Displays tabular data as a table.
    ConsoleModeTable ConsoleMode = "table"

    // Starts a timer you can use to track how long an operation takes. You give each timer a unique name, and may have up to 10,000 timers running on a given page. When you call console.timeEnd() with the same name, the browser will output the time, in milliseconds, that elapsed since the timer was started.
    ConsoleModeTime ConsoleMode = "time"

    // The console.timeEnd() static method stops a timer that was previously started by calling console.time().
    ConsoleModeTimeEnd ConsoleMode = "timeEnd"

    // Logs the current value of a timer that was previously started by calling console.time().
    ConsoleModeTimeLog ConsoleMode = "timeLog"

    // Outputs a stack trace to the console.
    ConsoleModeTrace ConsoleMode = "trace"

    // Outputs a warning message to the console at the &#39;warning&#39; log level.
    ConsoleModeWarn ConsoleMode = "warn"

)