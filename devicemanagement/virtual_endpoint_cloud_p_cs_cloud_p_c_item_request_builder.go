package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointCloudPCsCloudPCItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetQueryParameters read the properties and relationships of a specific cloudPC object.
type VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetQueryParameters
}
// VirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeUserAccountType provides operations to call the changeUserAccountType method.
// returns a *VirtualEndpointCloudPCsItemChangeUserAccountTypeRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) ChangeUserAccountType()(*VirtualEndpointCloudPCsItemChangeUserAccountTypeRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemChangeUserAccountTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal instantiates a new VirtualEndpointCloudPCsCloudPCItemRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    m := &VirtualEndpointCloudPCsCloudPCItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualEndpointCloudPCsCloudPCItemRequestBuilder instantiates a new VirtualEndpointCloudPCsCloudPCItemRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateSnapshot provides operations to call the createSnapshot method.
// returns a *VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) CreateSnapshot()(*VirtualEndpointCloudPCsItemCreateSnapshotRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemCreateSnapshotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property cloudPCs for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// EndGracePeriod provides operations to call the endGracePeriod method.
// returns a *VirtualEndpointCloudPCsItemEndGracePeriodRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) EndGracePeriod()(*VirtualEndpointCloudPCsItemEndGracePeriodRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemEndGracePeriodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a specific cloudPC object.
// returns a CloudPCable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-get?view=graph-rest-beta
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable), nil
}
// GetCloudPcConnectivityHistory provides operations to call the getCloudPcConnectivityHistory method.
// returns a *VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*VirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
// returns a *VirtualEndpointCloudPCsItemGetCloudPcLaunchInfoRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*VirtualEndpointCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFrontlineCloudPcAccessState provides operations to call the getFrontlineCloudPcAccessState method.
// returns a *VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetFrontlineCloudPcAccessState()(*VirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
// returns a *VirtualEndpointCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetShiftWorkCloudPcAccessState()(*VirtualEndpointCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
// returns a *VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) GetSupportedCloudPcRemoteActions()(*VirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property cloudPCs in deviceManagement
// returns a CloudPCable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable), nil
}
// PowerOff provides operations to call the powerOff method.
// returns a *VirtualEndpointCloudPCsItemPowerOffRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) PowerOff()(*VirtualEndpointCloudPCsItemPowerOffRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemPowerOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PowerOn provides operations to call the powerOn method.
// returns a *VirtualEndpointCloudPCsItemPowerOnRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) PowerOn()(*VirtualEndpointCloudPCsItemPowerOnRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemPowerOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reboot provides operations to call the reboot method.
// returns a *VirtualEndpointCloudPCsItemRebootRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Reboot()(*VirtualEndpointCloudPCsItemRebootRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemRebootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rename provides operations to call the rename method.
// returns a *VirtualEndpointCloudPCsItemRenameRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Rename()(*VirtualEndpointCloudPCsItemRenameRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemRenameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reprovision provides operations to call the reprovision method.
// returns a *VirtualEndpointCloudPCsItemReprovisionRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Reprovision()(*VirtualEndpointCloudPCsItemReprovisionRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemReprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resize provides operations to call the resize method.
// returns a *VirtualEndpointCloudPCsItemResizeRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Resize()(*VirtualEndpointCloudPCsItemResizeRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemResizeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *VirtualEndpointCloudPCsItemRestoreRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Restore()(*VirtualEndpointCloudPCsItemRestoreRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
// returns a *VirtualEndpointCloudPCsItemRetryPartnerAgentInstallationRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) RetryPartnerAgentInstallation()(*VirtualEndpointCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemRetryPartnerAgentInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start provides operations to call the start method.
// returns a *VirtualEndpointCloudPCsItemStartRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Start()(*VirtualEndpointCloudPCsItemStartRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *VirtualEndpointCloudPCsItemStopRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Stop()(*VirtualEndpointCloudPCsItemStopRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property cloudPCs for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a specific cloudPC object.
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property cloudPCs in deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Troubleshoot provides operations to call the troubleshoot method.
// returns a *VirtualEndpointCloudPCsItemTroubleshootRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Troubleshoot()(*VirtualEndpointCloudPCsItemTroubleshootRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemTroubleshootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualEndpointCloudPCsCloudPCItemRequestBuilder when successful
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) WithUrl(rawUrl string)(*VirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    return NewVirtualEndpointCloudPCsCloudPCItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
