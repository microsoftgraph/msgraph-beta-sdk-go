package privilegedsignupstatus

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CansignupCanSignUpRequestBuilder provides operations to call the canSignUp method.
type CansignupCanSignUpRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CansignupCanSignUpRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CansignupCanSignUpRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCansignupCanSignUpRequestBuilderInternal instantiates a new CansignupCanSignUpRequestBuilder and sets the default values.
func NewCansignupCanSignUpRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CansignupCanSignUpRequestBuilder) {
    m := &CansignupCanSignUpRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedSignupStatus/canSignUp()", pathParameters),
    }
    return m
}
// NewCansignupCanSignUpRequestBuilder instantiates a new CansignupCanSignUpRequestBuilder and sets the default values.
func NewCansignupCanSignUpRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CansignupCanSignUpRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCansignupCanSignUpRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function canSignUp
// Deprecated: This method is obsolete. Use GetAsCanSignUpGetResponse instead.
// returns a CansignupCanSignUpResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CansignupCanSignUpRequestBuilder) Get(ctx context.Context, requestConfiguration *CansignupCanSignUpRequestBuilderGetRequestConfiguration)(CansignupCanSignUpResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCansignupCanSignUpResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CansignupCanSignUpResponseable), nil
}
// GetAsCanSignUpGetResponse invoke function canSignUp
// returns a CansignupCanSignUpGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CansignupCanSignUpRequestBuilder) GetAsCanSignUpGetResponse(ctx context.Context, requestConfiguration *CansignupCanSignUpRequestBuilderGetRequestConfiguration)(CansignupCanSignUpGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCansignupCanSignUpGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CansignupCanSignUpGetResponseable), nil
}
// ToGetRequestInformation invoke function canSignUp
// returns a *RequestInformation when successful
func (m *CansignupCanSignUpRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CansignupCanSignUpRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CansignupCanSignUpRequestBuilder when successful
func (m *CansignupCanSignUpRequestBuilder) WithUrl(rawUrl string)(*CansignupCanSignUpRequestBuilder) {
    return NewCansignupCanSignUpRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
