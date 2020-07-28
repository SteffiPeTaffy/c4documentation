package diagrams

import "github.com/SteffiPeTaffy/c4documentation/elements"

var someWrappingBoundary = elements.
	NewSystemBoundary("Wrapping Boundary")

var someSystemBoundary = elements.
	NewSystemBoundary("boundary one").
	BelongsTo(someWrappingBoundary)

var someContainerDatabase = elements.
	NewDatabase("my database").
	Description("stores stuff").
	Owner("Postgres").
	BelongsTo(someSystemBoundary)

var someContainer = elements.
	NewContainer("my first container").
	Repo("www.google.com").
	DevEnvironment("dev.com").
	NonProdEnvironment("non-prod.com").
	ProdEnvironment("prod.com").
	RelatesTo(someContainerDatabase.C4BaseElement, "persists stuff", "REST/https").
	BelongsTo(someSystemBoundary)

var outerContainer = elements.
	NewContainer("my outer container").
	RelatesTo(someContainerDatabase.C4BaseElement, "persists stuff", "REST/https").
	BelongsTo(someWrappingBoundary)

var someOtherSystemBoundary = elements.
	NewSystemBoundary("boundary two")

var someOtherContainer = elements.
	NewContainer("my other service").
	Description("does also stuff").
	Owner("Go").
	RelatesTo(someContainer.C4BaseElement, "requests stuff", "REST/https").
	BelongsTo(someOtherSystemBoundary)

var myModel = &elements.C4Model{
	Elements: []elements.WritableElement{someContainer, someContainerDatabase, outerContainer, someOtherContainer},
}
