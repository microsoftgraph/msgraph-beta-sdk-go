package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CloudPCsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type CloudPCsCloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CloudPCsCloudPCItemRequestBuilderGetQueryParameters get cloudPCs from me
type CloudPCsCloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CloudPCsCloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCsCloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudPCsCloudPCItemRequestBuilderGetQueryParameters
}
// CloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCloudPCsCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, cloudPCId *string)(*CloudPCsCloudPCItemRequestBuilder) {
    m := &CloudPCsCloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/cloudPCs/{cloudPC%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if cloudPCId != nil {
        urlTplParams["cloudPC%2Did"] = *cloudPCId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewCloudPCsCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property cloudPCs for me
func (m *CloudPCsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get cloudPCs from me
func (m *CloudPCsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable), nil
}
// MicrosoftGraphChangeUserAccountType provides operations to call the changeUserAccountType method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphChangeUserAccountType()(*CloudPCsItemMicrosoftGraphChangeUserAccountTypeChangeUserAccountTypeRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphChangeUserAccountTypeChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndGracePeriod provides operations to call the endGracePeriod method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphEndGracePeriod()(*CloudPCsItemMicrosoftGraphEndGracePeriodEndGracePeriodRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphEndGracePeriodEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcConnectivityHistory provides operations to call the getCloudPcConnectivityHistory method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetCloudPcConnectivityHistory()(*CloudPCsItemMicrosoftGraphGetCloudPcConnectivityHistoryGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphGetCloudPcConnectivityHistoryGetCloudPcConnectivityHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetCloudPcLaunchInfo()(*CloudPCsItemMicrosoftGraphGetCloudPcLaunchInfoGetCloudPcLaunchInfoRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphGetCloudPcLaunchInfoGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetShiftWorkCloudPcAccessState()(*CloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetSupportedCloudPcRemoteActions()(*CloudPCsItemMicrosoftGraphGetSupportedCloudPcRemoteActionsGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphGetSupportedCloudPcRemoteActionsGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReboot provides operations to call the reboot method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphReboot()(*CloudPCsItemMicrosoftGraphRebootRebootRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphRebootRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRename provides operations to call the rename method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRename()(*CloudPCsItemMicrosoftGraphRenameRenameRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphRenameRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprovision provides operations to call the reprovision method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphReprovision()(*CloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphReprovisionReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRestore()(*CloudPCsItemMicrosoftGraphRestoreRestoreRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphRestoreRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRetryPartnerAgentInstallation()(*CloudPCsItemMicrosoftGraphRetryPartnerAgentInstallationRetryPartnerAgentInstallationRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphRetryPartnerAgentInstallationRetryPartnerAgentInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTroubleshoot provides operations to call the troubleshoot method.
func (m *CloudPCsCloudPCItemRequestBuilder) MicrosoftGraphTroubleshoot()(*CloudPCsItemMicrosoftGraphTroubleshootTroubleshootRequestBuilder) {
    return NewCloudPCsItemMicrosoftGraphTroubleshootTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property cloudPCs in me
func (m *CloudPCsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *CloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCloudPCFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable), nil
}
// ToDeleteRequestInformation delete navigation property cloudPCs for me
func (m *CloudPCsCloudPCItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *CloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get cloudPCs from me
func (m *CloudPCsCloudPCItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property cloudPCs in me
func (m *CloudPCsCloudPCItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *CloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
