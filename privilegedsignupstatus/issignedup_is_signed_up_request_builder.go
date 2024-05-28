package privilegedsignupstatus

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// IssignedupIsSignedUpRequestBuilder provides operations to call the isSignedUp method.
type IssignedupIsSignedUpRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IssignedupIsSignedUpRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IssignedupIsSignedUpRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIssignedupIsSignedUpRequestBuilderInternal instantiates a new IssignedupIsSignedUpRequestBuilder and sets the default values.
func NewIssignedupIsSignedUpRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IssignedupIsSignedUpRequestBuilder) {
    m := &IssignedupIsSignedUpRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedSignupStatus/isSignedUp()", pathParameters),
    }
    return m
}
// NewIssignedupIsSignedUpRequestBuilder instantiates a new IssignedupIsSignedUpRequestBuilder and sets the default values.
func NewIssignedupIsSignedUpRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IssignedupIsSignedUpRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIssignedupIsSignedUpRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function isSignedUp
// Deprecated: This method is obsolete. Use GetAsIsSignedUpGetResponse instead.
// returns a IssignedupIsSignedUpResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IssignedupIsSignedUpRequestBuilder) Get(ctx context.Context, requestConfiguration *IssignedupIsSignedUpRequestBuilderGetRequestConfiguration)(IssignedupIsSignedUpResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIssignedupIsSignedUpResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IssignedupIsSignedUpResponseable), nil
}
// GetAsIsSignedUpGetResponse invoke function isSignedUp
// returns a IssignedupIsSignedUpGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *IssignedupIsSignedUpRequestBuilder) GetAsIsSignedUpGetResponse(ctx context.Context, requestConfiguration *IssignedupIsSignedUpRequestBuilderGetRequestConfiguration)(IssignedupIsSignedUpGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateIssignedupIsSignedUpGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(IssignedupIsSignedUpGetResponseable), nil
}
// ToGetRequestInformation invoke function isSignedUp
// returns a *RequestInformation when successful
func (m *IssignedupIsSignedUpRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IssignedupIsSignedUpRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *IssignedupIsSignedUpRequestBuilder when successful
func (m *IssignedupIsSignedUpRequestBuilder) WithUrl(rawUrl string)(*IssignedupIsSignedUpRequestBuilder) {
    return NewIssignedupIsSignedUpRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
