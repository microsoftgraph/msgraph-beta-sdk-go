# msgraph-beta-sdk-go



### Random notes

Kiota generation command

```Shell
kiota --openapi https://raw.githubusercontent.com/microsoftgraph/msgraph-metadata/master/openapi/beta/openapi.yaml --language go -o C:/sources/github/msgraph-sdk-go -n github.com/microsoftgraph/msgraph-beta-sdk-go/ -c GraphServiceClient
```

Do not delete graph_request_adapter.go!