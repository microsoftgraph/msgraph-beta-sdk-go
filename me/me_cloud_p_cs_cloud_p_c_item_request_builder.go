package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// MeCloudPCsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type MeCloudPCsCloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// MeCloudPCsCloudPCItemRequestBuilderGetQueryParameters get cloudPCs from me
type MeCloudPCsCloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeCloudPCsCloudPCItemRequestBuilderGetQueryParameters
}
// MeCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeUserAccountType provides operations to call the changeUserAccountType method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) ChangeUserAccountType()(*MeCloudPCsItemChangeUserAccountTypeRequestBuilder) {
    return NewMeCloudPCsItemChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeCloudPCsCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewMeCloudPCsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCloudPCsCloudPCItemRequestBuilder) {
    m := &MeCloudPCsCloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/cloudPCs/{cloudPC%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeCloudPCsCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewMeCloudPCsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeCloudPCsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeCloudPCsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPCs for me
func (m *MeCloudPCsCloudPCItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *MeCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get cloudPCs from me
func (m *MeCloudPCsCloudPCItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property cloudPCs in me
func (m *MeCloudPCsCloudPCItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *MeCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property cloudPCs for me
func (m *MeCloudPCsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *MeCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *MeCloudPCsCloudPCItemRequestBuilder) EndGracePeriod()(*MeCloudPCsItemEndGracePeriodRequestBuilder) {
    return NewMeCloudPCsItemEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get cloudPCs from me
func (m *MeCloudPCsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
func (m *MeCloudPCsCloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*MeCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewMeCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*MeCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    return NewMeCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) GetShiftWorkCloudPcAccessState()(*MeCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewMeCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) GetSupportedCloudPcRemoteActions()(*MeCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewMeCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property cloudPCs in me
func (m *MeCloudPCsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *MeCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
func (m *MeCloudPCsCloudPCItemRequestBuilder) Reboot()(*MeCloudPCsItemRebootRequestBuilder) {
    return NewMeCloudPCsItemRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Rename provides operations to call the rename method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) Rename()(*MeCloudPCsItemRenameRequestBuilder) {
    return NewMeCloudPCsItemRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reprovision provides operations to call the reprovision method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) Reprovision()(*MeCloudPCsItemReprovisionRequestBuilder) {
    return NewMeCloudPCsItemReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) Restore()(*MeCloudPCsItemRestoreRequestBuilder) {
    return NewMeCloudPCsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) RetryPartnerAgentInstallation()(*MeCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) {
    return NewMeCloudPCsItemRetryPartnerAgentInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Troubleshoot provides operations to call the troubleshoot method.
func (m *MeCloudPCsCloudPCItemRequestBuilder) Troubleshoot()(*MeCloudPCsItemTroubleshootRequestBuilder) {
    return NewMeCloudPCsItemTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
