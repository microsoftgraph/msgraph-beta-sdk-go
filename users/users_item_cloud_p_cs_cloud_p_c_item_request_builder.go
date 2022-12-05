package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemCloudPCsCloudPCItemRequestBuilder provides operations to manage the cloudPCs property of the microsoft.graph.user entity.
type UsersItemCloudPCsCloudPCItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemCloudPCsCloudPCItemRequestBuilderGetQueryParameters get cloudPCs from users
type UsersItemCloudPCsCloudPCItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemCloudPCsCloudPCItemRequestBuilderGetQueryParameters
}
// UsersItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChangeUserAccountType provides operations to call the changeUserAccountType method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) ChangeUserAccountType()(*UsersItemCloudPCsItemChangeUserAccountTypeRequestBuilder) {
    return NewUsersItemCloudPCsItemChangeUserAccountTypeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemCloudPCsCloudPCItemRequestBuilderInternal instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewUsersItemCloudPCsCloudPCItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCloudPCsCloudPCItemRequestBuilder) {
    m := &UsersItemCloudPCsCloudPCItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemCloudPCsCloudPCItemRequestBuilder instantiates a new CloudPCItemRequestBuilder and sets the default values.
func NewUsersItemCloudPCsCloudPCItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemCloudPCsCloudPCItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemCloudPCsCloudPCItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property cloudPCs for users
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get cloudPCs from users
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property cloudPCs in users
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *UsersItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property cloudPCs for users
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemCloudPCsCloudPCItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) EndGracePeriod()(*UsersItemCloudPCsItemEndGracePeriodRequestBuilder) {
    return NewUsersItemCloudPCsItemEndGracePeriodRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get cloudPCs from users
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemCloudPCsCloudPCItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) GetCloudPcConnectivityHistory()(*UsersItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilder) {
    return NewUsersItemCloudPCsItemGetCloudPcConnectivityHistoryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetCloudPcLaunchInfo provides operations to call the getCloudPcLaunchInfo method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) GetCloudPcLaunchInfo()(*UsersItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilder) {
    return NewUsersItemCloudPCsItemGetCloudPcLaunchInfoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetShiftWorkCloudPcAccessState provides operations to call the getShiftWorkCloudPcAccessState method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) GetShiftWorkCloudPcAccessState()(*UsersItemCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilder) {
    return NewUsersItemCloudPCsItemGetShiftWorkCloudPcAccessStateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetSupportedCloudPcRemoteActions provides operations to call the getSupportedCloudPcRemoteActions method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) GetSupportedCloudPcRemoteActions()(*UsersItemCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilder) {
    return NewUsersItemCloudPCsItemGetSupportedCloudPcRemoteActionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property cloudPCs in users
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, requestConfiguration *UsersItemCloudPCsCloudPCItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CloudPCable, error) {
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
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) Reboot()(*UsersItemCloudPCsItemRebootRequestBuilder) {
    return NewUsersItemCloudPCsItemRebootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Rename provides operations to call the rename method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) Rename()(*UsersItemCloudPCsItemRenameRequestBuilder) {
    return NewUsersItemCloudPCsItemRenameRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reprovision provides operations to call the reprovision method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) Reprovision()(*UsersItemCloudPCsItemReprovisionRequestBuilder) {
    return NewUsersItemCloudPCsItemReprovisionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Restore provides operations to call the restore method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) Restore()(*UsersItemCloudPCsItemRestoreRequestBuilder) {
    return NewUsersItemCloudPCsItemRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RetryPartnerAgentInstallation provides operations to call the retryPartnerAgentInstallation method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) RetryPartnerAgentInstallation()(*UsersItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilder) {
    return NewUsersItemCloudPCsItemRetryPartnerAgentInstallationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Troubleshoot provides operations to call the troubleshoot method.
func (m *UsersItemCloudPCsCloudPCItemRequestBuilder) Troubleshoot()(*UsersItemCloudPCsItemTroubleshootRequestBuilder) {
    return NewUsersItemCloudPCsItemTroubleshootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
