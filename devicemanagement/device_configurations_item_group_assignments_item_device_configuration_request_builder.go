package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder provides operations to manage the deviceConfiguration property of the microsoft.graph.deviceConfigurationGroupAssignment entity.
type DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderGetQueryParameters the navigation link to the Device Configuration being targeted.
type DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderGetQueryParameters
}
// NewDeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderInternal instantiates a new DeviceConfigurationRequestBuilder and sets the default values.
func NewDeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder) {
    m := &DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceConfigurations/{deviceConfiguration%2Did}/groupAssignments/{deviceConfigurationGroupAssignment%2Did}/deviceConfiguration{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewDeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder instantiates a new DeviceConfigurationRequestBuilder and sets the default values.
func NewDeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceConfigurationable), nil
}
// ToGetRequestInformation the navigation link to the Device Configuration being targeted.
func (m *DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder) WithUrl(rawUrl string)(*DeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder) {
    return NewDeviceConfigurationsItemGroupAssignmentsItemDeviceConfigurationRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
