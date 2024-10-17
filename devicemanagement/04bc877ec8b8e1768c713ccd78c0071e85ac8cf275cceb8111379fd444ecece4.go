package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder provides operations to call the retrievePowerliftAppDiagnosticsDetails method.
type ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetQueryParameters invoke function retrievePowerliftAppDiagnosticsDetails
type ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetQueryParameters struct {
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
// ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetQueryParameters
}
// NewComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderInternal instantiates a new ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder and sets the default values.
func NewComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, userPrincipalName *string)(*ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) {
    m := &ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/comanagedDevices/retrievePowerliftAppDiagnosticsDetails(userPrincipalName='{userPrincipalName}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if userPrincipalName != nil {
        m.BaseRequestBuilder.PathParameters["userPrincipalName"] = *userPrincipalName
    }
    return m
}
// NewComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder instantiates a new ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder and sets the default values.
func NewComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function retrievePowerliftAppDiagnosticsDetails
// Deprecated: This method is obsolete. Use GetAsRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse instead.
// returns a ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) Get(ctx context.Context, requestConfiguration *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseable), nil
}
// GetAsRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse invoke function retrievePowerliftAppDiagnosticsDetails
// returns a ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) GetAsRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse(ctx context.Context, requestConfiguration *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable), nil
}
// ToGetRequestInformation invoke function retrievePowerliftAppDiagnosticsDetails
// returns a *RequestInformation when successful
func (m *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder when successful
func (m *ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) WithUrl(rawUrl string)(*ComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) {
    return NewComanagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
