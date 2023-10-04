package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsItemGetCustomizedSettingsRequestBuilder provides operations to call the getCustomizedSettings method.
type IntentsItemGetCustomizedSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntentsItemGetCustomizedSettingsRequestBuilderGetQueryParameters invoke function getCustomizedSettings
type IntentsItemGetCustomizedSettingsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// IntentsItemGetCustomizedSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemGetCustomizedSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntentsItemGetCustomizedSettingsRequestBuilderGetQueryParameters
}
// NewIntentsItemGetCustomizedSettingsRequestBuilderInternal instantiates a new GetCustomizedSettingsRequestBuilder and sets the default values.
func NewIntentsItemGetCustomizedSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemGetCustomizedSettingsRequestBuilder) {
    m := &IntentsItemGetCustomizedSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/getCustomizedSettings(){?%24top,%24skip,%24search,%24filter,%24count}", pathParameters),
    }
    return m
}
// NewIntentsItemGetCustomizedSettingsRequestBuilder instantiates a new GetCustomizedSettingsRequestBuilder and sets the default values.
func NewIntentsItemGetCustomizedSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemGetCustomizedSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsItemGetCustomizedSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCustomizedSettings
// Deprecated: This method is obsolete. Use GetAsGetCustomizedSettingsGetResponse instead.
func (m *IntentsItemGetCustomizedSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsItemGetCustomizedSettingsRequestBuilderGetRequestConfiguration)(IntentsItemGetCustomizedSettingsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIntentsItemGetCustomizedSettingsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IntentsItemGetCustomizedSettingsResponseable), nil
}
// GetAsGetCustomizedSettingsGetResponse invoke function getCustomizedSettings
func (m *IntentsItemGetCustomizedSettingsRequestBuilder) GetAsGetCustomizedSettingsGetResponse(ctx context.Context, requestConfiguration *IntentsItemGetCustomizedSettingsRequestBuilderGetRequestConfiguration)(IntentsItemGetCustomizedSettingsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIntentsItemGetCustomizedSettingsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IntentsItemGetCustomizedSettingsGetResponseable), nil
}
// ToGetRequestInformation invoke function getCustomizedSettings
func (m *IntentsItemGetCustomizedSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsItemGetCustomizedSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *IntentsItemGetCustomizedSettingsRequestBuilder) WithUrl(rawUrl string)(*IntentsItemGetCustomizedSettingsRequestBuilder) {
    return NewIntentsItemGetCustomizedSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
