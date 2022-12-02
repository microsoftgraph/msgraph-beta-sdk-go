package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.iosLobAppProvisioningConfiguration entity.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetQueryParameters the associated group assignments for IosLobAppProvisioningConfiguration.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetQueryParameters
}
// DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal instantiates a new IosLobAppProvisioningConfigurationAssignmentItemRequestBuilder and sets the default values.
func NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    m := &DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/iosLobAppProvisioningConfigurations/{iosLobAppProvisioningConfiguration%2Did}/assignments/{iosLobAppProvisioningConfigurationAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder instantiates a new IosLobAppProvisioningConfigurationAssignmentItemRequestBuilder and sets the default values.
func NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignments for deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the associated group assignments for IosLobAppProvisioningConfiguration.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property assignments in deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property assignments for deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the associated group assignments for IosLobAppProvisioningConfiguration.
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable), nil
}
// Patch update the navigation property assignments in deviceAppManagement
func (m *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, requestConfiguration *DeviceAppManagementIosLobAppProvisioningConfigurationsItemAssignmentsIosLobAppProvisioningConfigurationAssignmentItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateIosLobAppProvisioningConfigurationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.IosLobAppProvisioningConfigurationAssignmentable), nil
}
