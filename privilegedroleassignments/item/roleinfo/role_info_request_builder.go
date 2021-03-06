package roleinfo

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i14d3b1099c12de821235b9046c0fe443f0d08af3ef9196da3852f7c9096a9456 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/roleinfo/summary"
    i1be213330230666a7f78e848b631064650ba359f512cbfcef9054acfcd95000d "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/roleinfo/selfactivate"
    i5d9b0586636a3cb17b9fc6b4b07b708ddf8a4c94f548693caab3ae95df11297d "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/roleinfo/selfdeactivate"
    i7be4f06094555e6d1ed283d4812cebe44786b3a3896c5d67fcf9b8066bfaeb8e "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/roleinfo/settings"
    iddbd5b0e7375a8b731ea7c30727a2839cc1fe6936313d8f9fb36762191a4a258 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/roleinfo/assignments"
    id336737ac0abd6bb7b542bc7f6e733eee4b4936acef3152c3b4beee97311be81 "github.com/microsoftgraph/msgraph-beta-sdk-go/privilegedroleassignments/item/roleinfo/assignments/item"
)

// RoleInfoRequestBuilder provides operations to manage the roleInfo property of the microsoft.graph.privilegedRoleAssignment entity.
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
// RoleInfoRequestBuilderGetQueryParameters read-only. Nullable. The associated role information.
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
func (m *RoleInfoRequestBuilder) Assignments()(*iddbd5b0e7375a8b731ea7c30727a2839cc1fe6936313d8f9fb36762191a4a258.AssignmentsRequestBuilder) {
    return iddbd5b0e7375a8b731ea7c30727a2839cc1fe6936313d8f9fb36762191a4a258.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.privilegedRoleAssignments.item.roleInfo.assignments.item collection
func (m *RoleInfoRequestBuilder) AssignmentsById(id string)(*id336737ac0abd6bb7b542bc7f6e733eee4b4936acef3152c3b4beee97311be81.PrivilegedRoleAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["privilegedRoleAssignment%2Did1"] = id
    }
    return id336737ac0abd6bb7b542bc7f6e733eee4b4936acef3152c3b4beee97311be81.NewPrivilegedRoleAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRoleInfoRequestBuilderInternal instantiates a new RoleInfoRequestBuilder and sets the default values.
func NewRoleInfoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleInfoRequestBuilder) {
    m := &RoleInfoRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/privilegedRoleAssignments/{privilegedRoleAssignment%2Did}/roleInfo{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property roleInfo for privilegedRoleAssignments
func (m *RoleInfoRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property roleInfo for privilegedRoleAssignments
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
// CreateGetRequestInformation read-only. Nullable. The associated role information.
func (m *RoleInfoRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration read-only. Nullable. The associated role information.
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
// CreatePatchRequestInformation update the navigation property roleInfo in privilegedRoleAssignments
func (m *RoleInfoRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property roleInfo in privilegedRoleAssignments
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
// Delete delete navigation property roleInfo for privilegedRoleAssignments
func (m *RoleInfoRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property roleInfo for privilegedRoleAssignments
func (m *RoleInfoRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *RoleInfoRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable. The associated role information.
func (m *RoleInfoRequestBuilder) Get()(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler read-only. Nullable. The associated role information.
func (m *RoleInfoRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *RoleInfoRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreatePrivilegedRoleFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable), nil
}
// Patch update the navigation property roleInfo in privilegedRoleAssignments
func (m *RoleInfoRequestBuilder) Patch(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property roleInfo in privilegedRoleAssignments
func (m *RoleInfoRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.PrivilegedRoleable, requestConfiguration *RoleInfoRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SelfActivate the selfActivate property
func (m *RoleInfoRequestBuilder) SelfActivate()(*i1be213330230666a7f78e848b631064650ba359f512cbfcef9054acfcd95000d.SelfActivateRequestBuilder) {
    return i1be213330230666a7f78e848b631064650ba359f512cbfcef9054acfcd95000d.NewSelfActivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SelfDeactivate the selfDeactivate property
func (m *RoleInfoRequestBuilder) SelfDeactivate()(*i5d9b0586636a3cb17b9fc6b4b07b708ddf8a4c94f548693caab3ae95df11297d.SelfDeactivateRequestBuilder) {
    return i5d9b0586636a3cb17b9fc6b4b07b708ddf8a4c94f548693caab3ae95df11297d.NewSelfDeactivateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Settings the settings property
func (m *RoleInfoRequestBuilder) Settings()(*i7be4f06094555e6d1ed283d4812cebe44786b3a3896c5d67fcf9b8066bfaeb8e.SettingsRequestBuilder) {
    return i7be4f06094555e6d1ed283d4812cebe44786b3a3896c5d67fcf9b8066bfaeb8e.NewSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Summary the summary property
func (m *RoleInfoRequestBuilder) Summary()(*i14d3b1099c12de821235b9046c0fe443f0d08af3ef9196da3852f7c9096a9456.SummaryRequestBuilder) {
    return i14d3b1099c12de821235b9046c0fe443f0d08af3ef9196da3852f7c9096a9456.NewSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
