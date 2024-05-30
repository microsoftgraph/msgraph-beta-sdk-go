package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type ItemCloudpcsCloudPCItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCloudpcsCloudPCItemRequestBuilderGetQueryParameters get cloudPCs from users
type ItemCloudpcsCloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCloudpcsCloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsCloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudpcsCloudPCItemRequestBuilderGetQueryParameters
}
// ItemCloudpcsCloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsCloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeUserAccountType provides operations to call the changeUserAccountType method.
// returns a *ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) ChangeUserAccountType()(*ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder) {
    return NewItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCloudpcsCloudPCItemRequestBuilderInternal instantiates a new ItemCloudpcsCloudPCItemRequestBuilder and sets the default values.
func NewItemCloudpcsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsCloudPCItemRequestBuilder) {
    m := &ItemCloudpcsCloudPCItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCloudpcsCloudPCItemRequestBuilder instantiates a new ItemCloudpcsCloudPCItemRequestBuilder and sets the default values.
func NewItemCloudpcsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateSnapshot provides operations to call the createSnapshot method.
// returns a *ItemCloudpcsItemCreatesnapshotCreateSnapshotRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) CreateSnapshot()(*ItemCloudpcsItemCreatesnapshotCreateSnapshotRequestBuilder) {
    return NewItemCloudpcsItemCreatesnapshotCreateSnapshotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property cloudPCs for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCloudpcsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) EndGracePeriod()(*ItemCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilder) {
    return NewItemCloudpcsItemEndgraceperiodEndGracePeriodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get cloudPCs from users
// returns a CloudPCable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudpcsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
// returns a *ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*ItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewItemCloudpcsItemGetcloudpcconnectivityhistoryGetCloudPcConnectivityHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
// returns a *ItemCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*ItemCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilder) {
    return NewItemCloudpcsItemGetcloudpclaunchinfoGetCloudPcLaunchInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFrontlineCloudPcAccessState provides operations to call the getFrontlineCloudPcAccessState method.
// returns a *ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) GetFrontlineCloudPcAccessState()(*ItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilder) {
    return NewItemCloudpcsItemGetfrontlinecloudpcaccessstateGetFrontlineCloudPcAccessStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
// returns a *ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) GetShiftWorkCloudPcAccessState()(*ItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewItemCloudpcsItemGetshiftworkcloudpcaccessstateGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
// returns a *ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) GetSupportedCloudPcRemoteActions()(*ItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewItemCloudpcsItemGetsupportedcloudpcremoteactionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property cloudPCs in users
// returns a CloudPCable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *ItemCloudpcsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
// returns a *ItemCloudpcsItemPoweroffPowerOffRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) PowerOff()(*ItemCloudpcsItemPoweroffPowerOffRequestBuilder) {
    return NewItemCloudpcsItemPoweroffPowerOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PowerOn provides operations to call the powerOn method.
// returns a *ItemCloudpcsItemPoweronPowerOnRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) PowerOn()(*ItemCloudpcsItemPoweronPowerOnRequestBuilder) {
    return NewItemCloudpcsItemPoweronPowerOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reboot provides operations to call the reboot method.
// returns a *ItemCloudpcsItemRebootRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Reboot()(*ItemCloudpcsItemRebootRequestBuilder) {
    return NewItemCloudpcsItemRebootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rename provides operations to call the rename method.
// returns a *ItemCloudpcsItemRenameRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Rename()(*ItemCloudpcsItemRenameRequestBuilder) {
    return NewItemCloudpcsItemRenameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reprovision provides operations to call the reprovision method.
// returns a *ItemCloudpcsItemReprovisionRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Reprovision()(*ItemCloudpcsItemReprovisionRequestBuilder) {
    return NewItemCloudpcsItemReprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resize provides operations to call the resize method.
// returns a *ItemCloudpcsItemResizeRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Resize()(*ItemCloudpcsItemResizeRequestBuilder) {
    return NewItemCloudpcsItemResizeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *ItemCloudpcsItemRestoreRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Restore()(*ItemCloudpcsItemRestoreRequestBuilder) {
    return NewItemCloudpcsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveReviewStatus provides operations to call the retrieveReviewStatus method.
// returns a *ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) RetrieveReviewStatus()(*ItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilder) {
    return NewItemCloudpcsItemRetrievereviewstatusRetrieveReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
// returns a *ItemCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) RetryPartnerAgentInstallation()(*ItemCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilder) {
    return NewItemCloudpcsItemRetrypartneragentinstallationRetryPartnerAgentInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetReviewStatus provides operations to call the setReviewStatus method.
// returns a *ItemCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) SetReviewStatus()(*ItemCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilder) {
    return NewItemCloudpcsItemSetreviewstatusSetReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start provides operations to call the start method.
// returns a *ItemCloudpcsItemStartRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Start()(*ItemCloudpcsItemStartRequestBuilder) {
    return NewItemCloudpcsItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *ItemCloudpcsItemStopRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Stop()(*ItemCloudpcsItemStopRequestBuilder) {
    return NewItemCloudpcsItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property cloudPCs for users
// returns a *RequestInformation when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCloudpcsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get cloudPCs from users
// returns a *RequestInformation when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudpcsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property cloudPCs in users
// returns a *RequestInformation when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *ItemCloudpcsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudpcsItemTroubleshootRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) Troubleshoot()(*ItemCloudpcsItemTroubleshootRequestBuilder) {
    return NewItemCloudpcsItemTroubleshootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemCloudpcsCloudPCItemRequestBuilder when successful
func (m *ItemCloudpcsCloudPCItemRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsCloudPCItemRequestBuilder) {
    return NewItemCloudpcsCloudPCItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
