package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.user entity.
type ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters get enrollment configurations targeted to the user
type ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetQueryParameters
}
// ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.deviceEnrollmentConfiguration entity.
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Assignments()(*ItemDeviceEnrollmentConfigurationsItemAssignmentsRequestBuilder) {
    return NewItemDeviceEnrollmentConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.deviceEnrollmentConfiguration entity.
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) AssignmentsById(id string)(*ItemDeviceEnrollmentConfigurationsItemAssignmentsEnrollmentConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemDeviceEnrollmentConfigurationsItemAssignmentsEnrollmentConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// NewItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal instantiates a new DeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceEnrollmentConfigurationId *string)(*ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    m := &ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/deviceEnrollmentConfigurations/{deviceEnrollmentConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if deviceEnrollmentConfigurationId != nil {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = *deviceEnrollmentConfigurationId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder instantiates a new DeviceEnrollmentConfigurationItemRequestBuilder and sets the default values.
func NewItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property deviceEnrollmentConfigurations for users
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get enrollment configurations targeted to the user
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable), nil
}
// MicrosoftGraphAssign provides operations to call the assign method.
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) MicrosoftGraphAssign()(*ItemDeviceEnrollmentConfigurationsItemMicrosoftGraphAssignAssignRequestBuilder) {
    return NewItemDeviceEnrollmentConfigurationsItemMicrosoftGraphAssignAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphSetPriority provides operations to call the setPriority method.
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) MicrosoftGraphSetPriority()(*ItemDeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityRequestBuilder) {
    return NewItemDeviceEnrollmentConfigurationsItemMicrosoftGraphSetPrioritySetPriorityRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update the navigation property deviceEnrollmentConfigurations in users
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, requestConfiguration *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable), nil
}
// ToDeleteRequestInformation delete navigation property deviceEnrollmentConfigurations for users
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get enrollment configurations targeted to the user
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property deviceEnrollmentConfigurations in users
func (m *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceEnrollmentConfigurationable, requestConfiguration *ItemDeviceEnrollmentConfigurationsDeviceEnrollmentConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
