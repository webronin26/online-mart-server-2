package test

import "online-mart-server-2/pkg/presenter"

type Output struct {
	TestOutput testOutput `json:"test_output"`
}

type testOutput struct {
	ID    uint64 `json:"id"`
	Title string `json:"title"`
}

// Exec return default information
// Exec 回傳 默認值
func Exec() (Output, presenter.StatusCode, error) {

	var output Output

	output.TestOutput.ID = 1
	output.TestOutput.Title = "OK"

	return output, presenter.StatusSuccess, nil
}
