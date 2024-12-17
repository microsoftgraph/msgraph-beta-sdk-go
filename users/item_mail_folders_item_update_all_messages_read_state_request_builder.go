package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder provides operations to call the updateAllMessagesReadState method.
type ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilderInternal instantiates a new ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder and sets the default values.
func NewItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder) {
    m := &ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/mailFolders/{mailFolder%2Did}/updateAllMessagesReadState", pathParameters),
    }
    return m
}
// NewItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder instantiates a new ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder and sets the default values.
func NewItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action updateAllMessagesReadState
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder) Post(ctx context.Context, body ItemMailFoldersItemUpdateAllMessagesReadStatePostRequestBodyable, requestConfiguration *ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action updateAllMessagesReadState
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI
// returns a *RequestInformation when successful
func (m *ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemMailFoldersItemUpdateAllMessagesReadStatePostRequestBodyable, requestConfiguration *ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated:  as of 2024-04/PrivatePreview:updateAllMessagesReadStateAPI
// returns a *ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder when successful
func (m *ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder) WithUrl(rawUrl string)(*ItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder) {
    return NewItemMailFoldersItemUpdateAllMessagesReadStateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
