package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder provides operations to manage the userStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
type IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderGetQueryParameters the list of user installation states for this mobile app configuration.
type IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderGetQueryParameters struct {
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
// IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderGetQueryParameters
}
// IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedDeviceMobileAppConfigurationUserStatusId provides operations to manage the userStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
// returns a *IoslobappprovisioningconfigurationsItemUserstatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) ByManagedDeviceMobileAppConfigurationUserStatusId(managedDeviceMobileAppConfigurationUserStatusId string)(*IoslobappprovisioningconfigurationsItemUserstatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceMobileAppConfigurationUserStatusId != "" {
        urlTplParams["managedDeviceMobileAppConfigurationUserStatus%2Did"] = managedDeviceMobileAppConfigurationUserStatusId
    }
    return NewIoslobappprovisioningconfigurationsItemUserstatusesManagedDeviceMobileAppConfigurationUserStatusItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewIoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderInternal instantiates a new IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder and sets the default values.
func NewIoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) {
    m := &IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}/userStatuses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewIoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder instantiates a new IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder and sets the default values.
func NewIoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *IoslobappprovisioningconfigurationsItemUserstatusesCountRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) Count()(*IoslobappprovisioningconfigurationsItemUserstatusesCountRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemUserstatusesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of user installation states for this mobile app configuration.
// returns a ManagedDeviceMobileAppConfigurationUserStatusCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationUserStatusCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusCollectionResponseable), nil
}
// Post create new navigation property to userStatuses for deviceAppManagement
// returns a ManagedDeviceMobileAppConfigurationUserStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable, requestConfiguration *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationUserStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable), nil
}
// ToGetRequestInformation the list of user installation states for this mobile app configuration.
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userStatuses for deviceAppManagement
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationUserStatusable, requestConfiguration *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) WithUrl(rawUrl string)(*IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
