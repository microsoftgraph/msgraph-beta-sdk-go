package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemDeletePasswordSingleSignOnCredentialsRequestBuilder provides operations to call the deletePasswordSingleSignOnCredentials method.
type ItemDeletePasswordSingleSignOnCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemDeletePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemDeletePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemDeletePasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new ItemDeletePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeletePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &ItemDeletePasswordSingleSignOnCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/deletePasswordSingleSignOnCredentials", pathParameters),
    }
    return m
}
// NewItemDeletePasswordSingleSignOnCredentialsRequestBuilder instantiates a new ItemDeletePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemDeletePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemDeletePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemDeletePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post delete the password-based single sign-on credentials for a given user to a given service principal.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-deletepasswordsinglesignoncredentials?view=graph-rest-beta
func (m *ItemDeletePasswordSingleSignOnCredentialsRequestBuilder) Post(ctx context.Context, body ItemDeletePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemDeletePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation delete the password-based single sign-on credentials for a given user to a given service principal.
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *RequestInformation when successful
func (m *ItemDeletePasswordSingleSignOnCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemDeletePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemDeletePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-07/PrivatePreview:copilotExportAPI
// returns a *ItemDeletePasswordSingleSignOnCredentialsRequestBuilder when successful
func (m *ItemDeletePasswordSingleSignOnCredentialsRequestBuilder) WithUrl(rawUrl string)(*ItemDeletePasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemDeletePasswordSingleSignOnCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
