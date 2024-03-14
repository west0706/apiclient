package main

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-cd/v2/pkg/apiclient"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/v2/pkg/apiclient/applicationset"
)

func main() {
	fmt.Println("aaa")

	apiClient, err := apiclient.NewClient(&apiclient.ClientOptions{
		ServerAddr: "",
		Insecure:   true,
		AuthToken:  "",
	})
	fmt.Println(apiClient, err)

	_, appClient, err := apiClient.NewApplicationClient()
	fmt.Println(appClient, err)

	appClient.List(context.TODO(), &application.ApplicationQuery{})

	_, appSetClient, err := apiClient.NewApplicationSetClient()
	fmt.Println(appSetClient, err)

	appSetClient.List(context.TODO(), &applicationset.ApplicationSetListQuery{})

}
