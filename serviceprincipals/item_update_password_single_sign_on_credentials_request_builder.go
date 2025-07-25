// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder provides operations to call the updatePasswordSingleSignOnCredentials method.
type ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemUpdatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemUpdatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal instantiates a new ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    m := &ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/updatePasswordSingleSignOnCredentials", pathParameters),
    }
    return m
}
// NewItemUpdatePasswordSingleSignOnCredentialsRequestBuilder instantiates a new ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder and sets the default values.
func NewItemUpdatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUpdatePasswordSingleSignOnCredentialsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post update single sign-on credentials using a password for a user or group.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-updatepasswordsinglesignoncredentials?view=graph-rest-beta
func (m *ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder) Post(ctx context.Context, body ItemUpdatePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemUpdatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation update single sign-on credentials using a password for a user or group.
// returns a *RequestInformation when successful
func (m *ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemUpdatePasswordSingleSignOnCredentialsPostRequestBodyable, requestConfiguration *ItemUpdatePasswordSingleSignOnCredentialsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder when successful
func (m *ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder) WithUrl(rawUrl string)(*ItemUpdatePasswordSingleSignOnCredentialsRequestBuilder) {
    return NewItemUpdatePasswordSingleSignOnCredentialsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
