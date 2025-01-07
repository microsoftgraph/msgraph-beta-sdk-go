package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder provides operations to manage the deviceInstances property of the microsoft.graph.deviceTemplate entity.
type TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderGetQueryParameters collection of device objects created based on this template.
type TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderGetQueryParameters
}
// NewTemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderInternal instantiates a new TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder and sets the default values.
func NewTemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, deviceId *string)(*TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder) {
    m := &TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/templates/deviceTemplates/{deviceTemplate%2Did}/deviceInstances(deviceId='{deviceId}'){?%24expand,%24select}", pathParameters),
    }
    if deviceId != nil {
        m.BaseRequestBuilder.PathParameters["deviceId"] = *deviceId
    }
    return m
}
// NewTemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder instantiates a new TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder and sets the default values.
func NewTemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get collection of device objects created based on this template.
// returns a Deviceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder) Get(ctx context.Context, requestConfiguration *TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDeviceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Deviceable), nil
}
// ToGetRequestInformation collection of device objects created based on this template.
// returns a *RequestInformation when successful
func (m *TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder when successful
func (m *TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder) WithUrl(rawUrl string)(*TemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder) {
    return NewTemplatesDeviceTemplatesItemDeviceInstancesWithDeviceIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
