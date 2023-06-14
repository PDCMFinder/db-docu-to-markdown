package generation

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/PDCMFinder/db-descriptor/pkg/model"
)

func generateMarkdownContent(databaseDescription model.DatabaseDescription) {
	templateContent := readTemplate("singlePageTemplate.md")
	for _, schema := range databaseDescription.Schemas {
		generateContentForSchema(schema, templateContent)
	}
}

func generateContentForSchema(schema model.Schema, templateContent string) {
	fileContent := strings.ReplaceAll(templateContent, "[SCHEMA_NAME]", schema.Name)
	tablesSection := buildEntitiesContent(schema.GetEntitiesByType("table"))
	fileContent = strings.ReplaceAll(fileContent, "[TABLES_SECTION]", tablesSection)
	viewsSection := buildEntitiesContent(schema.GetEntitiesByType("view"))
	fileContent = strings.ReplaceAll(fileContent, "[VIEWS_SECTION]", viewsSection)
	saveOutput(fileContent, schema.Name)
}

func buildEntitiesContent(entities []model.Entity) string {
	sort.Slice(entities, func(i, j int) bool {
		return entities[i].Name < entities[j].Name
	})
	template := readTemplate("entityDescriptionTemplate.md")
	var entitiesSectionSb strings.Builder
	for _, entity := range entities {
		entityContent := buildEntityContent(entity, template)
		entitiesSectionSb.WriteString(entityContent + "\n")
	}
	return entitiesSectionSb.String()
}

func buildEntityContent(entity model.Entity, template string) string {
	entityContent := strings.ReplaceAll(template, "[ENTITY_NAME]", entity.Name)
	entityContent = strings.ReplaceAll(entityContent, "[ENTITY_COMMENT]", entity.Comment)
	columsTable := buildColumnsTable(entity.Columns)
	entityContent = strings.ReplaceAll(entityContent, "[COLUMNS]", columsTable)
	return entityContent
}

func buildColumnsTable(columns []model.Column) string {
	var sb strings.Builder
	headerList := []string{"Column Name", "Data Type", "Comment"}
	header := createTableHeader(headerList)
	var rowsSb strings.Builder
	for _, column := range columns {
		columnValues := make([]string, 0)
		columnName := column.Name
		if column.IsPrimaryKey {
			columnName = columnName + " " + "\U0001F511"
		}
		columnValues = append(columnValues, columnName)
		columnValues = append(columnValues, column.DataType)
		comment := column.Comment
		if comment == "" {
			comment = "-"
		}
		columnValues = append(columnValues, comment)
		rowsSb.WriteString(createTableRow(columnValues) + "\n")
	}

	sb.WriteString(header + "\n")
	sb.WriteString(rowsSb.String() + "\n")
	return sb.String()
}

func createTableHeader(headerList []string) string {
	header := createTableRow(headerList)
	var line []string
	for i := 0; i < len(headerList); i++ {
		line = append(line, "-----")
	}

	headerSeparator := createTableRow(line)
	header = header + "\n" + headerSeparator

	return header
}

func createTableRow(elements []string) string {
	return createMarkdownTableRow(elements)
}

func createMarkdownTableRow(elements []string) string {
	separator := "|"
	markdownRow := strings.Join(elements, separator)
	if len(markdownRow) > 0 {
		markdownRow = separator + markdownRow + separator
	}

	return markdownRow
}

func saveOutput(fileContent, schemaName string) {
	directory := "output"
	fileName := schemaName + ".md"

	// Create the directory if it doesn't exist
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		log.Fatal("Error creating directory:", err)
	}

	filePath := fmt.Sprintf("%s/%s", directory, fileName)

	f, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(fileContent)

	if err2 != nil {
		log.Fatal(err2)
	}

	log.Println("File", filePath, "created")
}

func readTemplate(templateName string) string {
	content, err := os.ReadFile("resources/" + templateName)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	return string(content)
}
