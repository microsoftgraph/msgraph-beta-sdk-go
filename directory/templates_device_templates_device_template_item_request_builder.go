package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder provides operations to manage the deviceTemplates property of the microsoft.graph.template entity.
type TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderGetQueryParameters get the properties and relationships of a deviceTemplate object.
type TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderGetQueryParameters
}
// TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderInternal instantiates a new TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder and sets the default values.
func NewTemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) {
    m := &TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/templates/deviceTemplates/{deviceTemplate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder instantiates a new TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder and sets the default values.
func NewTemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeviceFromTemplate provides operations to call the createDeviceFromTemplate method.
// returns a *TemplatesDeviceTemplatesItemCreateDeviceFromTemplateRequestBuilder when successful
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) CreateDeviceFromTemplate()(*TemplatesDeviceTemplatesItemCreateDeviceFromTemplateRequestBuilder) {
    return NewTemplatesDeviceTemplatesItemCreateDeviceFromTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a registered deviceTemplate. You must first delete all devices linked to the template before deleting the template itself. Only registered owners of the template can perform this operation.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/devicetemplate-delete?view=graph-rest-beta
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *TemplatesDeviceTemplatesItemDeviceInstancesRequestBuilder when successful
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) DeviceInstances()(*TemplatesDeviceTemplatesItemDeviceInstancesRequestBuilder) {
    return NewTemplatesDeviceTemplatesItemDeviceInstancesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DeviceInstancesWithDeviceId provides operations to manage the deviceInstances property of the microsoft.graph.deviceTemplate entity.
// returns a *TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder when successful
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) DeviceInstancesWithDeviceId(deviceId *string)(*TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder) {
    return NewTemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, deviceId)
}
// Get get the properties and relationships of a deviceTemplate object.
// returns a DeviceTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/devicetemplate-get?view=graph-rest-beta
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable, error) {
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
// returns a *TemplatesDeviceTemplatesItemOwnersRequestBuilder when successful
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) Owners()(*TemplatesDeviceTemplatesItemOwnersRequestBuilder) {
    return NewTemplatesDeviceTemplatesItemOwnersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property deviceTemplates in directory
// returns a DeviceTemplateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable, requestConfiguration *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable, error) {
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
// ToDeleteRequestInformation delete a registered deviceTemplate. You must first delete all devices linked to the template before deleting the template itself. Only registered owners of the template can perform this operation.
// returns a *RequestInformation when successful
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of a deviceTemplate object.
// returns a *RequestInformation when successful
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property deviceTemplates in directory
// returns a *RequestInformation when successful
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DeviceTemplateable, requestConfiguration *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder when successful
func (m *TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) WithUrl(rawUrl string)(*TemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder) {
    return NewTemplatesDeviceTemplatesDeviceTemplateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
