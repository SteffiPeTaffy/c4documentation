### Generates C4 plantUML diagram
Small lib that allows drawing C4-PlantUML digrams in golang

## How to run
``./do generate-c4``

## Questions
###How can I have a truly fluent API?
going from:
```go
	someOtherContainer := func() *elements.Container {
		c := elements.NewContainer(
			"my other service",
			"does also stuff",
			"Go")
		c.RelatesTo(someContainer, "requests stuff", "REST/https")
		c.RelatesTo(someContainerDatabase, "persists stuff", "REST/https")
		return c
	}()
```
to: 
```go
	someOtherContainer := elements.NewContainer("my other service", "does also stuff", "Go").
		RelatesTo(someContainer, "requests stuff", "REST/https").
		RelatesTo(someContainerDatabase, "persists stuff", "REST/https")
```

###How can I print all elements (of different types) before I print all outgoing relations?
* How to recursively visit all elements inside system boundaries in order to "find" all relations, given the current inheritance schema (visitor func is clear, just not how to "parse" NamedElement to SystemBoundary)
* Can I find a solution where every element has a pointer to it's writer to avoid this strugge? How could this look like?