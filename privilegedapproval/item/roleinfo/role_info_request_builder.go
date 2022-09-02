package roleinfo

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i09a67af6e9bae1dd8dfd5d4374836f24ecdb5be20e1fae73a258e777e97458b4 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval/item/roleinfo/selfdeactivate"
    i5cc87848e96d90dbe672f3e2fe6b8ec5e28e3cdc856ef730de729562e22b9524 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval/item/roleinfo/selfactivate"
    i8b106f89a5024f1de5fce1ceb14aee752150a35a2f3ead41de7233973c9e2ec7 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval/item/roleinfo/settings"
    i920978860cb85582d4cc403166c2e16a9faae13fa4b1bd4c748c659db6073a7f "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval/item/roleinfo/summary"
    ib5cd5b3c19268d246febcc379cad7c6c55d83a8e91bd1b806e4aeced1e9bcbf0 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval/item/roleinfo/assignments"
    i91831fb3bb284dcc082f0d50962d9b4761f0d7da6397b6aaa15e6aadeeef1c31 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedapproval/item/roleinfo/assignments/item"
)

// RoleInfoRequestBuilder provides operations to manage the roleInfo property of the microsoft.graph.privilegedApproval entity.
type RoleInfoRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleInfoRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleInfoRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RoleInfoRequestBuilderGetQueryParameters get roleInfo from privilegedApproval
type RoleInfoRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RoleInfoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleInfoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RoleInfoRequestBuilderGetQueryParameters
}
// RoleInfoRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleInfoRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments the assignments property
func (m *RoleInfoRequestBuilder) Assignments()(*ib5cd5b3c19268d246febcc379cad7c6c55d83a8e91bd1b806e4aeced1e9bcbf0.AssignmentsRequestBuilder) {
    return ib5cd5b3c19268d246febcc379cad7c6c55d83a8e91bd1b806e4aeced1e9bcbf0.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedApproval.item.roleInfo.assignments.item collection
func (m *RoleInfoRequestBuilder) AssignmentsById(id string)(*i91831fb3bb284dcc082f0d50962d9b4761f0d7da6397b6aaa15e6aadeeef1c31.PrivilegedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRoleAssignment%2Did"] = id
    }
    return i91831fb3bb284dcc082f0d50962d9b4761f0d7da6397b6aaa15e6aadeeef1c31.NewPrivilegedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRoleInfoRequestBuilderInternal instantiates a new RoleInfoRequestBuilder and sets the default values.
func NewRoleInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleInfoRequestBuilder) {
    m := &RoleInfoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedApproval/{privilegedApproval%2Did}/roleInfo{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleInfoRequestBuilder instantiates a new RoleInfoRequestBuilder and sets the default values.
func NewRoleInfoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleInfoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleInfoRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleInfo for privilegedApproval
func (m *RoleInfoRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleInfo for privilegedApproval
func (m *RoleInfoRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *RoleInfoRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get roleInfo from privilegedApproval
func (m *RoleInfoRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get roleInfo from privilegedApproval
func (m *RoleInfoRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RoleInfoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property roleInfo in privilegedApproval
func (m *RoleInfoRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleInfo in privilegedApproval
func (m *RoleInfoRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, requestConfiguration *RoleInfoRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property roleInfo for privilegedApproval
func (m *RoleInfoRequestBuilder) Delete(ctx context.Context, requestConfiguration *RoleInfoRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get roleInfo from privilegedApproval
func (m *RoleInfoRequestBuilder) Get(ctx context.Context, requestConfiguration *RoleInfoRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedRoleFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable), nil
}
// Patch update the navigation property roleInfo in privilegedApproval
func (m *RoleInfoRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, requestConfiguration *RoleInfoRequestBuilderPatchRequestConfiguration)(error) {
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
// SelfActivate the selfActivate property
func (m *RoleInfoRequestBuilder) SelfActivate()(*i5cc87848e96d90dbe672f3e2fe6b8ec5e28e3cdc856ef730de729562e22b9524.SelfActivateRequestBuilder) {
    return i5cc87848e96d90dbe672f3e2fe6b8ec5e28e3cdc856ef730de729562e22b9524.NewSelfActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SelfDeactivate the selfDeactivate property
func (m *RoleInfoRequestBuilder) SelfDeactivate()(*i09a67af6e9bae1dd8dfd5d4374836f24ecdb5be20e1fae73a258e777e97458b4.SelfDeactivateRequestBuilder) {
    return i09a67af6e9bae1dd8dfd5d4374836f24ecdb5be20e1fae73a258e777e97458b4.NewSelfDeactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *RoleInfoRequestBuilder) Settings()(*i8b106f89a5024f1de5fce1ceb14aee752150a35a2f3ead41de7233973c9e2ec7.SettingsRequestBuilder) {
    return i8b106f89a5024f1de5fce1ceb14aee752150a35a2f3ead41de7233973c9e2ec7.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Summary the summary property
func (m *RoleInfoRequestBuilder) Summary()(*i920978860cb85582d4cc403166c2e16a9faae13fa4b1bd4c748c659db6073a7f.SummaryRequestBuilder) {
    return i920978860cb85582d4cc403166c2e16a9faae13fa4b1bd4c748c659db6073a7f.NewSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
