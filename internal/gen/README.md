# Generating the Meta Model

Microsoft provides a meta model describing the Lanugage Server Protocol.

**Generating the Go types for reading the model**

	(cd $HOME; git clone https://github.com/microsoft/language-server-protocol.git)
	go install github.com/atombender/go-jsonschema/cmd/gojsonschema@latest
	gojsonschema -p main -v <(sed '/additionalProperties/d' <~/language-server-protocol/_specifications/lsp/3.17/metaModel/metaModel.schema.json) >metaModel.go

Then manually delete all `_1` and `_2` types and methods from metaModel.go
