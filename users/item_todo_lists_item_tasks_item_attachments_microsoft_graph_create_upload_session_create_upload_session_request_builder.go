package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder provides operations to call the createUploadSession method.
type ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderInternal instantiates a new CreateUploadSessionRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder) {
    m := &ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}/attachments/microsoft.graph.createUploadSession";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder instantiates a new CreateUploadSessionRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create an upload session to iteratively upload ranges of a file as an attachment to a todoTask. As part of the response, this action returns an upload URL that you can use in subsequent sequential `PUT` queries. The request headers for each `PUT` operation let you specify the exact range of bytes to be uploaded. This allows the transfer to be resumed, in case the network connection is dropped during the upload. The following are the steps to attach a file to a Microsoft To Do task using an upload session: For an example that describes the end-to-end attachment process, see attach files to a To Do task.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/taskfileattachment-createuploadsession?view=graph-rest-1.0
func (m *ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder) Post(ctx context.Context, body ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UploadSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateUploadSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.UploadSessionable), nil
}
// ToPostRequestInformation create an upload session to iteratively upload ranges of a file as an attachment to a todoTask. As part of the response, this action returns an upload URL that you can use in subsequent sequential `PUT` queries. The request headers for each `PUT` operation let you specify the exact range of bytes to be uploaded. This allows the transfer to be resumed, in case the network connection is dropped during the upload. The following are the steps to attach a file to a Microsoft To Do task using an upload session: For an example that describes the end-to-end attachment process, see attach files to a To Do task.
func (m *ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionPostRequestBodyable, requestConfiguration *ItemTodoListsItemTasksItemAttachmentsMicrosoftGraphCreateUploadSessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
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
