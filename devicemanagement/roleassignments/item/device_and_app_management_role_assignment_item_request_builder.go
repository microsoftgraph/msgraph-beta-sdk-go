package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i5cad8ef86dc95ddf4ad899618000e669a92e76bdb8d952086dd1b1ac1b69d619 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/roleassignments/item/rolescopetags"
    i0144dfcc833d31687db4d6e06640cdfd86f365134712a6504ff2ceeccb6a3aa6 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/roleassignments/item/rolescopetags/item"
)

// DeviceAndAppManagementRoleAssignmentItemRequestBuilder provides operations to manage the roleAssignments property of the microsoft.graph.deviceManagement entity.
type DeviceAndAppManagementRoleAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceAndAppManagementRoleAssignmentItemRequestBuilderDeleteOptions options for Delete
type DeviceAndAppManagementRoleAssignmentItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// DeviceAndAppManagementRoleAssignmentItemRequestBuilderGetOptions options for Get
type DeviceAndAppManagementRoleAssignmentItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceAndAppManagementRoleAssignmentItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// DeviceAndAppManagementRoleAssignmentItemRequestBuilderGetQueryParameters the Role Assignments.
type DeviceAndAppManagementRoleAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceAndAppManagementRoleAssignmentItemRequestBuilderPatchOptions options for Patch
type DeviceAndAppManagementRoleAssignmentItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal instantiates a new DeviceAndAppManagementRoleAssignmentItemRequestBuilder and sets the default values.
func NewDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAndAppManagementRoleAssignmentItemRequestBuilder) {
    m := &DeviceAndAppManagementRoleAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/roleAssignments/{deviceAndAppManagementRoleAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAndAppManagementRoleAssignmentItemRequestBuilder instantiates a new DeviceAndAppManagementRoleAssignmentItemRequestBuilder and sets the default values.
func NewDeviceAndAppManagementRoleAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAndAppManagementRoleAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAndAppManagementRoleAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property roleAssignments for deviceManagement
func (m *DeviceAndAppManagementRoleAssignmentItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceAndAppManagementRoleAssignmentItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the Role Assignments.
func (m *DeviceAndAppManagementRoleAssignmentItemRequestBuilder) CreateGetRequestInformation(options *DeviceAndAppManagementRoleAssignmentItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property roleAssignments in deviceManagement
func (m *DeviceAndAppManagementRoleAssignmentItemRequestBuilder) CreatePatchRequestInformation(options *DeviceAndAppManagementRoleAssignmentItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property roleAssignments for deviceManagement
func (m *DeviceAndAppManagementRoleAssignmentItemRequestBuilder) Delete(options *DeviceAndAppManagementRoleAssignmentItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the Role Assignments.
func (m *DeviceAndAppManagementRoleAssignmentItemRequestBuilder) Get(options *DeviceAndAppManagementRoleAssignmentItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceAndAppManagementRoleAssignmentFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceAndAppManagementRoleAssignmentable), nil
}
// Patch update the navigation property roleAssignments in deviceManagement
func (m *DeviceAndAppManagementRoleAssignmentItemRequestBuilder) Patch(options *DeviceAndAppManagementRoleAssignmentItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RoleScopeTags the roleScopeTags property
func (m *DeviceAndAppManagementRoleAssignmentItemRequestBuilder) RoleScopeTags()(*i5cad8ef86dc95ddf4ad899618000e669a92e76bdb8d952086dd1b1ac1b69d619.RoleScopeTagsRequestBuilder) {
    return i5cad8ef86dc95ddf4ad899618000e669a92e76bdb8d952086dd1b1ac1b69d619.NewRoleScopeTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleScopeTagsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.roleAssignments.item.roleScopeTags.item collection
func (m *DeviceAndAppManagementRoleAssignmentItemRequestBuilder) RoleScopeTagsById(id string)(*i0144dfcc833d31687db4d6e06640cdfd86f365134712a6504ff2ceeccb6a3aa6.RoleScopeTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["roleScopeTag%2Did"] = id
    }
    return i0144dfcc833d31687db4d6e06640cdfd86f365134712a6504ff2ceeccb6a3aa6.NewRoleScopeTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
