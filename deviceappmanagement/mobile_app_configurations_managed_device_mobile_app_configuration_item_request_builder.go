package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder provides operations to manage the mobileAppConfigurations property of the microsoft.graph.deviceAppManagement entity.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters the Managed Device Mobile Application Configurations.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetQueryParameters
}
// MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *MobileAppConfigurationsItemAssignRequestBuilder when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Assign()(*MobileAppConfigurationsItemAssignRequestBuilder) {
    return NewMobileAppConfigurationsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
// returns a *MobileAppConfigurationsItemAssignmentsRequestBuilder when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Assignments()(*MobileAppConfigurationsItemAssignmentsRequestBuilder) {
    return NewMobileAppConfigurationsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal instantiates a new MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    m := &MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder instantiates a new MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder and sets the default values.
func NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property mobileAppConfigurations for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
// returns a *MobileAppConfigurationsItemDeviceStatusesRequestBuilder when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatuses()(*MobileAppConfigurationsItemDeviceStatusesRequestBuilder) {
    return NewMobileAppConfigurationsItemDeviceStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceStatusSummary provides operations to manage the deviceStatusSummary property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
// returns a *MobileAppConfigurationsItemDeviceStatusSummaryRequestBuilder when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) DeviceStatusSummary()(*MobileAppConfigurationsItemDeviceStatusSummaryRequestBuilder) {
    return NewMobileAppConfigurationsItemDeviceStatusSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the Managed Device Mobile Application Configurations.
// returns a ManagedDeviceMobileAppConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable), nil
}
// Patch update the navigation property mobileAppConfigurations in deviceAppManagement
// returns a ManagedDeviceMobileAppConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedDeviceMobileAppConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property mobileAppConfigurations for deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the Managed Device Mobile Application Configurations.
// returns a *RequestInformation when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property mobileAppConfigurations in deviceAppManagement
// returns a *RequestInformation when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedDeviceMobileAppConfigurationable, requestConfiguration *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
// returns a *MobileAppConfigurationsItemUserStatusesRequestBuilder when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatuses()(*MobileAppConfigurationsItemUserStatusesRequestBuilder) {
    return NewMobileAppConfigurationsItemUserStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// UserStatusSummary provides operations to manage the userStatusSummary property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
// returns a *MobileAppConfigurationsItemUserStatusSummaryRequestBuilder when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) UserStatusSummary()(*MobileAppConfigurationsItemUserStatusSummaryRequestBuilder) {
    return NewMobileAppConfigurationsItemUserStatusSummaryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder when successful
func (m *MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*MobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    return NewMobileAppConfigurationsManagedDeviceMobileAppConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
