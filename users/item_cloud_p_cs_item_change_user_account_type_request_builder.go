package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemCloudPCsItemChangeUserAccountTypeRequestBuilder provides operations to call the changeUserAccountType method.
type ItemCloudPCsItemChangeUserAccountTypeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloudPCsItemChangeUserAccountTypeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloudPCsItemChangeUserAccountTypeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloudPCsItemChangeUserAccountTypeRequestBuilderInternal instantiates a new ChangeUserAccountTypeRequestBuilder and sets the default values.
func NewItemCloudPCsItemChangeUserAccountTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemChangeUserAccountTypeRequestBuilder) {
    m := &ItemCloudPCsItemChangeUserAccountTypeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/cloudPCs/{cloudPC%2Did}/changeUserAccountType", pathParameters),
    }
    return m
}
// NewItemCloudPCsItemChangeUserAccountTypeRequestBuilder instantiates a new ChangeUserAccountTypeRequestBuilder and sets the default values.
func NewItemCloudPCsItemChangeUserAccountTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloudPCsItemChangeUserAccountTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloudPCsItemChangeUserAccountTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post change the account type of the user on a specific Cloud PC.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/cloudpc-changeuseraccounttype?view=graph-rest-1.0
func (m *ItemCloudPCsItemChangeUserAccountTypeRequestBuilder) Post(ctx context.Context, body ItemCloudPCsItemChangeUserAccountTypePostRequestBodyable, requestConfiguration *ItemCloudPCsItemChangeUserAccountTypeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation change the account type of the user on a specific Cloud PC.
func (m *ItemCloudPCsItemChangeUserAccountTypeRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemCloudPCsItemChangeUserAccountTypePostRequestBodyable, requestConfiguration *ItemCloudPCsItemChangeUserAccountTypeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *ItemCloudPCsItemChangeUserAccountTypeRequestBuilder) WithUrl(rawUrl string)(*ItemCloudPCsItemChangeUserAccountTypeRequestBuilder) {
    return NewItemCloudPCsItemChangeUserAccountTypeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
