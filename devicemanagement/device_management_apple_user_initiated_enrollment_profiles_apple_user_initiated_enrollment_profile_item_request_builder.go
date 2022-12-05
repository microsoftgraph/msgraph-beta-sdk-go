package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder provides operations to manage the appleUserInitiatedEnrollmentProfiles property of the microsoft.graph.deviceManagement entity.
type DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters apple user initiated enrollment profiles
type DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetQueryParameters
}
// DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.appleUserInitiatedEnrollmentProfile entity.
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) Assignments()(*DeviceManagementAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilder) {
    return NewDeviceManagementAppleUserInitiatedEnrollmentProfilesItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.appleUserInitiatedEnrollmentProfile entity.
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["appleEnrollmentProfileAssignment%2Did"] = id
    }
    return NewDeviceManagementAppleUserInitiatedEnrollmentProfilesItemAssignmentsAppleEnrollmentProfileAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal instantiates a new AppleUserInitiatedEnrollmentProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    m := &DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/appleUserInitiatedEnrollmentProfiles/{appleUserInitiatedEnrollmentProfile%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder instantiates a new AppleUserInitiatedEnrollmentProfileItemRequestBuilder and sets the default values.
func NewDeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation apple user initiated enrollment profiles
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property appleUserInitiatedEnrollmentProfiles for deviceManagement
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get apple user initiated enrollment profiles
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable), nil
}
// Patch update the navigation property appleUserInitiatedEnrollmentProfiles in deviceManagement
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, requestConfiguration *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateAppleUserInitiatedEnrollmentProfileFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.AppleUserInitiatedEnrollmentProfileable), nil
}
// SetPriority provides operations to call the setPriority method.
func (m *DeviceManagementAppleUserInitiatedEnrollmentProfilesAppleUserInitiatedEnrollmentProfileItemRequestBuilder) SetPriority()(*DeviceManagementAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilder) {
    return NewDeviceManagementAppleUserInitiatedEnrollmentProfilesItemSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
