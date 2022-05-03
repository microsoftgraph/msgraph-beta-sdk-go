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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters collection of resource access settings associated with account.
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters
}
// DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
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
        urlTplParams["deviceManagementResourceAccessProfileAssignment%2Did"] = id
    }
    return i4c6649ed921058137d23369f25fb530c052e741c6026b182c9c66a0903b527cd.NewDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal instantiates a new DeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    m := &DeviceManagementResourceAccessProfileBaseItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/resourceAccessProfiles/{deviceManagementResourceAccessProfileBase%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property resourceAccessProfiles for deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property resourceAccessProfiles for deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformationWithRequestConfiguration collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property resourceAccessProfiles in deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property resourceAccessProfiles in deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DeleteWithResponseHandler delete navigation property resourceAccessProfiles for deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration)(error) {
    return m.DeleteWithResponseHandler(requestConfiguration, nil);
}
// DeleteWithResponseHandler delete navigation property resourceAccessProfiles for deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) DeleteWithResponseHandler(requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
// GetWithResponseHandler collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) GetWithResponseHandler(requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, error) {
    return m.GetWithResponseHandler(requestConfiguration, nil);
}
// GetWithResponseHandler collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) GetWithResponseHandler(requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable), nil
}
// PatchWithResponseHandler update the navigation property resourceAccessProfiles in deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration)(error) {
    return m.PatchWithResponseHandler(body, requestConfiguration, nil);
}
// PatchWithResponseHandler update the navigation property resourceAccessProfiles in deviceManagement
func (m *DeviceManagementResourceAccessProfileBaseItemRequestBuilder) PatchWithResponseHandler(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *DeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
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
