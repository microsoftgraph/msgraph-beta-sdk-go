package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder provides operations to manage the resourceAccessProfiles property of the microsoft.graph.deviceManagement entity.
type DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters collection of resource access settings associated with account.
type DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetQueryParameters
}
// DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Assign()(*DeviceManagementResourceAccessProfilesItemAssignRequestBuilder) {
    return NewDeviceManagementResourceAccessProfilesItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceManagementResourceAccessProfileBase entity.
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Assignments()(*DeviceManagementResourceAccessProfilesItemAssignmentsRequestBuilder) {
    return NewDeviceManagementResourceAccessProfilesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceManagementResourceAccessProfileBase entity.
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementResourceAccessProfilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementResourceAccessProfileAssignment%2Did"] = id
    }
    return NewDeviceManagementResourceAccessProfilesItemAssignmentsDeviceManagementResourceAccessProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal instantiates a new DeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewDeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    m := &DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder{
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
// NewDeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder instantiates a new DeviceManagementResourceAccessProfileBaseItemRequestBuilder and sets the default values.
func NewDeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property resourceAccessProfiles for deviceManagement
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property resourceAccessProfiles in deviceManagement
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property resourceAccessProfiles for deviceManagement
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of resource access settings associated with account.
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable), nil
}
// Patch update the navigation property resourceAccessProfiles in deviceManagement
func (m *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, requestConfiguration *DeviceManagementResourceAccessProfilesDeviceManagementResourceAccessProfileBaseItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementResourceAccessProfileBaseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementResourceAccessProfileBaseable), nil
}
