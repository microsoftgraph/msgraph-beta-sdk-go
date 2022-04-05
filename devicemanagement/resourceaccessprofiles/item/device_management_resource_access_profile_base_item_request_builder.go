package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1aa4fb44cf4940a0794e33de635a58953960c975237235af78452d02bffe2854 "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/resourceaccessprofiles/item/assignments"
    iaac20666e49cb2032811a5fbdaae8facc6a32715040bd58f314cf8d5f006f5bd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/resourceaccessprofiles/item/assign"
    i4c6649ed921058137d23369f25fb530c052e741c6026b182c9c66a0903b527cd "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/resourceaccessprofiles/item/assignments/item"
)

// DeviceManagementResourceAccessProfileBaseItemRequestBuilder provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
type DeviceManagementResourceAccessProfileBaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteOptions options for Delete
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetOptions options for Get
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters collection of resource access settings associated with account.
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchOptions options for Patch
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Assign the assign property
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Assign()(*iaac20666e49cb2032811a5fbdaae8facc6a32715040bd58f314cf8d5f006f5bd.AssignRequestBuilder) {
    return iaac20666e49cb2032811a5fbdaae8facc6a32715040bd58f314cf8d5f006f5bd.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Assignments()(*i1aa4fb44cf4940a0794e33de635a58953960c975237235af78452d02bffe2854.AssignmentsRequestBuilder) {
    return i1aa4fb44cf4940a0794e33de635a58953960c975237235af78452d02bffe2854.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.deviceManagement.resourceAccessProfiles.item.assignments.item collection
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) AssignmentsById(id string)(*i4c6649ed921058137d23369f25fb530c052e741c6026b182c9c66a0903b527cd.DeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementResourceAccessProfileAssignment_id"] = id
    }
    return i4c6649ed921058137d23369f25fb530c052e741c6026b182c9c66a0903b527cd.NewDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal instantiates a new DeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    m := &DeviceManagementResourceAccessProfileBaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/resourceAccessProfiles/{deviceManagementResourceAccessProfileBase_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementResourceAccessProfileBaseItemRequestBuilder instantiates a new DeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewDeviceManagementResourceAccessProfileBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resourceAccessProfiles for deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateGetRequestInformation(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property resourceAccessProfiles in deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property resourceAccessProfiles for deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Delete(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteOptions)(error) {
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
// Get collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Get(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable), nil
}
// Patch update the navigation property resourceAccessProfiles in deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) Patch(options *DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchOptions)(error) {
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
