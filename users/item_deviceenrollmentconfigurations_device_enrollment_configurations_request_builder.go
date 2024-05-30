package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
type ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters get enrollment configurations targeted to the user
type ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
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
// ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters
}
// ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceEnrollmentConfigurationId provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
// returns a *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder when successful
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) ByDeviceEnrollmentConfigurationId(deviceEnrollmentConfigurationId string)(*ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceEnrollmentConfigurationId != "" {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = deviceEnrollmentConfigurationId
    }
    return NewItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal instantiates a new ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    m := &ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/deviceEnrollmentConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder instantiates a new ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemDeviceenrollmentconfigurationsCountRequestBuilder when successful
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) Count()(*ItemDeviceenrollmentconfigurationsCountRequestBuilder) {
    return NewItemDeviceenrollmentconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateEnrollmentNotificationConfiguration provides operations to call the createEnrollmentNotificationConfiguration method.
// returns a *ItemDeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationRequestBuilder when successful
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) CreateEnrollmentNotificationConfiguration()(*ItemDeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationRequestBuilder) {
    return NewItemDeviceenrollmentconfigurationsCreateenrollmentnotificationconfigurationCreateEnrollmentNotificationConfigurationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get enrollment configurations targeted to the user
// returns a DeviceEnrollmentConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationCollectionResponseable), nil
}
// HasPayloadLinks provides operations to call the hasPayloadLinks method.
// returns a *ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder when successful
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) HasPayloadLinks()(*ItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilder) {
    return NewItemDeviceenrollmentconfigurationsHaspayloadlinksHasPayloadLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Post create new navigation property to deviceEnrollmentConfigurations for users
// returns a DeviceEnrollmentConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, requestConfiguration *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable), nil
}
// ToGetRequestInformation get enrollment configurations targeted to the user
// returns a *RequestInformation when successful
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to deviceEnrollmentConfigurations for users
// returns a *RequestInformation when successful
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, requestConfiguration *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder when successful
func (m *ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) WithUrl(rawUrl string)(*ItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewItemDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
