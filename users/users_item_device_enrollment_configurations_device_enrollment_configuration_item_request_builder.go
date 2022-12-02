package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
type UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters get enrollment configurations targeted to the user
type UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters
}
// UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Assign()(*UsersItemDeviceEnrollmentConfigurationsItemAssignRequestBuilder) {
    return NewUsersItemDeviceEnrollmentConfigurationsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceEnrollmentConfiguration entity.
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Assignments()(*UsersItemDeviceEnrollmentConfigurationsItemAssignmentsRequestBuilder) {
    return NewUsersItemDeviceEnrollmentConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceEnrollmentConfiguration entity.
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) AssignmentsById(id string)(*UsersItemDeviceEnrollmentConfigurationsItemAssignmentsEnrollmentConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentConfigurationAssignment%2Did"] = id
    }
    return NewUsersItemDeviceEnrollmentConfigurationsItemAssignmentsEnrollmentConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewUsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal instantiates a new DeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewUsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    m := &UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder instantiates a new DeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewUsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceEnrollmentConfigurations for users
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get enrollment configurations targeted to the user
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceEnrollmentConfigurations in users
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, requestConfiguration *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceEnrollmentConfigurations for users
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get enrollment configurations targeted to the user
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable), nil
}
// Patch update the navigation property deviceEnrollmentConfigurations in users
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, requestConfiguration *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable), nil
}
// SetPriority provides operations to call the setPriority method.
func (m *UsersItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) SetPriority()(*UsersItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilder) {
    return NewUsersItemDeviceEnrollmentConfigurationsItemSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
