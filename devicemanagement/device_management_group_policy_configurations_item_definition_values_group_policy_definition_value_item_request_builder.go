package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder provides operations to manage the definitionValues property of the microsoft.graph.groupPolicyConfiguration entity.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters the list of enabled or disabled group policy definition values for the configuration.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderInternal instantiates a new GroupPolicyDefinitionValueItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}/definitionValues/{groupPolicyDefinitionValue%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder instantiates a new GroupPolicyDefinitionValueItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property definitionValues for deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the list of enabled or disabled group policy definition values for the configuration.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property definitionValues in deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Definition provides operations to manage the definition property of the microsoft.graph.groupPolicyDefinitionValue entity.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) Definition()(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemDefinitionRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemDefinitionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property definitionValues for deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the list of enabled or disabled group policy definition values for the configuration.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable), nil
}
// Patch update the navigation property definitionValues in deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyDefinitionValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyDefinitionValueable), nil
}
// PresentationValues provides operations to manage the presentationValues property of the microsoft.graph.groupPolicyDefinitionValue entity.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) PresentationValues()(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PresentationValuesById provides operations to manage the presentationValues property of the microsoft.graph.groupPolicyDefinitionValue entity.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesGroupPolicyDefinitionValueItemRequestBuilder) PresentationValuesById(id string)(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["groupPolicyPresentationValue%2Did"] = id
    }
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
