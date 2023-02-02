package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder provides operations to call the isManagedAppUserBlocked method.
type MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilderInternal instantiates a new IsManagedAppUserBlockedRequestBuilder and sets the default values.
func NewMicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder) {
    m := &MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.isManagedAppUserBlocked()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder instantiates a new IsManagedAppUserBlockedRequestBuilder and sets the default values.
func NewMicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets the blocked state of a managed app user.
func (m *MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder) Get(ctx context.Context, requestConfiguration *MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration)(MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, CreateMicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedResponseable), nil
}
// ToGetRequestInformation gets the blocked state of a managed app user.
func (m *MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphIsManagedAppUserBlockedIsManagedAppUserBlockedRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
