package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
type IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters the IOS Lob App Provisioning Configurations.
type IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters
}
// IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *IoslobappprovisioningconfigurationsItemAssignRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Assign()(*IoslobappprovisioningconfigurationsItemAssignRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
// returns a *IoslobappprovisioningconfigurationsItemAssignmentsRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Assignments()(*IoslobappprovisioningconfigurationsItemAssignmentsRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal instantiates a new IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    m := &IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder instantiates a new IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceStatuses provides operations to manage the deviceStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
// returns a *IoslobappprovisioningconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatuses()(*IoslobappprovisioningconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the IOS Lob App Provisioning Configurations.
// returns a IosLobAppProvisioningConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable), nil
}
// GroupAssignments provides operations to manage the groupAssignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
// returns a *IoslobappprovisioningconfigurationsItemGroupassignmentsGroupAssignmentsRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignments()(*IoslobappprovisioningconfigurationsItemGroupassignmentsGroupAssignmentsRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemGroupassignmentsGroupAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
// returns a IosLobAppProvisioningConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the IOS Lob App Provisioning Configurations.
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
// returns a *RequestInformation when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserStatuses provides operations to manage the userStatuses property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
// returns a *IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) UserStatuses()(*IoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsItemUserstatusesUserStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder when successful
func (m *IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*IoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    return NewIoslobappprovisioningconfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
