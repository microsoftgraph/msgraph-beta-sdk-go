package privilegedsignupstatus

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CompletesetupCompleteSetupRequestBuilder provides operations to call the completeSetup method.
type CompletesetupCompleteSetupRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompletesetupCompleteSetupRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompletesetupCompleteSetupRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewCompletesetupCompleteSetupRequestBuilderInternal instantiates a new CompletesetupCompleteSetupRequestBuilder and sets the default values.
func NewCompletesetupCompleteSetupRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompletesetupCompleteSetupRequestBuilder) {
    m := &CompletesetupCompleteSetupRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/privilegedSignupStatus/completeSetup", pathParameters),
    }
    return m
}
// NewCompletesetupCompleteSetupRequestBuilder instantiates a new CompletesetupCompleteSetupRequestBuilder and sets the default values.
func NewCompletesetupCompleteSetupRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompletesetupCompleteSetupRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompletesetupCompleteSetupRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action completeSetup
// Deprecated: This method is obsolete. Use PostAsCompleteSetupPostResponse instead.
// returns a CompletesetupCompleteSetupResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompletesetupCompleteSetupRequestBuilder) Post(ctx context.Context, body CompletesetupCompleteSetupPostRequestBodyable, requestConfiguration *CompletesetupCompleteSetupRequestBuilderPostRequestConfiguration)(CompletesetupCompleteSetupResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompletesetupCompleteSetupResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CompletesetupCompleteSetupResponseable), nil
}
// PostAsCompleteSetupPostResponse invoke action completeSetup
// returns a CompletesetupCompleteSetupPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CompletesetupCompleteSetupRequestBuilder) PostAsCompleteSetupPostResponse(ctx context.Context, body CompletesetupCompleteSetupPostRequestBodyable, requestConfiguration *CompletesetupCompleteSetupRequestBuilderPostRequestConfiguration)(CompletesetupCompleteSetupPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateCompletesetupCompleteSetupPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(CompletesetupCompleteSetupPostResponseable), nil
}
// ToPostRequestInformation invoke action completeSetup
// returns a *RequestInformation when successful
func (m *CompletesetupCompleteSetupRequestBuilder) ToPostRequestInformation(ctx context.Context, body CompletesetupCompleteSetupPostRequestBodyable, requestConfiguration *CompletesetupCompleteSetupRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *CompletesetupCompleteSetupRequestBuilder when successful
func (m *CompletesetupCompleteSetupRequestBuilder) WithUrl(rawUrl string)(*CompletesetupCompleteSetupRequestBuilder) {
    return NewCompletesetupCompleteSetupRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
