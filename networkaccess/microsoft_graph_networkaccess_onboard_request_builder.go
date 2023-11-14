package networkaccess

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MicrosoftGraphNetworkaccessOnboardRequestBuilder provides operations to call the onboard method.
type MicrosoftGraphNetworkaccessOnboardRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MicrosoftGraphNetworkaccessOnboardRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MicrosoftGraphNetworkaccessOnboardRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMicrosoftGraphNetworkaccessOnboardRequestBuilderInternal instantiates a new MicrosoftGraphNetworkaccessOnboardRequestBuilder and sets the default values.
func NewMicrosoftGraphNetworkaccessOnboardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphNetworkaccessOnboardRequestBuilder) {
    m := &MicrosoftGraphNetworkaccessOnboardRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/networkAccess/microsoft.graph.networkaccess.onboard", pathParameters),
    }
    return m
}
// NewMicrosoftGraphNetworkaccessOnboardRequestBuilder instantiates a new MicrosoftGraphNetworkaccessOnboardRequestBuilder and sets the default values.
func NewMicrosoftGraphNetworkaccessOnboardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MicrosoftGraphNetworkaccessOnboardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMicrosoftGraphNetworkaccessOnboardRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action onboard
func (m *MicrosoftGraphNetworkaccessOnboardRequestBuilder) Post(ctx context.Context, requestConfiguration *MicrosoftGraphNetworkaccessOnboardRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action onboard
func (m *MicrosoftGraphNetworkaccessOnboardRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MicrosoftGraphNetworkaccessOnboardRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *MicrosoftGraphNetworkaccessOnboardRequestBuilder) WithUrl(rawUrl string)(*MicrosoftGraphNetworkaccessOnboardRequestBuilder) {
    return NewMicrosoftGraphNetworkaccessOnboardRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
