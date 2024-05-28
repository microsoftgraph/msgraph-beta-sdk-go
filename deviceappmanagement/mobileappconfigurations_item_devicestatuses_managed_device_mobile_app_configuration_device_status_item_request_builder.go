package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder provides operations to manage the deviceStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
type MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderGetQueryParameters list of ManagedDeviceMobileAppConfigurationDeviceStatus.
type MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderGetQueryParameters
}
// MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal instantiates a new MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder and sets the default values.
func NewMobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    m := &MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}/deviceStatuses/{managedDeviceMobileAppConfigurationDeviceStatus%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder instantiates a new MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder and sets the default values.
func NewMobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property deviceStatuses for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of ManagedDeviceMobileAppConfigurationDeviceStatus.
// returns a ManagedDeviceMobileAppConfigurationDeviceStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationDeviceStatusable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationDeviceStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationDeviceStatusable), nil
}
// Patch update the navigation property deviceStatuses in deviceAppManagement
// returns a ManagedDeviceMobileAppConfigurationDeviceStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationDeviceStatusable, requestConfiguration *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationDeviceStatusable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationDeviceStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationDeviceStatusable), nil
}
// ToDeleteRequestInformation delete navigation property deviceStatuses for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of ManagedDeviceMobileAppConfigurationDeviceStatus.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceStatuses in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationDeviceStatusable, requestConfiguration *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder when successful
func (m *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) WithUrl(rawUrl string)(*MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    return NewMobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
