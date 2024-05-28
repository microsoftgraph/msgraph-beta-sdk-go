package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder provides operations to manage the reusablePolicySettings property of the microsoft.graph.deviceManagement entity.
type ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters list of all reusable settings that can be referred in a policy
type ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetQueryParameters
}
// ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Clone provides operations to call the clone method.
// returns a *ReusablepolicysettingsItemCloneRequestBuilder when successful
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) Clone()(*ReusablepolicysettingsItemCloneRequestBuilder) {
    return NewReusablepolicysettingsItemCloneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderInternal instantiates a new ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder and sets the default values.
func NewReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) {
    m := &ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/reusablePolicySettings/{deviceManagementReusablePolicySetting%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder instantiates a new ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder and sets the default values.
func NewReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property reusablePolicySettings for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of all reusable settings that can be referred in a policy
// returns a DeviceManagementReusablePolicySettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable), nil
}
// Patch update the navigation property reusablePolicySettings in deviceManagement
// returns a DeviceManagementReusablePolicySettingable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable, requestConfiguration *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementReusablePolicySettingFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable), nil
}
// ReferencingConfigurationPolicies provides operations to manage the referencingConfigurationPolicies property of the microsoft.graph.deviceManagementReusablePolicySetting entity.
// returns a *ReusablepolicysettingsItemReferencingconfigurationpoliciesReferencingConfigurationPoliciesRequestBuilder when successful
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) ReferencingConfigurationPolicies()(*ReusablepolicysettingsItemReferencingconfigurationpoliciesReferencingConfigurationPoliciesRequestBuilder) {
    return NewReusablepolicysettingsItemReferencingconfigurationpoliciesReferencingConfigurationPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property reusablePolicySettings for deviceManagement
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of all reusable settings that can be referred in a policy
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property reusablePolicySettings in deviceManagement
// returns a *RequestInformation when successful
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementReusablePolicySettingable, requestConfiguration *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder when successful
func (m *ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) WithUrl(rawUrl string)(*ReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder) {
    return NewReusablepolicysettingsDeviceManagementReusablePolicySettingItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
