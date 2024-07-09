package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder provides operations to manage the iosLobAppProvisioningConfigurations property of the microsoft.graph.deviceAppManagement entity.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters the IOS Lob App Provisioning Configurations.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetQueryParameters
}
// IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
// returns a *IosLobAppProvisioningConfigurationsItemAssignRequestBuilder when successful
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Assign()(*IosLobAppProvisioningConfigurationsItemAssignRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemAssignRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
// returns a *IosLobAppProvisioningConfigurationsItemAssignmentsRequestBuilder when successful
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Assignments()(*IosLobAppProvisioningConfigurationsItemAssignmentsRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal instantiates a new IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    m := &IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder instantiates a new IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder and sets the default values.
func NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property iosLobAppProvisioningConfigurations for deviceAppManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *IosLobAppProvisioningConfigurationsItemDeviceStatusesRequestBuilder when successful
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) DeviceStatuses()(*IosLobAppProvisioningConfigurationsItemDeviceStatusesRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemDeviceStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the IOS Lob App Provisioning Configurations.
// returns a IosLobAppProvisioningConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
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
// returns a *IosLobAppProvisioningConfigurationsItemGroupAssignmentsRequestBuilder when successful
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) GroupAssignments()(*IosLobAppProvisioningConfigurationsItemGroupAssignmentsRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemGroupAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property iosLobAppProvisioningConfigurations in deviceAppManagement
// returns a IosLobAppProvisioningConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, error) {
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
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationable, requestConfiguration *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *IosLobAppProvisioningConfigurationsItemUserStatusesRequestBuilder when successful
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) UserStatuses()(*IosLobAppProvisioningConfigurationsItemUserStatusesRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsItemUserStatusesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder when successful
func (m *IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*IosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder) {
    return NewIosLobAppProvisioningConfigurationsIosLobAppProvisioningConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
