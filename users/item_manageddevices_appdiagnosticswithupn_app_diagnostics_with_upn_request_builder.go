package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder provides operations to call the appDiagnostics method.
type ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetQueryParameters invoke function appDiagnostics
type ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetQueryParameters struct {
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
// ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetQueryParameters
}
// NewItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderInternal instantiates a new ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder and sets the default values.
func NewItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, upn *string)(*ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) {
    m := &ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/managedDevices/appDiagnostics(upn='{upn}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if upn != nil {
        m.BaseRequestBuilder.PathParameters["upn"] = *upn
    }
    return m
}
// NewItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder instantiates a new ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder and sets the default values.
func NewItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function appDiagnostics
// Deprecated: This method is obsolete. Use GetAsAppDiagnosticsWithUpnGetResponse instead.
// returns a ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnResponseable), nil
}
// GetAsAppDiagnosticsWithUpnGetResponse invoke function appDiagnostics
// returns a ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) GetAsAppDiagnosticsWithUpnGetResponse(ctx context.Context, requestConfiguration *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnGetResponseable), nil
}
// ToGetRequestInformation invoke function appDiagnostics
// returns a *RequestInformation when successful
func (m *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder when successful
func (m *ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) WithUrl(rawUrl string)(*ItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder) {
    return NewItemManageddevicesAppdiagnosticswithupnAppDiagnosticsWithUpnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
