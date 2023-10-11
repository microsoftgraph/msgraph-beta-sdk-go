package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder provides operations to call the complete method.
type ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderInternal instantiates a new CompleteRequestBuilder and sets the default values.
func NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) {
    m := &ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook/taskGroups/{outlookTaskGroup%2Did}/taskFolders/{outlookTaskFolder%2Did}/tasks/{outlookTask%2Did}/complete", pathParameters),
    }
    return m
}
// NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder instantiates a new CompleteRequestBuilder and sets the default values.
func NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderInternal(urlParams, requestAdapter)
}
// Post complete an Outlook task which sets the completedDateTime property to the current date, and the status property to completed. If you are completing a task in a recurring series, in the response, the task collection will contain the completed task in the series, and the next task in the series. The completedDateTime property represents the date when the task is finished. The time portion of completedDateTime is set to midnight UTC by default. By default, this operation (and the POST, GET, and PATCH task operations) returns date-related properties in UTC. You can use the Prefer: outlook.timezone header to have all the date-related properties in the response represented in a time zone different than UTC. This API is available in the following national cloud deployments.
// Deprecated: This method is obsolete. Use PostAsCompletePostResponse instead.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/outlooktask-complete?view=graph-rest-1.0
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration)(ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteResponseable), nil
}
// PostAsCompletePostResponse complete an Outlook task which sets the completedDateTime property to the current date, and the status property to completed. If you are completing a task in a recurring series, in the response, the task collection will contain the completed task in the series, and the next task in the series. The completedDateTime property represents the date when the task is finished. The time portion of completedDateTime is set to midnight UTC by default. By default, this operation (and the POST, GET, and PATCH task operations) returns date-related properties in UTC. You can use the Prefer: outlook.timezone header to have all the date-related properties in the response represented in a time zone different than UTC. This API is available in the following national cloud deployments.
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/outlooktask-complete?view=graph-rest-1.0
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) PostAsCompletePostResponse(ctx context.Context, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration)(ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompletePostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompletePostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompletePostResponseable), nil
}
// ToPostRequestInformation complete an Outlook task which sets the completedDateTime property to the current date, and the status property to completed. If you are completing a task in a recurring series, in the response, the task collection will contain the completed task in the series, and the next task in the series. The completedDateTime property represents the date when the task is finished. The time portion of completedDateTime is set to midnight UTC by default. By default, this operation (and the POST, GET, and PATCH task operations) returns date-related properties in UTC. You can use the Prefer: outlook.timezone header to have all the date-related properties in the response represented in a time zone different than UTC. This API is available in the following national cloud deployments.
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The Outlook tasks API is deprecated and will stop returning data on February 20, 2023. Please use the new To Do API. For more details, please visit https://developer.microsoft.com/en-us/office/blogs/announcing-the-general-availability-of-microsoft-to-do-apis-on-graph/ as of 2020-08/Outlook_Tasks on 2021-02-20 and will be removed 2023-02-20
func (m *ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) WithUrl(rawUrl string)(*ItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder) {
    return NewItemOutlookTaskGroupsItemTaskFoldersItemTasksItemCompleteRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
