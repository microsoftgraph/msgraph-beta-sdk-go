package auditlogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SigninsConfirmsafeConfirmSafeRequestBuilder provides operations to call the confirmSafe method.
type SigninsConfirmsafeConfirmSafeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SigninsConfirmsafeConfirmSafeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SigninsConfirmsafeConfirmSafeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSigninsConfirmsafeConfirmSafeRequestBuilderInternal instantiates a new SigninsConfirmsafeConfirmSafeRequestBuilder and sets the default values.
func NewSigninsConfirmsafeConfirmSafeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SigninsConfirmsafeConfirmSafeRequestBuilder) {
    m := &SigninsConfirmsafeConfirmSafeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/auditLogs/signIns/confirmSafe", pathParameters),
    }
    return m
}
// NewSigninsConfirmsafeConfirmSafeRequestBuilder instantiates a new SigninsConfirmsafeConfirmSafeRequestBuilder and sets the default values.
func NewSigninsConfirmsafeConfirmSafeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SigninsConfirmsafeConfirmSafeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSigninsConfirmsafeConfirmSafeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post allow admins to mark an event in Microsoft Entra sign-in logs as safe. Admins can either mark the events flagged as risky by Microsoft Entra ID Protection as safe, or they can mark unflagged events as safe. For details about investigating Identity Protection risks, see How to investigate risk.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/signin-confirmsafe?view=graph-rest-beta
func (m *SigninsConfirmsafeConfirmSafeRequestBuilder) Post(ctx context.Context, body SigninsConfirmsafeConfirmSafePostRequestBodyable, requestConfiguration *SigninsConfirmsafeConfirmSafeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation allow admins to mark an event in Microsoft Entra sign-in logs as safe. Admins can either mark the events flagged as risky by Microsoft Entra ID Protection as safe, or they can mark unflagged events as safe. For details about investigating Identity Protection risks, see How to investigate risk.
// returns a *RequestInformation when successful
func (m *SigninsConfirmsafeConfirmSafeRequestBuilder) ToPostRequestInformation(ctx context.Context, body SigninsConfirmsafeConfirmSafePostRequestBodyable, requestConfiguration *SigninsConfirmsafeConfirmSafeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SigninsConfirmsafeConfirmSafeRequestBuilder when successful
func (m *SigninsConfirmsafeConfirmSafeRequestBuilder) WithUrl(rawUrl string)(*SigninsConfirmsafeConfirmSafeRequestBuilder) {
    return NewSigninsConfirmsafeConfirmSafeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
