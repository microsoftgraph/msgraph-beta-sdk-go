package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder provides operations to manage the presentationValues property of the microsoft.graph.groupPolicyDefinitionValue entity.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetQueryParameters the associated group policy presentation values with the definition value.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetQueryParameters
}
// DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderInternal instantiates a new GroupPolicyPresentationValueItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) {
    m := &DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/groupPolicyConfigurations/{groupPolicyConfiguration%2Did}/definitionValues/{groupPolicyDefinitionValue%2Did}/presentationValues/{groupPolicyPresentationValue%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder instantiates a new GroupPolicyPresentationValueItemRequestBuilder and sets the default values.
func NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property presentationValues for deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the associated group policy presentation values with the definition value.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property presentationValues in deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// DefinitionValue provides operations to manage the definitionValue property of the microsoft.graph.groupPolicyPresentationValue entity.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) DefinitionValue()(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemDefinitionValueRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Delete delete navigation property presentationValues for deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the associated group policy presentation values with the definition value.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable), nil
}
// Patch update the navigation property presentationValues in deviceManagement
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, requestConfiguration *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateGroupPolicyPresentationValueFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.GroupPolicyPresentationValueable), nil
}
// Presentation provides operations to manage the presentation property of the microsoft.graph.groupPolicyPresentationValue entity.
func (m *DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesGroupPolicyPresentationValueItemRequestBuilder) Presentation()(*DeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemPresentationRequestBuilder) {
    return NewDeviceManagementGroupPolicyConfigurationsItemDefinitionValuesItemPresentationValuesItemPresentationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
