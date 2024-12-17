package templates

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceTemplatesDeviceTemplateItemRequestBuilder provides operations to manage the deviceTemplates property of the microsoft.graph.template entity.
type DeviceTemplatesDeviceTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceTemplatesDeviceTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceTemplatesDeviceTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceTemplatesDeviceTemplateItemRequestBuilderGetQueryParameters get deviceTemplates from templates
type DeviceTemplatesDeviceTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceTemplatesDeviceTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceTemplatesDeviceTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceTemplatesDeviceTemplateItemRequestBuilderGetQueryParameters
}
// DeviceTemplatesDeviceTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceTemplatesDeviceTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceTemplatesDeviceTemplateItemRequestBuilderInternal instantiates a new DeviceTemplatesDeviceTemplateItemRequestBuilder and sets the default values.
func NewDeviceTemplatesDeviceTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceTemplatesDeviceTemplateItemRequestBuilder) {
    m := &DeviceTemplatesDeviceTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/templates/deviceTemplates/{deviceTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeviceTemplatesDeviceTemplateItemRequestBuilder instantiates a new DeviceTemplatesDeviceTemplateItemRequestBuilder and sets the default values.
func NewDeviceTemplatesDeviceTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceTemplatesDeviceTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceTemplatesDeviceTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceFromTemplate provides operations to call the createDeviceFromTemplate method.
// returns a *DeviceTemplatesItemCreateDeviceFromTemplateRequestBuilder when successful
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) CreateDeviceFromTemplate()(*DeviceTemplatesItemCreateDeviceFromTemplateRequestBuilder) {
    return NewDeviceTemplatesItemCreateDeviceFromTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property deviceTemplates for templates
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeviceTemplatesDeviceTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DeviceInstances provides operations to manage the deviceInstances property of the microsoft.graph.deviceTemplate entity.
// returns a *DeviceTemplatesItemDeviceInstancesRequestBuilder when successful
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) DeviceInstances()(*DeviceTemplatesItemDeviceInstancesRequestBuilder) {
    return NewDeviceTemplatesItemDeviceInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceInstancesWithDeviceId provides operations to manage the deviceInstances property of the microsoft.graph.deviceTemplate entity.
// returns a *DeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder when successful
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) DeviceInstancesWithDeviceId(deviceId *string)(*DeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder) {
    return NewDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, deviceId)
}
// Get get deviceTemplates from templates
// returns a DeviceTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceTemplatesDeviceTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable), nil
}
// Owners provides operations to manage the owners property of the microsoft.graph.deviceTemplate entity.
// returns a *DeviceTemplatesItemOwnersRequestBuilder when successful
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) Owners()(*DeviceTemplatesItemOwnersRequestBuilder) {
    return NewDeviceTemplatesItemOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceTemplates in templates
// returns a DeviceTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable, requestConfiguration *DeviceTemplatesDeviceTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceTemplateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable), nil
}
// ToDeleteRequestInformation delete navigation property deviceTemplates for templates
// returns a *RequestInformation when successful
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeviceTemplatesDeviceTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get deviceTemplates from templates
// returns a *RequestInformation when successful
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceTemplatesDeviceTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceTemplates in templates
// returns a *RequestInformation when successful
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable, requestConfiguration *DeviceTemplatesDeviceTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceTemplatesDeviceTemplateItemRequestBuilder when successful
func (m *DeviceTemplatesDeviceTemplateItemRequestBuilder) WithUrl(rawUrl string)(*DeviceTemplatesDeviceTemplateItemRequestBuilder) {
    return NewDeviceTemplatesDeviceTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
