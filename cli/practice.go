package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const leetCodeGraphQL = "https://leetcode.com/graphql"

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

func createStubbyFile(problemID, language string) error {
	fmt.Println("Creating stub")
	_, functionName, err := getProblemData(problemID, language)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.go", functionName)
	filePath := fmt.Sprintf("%s/%s", language, fileName)
	err = os.MkdirAll(fmt.Sprintf("%s", language), os.ModePerm)
	if err != nil {
		return err
	}

	content := fmt.Sprintf(`import unittest
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
`, functionName, functionName)

	err = os.WriteFile(filePath, []byte(content), os.ModePerm)
	if err != nil {
		return err
	}

	fmt.Printf("Stub file '%s' created in '%s'\n", fileName, filePath)
	return nil
}

var extMapping = map[string]string{"python": ".py", "go": ".go"}

func createStubFiles(problem, language string) error {
	_, snippet, err := getProblemData(problem, language)
	if err != nil {
		return fmt.Errorf("Failed to fetch problem data: %v", err)
	}

	ext := string(extMapping[language])
	fileName := fmt.Sprintf("%s%s", problem, ext) 
	filePath := fmt.Sprintf("%s/%s", language, fileName)
	fileTestName := fmt.Sprintf("%s_test%s", problem, ext)
	fileTestPath := fmt.Sprintf("%s/%s", language, fileTestName)
	fmt.Println("FileName", filePath)

	err = os.MkdirAll(fmt.Sprintf("%s", language), os.ModePerm)
	if err != nil {
		return err
	}


	content := fmt.Sprintf(`import unittest
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
`, problem, problem)

	err = os.WriteFile(filePath, []byte(snippet), os.ModePerm)
	if err != nil {
		return err
	}
	err = os.WriteFile(fileTestPath, []byte(content), os.ModePerm)
	if err != nil {
		return err
	}
	fmt.Printf("Stub file '%s' created in '%s'\n", fileName, filePath)
	return nil
}

func downloadFunc(cmd *cobra.Command, args []string) error {
	problemID, err := cmd.Flags().GetString("problem")
	fmt.Println("Problem", problemID)
	if err != nil {
		return err
	}
	language, err := cmd.Flags().GetString("language")
	fmt.Println("Language", language)
	if err != nil {
		return err
	}

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
	downloadCmd.Flags().StringP("language", "l", "python", "Programming language (default: python)")
	cmd.AddCommand(downloadCmd)

	viper.BindPFlag("problem", cmd.Flags().Lookup("problem"))
	viper.BindPFlag("language", cmd.Flags().Lookup("language"))

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
