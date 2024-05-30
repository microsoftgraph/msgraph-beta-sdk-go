package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualendpointCloudpcsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
type VirtualendpointCloudpcsCloudPCItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VirtualendpointCloudpcsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualendpointCloudpcsCloudPCItemRequestBuilderGetQueryParameters read the properties and relationships of a specific cloudPC object.
type VirtualendpointCloudpcsCloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// VirtualendpointCloudpcsCloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsCloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VirtualendpointCloudpcsCloudPCItemRequestBuilderGetQueryParameters
}
// VirtualendpointCloudpcsCloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualendpointCloudpcsCloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeUserAccountType provides operations to call the changeUserAccountType method.
// returns a *VirtualendpointCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) ChangeUserAccountType()(*VirtualendpointCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder) {
    return NewVirtualendpointCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewVirtualendpointCloudpcsCloudPCItemRequestBuilderInternal instantiates a new VirtualendpointCloudpcsCloudPCItemRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsCloudPCItemRequestBuilder) {
    m := &VirtualendpointCloudpcsCloudPCItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewVirtualendpointCloudpcsCloudPCItemRequestBuilder instantiates a new VirtualendpointCloudpcsCloudPCItemRequestBuilder and sets the default values.
func NewVirtualendpointCloudpcsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualendpointCloudpcsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualendpointCloudpcsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateSnapshot provides operations to call the createSnapshot method.
// returns a *VirtualendpointCloudpcsItemCreatesnapshotCreateSnapshotRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) CreateSnapshot()(*VirtualendpointCloudpcsItemCreatesnapshotCreateSnapshotRequestBuilder) {
    return NewVirtualendpointCloudpcsItemCreatesnapshotCreateSnapshotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property cloudPCs for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) EndGracePeriod()(*VirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder) {
    return NewVirtualendpointCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a specific cloudPC object.
// returns a CloudPCable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-get?view=graph-rest-beta
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
// returns a *VirtualendpointCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*VirtualendpointCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewVirtualendpointCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
// returns a *VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*VirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder) {
    return NewVirtualendpointCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFrontlineCloudPcAccessState provides operations to call the getFrontlineCloudPcAccessState method.
// returns a *VirtualendpointCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) GetFrontlineCloudPcAccessState()(*VirtualendpointCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder) {
    return NewVirtualendpointCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
// returns a *VirtualendpointCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) GetShiftWorkCloudPcAccessState()(*VirtualendpointCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewVirtualendpointCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
// returns a *VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) GetSupportedCloudPcRemoteActions()(*VirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewVirtualendpointCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property cloudPCs in deviceManagement
// returns a CloudPCable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *VirtualendpointCloudpcsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
// returns a *VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) PowerOff()(*VirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilder) {
    return NewVirtualendpointCloudpcsItemPoweroffPowerOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PowerOn provides operations to call the powerOn method.
// returns a *VirtualendpointCloudpcsItemPoweronPowerOnRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) PowerOn()(*VirtualendpointCloudpcsItemPoweronPowerOnRequestBuilder) {
    return NewVirtualendpointCloudpcsItemPoweronPowerOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reboot provides operations to call the reboot method.
// returns a *VirtualendpointCloudpcsItemRebootRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Reboot()(*VirtualendpointCloudpcsItemRebootRequestBuilder) {
    return NewVirtualendpointCloudpcsItemRebootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rename provides operations to call the rename method.
// returns a *VirtualendpointCloudpcsItemRenameRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Rename()(*VirtualendpointCloudpcsItemRenameRequestBuilder) {
    return NewVirtualendpointCloudpcsItemRenameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reprovision provides operations to call the reprovision method.
// returns a *VirtualendpointCloudpcsItemReprovisionRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Reprovision()(*VirtualendpointCloudpcsItemReprovisionRequestBuilder) {
    return NewVirtualendpointCloudpcsItemReprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resize provides operations to call the resize method.
// returns a *VirtualendpointCloudpcsItemResizeRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Resize()(*VirtualendpointCloudpcsItemResizeRequestBuilder) {
    return NewVirtualendpointCloudpcsItemResizeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *VirtualendpointCloudpcsItemRestoreRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Restore()(*VirtualendpointCloudpcsItemRestoreRequestBuilder) {
    return NewVirtualendpointCloudpcsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveReviewStatus provides operations to call the retrieveReviewStatus method.
// returns a *VirtualendpointCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) RetrieveReviewStatus()(*VirtualendpointCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder) {
    return NewVirtualendpointCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
// returns a *VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) RetryPartnerAgentInstallation()(*VirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder) {
    return NewVirtualendpointCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetReviewStatus provides operations to call the setReviewStatus method.
// returns a *VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) SetReviewStatus()(*VirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder) {
    return NewVirtualendpointCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start provides operations to call the start method.
// returns a *VirtualendpointCloudpcsItemStartRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Start()(*VirtualendpointCloudpcsItemStartRequestBuilder) {
    return NewVirtualendpointCloudpcsItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *VirtualendpointCloudpcsItemStopRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Stop()(*VirtualendpointCloudpcsItemStopRequestBuilder) {
    return NewVirtualendpointCloudpcsItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property cloudPCs for deviceManagement
// returns a *RequestInformation when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualendpointCloudpcsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *VirtualendpointCloudpcsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VirtualendpointCloudpcsItemTroubleshootRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) Troubleshoot()(*VirtualendpointCloudpcsItemTroubleshootRequestBuilder) {
    return NewVirtualendpointCloudpcsItemTroubleshootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VirtualendpointCloudpcsCloudPCItemRequestBuilder when successful
func (m *VirtualendpointCloudpcsCloudPCItemRequestBuilder) WithUrl(rawUrl string)(*VirtualendpointCloudpcsCloudPCItemRequestBuilder) {
    return NewVirtualendpointCloudpcsCloudPCItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
