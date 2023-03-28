package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder provides operations to call the getEffectiveDeviceEnrollmentConfigurations method.
type ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters invoke function getEffectiveDeviceEnrollmentConfigurations
type ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters
}
// NewItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal instantiates a new GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    m := &ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/getEffectiveDeviceEnrollmentConfigurations(){?%24top,%24skip,%24search,%24filter,%24count,%24select,%24orderby}", pathParameters),
    }
    return m
}
// NewItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder instantiates a new GetEffectiveDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getEffectiveDeviceEnrollmentConfigurations
func (m *ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(ItemGetEffectiveDeviceEnrollmentConfigurationsResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetEffectiveDeviceEnrollmentConfigurationsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetEffectiveDeviceEnrollmentConfigurationsResponseable), nil
}
// ToGetRequestInformation invoke function getEffectiveDeviceEnrollmentConfigurations
func (m *ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetEffectiveDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
