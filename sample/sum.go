package main

import (
	"context"
	"fmt"
	"os"

	calcssdkgo "github.com/sec4you/go_client_sdk"
)

func Sum() {
	body := *calcssdkgo.NewSumPostRequest()
	body.Numbers = []float32{1.0, 2.0, 3.0, 4.0, 5.0}

	configuration := calcssdkgo.NewConfiguration()
	apiClient := calcssdkgo.NewAPIClient(configuration)

	resp, r, err := apiClient.ArithmeticOperationsAPI.SumPost(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArithmeticOperationsAPI.SumPost`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	fmt.Fprintf(os.Stdout, "Response from `ArithmeticOperationsAPI`: %v\n", resp.GetSum())
}
