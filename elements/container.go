package elements

import (
	"fmt"
	"strings"
)

type Target string

const (
	Cloud            Target = "cloud"
	DeviceCluster    Target = "device"
	TestBench        Target = "cloud"
	DeveloperMachine Target = "monitor"
	Other            Target = "other"
)

type Container struct {
	*C4Element
	description      string
	owner            string
	deploymentTarget Target
	swaggerUrl       string
	pactUrl          string
	repoUrl          string
	pipelineUrl      string
}

func NewContainer(name string) *Container {
	container := Container{
		C4Element: &C4Element{
			C4BaseElement: &C4BaseElement{
				Name:              name,
				OutgoingRelations: []*C4Relation{},
			},
		},
	}
	container.C4Writer = container.toC4PlantUMLString
	return &container
}

func (c *Container) Description(description string) *Container {
	c.description = description
	return c
}

func (c *Container) Owner(owner string) *Container {
	c.owner = owner
	return c
}

func (c *Container) Target(deploymentTarget Target) *Container {
	c.deploymentTarget = deploymentTarget
	return c
}

func (c *Container) Swagger(url string) *Container {
	c.swaggerUrl = url
	return c
}

func (c *Container) Pact(url string) *Container {
	c.pactUrl = url
	return c
}

func (c *Container) Repo(url string) *Container {
	c.repoUrl = url
	return c
}

func (c *Container) Pipeline(url string) *Container {
	c.pipelineUrl = url
	return c
}

func (c *Container) toC4PlantUMLString() string {
	repoUrl := toPlantUMLLinkString(c.repoUrl, "repo")
	pipelineUrl := toPlantUMLLinkString(c.pipelineUrl, "pipeline")
	swaggerUrl := toPlantUMLLinkString(c.swaggerUrl, "swagger")
	details := concatUrl(repoUrl, pipelineUrl, swaggerUrl)
	return fmt.Sprintf("Container(%v, '%s', '%s', '%s', '%s')\n", c.Alias(), c.Name, c.owner, c.description, details)
}

func toPlantUMLLinkString(url string, label string) string {
	if url != "" {
		return fmt.Sprintf("[[%s %s]]", url, label)
	}
	return ""
}

func concatUrl(urls ...string) string {
	var details []string

	for _, url := range urls {
		if url != "" {
			details = append(details, url)
		}
	}
	return strings.Join(details, " | ")
}