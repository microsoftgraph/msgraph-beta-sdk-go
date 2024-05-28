package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder provides operations to call the getCustomizedSettings method.
type IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderGetQueryParameters invoke function getCustomizedSettings
type IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderGetQueryParameters struct {
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
// IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderGetQueryParameters
}
// NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderInternal instantiates a new IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder and sets the default values.
func NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder) {
    m := &IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/intents/{deviceManagementIntent%2Did}/getCustomizedSettings(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder instantiates a new IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder and sets the default values.
func NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getCustomizedSettings
// Deprecated: This method is obsolete. Use GetAsGetCustomizedSettingsGetResponse instead.
// returns a IntentsItemGetcustomizedsettingsGetCustomizedSettingsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderGetRequestConfiguration)(IntentsItemGetcustomizedsettingsGetCustomizedSettingsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIntentsItemGetcustomizedsettingsGetCustomizedSettingsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IntentsItemGetcustomizedsettingsGetCustomizedSettingsResponseable), nil
}
// GetAsGetCustomizedSettingsGetResponse invoke function getCustomizedSettings
// returns a IntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder) GetAsGetCustomizedSettingsGetResponse(ctx context.Context, requestConfiguration *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderGetRequestConfiguration)(IntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IntentsItemGetcustomizedsettingsGetCustomizedSettingsGetResponseable), nil
}
// ToGetRequestInformation invoke function getCustomizedSettings
// returns a *RequestInformation when successful
func (m *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder when successful
func (m *IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder) WithUrl(rawUrl string)(*IntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder) {
    return NewIntentsItemGetcustomizedsettingsGetCustomizedSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
