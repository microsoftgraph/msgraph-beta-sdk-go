package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i13ec1c9a15a271cc8f1a63c29667b2eb16fa66895eadeae5018091e29071ea9d "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/changeuseraccounttype"
    i170388c68b05b462f70fd530877e5751fb8c26c036b698799100779eefe3754a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/rename"
    i3ea7449e4e31fbb2dba44ba96548c03bf21e30fa0fbc867460aef4206de9a3ab "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/reboot"
    i6e8302b9072d34b3ea0f98fa81a45d53c8e894280a606660773b4ade2bc98d42 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/getcloudpclaunchinfo"
    i7174971b3bcb6f4ec85df581ff277b07bbaf56b391c32a3f561718b3a00346fe "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/troubleshoot"
    ia073266fa87f992d3a858948753dc53b6ac9007c5febc8b5d74af59df63531cf "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/reprovision"
    iafdeeea6f76d4e54cc8d64182d5ef5738139e7487d31b19e0f8cd67313dd907f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/getcloudpcconnectivityhistory"
    ic5c3712fbfc3279ed202453cc12bb33fc813457da47b3438a76bf9022c5b2f4a "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/endgraceperiod"
    idf1c60771e98e3ab68c53c8c8b81360b0ea19beb273552bc7b02f0d6a691436f "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/virtualendpoint/cloudpcs/item/restore"
)

// CloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.virtualEndpoint entity.
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
// CloudPCItemRequestBuilderGetQueryParameters cloud managed virtual desktops.
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
// ChangeUserAccountType the changeUserAccountType property
func (m *CloudPCItemRequestBuilder) ChangeUserAccountType()(*i13ec1c9a15a271cc8f1a63c29667b2eb16fa66895eadeae5018091e29071ea9d.ChangeUserAccountTypeRequestBuilder) {
    return i13ec1c9a15a271cc8f1a63c29667b2eb16fa66895eadeae5018091e29071ea9d.NewChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCItemRequestBuilder) {
    m := &CloudPCItemRequestBuilder{
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
// NewCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPCs for deviceManagement
func (m *CloudPCItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property cloudPCs for deviceManagement
func (m *CloudPCItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *CloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CloudPCItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration cloud managed virtual desktops.
func (m *CloudPCItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CloudPCItemRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property cloudPCs in deviceManagement
func (m *CloudPCItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *CloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property cloudPCs for deviceManagement
func (m *CloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
// EndGracePeriod the endGracePeriod property
func (m *CloudPCItemRequestBuilder) EndGracePeriod()(*ic5c3712fbfc3279ed202453cc12bb33fc813457da47b3438a76bf9022c5b2f4a.EndGracePeriodRequestBuilder) {
    return ic5c3712fbfc3279ed202453cc12bb33fc813457da47b3438a76bf9022c5b2f4a.NewEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get cloud managed virtual desktops.
func (m *CloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *CloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*iafdeeea6f76d4e54cc8d64182d5ef5738139e7487d31b19e0f8cd67313dd907f.GetCloudPcConnectivityHistoryRequestBuilder) {
    return iafdeeea6f76d4e54cc8d64182d5ef5738139e7487d31b19e0f8cd67313dd907f.NewGetCloudPcConnectivityHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *CloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*i6e8302b9072d34b3ea0f98fa81a45d53c8e894280a606660773b4ade2bc98d42.GetCloudPcLaunchInfoRequestBuilder) {
    return i6e8302b9072d34b3ea0f98fa81a45d53c8e894280a606660773b4ade2bc98d42.NewGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property cloudPCs in deviceManagement
func (m *CloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *CloudPCItemRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Reboot the reboot property
func (m *CloudPCItemRequestBuilder) Reboot()(*i3ea7449e4e31fbb2dba44ba96548c03bf21e30fa0fbc867460aef4206de9a3ab.RebootRequestBuilder) {
    return i3ea7449e4e31fbb2dba44ba96548c03bf21e30fa0fbc867460aef4206de9a3ab.NewRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Rename the rename property
func (m *CloudPCItemRequestBuilder) Rename()(*i170388c68b05b462f70fd530877e5751fb8c26c036b698799100779eefe3754a.RenameRequestBuilder) {
    return i170388c68b05b462f70fd530877e5751fb8c26c036b698799100779eefe3754a.NewRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reprovision the reprovision property
func (m *CloudPCItemRequestBuilder) Reprovision()(*ia073266fa87f992d3a858948753dc53b6ac9007c5febc8b5d74af59df63531cf.ReprovisionRequestBuilder) {
    return ia073266fa87f992d3a858948753dc53b6ac9007c5febc8b5d74af59df63531cf.NewReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore the restore property
func (m *CloudPCItemRequestBuilder) Restore()(*idf1c60771e98e3ab68c53c8c8b81360b0ea19beb273552bc7b02f0d6a691436f.RestoreRequestBuilder) {
    return idf1c60771e98e3ab68c53c8c8b81360b0ea19beb273552bc7b02f0d6a691436f.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Troubleshoot the troubleshoot property
func (m *CloudPCItemRequestBuilder) Troubleshoot()(*i7174971b3bcb6f4ec85df581ff277b07bbaf56b391c32a3f561718b3a00346fe.TroubleshootRequestBuilder) {
    return i7174971b3bcb6f4ec85df581ff277b07bbaf56b391c32a3f561718b3a00346fe.NewTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
