package generation

import (
	_ "embed"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/PDCMFinder/db-descriptor/pkg/model"
)

//go:embed templates/singlePageTemplate.md
var singlePageTemplate string

//go:embed templates/entityDescriptionTemplate.md
var entityDescriptionTemplate string

func generateMarkdownContent(databaseDescription model.DatabaseDescription) {
	templateContent := singlePageTemplate
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
	template := entityDescriptionTemplate
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
	columnsTable := buildColumnsTable(entity.Columns)
	entityContent = strings.ReplaceAll(entityContent, "[COLUMNS]", columnsTable)
	if len(entity.Relations) > 0 {
		relationsTable := buildRelationsTable(entity.Relations)
		entityContent = strings.ReplaceAll(entityContent, "[RELATIONS]", relationsTable)
	} else {
		entityContent = strings.ReplaceAll(entityContent, "#### Relations", "")
		entityContent = strings.ReplaceAll(entityContent, "[RELATIONS]", "")
	}
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

func buildRelationsTable(relations []model.Relation) string {
	var sb strings.Builder
	headerList := []string{"Column Name", "Foreign Table", "Foreign Table Primary Key", "Foreign Key Name"}
	header := createTableHeader(headerList)
	var rowsSb strings.Builder
	for _, relation := range relations {
		relationValues := make([]string, 0)
		columnName := relation.ColumnName
		relationValues = append(relationValues, columnName)
		relationValues = append(relationValues, relation.ForeignEntityName)
		relationValues = append(relationValues, relation.ForeignColumnName)
		relationValues = append(relationValues, relation.RelationName)
		rowsSb.WriteString(createTableRow(relationValues) + "\n")
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
