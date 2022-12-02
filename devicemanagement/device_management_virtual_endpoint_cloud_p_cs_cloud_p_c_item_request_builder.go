package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
type DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderGetQueryParameters cloud managed virtual desktops.
type DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderGetQueryParameters
}
// DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeUserAccountType provides operations to call the changeUserAccountType method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) ChangeUserAccountType()(*DeviceManagementVirtualEndpointCloudPCsItemChangeUserAccountTypeRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewDeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    m := &DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewDeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPCs for deviceManagement
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation cloud managed virtual desktops.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property cloudPCs in deviceManagement
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property cloudPCs for deviceManagement
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// EndGracePeriod provides operations to call the endGracePeriod method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) EndGracePeriod()(*DeviceManagementVirtualEndpointCloudPCsItemEndGracePeriodRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get cloud managed virtual desktops.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable), nil
}
// GetCloudPcConnectivityHistory provides operations to call the getCloudPcConnectivityHistory method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*DeviceManagementVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*DeviceManagementVirtualEndpointCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetShiftWorkCloudPcAccessState()(*DeviceManagementVirtualEndpointCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetSupportedCloudPcRemoteActions()(*DeviceManagementVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property cloudPCs in deviceManagement
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable), nil
}
// Reboot provides operations to call the reboot method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) Reboot()(*DeviceManagementVirtualEndpointCloudPCsItemRebootRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Rename provides operations to call the rename method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) Rename()(*DeviceManagementVirtualEndpointCloudPCsItemRenameRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reprovision provides operations to call the reprovision method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) Reprovision()(*DeviceManagementVirtualEndpointCloudPCsItemReprovisionRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) Restore()(*DeviceManagementVirtualEndpointCloudPCsItemRestoreRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) RetryPartnerAgentInstallation()(*DeviceManagementVirtualEndpointCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemRetryPartnerAgentInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Troubleshoot provides operations to call the troubleshoot method.
func (m *DeviceManagementVirtualEndpointCloudPCsCloudPCItemRequestBuilder) Troubleshoot()(*DeviceManagementVirtualEndpointCloudPCsItemTroubleshootRequestBuilder) {
    return NewDeviceManagementVirtualEndpointCloudPCsItemTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
