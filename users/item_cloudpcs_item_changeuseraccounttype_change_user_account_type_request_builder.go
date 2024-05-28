package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder provides operations to call the changeUserAccountType method.
type ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderInternal instantiates a new ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder and sets the default values.
func NewItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder) {
    m := &ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/changeUserAccountType", pathParameters),
    }
    return m
}
// NewItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder instantiates a new ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder and sets the default values.
func NewItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post change the account type of the user on a specific Cloud PC.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-changeuseraccounttype?view=graph-rest-beta
func (m *ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder) Post(ctx context.Context, body ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypePostRequestBodyable, requestConfiguration *ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation change the account type of the user on a specific Cloud PC.
// returns a *RequestInformation when successful
func (m *ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypePostRequestBodyable, requestConfiguration *ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder when successful
func (m *ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder) WithUrl(rawUrl string)(*ItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder) {
    return NewItemCloudpcsItemChangeuseraccounttypeChangeUserAccountTypeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
