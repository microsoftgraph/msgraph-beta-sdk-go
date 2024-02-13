package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder provides operations to manage the changes property of the microsoft.graph.workbookDocumentTask entity.
type ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderGetQueryParameters a collection of task change histories.
type ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderGetQueryParameters
}
// ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tasks/{workbookDocumentTask%2Did}/comment/replies/{workbookCommentReply%2Did}/task/changes/{workbookDocumentTaskChange%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property changes for drives
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get a collection of task change histories.
// returns a WorkbookDocumentTaskChangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookDocumentTaskChangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable), nil
}
// Patch update the navigation property changes in drives
// returns a WorkbookDocumentTaskChangeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateWorkbookDocumentTaskChangeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable), nil
}
// ToDeleteRequestInformation delete navigation property changes for drives
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tasks/{workbookDocumentTask%2Did}/comment/replies/{workbookCommentReply%2Did}/task/changes/{workbookDocumentTaskChange%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of task change histories.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property changes in drives
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.WorkbookDocumentTaskChangeable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tasks/{workbookDocumentTask%2Did}/comment/replies/{workbookCommentReply%2Did}/task/changes/{workbookDocumentTaskChange%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemTasksItemCommentRepliesItemTaskChangesWorkbookDocumentTaskChangeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
