package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder provides operations to count the resources in the collection.
type DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderGetQueryParameters get the number of the resource
type DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderGetQueryParameters struct {
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
}
// DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderGetQueryParameters
}
// NewDevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderInternal instantiates a new DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder) {
    m := &DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceCustomAttributeShellScripts/{deviceCustomAttributeShellScript%2Did}/userRunStates/{deviceManagementScriptUserState%2Did}/deviceRunStates/$count{?%24filter,%24search}", pathParameters),
    }
    return m
}
// NewDevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder instantiates a new DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder and sets the default values.
func NewDevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the number of the resource
// returns a *int32 when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder) Get(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
// ToGetRequestInformation get the number of the resource
// returns a *RequestInformation when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "text/plain;q=0.9")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder when successful
func (m *DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder) WithUrl(rawUrl string)(*DevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder) {
    return NewDevicecustomattributeshellscriptsItemUserrunstatesItemDevicerunstatesCountRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
