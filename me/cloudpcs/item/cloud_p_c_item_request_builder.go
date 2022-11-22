package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i0543ab08206c491078113b4b6ec824e6ea4dc3659ea1392f05efc163b50bac03 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/changeuseraccounttype"
    i3fc1b0e9b520291a737c5354479f5ce10f5650dd33030dfefdc4b6ee1d1af0a6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/reboot"
    i48833d47a3b49cc67b08b792c187d1d2e754160e7d01e8291dfc3500786630e6 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/troubleshoot"
    i88cd15ed11dc17187b0e27e7a33a8beec99e35253eff3491796a0b40e3081054 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/restore"
    ib05f8eb4a398c1197e5500e337b98f29806e31c0d85a4c4d513aa425dcdbbc08 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/getcloudpcconnectivityhistory"
    ibd3ed8e86fd7912f38df4a5ee348162bbd2f3faee2c4688adbdf41aed4563b2d "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/getcloudpclaunchinfo"
    ic2801aae9bbfd9e2fc587c7534a8202e1104a4e45684db61ac6d578a3aa2d66c "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/getshiftworkcloudpcaccessstate"
    id59d5f3bf71ded12222a12943edebd3dcac00295833811722cfa0838178ac8ca "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/retrypartneragentinstallation"
    idd692d448454691ef652a1555dc60f805306ee63076a45f51d1f0e45253c9964 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/rename"
    ie1b3b83917cdcc75511ad5eb6d6dbf8d5c29ca730fb1a3e4c356599c7261e842 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/endgraceperiod"
    ie4696cde156288c3f3bd9ecdc93adfc93dc5d8beeb9cceac7e81177b09b725a3 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/reprovision"
    ie88847fb120b95088657195eed01ca0ea1c02640776d8ba96966309669339590 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/cloudpcs/item/getsupportedcloudpcremoteactions"
)

// CloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type CloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CloudPCItemRequestBuilderGetQueryParameters get cloudPCs from me
type CloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CloudPCItemRequestBuilderGetQueryParameters
}
// CloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeUserAccountType provides operations to call the changeUserAccountType method.
func (m *CloudPCItemRequestBuilder) ChangeUserAccountType()(*i0543ab08206c491078113b4b6ec824e6ea4dc3659ea1392f05efc163b50bac03.ChangeUserAccountTypeRequestBuilder) {
    return i0543ab08206c491078113b4b6ec824e6ea4dc3659ea1392f05efc163b50bac03.NewChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCItemRequestBuilder) {
    m := &CloudPCItemRequestBuilder{
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
// NewCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPCs for me
func (m *CloudPCItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *CloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CloudPCItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CloudPCItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *CloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *CloudPCItemRequestBuilder) EndGracePeriod()(*ie1b3b83917cdcc75511ad5eb6d6dbf8d5c29ca730fb1a3e4c356599c7261e842.EndGracePeriodRequestBuilder) {
    return ie1b3b83917cdcc75511ad5eb6d6dbf8d5c29ca730fb1a3e4c356599c7261e842.NewEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get cloudPCs from me
func (m *CloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
func (m *CloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*ib05f8eb4a398c1197e5500e337b98f29806e31c0d85a4c4d513aa425dcdbbc08.GetCloudPcConnectivityHistoryRequestBuilder) {
    return ib05f8eb4a398c1197e5500e337b98f29806e31c0d85a4c4d513aa425dcdbbc08.NewGetCloudPcConnectivityHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *CloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*ibd3ed8e86fd7912f38df4a5ee348162bbd2f3faee2c4688adbdf41aed4563b2d.GetCloudPcLaunchInfoRequestBuilder) {
    return ibd3ed8e86fd7912f38df4a5ee348162bbd2f3faee2c4688adbdf41aed4563b2d.NewGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
func (m *CloudPCItemRequestBuilder) GetShiftWorkCloudPcAccessState()(*ic2801aae9bbfd9e2fc587c7534a8202e1104a4e45684db61ac6d578a3aa2d66c.GetShiftWorkCloudPcAccessStateRequestBuilder) {
    return ic2801aae9bbfd9e2fc587c7534a8202e1104a4e45684db61ac6d578a3aa2d66c.NewGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
func (m *CloudPCItemRequestBuilder) GetSupportedCloudPcRemoteActions()(*ie88847fb120b95088657195eed01ca0ea1c02640776d8ba96966309669339590.GetSupportedCloudPcRemoteActionsRequestBuilder) {
    return ie88847fb120b95088657195eed01ca0ea1c02640776d8ba96966309669339590.NewGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property cloudPCs in me
func (m *CloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *CloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
func (m *CloudPCItemRequestBuilder) Reboot()(*i3fc1b0e9b520291a737c5354479f5ce10f5650dd33030dfefdc4b6ee1d1af0a6.RebootRequestBuilder) {
    return i3fc1b0e9b520291a737c5354479f5ce10f5650dd33030dfefdc4b6ee1d1af0a6.NewRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Rename provides operations to call the rename method.
func (m *CloudPCItemRequestBuilder) Rename()(*idd692d448454691ef652a1555dc60f805306ee63076a45f51d1f0e45253c9964.RenameRequestBuilder) {
    return idd692d448454691ef652a1555dc60f805306ee63076a45f51d1f0e45253c9964.NewRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reprovision provides operations to call the reprovision method.
func (m *CloudPCItemRequestBuilder) Reprovision()(*ie4696cde156288c3f3bd9ecdc93adfc93dc5d8beeb9cceac7e81177b09b725a3.ReprovisionRequestBuilder) {
    return ie4696cde156288c3f3bd9ecdc93adfc93dc5d8beeb9cceac7e81177b09b725a3.NewReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *CloudPCItemRequestBuilder) Restore()(*i88cd15ed11dc17187b0e27e7a33a8beec99e35253eff3491796a0b40e3081054.RestoreRequestBuilder) {
    return i88cd15ed11dc17187b0e27e7a33a8beec99e35253eff3491796a0b40e3081054.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
func (m *CloudPCItemRequestBuilder) RetryPartnerAgentInstallation()(*id59d5f3bf71ded12222a12943edebd3dcac00295833811722cfa0838178ac8ca.RetryPartnerAgentInstallationRequestBuilder) {
    return id59d5f3bf71ded12222a12943edebd3dcac00295833811722cfa0838178ac8ca.NewRetryPartnerAgentInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Troubleshoot provides operations to call the troubleshoot method.
func (m *CloudPCItemRequestBuilder) Troubleshoot()(*i48833d47a3b49cc67b08b792c187d1d2e754160e7d01e8291dfc3500786630e6.TroubleshootRequestBuilder) {
    return i48833d47a3b49cc67b08b792c187d1d2e754160e7d01e8291dfc3500786630e6.NewTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
