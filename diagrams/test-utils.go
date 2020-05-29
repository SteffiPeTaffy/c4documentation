package diagrams

import "github.com/SteffiPeTaffy/c4documentation/elements"

var someWrappingBoundary = elements.
	NewSystemBoundary("Wrapping Boundary").
	Build()

var someSystemBoundary = elements.
	NewSystemBoundary("boundary one").
	BelongsTo(someWrappingBoundary).
	Build()

var someContainerDatabase = elements.
	NewDatabase("my database").
	Description("stores stuff").
	Owner("Postgres").
	BelongsTo(someSystemBoundary).
	Build()

var someContainer = elements.
	NewContainer("my first container").
	RelatesTo(someContainerDatabase.C4BaseElement, "persists stuff", "REST/https").
	BelongsTo(someSystemBoundary).
	Build()

var outerContainer = elements.
	NewContainer("my outer container").
	RelatesTo(someContainerDatabase.C4BaseElement, "persists stuff", "REST/https").
	BelongsTo(someWrappingBoundary).
	Build()

var someOtherSystemBoundary = elements.
	NewSystemBoundary("boundary two").
	Build()

var someOtherContainer = elements.
	NewContainer("my other service").
	Description("does also stuff").
	Owner("Go").
	RelatesTo(someContainer.C4BaseElement, "requests stuff", "REST/https").
	BelongsTo(someOtherSystemBoundary).
	Build()

var myModel = &elements.C4Model{
	Elements: []*elements.C4Element{someContainer, someContainerDatabase, outerContainer, someOtherContainer},
}
