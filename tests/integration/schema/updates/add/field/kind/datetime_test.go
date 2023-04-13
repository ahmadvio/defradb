// Copyright 2023 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package kind

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration"
)

func TestSchemaUpdatesAddFieldKindDateTime(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, add field with kind datetime (10)",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "add", "path": "/Users/Schema/Fields/-", "value": {"Name": "Foo", "Kind": 10} }
					]
				`,
			},
			testUtils.Request{
				Request: `query {
					Users {
						Name
						Foo
					}
				}`,
				Results: []map[string]any{},
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesAddFieldKindDateTimeWithCreate(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, add field with kind datetime (10) with create",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "add", "path": "/Users/Schema/Fields/-", "value": {"Name": "Foo", "Kind": 4} }
					]
				`,
			},
			testUtils.CreateDoc{
				CollectionID: 0,
				Doc: `{
					"Name": "John",
					"Foo": "2017-07-23T03:46:56.647Z"
				}`,
			},
			testUtils.Request{
				Request: `query {
					Users {
						Name
						Foo
					}
				}`,
				Results: []map[string]any{
					{
						"Name": "John",
						"Foo":  "2017-07-23T03:46:56.647Z",
					},
				},
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}

func TestSchemaUpdatesAddFieldKindDateTimeSubstitutionWithCreate(t *testing.T) {
	test := testUtils.TestCase{
		Description: "Test schema update, add field with kind datetime substitution with create",
		Actions: []any{
			testUtils.SchemaUpdate{
				Schema: `
					type Users {
						Name: String
					}
				`,
			},
			testUtils.SchemaPatch{
				Patch: `
					[
						{ "op": "add", "path": "/Users/Schema/Fields/-", "value": {"Name": "Foo", "Kind": "DateTime"} }
					]
				`,
			},
			testUtils.CreateDoc{
				CollectionID: 0,
				Doc: `{
					"Name": "John",
					"Foo": "2017-07-23T03:46:56.647Z"
				}`,
			},
			testUtils.Request{
				Request: `query {
					Users {
						Name
						Foo
					}
				}`,
				Results: []map[string]any{
					{
						"Name": "John",
						"Foo":  "2017-07-23T03:46:56.647Z",
					},
				},
			},
		},
	}
	testUtils.ExecuteTestCase(t, []string{"Users"}, test)
}