package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder provides operations to manage the templateInsights property of the microsoft.graph.deviceManagement entity.
type TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderGetQueryParameters list of setting insights in a template
type TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderGetQueryParameters
}
// TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderInternal instantiates a new TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder and sets the default values.
func NewTemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) {
    m := &TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/templateInsights/{deviceManagementTemplateInsightsDefinition%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder instantiates a new TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder and sets the default values.
func NewTemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property templateInsights for deviceManagement
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get list of setting insights in a template
// returns a DeviceManagementTemplateInsightsDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateInsightsDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplateInsightsDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateInsightsDefinitionable), nil
}
// Patch update the navigation property templateInsights in deviceManagement
// returns a DeviceManagementTemplateInsightsDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateInsightsDefinitionable, requestConfiguration *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateInsightsDefinitionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceManagementTemplateInsightsDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateInsightsDefinitionable), nil
}
// ToDeleteRequestInformation delete navigation property templateInsights for deviceManagement
// returns a *RequestInformation when successful
func (m *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/deviceManagement/templateInsights/{deviceManagementTemplateInsightsDefinition%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation list of setting insights in a template
// returns a *RequestInformation when successful
func (m *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property templateInsights in deviceManagement
// returns a *RequestInformation when successful
func (m *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceManagementTemplateInsightsDefinitionable, requestConfiguration *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/deviceManagement/templateInsights/{deviceManagementTemplateInsightsDefinition%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder when successful
func (m *TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) WithUrl(rawUrl string)(*TemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder) {
    return NewTemplateInsightsDeviceManagementTemplateInsightsDefinitionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
