package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VirtualEndpointCloudPCsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
type VirtualEndpointCloudPCsCloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// VirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetQueryParameters cloud managed virtual desktops.
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
// NewVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    m := &VirtualEndpointCloudPCsCloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/virtualEndpoint/cloudPCs/{cloudPC%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewVirtualEndpointCloudPCsCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewVirtualEndpointCloudPCsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VirtualEndpointCloudPCsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVirtualEndpointCloudPCsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property cloudPCs for deviceManagement
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get cloud managed virtual desktops.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphChangeUserAccountType()(*VirtualEndpointCloudPCsItemMicrosoftGraphChangeUserAccountTypeRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphEndGracePeriod provides operations to call the endGracePeriod method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphEndGracePeriod()(*VirtualEndpointCloudPCsItemMicrosoftGraphEndGracePeriodRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcConnectivityHistory provides operations to call the getCloudPcConnectivityHistory method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetCloudPcConnectivityHistory()(*VirtualEndpointCloudPCsItemMicrosoftGraphGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphGetCloudPcConnectivityHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetCloudPcLaunchInfo()(*VirtualEndpointCloudPCsItemMicrosoftGraphGetCloudPcLaunchInfoRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetShiftWorkCloudPcAccessState()(*VirtualEndpointCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphGetSupportedCloudPcRemoteActions()(*VirtualEndpointCloudPCsItemMicrosoftGraphGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReboot provides operations to call the reboot method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphReboot()(*VirtualEndpointCloudPCsItemMicrosoftGraphRebootRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRename provides operations to call the rename method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRename()(*VirtualEndpointCloudPCsItemMicrosoftGraphRenameRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphReprovision provides operations to call the reprovision method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphReprovision()(*VirtualEndpointCloudPCsItemMicrosoftGraphReprovisionRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRestore provides operations to call the restore method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRestore()(*VirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphRetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphRetryPartnerAgentInstallation()(*VirtualEndpointCloudPCsItemMicrosoftGraphRetryPartnerAgentInstallationRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphRetryPartnerAgentInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphTroubleshoot provides operations to call the troubleshoot method.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) MicrosoftGraphTroubleshoot()(*VirtualEndpointCloudPCsItemMicrosoftGraphTroubleshootRequestBuilder) {
    return NewVirtualEndpointCloudPCsItemMicrosoftGraphTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property cloudPCs in deviceManagement
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
// ToDeleteRequestInformation delete navigation property cloudPCs for deviceManagement
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation cloud managed virtual desktops.
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property cloudPCs in deviceManagement
func (m *VirtualEndpointCloudPCsCloudPCItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *VirtualEndpointCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
