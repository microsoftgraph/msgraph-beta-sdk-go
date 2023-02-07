package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type ItemCloudPCsCloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
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
// NewItemCloudPCsCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewItemCloudPCsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsCloudPCItemRequestBuilder) {
    m := &ItemCloudPCsCloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemCloudPCsCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewItemCloudPCsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property cloudPCs for users
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get cloudPCs from users
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphChangeUserAccountType()(*ItemCloudPCsItemMicrosoftGraphChangeUserAccountTypeRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndGracePeriod provides operations to call the endGracePeriod method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphEndGracePeriod()(*ItemCloudPCsItemMicrosoftGraphEndGracePeriodRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcConnectivityHistory provides operations to call the getCloudPcConnectivityHistory method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetCloudPcConnectivityHistory()(*ItemCloudPCsItemMicrosoftGraphGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphGetCloudPcConnectivityHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetCloudPcLaunchInfo()(*ItemCloudPCsItemMicrosoftGraphGetCloudPcLaunchInfoRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetShiftWorkCloudPcAccessState()(*ItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetSupportedCloudPcRemoteActions()(*ItemCloudPCsItemMicrosoftGraphGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReboot provides operations to call the reboot method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphReboot()(*ItemCloudPCsItemMicrosoftGraphRebootRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRename provides operations to call the rename method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRename()(*ItemCloudPCsItemMicrosoftGraphRenameRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprovision provides operations to call the reprovision method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphReprovision()(*ItemCloudPCsItemMicrosoftGraphReprovisionRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRestore()(*ItemCloudPCsItemMicrosoftGraphRestoreRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRetryPartnerAgentInstallation()(*ItemCloudPCsItemMicrosoftGraphRetryPartnerAgentInstallationRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphRetryPartnerAgentInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTroubleshoot provides operations to call the troubleshoot method.
func (m *ItemCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphTroubleshoot()(*ItemCloudPCsItemMicrosoftGraphTroubleshootRequestBuilder) {
    return NewItemCloudPCsItemMicrosoftGraphTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property cloudPCs in users
func (m *ItemCloudPCsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
// ToDeleteRequestInformation delete navigation property cloudPCs for users
func (m *ItemCloudPCsCloudPCItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation get cloudPCs from users
func (m *ItemCloudPCsCloudPCItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property cloudPCs in users
func (m *ItemCloudPCsCloudPCItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *ItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
