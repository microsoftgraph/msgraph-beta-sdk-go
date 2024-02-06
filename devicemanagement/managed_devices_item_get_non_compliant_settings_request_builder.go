package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesItemGetNonCompliantSettingsRequestBuilder provides operations to call the getNonCompliantSettings method.
type ManagedDevicesItemGetNonCompliantSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetQueryParameters invoke function getNonCompliantSettings
type ManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetQueryParameters struct {
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
// ManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetQueryParameters
}
// NewManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal instantiates a new GetNonCompliantSettingsRequestBuilder and sets the default values.
func NewManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    m := &ManagedDevicesItemGetNonCompliantSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/{managedDevice%2Did}/getNonCompliantSettings(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedDevicesItemGetNonCompliantSettingsRequestBuilder instantiates a new GetNonCompliantSettingsRequestBuilder and sets the default values.
func NewManagedDevicesItemGetNonCompliantSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesItemGetNonCompliantSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getNonCompliantSettings
// Deprecated: This method is obsolete. Use GetAsGetNonCompliantSettingsGetResponse instead.
func (m *ManagedDevicesItemGetNonCompliantSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration)(ManagedDevicesItemGetNonCompliantSettingsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedDevicesItemGetNonCompliantSettingsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesItemGetNonCompliantSettingsResponseable), nil
}
// GetAsGetNonCompliantSettingsGetResponse invoke function getNonCompliantSettings
func (m *ManagedDevicesItemGetNonCompliantSettingsRequestBuilder) GetAsGetNonCompliantSettingsGetResponse(ctx context.Context, requestConfiguration *ManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration)(ManagedDevicesItemGetNonCompliantSettingsGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedDevicesItemGetNonCompliantSettingsGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesItemGetNonCompliantSettingsGetResponseable), nil
}
// ToGetRequestInformation invoke function getNonCompliantSettings
func (m *ManagedDevicesItemGetNonCompliantSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesItemGetNonCompliantSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ManagedDevicesItemGetNonCompliantSettingsRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesItemGetNonCompliantSettingsRequestBuilder) {
    return NewManagedDevicesItemGetNonCompliantSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
