package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder provides operations to call the appDiagnostics method.
type ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetQueryParameters invoke function appDiagnostics
type ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetQueryParameters struct {
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
// ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetQueryParameters
}
// NewManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderInternal instantiates a new ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder and sets the default values.
func NewManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, upn *string)(*ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) {
    m := &ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/appDiagnostics(upn='{upn}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if upn != nil {
        m.BaseRequestBuilder.PathParameters["upn"] = *upn
    }
    return m
}
// NewManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder instantiates a new ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder and sets the default values.
func NewManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function appDiagnostics
// Deprecated: This method is obsolete. Use GetAsAppDiagnosticsWithUpnGetResponse instead.
// returns a ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) Get(ctx context.Context, requestConfiguration *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnResponseable), nil
}
// GetAsAppDiagnosticsWithUpnGetResponse invoke function appDiagnostics
// returns a ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) GetAsAppDiagnosticsWithUpnGetResponse(ctx context.Context, requestConfiguration *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnGetResponseable), nil
}
// ToGetRequestInformation invoke function appDiagnostics
// returns a *RequestInformation when successful
func (m *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder when successful
func (m *ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) WithUrl(rawUrl string)(*ManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) {
    return NewManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
