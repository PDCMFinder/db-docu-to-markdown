package generation

import (
	"github.com/PDCMFinder/db-descriptor/pkg/connector"
	"github.com/PDCMFinder/db-descriptor/pkg/service"
)

func GenerateMarkdown(dbDescriptorInput connector.Input) {
	// Get the database description
	databaseDescription := service.GetDbDescription(dbDescriptorInput)
	generateMarkdownContent(databaseDescription)
}
