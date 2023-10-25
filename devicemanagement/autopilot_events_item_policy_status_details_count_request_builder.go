package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder provides operations to count the resources in the collection.
type AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AutopilotEventsItemPolicyStatusDetailsCountRequestBuilderGetQueryParameters get the number of the resource
type AutopilotEventsItemPolicyStatusDetailsCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// AutopilotEventsItemPolicyStatusDetailsCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AutopilotEventsItemPolicyStatusDetailsCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AutopilotEventsItemPolicyStatusDetailsCountRequestBuilderGetQueryParameters
}
// NewAutopilotEventsItemPolicyStatusDetailsCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewAutopilotEventsItemPolicyStatusDetailsCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder) {
    m := &AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/autopilotEvents/{deviceManagementAutopilotEvent%2Did}/policyStatusDetails/$count{?%24search,%24filter}", pathParameters),
    }
    return m
}
// NewAutopilotEventsItemPolicyStatusDetailsCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewAutopilotEventsItemPolicyStatusDetailsCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAutopilotEventsItemPolicyStatusDetailsCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
func (m *AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder) Get(ctx context.Context, requestConfiguration *AutopilotEventsItemPolicyStatusDetailsCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
func (m *AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AutopilotEventsItemPolicyStatusDetailsCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder) WithUrl(rawUrl string)(*AutopilotEventsItemPolicyStatusDetailsCountRequestBuilder) {
    return NewAutopilotEventsItemPolicyStatusDetailsCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
