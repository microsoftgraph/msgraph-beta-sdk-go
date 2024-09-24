package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type ItemCloudPCsCloudPCItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemCloudPCsCloudPCItemRequestBuilderGetQueryParameters get cloudPCs from users
type ItemCloudPCsCloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemCloudPCsCloudPCItemRequestBuilderGetQueryParameters
}
// ItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeUserAccountType provides operations to call the changeUserAccountType method.
// returns a *ItemCloudPCsItemChangeUserAccountTypeRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) ChangeUserAccountType()(*ItemCloudPCsItemChangeUserAccountTypeRequestBuilder) {
    return NewItemCloudPCsItemChangeUserAccountTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemCloudPCsCloudPCItemRequestBuilderInternal instantiates a new ItemCloudPCsCloudPCItemRequestBuilder and sets the default values.
func NewItemCloudPCsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsCloudPCItemRequestBuilder) {
    m := &ItemCloudPCsCloudPCItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemCloudPCsCloudPCItemRequestBuilder instantiates a new ItemCloudPCsCloudPCItemRequestBuilder and sets the default values.
func NewItemCloudPCsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateSnapshot provides operations to call the createSnapshot method.
// returns a *ItemCloudPCsItemCreateSnapshotRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) CreateSnapshot()(*ItemCloudPCsItemCreateSnapshotRequestBuilder) {
    return NewItemCloudPCsItemCreateSnapshotRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property cloudPCs for users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemCloudPCsItemEndGracePeriodRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) EndGracePeriod()(*ItemCloudPCsItemEndGracePeriodRequestBuilder) {
    return NewItemCloudPCsItemEndGracePeriodRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get cloudPCs from users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a CloudPCable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
// returns a *ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*ItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
// returns a *ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*ItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    return NewItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetFrontlineCloudPcAccessState provides operations to call the getFrontlineCloudPcAccessState method.
// returns a *ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) GetFrontlineCloudPcAccessState()(*ItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilder) {
    return NewItemCloudPCsItemGetFrontlineCloudPcAccessStateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
// returns a *ItemCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) GetSupportedCloudPcRemoteActions()(*ItemCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewItemCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property cloudPCs in users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a CloudPCable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
// returns a *ItemCloudPCsItemPowerOffRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) PowerOff()(*ItemCloudPCsItemPowerOffRequestBuilder) {
    return NewItemCloudPCsItemPowerOffRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PowerOn provides operations to call the powerOn method.
// returns a *ItemCloudPCsItemPowerOnRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) PowerOn()(*ItemCloudPCsItemPowerOnRequestBuilder) {
    return NewItemCloudPCsItemPowerOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reboot provides operations to call the reboot method.
// returns a *ItemCloudPCsItemRebootRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Reboot()(*ItemCloudPCsItemRebootRequestBuilder) {
    return NewItemCloudPCsItemRebootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Rename provides operations to call the rename method.
// returns a *ItemCloudPCsItemRenameRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Rename()(*ItemCloudPCsItemRenameRequestBuilder) {
    return NewItemCloudPCsItemRenameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Reprovision provides operations to call the reprovision method.
// returns a *ItemCloudPCsItemReprovisionRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Reprovision()(*ItemCloudPCsItemReprovisionRequestBuilder) {
    return NewItemCloudPCsItemReprovisionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resize provides operations to call the resize method.
// returns a *ItemCloudPCsItemResizeRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Resize()(*ItemCloudPCsItemResizeRequestBuilder) {
    return NewItemCloudPCsItemResizeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Restore provides operations to call the restore method.
// returns a *ItemCloudPCsItemRestoreRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Restore()(*ItemCloudPCsItemRestoreRequestBuilder) {
    return NewItemCloudPCsItemRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveCloudPCRemoteActionResults provides operations to call the retrieveCloudPCRemoteActionResults method.
// returns a *ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) RetrieveCloudPCRemoteActionResults()(*ItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilder) {
    return NewItemCloudPCsItemRetrieveCloudPCRemoteActionResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveReviewStatus provides operations to call the retrieveReviewStatus method.
// returns a *ItemCloudPCsItemRetrieveReviewStatusRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) RetrieveReviewStatus()(*ItemCloudPCsItemRetrieveReviewStatusRequestBuilder) {
    return NewItemCloudPCsItemRetrieveReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetrieveSnapshots provides operations to call the retrieveSnapshots method.
// returns a *ItemCloudPCsItemRetrieveSnapshotsRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) RetrieveSnapshots()(*ItemCloudPCsItemRetrieveSnapshotsRequestBuilder) {
    return NewItemCloudPCsItemRetrieveSnapshotsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
// returns a *ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) RetryPartnerAgentInstallation()(*ItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) {
    return NewItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SetReviewStatus provides operations to call the setReviewStatus method.
// returns a *ItemCloudPCsItemSetReviewStatusRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) SetReviewStatus()(*ItemCloudPCsItemSetReviewStatusRequestBuilder) {
    return NewItemCloudPCsItemSetReviewStatusRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Start provides operations to call the start method.
// returns a *ItemCloudPCsItemStartRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Start()(*ItemCloudPCsItemStartRequestBuilder) {
    return NewItemCloudPCsItemStartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Stop provides operations to call the stop method.
// returns a *ItemCloudPCsItemStopRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Stop()(*ItemCloudPCsItemStopRequestBuilder) {
    return NewItemCloudPCsItemStopRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property cloudPCs for users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get cloudPCs from users
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudPCsItemTroubleshootRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Troubleshoot()(*ItemCloudPCsItemTroubleshootRequestBuilder) {
    return NewItemCloudPCsItemTroubleshootRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemCloudPCsCloudPCItemRequestBuilder when successful
func (m *ItemCloudPCsCloudPCItemRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsCloudPCItemRequestBuilder) {
    return NewItemCloudPCsCloudPCItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
