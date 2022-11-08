package item

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i6170658a0d5b58db8b79dacb601c4d22e30a4c4ae87b850f6f7f68c09f444c89 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/deviceenrollmentconfigurations/item/assign"
    iaa0d710f90b1a32f5cd84646584587a42f6739141a218233692f0fa578073ccc "github.com/microsoftgraph/msgraph-beta-sdk-go/me/deviceenrollmentconfigurations/item/setpriority"
    if2a3d3b427cb31b340e4c2412c17e0640a24dc599137ce6be0cc982589e4deb0 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/deviceenrollmentconfigurations/item/assignments"
    i934f161901d6bdad15cab7b3c6bac23e288fe9d7ada1f9c5697626905ea96d9b "github.com/microsoftgraph/msgraph-beta-sdk-go/me/deviceenrollmentconfigurations/item/assignments/item"
)

// DeviceEnrollmentConfigurationItemRequestBuilder provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
type DeviceEnrollmentConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters get enrollment configurations targeted to the user
type DeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Assign()(*i6170658a0d5b58db8b79dacb601c4d22e30a4c4ae87b850f6f7f68c09f444c89.AssignRequestBuilder) {
    return i6170658a0d5b58db8b79dacb601c4d22e30a4c4ae87b850f6f7f68c09f444c89.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceEnrollmentConfiguration entity.
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Assignments()(*if2a3d3b427cb31b340e4c2412c17e0640a24dc599137ce6be0cc982589e4deb0.AssignmentsRequestBuilder) {
    return if2a3d3b427cb31b340e4c2412c17e0640a24dc599137ce6be0cc982589e4deb0.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceEnrollmentConfiguration entity.
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) AssignmentsById(id string)(*i934f161901d6bdad15cab7b3c6bac23e288fe9d7ada1f9c5697626905ea96d9b.EnrollmentConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["enrollmentConfigurationAssignment%2Did"] = id
    }
    return i934f161901d6bdad15cab7b3c6bac23e288fe9d7ada1f9c5697626905ea96d9b.NewEnrollmentConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceEnrollmentConfigurationItemRequestBuilderInternal instantiates a new DeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceEnrollmentConfigurationItemRequestBuilder) {
    m := &DeviceEnrollmentConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceEnrollmentConfigurationItemRequestBuilder instantiates a new DeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewDeviceEnrollmentConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceEnrollmentConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property deviceEnrollmentConfigurations for me
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property deviceEnrollmentConfigurations in me
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, requestConfiguration *DeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property deviceEnrollmentConfigurations for me
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, error) {
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
// Patch update the navigation property deviceEnrollmentConfigurations in me
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, requestConfiguration *DeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, error) {
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
func (m *DeviceEnrollmentConfigurationItemRequestBuilder) SetPriority()(*iaa0d710f90b1a32f5cd84646584587a42f6739141a218233692f0fa578073ccc.SetPriorityRequestBuilder) {
    return iaa0d710f90b1a32f5cd84646584587a42f6739141a218233692f0fa578073ccc.NewSetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
