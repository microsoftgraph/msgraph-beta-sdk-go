package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder provides operations to call the getGlobalScriptHighestAvailableVersion method.
type DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderInternal instantiates a new GetGlobalScriptHighestAvailableVersionRequestBuilder and sets the default values.
func NewDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    m := &DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceHealthScripts/{deviceHealthScript%2Did}/getGlobalScriptHighestAvailableVersion", pathParameters),
    }
    return m
}
// NewDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder instantiates a new GetGlobalScriptHighestAvailableVersionRequestBuilder and sets the default values.
func NewDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update the Proprietary Device Health Script
// Deprecated: This method is obsolete. Use PostAsGetGlobalScriptHighestAvailableVersionPostResponse instead.
func (m *DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) Post(ctx context.Context, requestConfiguration *DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration)(DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionResponseable), nil
}
// PostAsGetGlobalScriptHighestAvailableVersionPostResponse update the Proprietary Device Health Script
func (m *DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) PostAsGetGlobalScriptHighestAvailableVersionPostResponse(ctx context.Context, requestConfiguration *DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration)(DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionPostResponseable), nil
}
// ToPostRequestInformation update the Proprietary Device Health Script
func (m *DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) WithUrl(rawUrl string)(*DeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder) {
    return NewDeviceHealthScriptsItemGetGlobalScriptHighestAvailableVersionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
