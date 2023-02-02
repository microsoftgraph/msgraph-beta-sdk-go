package auditlogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder provides operations to call the confirmSafe method.
type SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilderInternal instantiates a new ConfirmSafeRequestBuilder and sets the default values.
func NewSignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder) {
    m := &SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/auditLogs/signIns/microsoft.graph.confirmSafe";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder instantiates a new ConfirmSafeRequestBuilder and sets the default values.
func NewSignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post allow admins to mark an event in Azure AD sign in logs as safe. Admins can either mark the events flagged as risky by Azure AD Identity Protection as safe, or they can mark unflagged events as safe. For details about investigating Identity Protection risks, see How to investigate risk.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/signin-confirmsafe?view=graph-rest-1.0
func (m *SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder) Post(ctx context.Context, body SignInsMicrosoftGraphConfirmSafeConfirmSafePostRequestBodyable, requestConfiguration *SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation allow admins to mark an event in Azure AD sign in logs as safe. Admins can either mark the events flagged as risky by Azure AD Identity Protection as safe, or they can mark unflagged events as safe. For details about investigating Identity Protection risks, see How to investigate risk.
func (m *SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilder) ToPostRequestInformation(ctx context.Context, body SignInsMicrosoftGraphConfirmSafeConfirmSafePostRequestBodyable, requestConfiguration *SignInsMicrosoftGraphConfirmSafeConfirmSafeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
