package main

import (
	"context"

	"buf.build/go/bufplugin/check"
	"buf.build/go/bufplugin/check/checkutil"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	check.Main(
		&check.Spec{
			Rules: []*check.RuleSpec{
				{
					ID:      "PLUGIN_FIELD_CAMEL_CASE",
					Default: true,
					Purpose: "Checks that all field names are camelCase.",
					Type:    check.RuleTypeLint,
					Handler: checkutil.NewFieldRuleHandler(checkFieldCamelCase, checkutil.WithoutImports()),
				},
			},
		},
	)
}

func checkFieldCamelCase(
	_ context.Context,
	responseWriter check.ResponseWriter,
	_ check.Request,
	fieldDescriptor protoreflect.FieldDescriptor,
) error {
	fieldName := string(fieldDescriptor.Name())
	fieldNameToCamelCase := toCamelCase(fieldName)
	if fieldName != fieldNameToCamelCase {
		responseWriter.AddAnnotation(
			check.WithMessagef(
				"Field name %q should be camelCase, such as %q.",
				fieldName,
				fieldNameToCamelCase,
			),
			check.WithDescriptor(fieldDescriptor),
		)
	}
	return nil
}

func toCamelCase(fieldName string) string {
	// The actual logic for toCamelCase would go here.
	return "TODO"
}
