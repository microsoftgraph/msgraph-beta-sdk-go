package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder provides operations to call the retrievePowerliftAppDiagnosticsDetails method.
type ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetQueryParameters invoke function retrievePowerliftAppDiagnosticsDetails
type ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetQueryParameters struct {
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
// ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetQueryParameters
}
// NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderInternal instantiates a new ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder and sets the default values.
func NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, userPrincipalName *string)(*ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) {
    m := &ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/managedDevices/retrievePowerliftAppDiagnosticsDetails(userPrincipalName='{userPrincipalName}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if userPrincipalName != nil {
        m.BaseRequestBuilder.PathParameters["userPrincipalName"] = *userPrincipalName
    }
    return m
}
// NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder instantiates a new ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder and sets the default values.
func NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function retrievePowerliftAppDiagnosticsDetails
// Deprecated: This method is obsolete. Use GetAsRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse instead.
// returns a ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameResponseable), nil
}
// GetAsRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse invoke function retrievePowerliftAppDiagnosticsDetails
// returns a ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) GetAsRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponse(ctx context.Context, requestConfiguration *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameGetResponseable), nil
}
// ToGetRequestInformation invoke function retrievePowerliftAppDiagnosticsDetails
// returns a *RequestInformation when successful
func (m *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder when successful
func (m *ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) WithUrl(rawUrl string)(*ManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder) {
    return NewManagedDevicesRetrievePowerliftAppDiagnosticsDetailsWithUserPrincipalNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
