package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudPCsItemMicrosoftGraphTroubleshootRequestBuilder provides operations to call the troubleshoot method.
type CloudPCsItemMicrosoftGraphTroubleshootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CloudPCsItemMicrosoftGraphTroubleshootRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCsItemMicrosoftGraphTroubleshootRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudPCsItemMicrosoftGraphTroubleshootRequestBuilderInternal instantiates a new MicrosoftGraphTroubleshootRequestBuilder and sets the default values.
func NewCloudPCsItemMicrosoftGraphTroubleshootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCsItemMicrosoftGraphTroubleshootRequestBuilder) {
    m := &CloudPCsItemMicrosoftGraphTroubleshootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/cloudPCs/{cloudPC%2Did}/microsoft.graph.troubleshoot";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCloudPCsItemMicrosoftGraphTroubleshootRequestBuilder instantiates a new MicrosoftGraphTroubleshootRequestBuilder and sets the default values.
func NewCloudPCsItemMicrosoftGraphTroubleshootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCsItemMicrosoftGraphTroubleshootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCsItemMicrosoftGraphTroubleshootRequestBuilderInternal(urlParams, requestAdapter)
}
// Post troubleshoot a specific Cloud PC. Use this API to check the health status of the Cloud PC and the session host.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/cloudpc-troubleshoot?view=graph-rest-1.0
func (m *CloudPCsItemMicrosoftGraphTroubleshootRequestBuilder) Post(ctx context.Context, requestConfiguration *CloudPCsItemMicrosoftGraphTroubleshootRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation troubleshoot a specific Cloud PC. Use this API to check the health status of the Cloud PC and the session host.
func (m *CloudPCsItemMicrosoftGraphTroubleshootRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *CloudPCsItemMicrosoftGraphTroubleshootRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
