package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder provides operations to call the getLicensesForApp method.
type VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderGetQueryParameters invoke function getLicensesForApp
type VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderGetQueryParameters struct {
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
// VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderGetQueryParameters
}
// NewVpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderInternal instantiates a new VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder and sets the default values.
func NewVpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, bundleId *string)(*VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder) {
    m := &VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/vppTokens/getLicensesForApp(bundleId='{bundleId}'){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    if bundleId != nil {
        m.BaseRequestBuilder.PathParameters["bundleId"] = *bundleId
    }
    return m
}
// NewVpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder instantiates a new VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder and sets the default values.
func NewVpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getLicensesForApp
// Deprecated: This method is obsolete. Use GetAsGetLicensesForAppWithBundleIdGetResponse instead.
// returns a VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder) Get(ctx context.Context, requestConfiguration *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration)(VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdResponseable), nil
}
// GetAsGetLicensesForAppWithBundleIdGetResponse invoke function getLicensesForApp
// returns a VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder) GetAsGetLicensesForAppWithBundleIdGetResponse(ctx context.Context, requestConfiguration *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration)(VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateVpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdGetResponseable), nil
}
// ToGetRequestInformation invoke function getLicensesForApp
// returns a *RequestInformation when successful
func (m *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder when successful
func (m *VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder) WithUrl(rawUrl string)(*VpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder) {
    return NewVpptokensGetlicensesforappwithbundleidGetLicensesForAppWithBundleIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
