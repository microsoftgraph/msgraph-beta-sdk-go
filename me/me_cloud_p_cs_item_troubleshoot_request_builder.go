package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCloudPCsItemTroubleshootRequestBuilder provides operations to call the troubleshoot method.
type MeCloudPCsItemTroubleshootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCloudPCsItemTroubleshootRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCloudPCsItemTroubleshootRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeCloudPCsItemTroubleshootRequestBuilderInternal instantiates a new TroubleshootRequestBuilder and sets the default values.
func NewMeCloudPCsItemTroubleshootRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCloudPCsItemTroubleshootRequestBuilder) {
    m := &MeCloudPCsItemTroubleshootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/cloudPCs/{cloudPC%2Did}/microsoft.graph.troubleshoot";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCloudPCsItemTroubleshootRequestBuilder instantiates a new TroubleshootRequestBuilder and sets the default values.
func NewMeCloudPCsItemTroubleshootRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCloudPCsItemTroubleshootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCloudPCsItemTroubleshootRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation troubleshoot a specific Cloud PC. Use this API to check the health status of the Cloud PC and the session host.
func (m *MeCloudPCsItemTroubleshootRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *MeCloudPCsItemTroubleshootRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post troubleshoot a specific Cloud PC. Use this API to check the health status of the Cloud PC and the session host.
func (m *MeCloudPCsItemTroubleshootRequestBuilder) Post(ctx context.Context, requestConfiguration *MeCloudPCsItemTroubleshootRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
