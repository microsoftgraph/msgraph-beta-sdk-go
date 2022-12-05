package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder provides operations to manage the groupPolicyConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderGetQueryParameters the group policy configurations created by this account.
type DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Assign provides operations to call the assign method.
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) Assign()(*DeviceManagementGroupPolicyConfigurationsItemAssignRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsItemAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.groupPolicyConfiguration entity.
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) Assignments()(*DeviceManagementGroupPolicyConfigurationsItemAssignmentsRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsItemAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.groupPolicyConfiguration entity.
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) AssignmentsById(id string)(*DeviceManagementGroupPolicyConfigurationsItemAssignmentsGroupPolicyConfigurationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyConfigurationAssignment%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyConfigurationsItemAssignmentsGroupPolicyConfigurationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderInternal instantiates a new GroupPolicyConfigurationItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder instantiates a new GroupPolicyConfigurationItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property groupPolicyConfigurations for deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the group policy configurations created by this account.
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property groupPolicyConfigurations in deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable, requestConfiguration *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DefinitionValues provides operations to manage the definitionValues property of the microsoft.graph.groupPolicyConfiguration entity.
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) DefinitionValues()(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefinitionValuesById provides operations to manage the definitionValues property of the microsoft.graph.groupPolicyConfiguration entity.
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) DefinitionValuesById(id string)(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyDefinitionValue%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Delete delete navigation property groupPolicyConfigurations for deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the group policy configurations created by this account.
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable), nil
}
// Patch update the navigation property groupPolicyConfigurations in deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable, requestConfiguration *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyConfigurationable), nil
}
// UpdateDefinitionValues provides operations to call the updateDefinitionValues method.
func (m *DeviceManagementGroupPolicyConfigurationsGroupPolicyConfigurationItemRequestBuilder) UpdateDefinitionValues()(*DeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsItemUpdateDefinitionValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
