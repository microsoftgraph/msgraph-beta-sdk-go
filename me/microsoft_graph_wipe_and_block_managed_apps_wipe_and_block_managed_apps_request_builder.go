package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder provides operations to call the wipeAndBlockManagedApps method.
type MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilderInternal instantiates a new WipeAndBlockManagedAppsRequestBuilder and sets the default values.
func NewMicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder) {
    m := &MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/microsoft.graph.wipeAndBlockManagedApps";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewMicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder instantiates a new WipeAndBlockManagedAppsRequestBuilder and sets the default values.
func NewMicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post blocks the managed app user from app check-in.
func (m *MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder) Post(ctx context.Context, requestConfiguration *MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation blocks the managed app user from app check-in.
func (m *MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphWipeAndBlockManagedAppsWipeAndBlockManagedAppsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
