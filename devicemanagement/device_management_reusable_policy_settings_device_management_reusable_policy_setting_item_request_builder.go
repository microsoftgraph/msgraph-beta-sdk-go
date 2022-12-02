package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
type DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters list of all reusable settings that can be referred in a policy
type DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters
}
// DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Clone provides operations to call the clone method.
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) Clone()(*DeviceManagementReusablePolicySettingsItemCloneRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsItemCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderInternal instantiates a new DeviceManagementReusablePolicySettingItemRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) {
    m := &DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder instantiates a new DeviceManagementReusablePolicySettingItemRequestBuilder and sets the default values.
func NewDeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property reusablePolicySettings for deviceManagement
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property reusablePolicySettings in deviceManagement
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable, requestConfiguration *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property reusablePolicySettings for deviceManagement
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of all reusable settings that can be referred in a policy
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable), nil
}
// Patch update the navigation property reusablePolicySettings in deviceManagement
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable, requestConfiguration *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable), nil
}
// ReferencingConfigurationPolicies provides operations to manage the referencingConfigurationPolicies property of the microsoft.graph.deviceManagementReusablePolicySetting entity.
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) ReferencingConfigurationPolicies()(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesRequestBuilder) {
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ReferencingConfigurationPoliciesById provides operations to manage the referencingConfigurationPolicies property of the microsoft.graph.deviceManagementReusablePolicySetting entity.
func (m *DeviceManagementReusablePolicySettingsDeviceManagementReusablePolicySettingItemRequestBuilder) ReferencingConfigurationPoliciesById(id string)(*DeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementConfigurationPolicy%2Did"] = id
    }
    return NewDeviceManagementReusablePolicySettingsItemReferencingConfigurationPoliciesDeviceManagementConfigurationPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
