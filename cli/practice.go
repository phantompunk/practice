package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const leetCodeGraphQL = "https://leetcode.com/graphql"
const pythonTestBase = `import unittest
from dataclasses import dataclass

from %s import Solution


class TestSolution(unittest.TestCase):
    def cases(self):
        @dataclass
        class Case:
            given: int
            expected: int

        return [
            Case(given=1, expected=1),
        ]

    def test_%s(self):
        for case in self.cases():
            sol = Solution()
            actual = sol  
            assert actual == case.expected  
`

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}
type CodeSnippet struct {
	Code     string
	LangSlug string
}
type GraphQLResponse struct {
	Data struct {
		Question struct {
			QuestionID   string
			CodeSnippets []CodeSnippet
		}
	}
}

func getProblemData(id, language string) (string, string, error) {
	query := `query questionEditorData($titleSlug: String!) {
  question(titleSlug: $titleSlug) {
    questionId
    codeSnippets {
      langSlug
      code
    }
  }
}`
	reqBody, err := json.Marshal(GraphQLRequest{Query: query, Variables: map[string]interface{}{
		"titleSlug": id,
	}})
	if err != nil {
		return "", "", fmt.Errorf("Failed to markshall request body: %v", err)
	}
	req, err := http.NewRequest("POST", leetCodeGraphQL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", "", fmt.Errorf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", fmt.Errorf("Failed to read response body: %v", err)
	}
	var gqlResp GraphQLResponse
	err = json.Unmarshal(body, &gqlResp)
	if err != nil {
		return "", "", fmt.Errorf("Failed to unmarsharl response: %v", err)
	}
	fmt.Println(gqlResp.Data.Question.QuestionID)

	var snippet string
	for _, v := range gqlResp.Data.Question.CodeSnippets {
		if v.LangSlug == language {
			fmt.Println(v.Code)
			snippet = v.Code
			break
		}
	}
	return "", snippet, nil
}

var extMapping = map[string]string{"python": ".py", "python3": ".py", "go": ".go"}

func createStubFiles(problem, language string) error {
	_, snippet, err := getProblemData(problem, language)
	if err != nil {
		return fmt.Errorf("Failed to fetch problem data: %v", err)
	}

	ext := string(extMapping[language])
	problem = formatProblemName(problem)
	language = formatLanguageName(language)
	fileName := fmt.Sprintf("%s%s", problem, ext)
	filePath := fmt.Sprintf("%s/%s", language, fileName)
	fileTestName := fmt.Sprintf("%s_test%s", problem, ext)
	fileTestPath := fmt.Sprintf("%s/%s", language, fileTestName)
	fmt.Println("FileName", filePath)

	err = os.MkdirAll(fmt.Sprintf("%s", language), os.ModePerm)
	if err != nil {
		return err
	}

	var testContent string = ""
	switch language {
	case "python":
		testContent = fmt.Sprintf(pythonTestBase, problem, problem)
	}

	err = os.WriteFile(filePath, []byte(snippet), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileTestPath, []byte(testContent), os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Stub file '%s' created in '%s'\n", fileName, filePath)
	return nil
}

var numberToString = map[string]string{"1": "one", "2": "two", "3": "three", "4": "four"}

func hasNumber(name string) bool {
	for _, char := range name {
		fmt.Println("Char:", char)
		if '0' <= char && char <= '9' {
			fmt.Println("Found")
			return true
		}
	}
	return false
}

func convertNumberToWritten(name string) string {
	letters := strings.Split(name, "")
	for i, letter := range letters {
		if hasNumber(letter) {
			written := numberToString[letter]
			letters[i] = written
		}
	}
	return strings.Join(letters, "")
}

func formatLanguageName(name string) string {
	if name == "python3" {
		name = "python"
	}
	return name
}

func formatProblemName(name string) string {
	if hasNumber(name) {
		return convertNumberToWritten(name)
	}

	return strings.ReplaceAll(name, "-", "_")
}

func downloadFunc(cmd *cobra.Command, args []string) error {
	problemID, err := cmd.Flags().GetString("problem")
	if err != nil {
		return err
	}
	language, err := cmd.Flags().GetString("language")
	if err != nil {
		return err
	}

	fmt.Println("Problem", problemID)
	fmt.Println("Language", language)
	return createStubFiles(problemID, language)
}

func main() {
	var cmd = &cobra.Command{
		Use:   "practice",
		Short: "CLI for LeetCode",
		Long:  "A CLI tool to generate stub files for LeetCode problems.",
	}

	var downloadCmd = &cobra.Command{
		Use:   "download",
		Short: "Download a leetcode problem",
		RunE:  downloadFunc,
	}

	downloadCmd.Flags().StringP("problem", "p", "", "LeetCode problem ID")
	downloadCmd.Flags().StringP("language", "l", "python3", "Programming language (default: python)")
	cmd.AddCommand(downloadCmd)

	viper.BindPFlag("problem", cmd.Flags().Lookup("problem"))
	viper.BindPFlag("language", cmd.Flags().Lookup("language"))

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
