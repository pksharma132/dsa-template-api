package generator

import (
	"os"
	"path/filepath"
	"testing"
	"dsa-template-api/internal/models"
	
)

func TestGenerateTemplate_Golden(t *testing.T) {
	input := models.TemplateRequest{
		QuestionID:  "two-sum",
		Title:       "Two Sum",
		Description: "Given an integer array...",
		Signature: models.Signature{
			FunctionName: "twoSum",
			Parameters: []models.Param{
				{Name: "nums", Type: "int[]"},
				{Name: "target", Type: "int"},
			},
			Returns: models.Return{Type: "int[]"},
		},
		Language: "python",
	}

	got, err := GenerateTemplate(input)
	if err != nil {
		t.Fatalf("GenerateTemplate() error = %v", err)
	}

	// Load expected output from golden file
	goldenPath := filepath.Join("templates", "two_sum.python.golden")
	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("failed to read golden file: %v", err)
	}

	if string(want) != got {
		t.Errorf("generated template does not match golden file\n\nExpected:\n%s\n\nGot:\n%s\n", want, got)
	}
}
