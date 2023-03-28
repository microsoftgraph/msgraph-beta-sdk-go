package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder provides operations to call the complete method.
type ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilderInternal instantiates a new CompleteRequestBuilder and sets the default values.
func NewItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder) {
    m := &ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook/taskFolders/{outlookTaskFolder%2Did}/tasks/{outlookTask%2Did}/complete", pathParameters),
    }
    return m
}
// NewItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder instantiates a new CompleteRequestBuilder and sets the default values.
func NewItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action complete
func (m *ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration)(ItemOutlookTaskFoldersItemTasksItemCompleteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookTaskFoldersItemTasksItemCompleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookTaskFoldersItemTasksItemCompleteResponseable), nil
}
// ToPostRequestInformation invoke action complete
func (m *ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
