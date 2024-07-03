package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder provides operations to call the appDiagnostics method.
type ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesAppDiagnosticsWithUpnRequestBuilderGetQueryParameters invoke function appDiagnostics
type ComanagedDevicesAppDiagnosticsWithUpnRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ComanagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilderGetQueryParameters
}
// NewComanagedDevicesAppDiagnosticsWithUpnRequestBuilderInternal instantiates a new ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder and sets the default values.
func NewComanagedDevicesAppDiagnosticsWithUpnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, upn *string)(*ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) {
    m := &ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/appDiagnostics(upn='{upn}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if upn != nil {
        m.BaseRequestBuilder.PathParameters["upn"] = *upn
    }
    return m
}
// NewComanagedDevicesAppDiagnosticsWithUpnRequestBuilder instantiates a new ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder and sets the default values.
func NewComanagedDevicesAppDiagnosticsWithUpnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesAppDiagnosticsWithUpnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function appDiagnostics
// Deprecated: This method is obsolete. Use GetAsAppDiagnosticsWithUpnGetResponse instead.
// returns a ComanagedDevicesAppDiagnosticsWithUpnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(ComanagedDevicesAppDiagnosticsWithUpnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateComanagedDevicesAppDiagnosticsWithUpnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ComanagedDevicesAppDiagnosticsWithUpnResponseable), nil
}
// GetAsAppDiagnosticsWithUpnGetResponse invoke function appDiagnostics
// returns a ComanagedDevicesAppDiagnosticsWithUpnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) GetAsAppDiagnosticsWithUpnGetResponse(ctx context.Context, requestConfiguration *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(ComanagedDevicesAppDiagnosticsWithUpnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateComanagedDevicesAppDiagnosticsWithUpnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ComanagedDevicesAppDiagnosticsWithUpnGetResponseable), nil
}
// ToGetRequestInformation invoke function appDiagnostics
// returns a *RequestInformation when successful
func (m *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder when successful
func (m *ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesAppDiagnosticsWithUpnRequestBuilder) {
    return NewComanagedDevicesAppDiagnosticsWithUpnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
