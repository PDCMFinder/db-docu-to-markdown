// Package generation orchestrates the logic to connect to the database and create the markdown files with the extracted documentation.
package generation

import (
	"github.com/PDCMFinder/db-descriptor/pkg/connector"
	"github.com/PDCMFinder/db-descriptor/pkg/service"
)

// Extracts the documentation from the database and creates the markdown files
func GenerateMarkdown(dbDescriptorInput connector.Input) {
	// Get the database description
	databaseDescription := service.GetDbDescription(dbDescriptorInput)
	generateMarkdownContent(databaseDescription)
}
