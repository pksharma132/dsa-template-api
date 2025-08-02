package generator

import (
	"dsa-template-api/internal/models"
	"errors"
	"fmt"
	"strings"
)

func GenerateTemplate(req models.TemplateRequest) (string, error) {
	switch req.Language {
	case "python":
		return generatePythonTemplate(req), nil
	default:
		return "", errors.New("unsupported language")
	}
}

func generatePythonTemplate(req models.TemplateRequest) string {
	var sb strings.Builder

	sb.WriteString("from typing import List, Dict, Any\n\n")
	sb.WriteString("class Solution:\n")
	sb.WriteString(fmt.Sprintf("    def %s(self,", req.Signature.FunctionName))

	for i, p := range req.Signature.Parameters {
		param := fmt.Sprintf(" %s: %s", p.Name, mapPythonType(p.Type))
		if i < len(req.Signature.Parameters)-1 {
			param += ","
		}
		sb.WriteString(param)
	}

	sb.WriteString(fmt.Sprintf(") -> %s:\n", mapPythonType(req.Signature.Returns.Type)))
	sb.WriteString("        # Write your logic here\n")
	sb.WriteString("        pass\n\n")
	sb.WriteString("if __name__ == \"__main__\":\n")
	sb.WriteString("    import sys, json\n")
	sb.WriteString("    data = json.loads(sys.stdin.read())\n")
	sb.WriteString(fmt.Sprintf("    print(json.dumps(Solution().%s(**data)))\n", req.Signature.FunctionName))

	return sb.String()
}

func mapPythonType(dsl string) string {
	mapping := map[string]string{
		"int":       "int",
		"long":      "int",
		"float":     "float",
		"double":    "float",
		"bool":      "bool",
		"string":    "str",
		"int[]":     "List[int]",
		"string[]":  "List[str]",
		"Tree<int>": "TreeNode",
		"Graph":     "Dict[int, List[int]]",
	}
	if val, ok := mapping[dsl]; ok {
		return val
	}
	return "Any"
}
