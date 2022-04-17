package main

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
	"io/ioutil"
	"net/http"
)

var graphqlResponse interface {
}

func main() {
	response, error := http.Post("https://tedhuxdev.sandbox.api.fluentretail.com/oauth/token?username=ukret_admin&password=2004IW&client_id=TEDHUXDEV&client_secret=428d5b18-5290-48e8-9e4a-259e4cb46d09&grant_type=password&scope=api", "application/json", nil)
	if error != nil {
		err := Methods{
			error,
			nil,
		}
		panic(err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	graphqlClient := graphql.NewClient("https://tedhuxdev.sandbox.api.fluentretail.com/graphql")
	graphqlRequest := graphql.NewRequest(`
        mutation{
  createFinancialTransaction(input:{
    ref:"ABC12345",
     type:"PAYMENT",
     amount:100, currency:"USD",
     paymentMethod:
    "GIFTCARD",
     order:{
      id:1203
    }
  }){
    id
  }
}
    `)
	graphqlRequest.Header.Set("Authorization", "bearer aadb50e6-213f-4dc2-8cde-3b42d5eef761")
	error = graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse)
	if error != nil {
		err := Methods{
			error,
			nil,
		}
		panic(err.catch())
	} else {
		response := Methods{
			error,
			graphqlResponse,
		}
		fmt.Println(response.then())
		fmt.Println(response.finally())
	}
}
