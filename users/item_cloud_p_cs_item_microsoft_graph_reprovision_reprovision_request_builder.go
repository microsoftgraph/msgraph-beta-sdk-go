package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder provides operations to call the reprovision method.
type ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderInternal instantiates a new ReprovisionRequestBuilder and sets the default values.
func NewItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) {
    m := &ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/microsoft.graph.reprovision";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder instantiates a new ReprovisionRequestBuilder and sets the default values.
func NewItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post reprovision a specific Cloud PC.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpc-reprovision?view=graph-rest-1.0
func (m *ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) Post(ctx context.Context, body ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBodyable, requestConfiguration *ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation reprovision a specific Cloud PC.
func (m *ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionPostRequestBodyable, requestConfiguration *ItemCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
