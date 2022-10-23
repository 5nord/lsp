package lsp

// The definition of a symbol represented as one or many [locations](#Location).
// For most programming languages there is only one location at which a symbol is
// defined.
//
// Servers should prefer returning `DefinitionLink` over `Definition` if supported
// by the client.
type Definition int //string, []interface {}

// Information about where a symbol is defined.
//
// Provides additional metadata over normal [location](#Location) definitions, including the range of
// the defining symbol
type DefinitionLink int //string, string

// LSP arrays.
// @since 3.17.0
type LSPArray int //string, map[string]interface {}

// The LSP any type.
// Please note that strictly speaking a property with the value `undefined`
// can't be converted into JSON preserving the property name. However for
// convenience it is allowed and assumed that all these properties are
// optional as well.
// @since 3.17.0
type LSPAny int //string, []interface {}

// The declaration of a symbol representation as one or many [locations](#Location).
type Declaration int //string, []interface {}

// Information about where a symbol is declared.
//
// Provides additional metadata over normal [location](#Location) declarations, including the range of
// the declaring symbol.
//
// Servers should prefer returning `DeclarationLink` over `Declaration` if supported
// by the client.
type DeclarationLink int //string, string

// Inline value information can be provided by different means:
// - directly as a text value (class InlineValueText).
// - as a name to use for a variable lookup (class InlineValueVariableLookup)
// - as an evaluatable expression (class InlineValueEvaluatableExpression)
// The InlineValue types combines all inline value types into one type.
//
// @since 3.17.0
type InlineValue int //string, []interface {}

// The result of a document diagnostic pull request. A report can
// either be a full report containing all diagnostics for the
// requested document or an unchanged report indicating that nothing
// has changed in terms of diagnostics in comparison to the last
// pull request.
//
// @since 3.17.0
type DocumentDiagnosticReport int //string, []interface {}

type PrepareRenameResult int //string, []interface {}

type ProgressToken int //string, []interface {}

// A document selector is the combination of one or many document filters.
//
// @sample `let sel:DocumentSelector = [{ language: 'typescript' }, { language: 'json', pattern: '**∕tsconfig.json' }]`;
//
// The use of a string as a document filter is deprecated @since 3.16.0.
type DocumentSelector int //string, map[string]interface {}

// An identifier to refer to a change annotation stored with a workspace edit.
type ChangeAnnotationIdentifier int //string, string

// A workspace diagnostic document report.
//
// @since 3.17.0
type WorkspaceDocumentDiagnosticReport int //string, []interface {}

// An event describing a change to a text document. If only a text is provided
// it is considered to be the full content of the document.
type TextDocumentContentChangeEvent int //string, []interface {}

// MarkedString can be used to render human readable text. It is either a markdown string
// or a code-block that provides a language and a code snippet. The language identifier
// is semantically equal to the optional language identifier in fenced code blocks in GitHub
// issues. See https://help.github.com/articles/creating-and-highlighting-code-blocks/#syntax-highlighting
//
// The pair of a language and a value is an equivalent to markdown:
// ```${language}
// ${value}
// ```
//
// Note that markdown strings will be sanitized - that means html will be escaped.
// @deprecated use MarkupContent instead.
type MarkedString int //string, []interface {}

// A document filter describes a top level text document or
// a notebook cell document.
//
// @since 3.17.0 - proposed support for NotebookCellTextDocumentFilter.
type DocumentFilter int //string, []interface {}

// The glob pattern. Either a string pattern or a relative pattern.
//
// @since 3.17.0
type GlobPattern int //string, []interface {}

// A document filter denotes a document by different properties like
// the [language](#TextDocument.languageId), the [scheme](#Uri.scheme) of
// its resource, or a glob-pattern that is applied to the [path](#TextDocument.fileName).
//
// Glob patterns can have the following syntax:
// - `*` to match one or more characters in a path segment
// - `?` to match on one character in a path segment
// - `**` to match any number of path segments, including none
// - `{}` to group sub patterns into an OR expression. (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
//
// @sample A language filter that applies to typescript files on disk: `{ language: 'typescript', scheme: 'file' }`
// @sample A language filter that applies to all package.json paths: `{ language: 'json', pattern: '**package.json' }`
//
// @since 3.17.0
type TextDocumentFilter int //string, []interface {}

// A notebook document filter denotes a notebook document by
// different properties. The properties will be match
// against the notebook's URI (same as with documents)
//
// @since 3.17.0
type NotebookDocumentFilter int //string, []interface {}

// The glob pattern to watch relative to the base path. Glob patterns can have the following syntax:
// - `*` to match one or more characters in a path segment
// - `?` to match on one character in a path segment
// - `**` to match any number of path segments, including none
// - `{}` to group conditions (e.g. `**​/*.{ts,js}` matches all TypeScript and JavaScript files)
// - `[]` to declare a range of characters to match in a path segment (e.g., `example.[0-9]` to match on `example.0`, `example.1`, …)
// - `[!...]` to negate a range of characters to match in a path segment (e.g., `example.[!0-9]` to match on `example.a`, `example.b`, but not `example.0`)
//
// @since 3.17.0
type Pattern int //string, string

// A set of predefined token types. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
//
// @since 3.16.0
type SemanticTokenTypes string

const (
	Namespace SemanticTokenTypes = "namespace"

	// Represents a generic type. Acts as a fallback for types which can't be mapped to
	// a specific type like class or enum.
	Type          SemanticTokenTypes = "type"
	Class         SemanticTokenTypes = "class"
	Enum          SemanticTokenTypes = "enum"
	Interface     SemanticTokenTypes = "interface"
	Struct        SemanticTokenTypes = "struct"
	TypeParameter SemanticTokenTypes = "typeParameter"
	Parameter     SemanticTokenTypes = "parameter"
	Variable      SemanticTokenTypes = "variable"
	Property      SemanticTokenTypes = "property"
	EnumMember    SemanticTokenTypes = "enumMember"
	Event         SemanticTokenTypes = "event"
	Function      SemanticTokenTypes = "function"
	Method        SemanticTokenTypes = "method"
	Macro         SemanticTokenTypes = "macro"
	Keyword       SemanticTokenTypes = "keyword"
	Modifier      SemanticTokenTypes = "modifier"
	Comment       SemanticTokenTypes = "comment"
	String        SemanticTokenTypes = "string"
	Number        SemanticTokenTypes = "number"
	Regexp        SemanticTokenTypes = "regexp"
	Operator      SemanticTokenTypes = "operator"

	// @since 3.17.0
	Decorator SemanticTokenTypes = "decorator"
)

// A set of predefined token modifiers. This set is not fixed
// an clients can specify additional token types via the
// corresponding client capabilities.
//
// @since 3.16.0
type SemanticTokenModifiers string

const (
	Declaration    SemanticTokenModifiers = "declaration"
	Definition     SemanticTokenModifiers = "definition"
	Readonly       SemanticTokenModifiers = "readonly"
	Static         SemanticTokenModifiers = "static"
	Deprecated     SemanticTokenModifiers = "deprecated"
	Abstract       SemanticTokenModifiers = "abstract"
	Async          SemanticTokenModifiers = "async"
	Modification   SemanticTokenModifiers = "modification"
	Documentation  SemanticTokenModifiers = "documentation"
	DefaultLibrary SemanticTokenModifiers = "defaultLibrary"
)

// The document diagnostic report kinds.
//
// @since 3.17.0
type DocumentDiagnosticReportKind string

const (

	// A diagnostic report with a full
	// set of problems.
	Full DocumentDiagnosticReportKind = "full"

	// A report indicating that the last
	// returned report is still accurate.
	Unchanged DocumentDiagnosticReportKind = "unchanged"
)

// Predefined error codes.
type ErrorCodes integer

const (
	ParseError     ErrorCodes = -32700
	InvalidRequest ErrorCodes = -32600
	MethodNotFound ErrorCodes = -32601
	InvalidParams  ErrorCodes = -32602
	InternalError  ErrorCodes = -32603

	// Error code indicating that a server received a notification or
	// request before the server has received the `initialize` request.
	ServerNotInitialized ErrorCodes = -32002
	UnknownErrorCode     ErrorCodes = -32001
)

type LSPErrorCodes integer

const (

	// A request failed but it was syntactically correct, e.g the
	// method name was known and the parameters were valid. The error
	// message should contain human readable information about why
	// the request failed.
	//
	// @since 3.17.0
	RequestFailed LSPErrorCodes = -32803

	// The server cancelled the request. This error code should
	// only be used for requests that explicitly support being
	// server cancellable.
	//
	// @since 3.17.0
	ServerCancelled LSPErrorCodes = -32802

	// The server detected that the content of a document got
	// modified outside normal conditions. A server should
	// NOT send this error code if it detects a content change
	// in it unprocessed messages. The result even computed
	// on an older state might still be useful for the client.
	//
	// If a client decides that a result is not of any use anymore
	// the client should cancel the request.
	ContentModified LSPErrorCodes = -32801

	// The client has canceled a request and a server as detected
	// the cancel.
	RequestCancelled LSPErrorCodes = -32800
)

// A set of predefined range kinds.
type FoldingRangeKind string

const (

	// Folding range for a comment
	Comment FoldingRangeKind = "comment"

	// Folding range for an import or include
	Imports FoldingRangeKind = "imports"

	// Folding range for a region (e.g. `#region`)
	Region FoldingRangeKind = "region"
)

// A symbol kind.
type SymbolKind uinteger

const (
	File          SymbolKind = 1
	Module        SymbolKind = 2
	Namespace     SymbolKind = 3
	Package       SymbolKind = 4
	Class         SymbolKind = 5
	Method        SymbolKind = 6
	Property      SymbolKind = 7
	Field         SymbolKind = 8
	Constructor   SymbolKind = 9
	Enum          SymbolKind = 10
	Interface     SymbolKind = 11
	Function      SymbolKind = 12
	Variable      SymbolKind = 13
	Constant      SymbolKind = 14
	String        SymbolKind = 15
	Number        SymbolKind = 16
	Boolean       SymbolKind = 17
	Array         SymbolKind = 18
	Object        SymbolKind = 19
	Key           SymbolKind = 20
	Null          SymbolKind = 21
	EnumMember    SymbolKind = 22
	Struct        SymbolKind = 23
	Event         SymbolKind = 24
	Operator      SymbolKind = 25
	TypeParameter SymbolKind = 26
)

// Symbol tags are extra annotations that tweak the rendering of a symbol.
//
// @since 3.16
type SymbolTag uinteger

const (

	// Render a symbol as obsolete, usually using a strike-out.
	Deprecated SymbolTag = 1
)

// Moniker uniqueness level to define scope of the moniker.
//
// @since 3.16.0
type UniquenessLevel string

const (

	// The moniker is only unique inside a document
	Document UniquenessLevel = "document"

	// The moniker is unique inside a project for which a dump got created
	Project UniquenessLevel = "project"

	// The moniker is unique inside the group to which a project belongs
	Group UniquenessLevel = "group"

	// The moniker is unique inside the moniker scheme.
	Scheme UniquenessLevel = "scheme"

	// The moniker is globally unique
	Global UniquenessLevel = "global"
)

// The moniker kind.
//
// @since 3.16.0
type MonikerKind string

const (

	// The moniker represent a symbol that is imported into a project
	Import MonikerKind = "import"

	// The moniker represents a symbol that is exported from a project
	Export MonikerKind = "export"

	// The moniker represents a symbol that is local to a project (e.g. a local
	// variable of a function, a class not visible outside the project, ...)
	Local MonikerKind = "local"
)

// Inlay hint kinds.
//
// @since 3.17.0
type InlayHintKind uinteger

const (

	// An inlay hint that for a type annotation.
	Type InlayHintKind = 1

	// An inlay hint that is for a parameter.
	Parameter InlayHintKind = 2
)

// The message type
type MessageType uinteger

const (

	// An error message.
	Error MessageType = 1

	// A warning message.
	Warning MessageType = 2

	// An information message.
	Info MessageType = 3

	// A log message.
	Log MessageType = 4
)

// Defines how the host (editor) should sync
// document changes to the language server.
type TextDocumentSyncKind uinteger

const (

	// Documents should not be synced at all.
	None TextDocumentSyncKind = 0

	// Documents are synced by always sending the full content
	// of the document.
	Full TextDocumentSyncKind = 1

	// Documents are synced by sending the full content on open.
	// After that only incremental updates to the document are
	// send.
	Incremental TextDocumentSyncKind = 2
)

// Represents reasons why a text document is saved.
type TextDocumentSaveReason uinteger

const (

	// Manually triggered, e.g. by the user pressing save, by starting debugging,
	// or by an API call.
	Manual TextDocumentSaveReason = 1

	// Automatic after a delay.
	AfterDelay TextDocumentSaveReason = 2

	// When the editor lost focus.
	FocusOut TextDocumentSaveReason = 3
)

// The kind of a completion entry.
type CompletionItemKind uinteger

const (
	Text          CompletionItemKind = 1
	Method        CompletionItemKind = 2
	Function      CompletionItemKind = 3
	Constructor   CompletionItemKind = 4
	Field         CompletionItemKind = 5
	Variable      CompletionItemKind = 6
	Class         CompletionItemKind = 7
	Interface     CompletionItemKind = 8
	Module        CompletionItemKind = 9
	Property      CompletionItemKind = 10
	Unit          CompletionItemKind = 11
	Value         CompletionItemKind = 12
	Enum          CompletionItemKind = 13
	Keyword       CompletionItemKind = 14
	Snippet       CompletionItemKind = 15
	Color         CompletionItemKind = 16
	File          CompletionItemKind = 17
	Reference     CompletionItemKind = 18
	Folder        CompletionItemKind = 19
	EnumMember    CompletionItemKind = 20
	Constant      CompletionItemKind = 21
	Struct        CompletionItemKind = 22
	Event         CompletionItemKind = 23
	Operator      CompletionItemKind = 24
	TypeParameter CompletionItemKind = 25
)

// Completion item tags are extra annotations that tweak the rendering of a completion
// item.
//
// @since 3.15.0
type CompletionItemTag uinteger

const (

	// Render a completion as obsolete, usually using a strike-out.
	Deprecated CompletionItemTag = 1
)

// Defines whether the insert text in a completion item should be interpreted as
// plain text or a snippet.
type InsertTextFormat uinteger

const (

	// The primary text to be inserted is treated as a plain string.
	PlainText InsertTextFormat = 1

	// The primary text to be inserted is treated as a snippet.
	//
	// A snippet can define tab stops and placeholders with `$1`, `$2`
	// and `${3:foo}`. `$0` defines the final tab stop, it defaults to
	// the end of the snippet. Placeholders with equal identifiers are linked,
	// that is typing in one will update others too.
	//
	// See also: https://microsoft.github.io/language-server-protocol/specifications/specification-current/#snippet_syntax
	Snippet InsertTextFormat = 2
)

// How whitespace and indentation is handled during completion
// item insertion.
//
// @since 3.16.0
type InsertTextMode uinteger

const (

	// The insertion or replace strings is taken as it is. If the
	// value is multi line the lines below the cursor will be
	// inserted using the indentation defined in the string value.
	// The client will not apply any kind of adjustments to the
	// string.
	AsIs InsertTextMode = 1

	// The editor adjusts leading whitespace of new lines so that
	// they match the indentation up to the cursor of the line for
	// which the item is accepted.
	//
	// Consider a line like this: <2tabs><cursor><3tabs>foo. Accepting a
	// multi line completion item is indented using 2 tabs and all
	// following lines inserted will be indented using 2 tabs as well.
	AdjustIndentation InsertTextMode = 2
)

// A document highlight kind.
type DocumentHighlightKind uinteger

const (

	// A textual occurrence.
	Text DocumentHighlightKind = 1

	// Read-access of a symbol, like reading a variable.
	Read DocumentHighlightKind = 2

	// Write-access of a symbol, like writing to a variable.
	Write DocumentHighlightKind = 3
)

// A set of predefined code action kinds
type CodeActionKind string

const (

	// Empty kind.
	Empty CodeActionKind = ""

	// Base kind for quickfix actions: 'quickfix'
	QuickFix CodeActionKind = "quickfix"

	// Base kind for refactoring actions: 'refactor'
	Refactor CodeActionKind = "refactor"

	// Base kind for refactoring extraction actions: 'refactor.extract'
	//
	// Example extract actions:
	//
	// - Extract method
	// - Extract function
	// - Extract variable
	// - Extract interface from class
	// - ...
	RefactorExtract CodeActionKind = "refactor.extract"

	// Base kind for refactoring inline actions: 'refactor.inline'
	//
	// Example inline actions:
	//
	// - Inline function
	// - Inline variable
	// - Inline constant
	// - ...
	RefactorInline CodeActionKind = "refactor.inline"

	// Base kind for refactoring rewrite actions: 'refactor.rewrite'
	//
	// Example rewrite actions:
	//
	// - Convert JavaScript function to class
	// - Add or remove parameter
	// - Encapsulate field
	// - Make method static
	// - Move method to base class
	// - ...
	RefactorRewrite CodeActionKind = "refactor.rewrite"

	// Base kind for source actions: `source`
	//
	// Source code actions apply to the entire file.
	Source CodeActionKind = "source"

	// Base kind for an organize imports source action: `source.organizeImports`
	SourceOrganizeImports CodeActionKind = "source.organizeImports"

	// Base kind for auto-fix source actions: `source.fixAll`.
	//
	// Fix all actions automatically fix errors that have a clear fix that do not require user input.
	// They should not suppress errors or perform unsafe fixes such as generating new types or classes.
	//
	// @since 3.15.0
	SourceFixAll CodeActionKind = "source.fixAll"
)

type TraceValues string

const (

	// Turn tracing off.
	Off TraceValues = "off"

	// Trace messages only.
	Messages TraceValues = "messages"

	// Verbose message tracing.
	Verbose TraceValues = "verbose"
)

// Describes the content type that a client supports in various
// result literals like `Hover`, `ParameterInfo` or `CompletionItem`.
//
// Please note that `MarkupKinds` must not start with a `$`. This kinds
// are reserved for internal usage.
type MarkupKind string

const (

	// Plain text is supported as a content format
	PlainText MarkupKind = "plaintext"

	// Markdown is supported as a content format
	Markdown MarkupKind = "markdown"
)

// A set of predefined position encoding kinds.
//
// @since 3.17.0
type PositionEncodingKind string

const (

	// Character offsets count UTF-8 code units.
	UTF8 PositionEncodingKind = "utf-8"

	// Character offsets count UTF-16 code units.
	//
	// This is the default and must always be supported
	// by servers
	UTF16 PositionEncodingKind = "utf-16"

	// Character offsets count UTF-32 code units.
	//
	// Implementation note: these are the same as Unicode code points,
	// so this `PositionEncodingKind` may also be used for an
	// encoding-agnostic representation of character offsets.
	UTF32 PositionEncodingKind = "utf-32"
)

// The file event type
type FileChangeType uinteger

const (

	// The file got created.
	Created FileChangeType = 1

	// The file got changed.
	Changed FileChangeType = 2

	// The file got deleted.
	Deleted FileChangeType = 3
)

type WatchKind uinteger

const (

	// Interested in create events.
	Create WatchKind = 1

	// Interested in change events
	Change WatchKind = 2

	// Interested in delete events
	Delete WatchKind = 4
)

// The diagnostic's severity.
type DiagnosticSeverity uinteger

const (

	// Reports an error.
	Error DiagnosticSeverity = 1

	// Reports a warning.
	Warning DiagnosticSeverity = 2

	// Reports an information.
	Information DiagnosticSeverity = 3

	// Reports a hint.
	Hint DiagnosticSeverity = 4
)

// The diagnostic tags.
//
// @since 3.15.0
type DiagnosticTag uinteger

const (

	// Unused or unnecessary code.
	//
	// Clients are allowed to render diagnostics with this tag faded out instead of having
	// an error squiggle.
	Unnecessary DiagnosticTag = 1

	// Deprecated or obsolete code.
	//
	// Clients are allowed to rendered diagnostics with this tag strike through.
	Deprecated DiagnosticTag = 2
)

// How a completion was triggered
type CompletionTriggerKind uinteger

const (

	// Completion was triggered by typing an identifier (24x7 code
	// complete), manual invocation (e.g Ctrl+Space) or via API.
	Invoked CompletionTriggerKind = 1

	// Completion was triggered by a trigger character specified by
	// the `triggerCharacters` properties of the `CompletionRegistrationOptions`.
	TriggerCharacter CompletionTriggerKind = 2

	// Completion was re-triggered as current completion list is incomplete
	TriggerForIncompleteCompletions CompletionTriggerKind = 3
)

// How a signature help was triggered.
//
// @since 3.15.0
type SignatureHelpTriggerKind uinteger

const (

	// Signature help was invoked manually by the user or by a command.
	Invoked SignatureHelpTriggerKind = 1

	// Signature help was triggered by a trigger character.
	TriggerCharacter SignatureHelpTriggerKind = 2

	// Signature help was triggered by the cursor moving or by the document content changing.
	ContentChange SignatureHelpTriggerKind = 3
)

// The reason why code actions were requested.
//
// @since 3.17.0
type CodeActionTriggerKind uinteger

const (

	// Code actions were explicitly requested by the user or by an extension.
	Invoked CodeActionTriggerKind = 1

	// Code actions were requested automatically.
	//
	// This typically happens when current selection in a file changes, but can
	// also be triggered when file content changes.
	Automatic CodeActionTriggerKind = 2
)

// A pattern kind describing if a glob pattern matches a file a folder or
// both.
//
// @since 3.16.0
type FileOperationPatternKind string

const (

	// The pattern matches a file only.
	File FileOperationPatternKind = "file"

	// The pattern matches a folder only.
	Folder FileOperationPatternKind = "folder"
)

// A notebook cell kind.
//
// @since 3.17.0
type NotebookCellKind uinteger

const (

	// A markup-cell is formatted source that is used for display.
	Markup NotebookCellKind = 1

	// A code-cell is source code.
	Code NotebookCellKind = 2
)

type ResourceOperationKind string

const (

	// Supports creating new files and folders.
	Create ResourceOperationKind = "create"

	// Supports renaming existing files and folders.
	Rename ResourceOperationKind = "rename"

	// Supports deleting existing files and folders.
	Delete ResourceOperationKind = "delete"
)

type FailureHandlingKind string

const (

	// Applying the workspace change is simply aborted if one of the changes provided
	// fails. All operations executed before the failing operation stay executed.
	Abort FailureHandlingKind = "abort"

	// All operations are executed transactional. That means they either all
	// succeed or no changes at all are applied to the workspace.
	Transactional FailureHandlingKind = "transactional"

	// If the workspace edit contains only textual file changes they are executed transactional.
	// If resource changes (create, rename or delete file) are part of the change the failure
	// handling strategy is abort.
	TextOnlyTransactional FailureHandlingKind = "textOnlyTransactional"

	// The client tries to undo the operations already executed. But there is no
	// guarantee that this is succeeding.
	Undo FailureHandlingKind = "undo"
)

type PrepareSupportDefaultBehavior uinteger

const (

	// The client's default behavior is to select the identifier
	// according the to language's syntax rule.
	Identifier PrepareSupportDefaultBehavior = 1
)

type TokenFormat string

const (
	Relative TokenFormat = "relative"
)
