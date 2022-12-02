package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder provides operations to call the restore method.
type GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal instantiates a new RestoreRequestBuilder and sets the default values.
func NewGroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) {
    m := &GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}/documentSetVersions/{documentSetVersion%2Did}/microsoft.graph.restore";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder instantiates a new RestoreRequestBuilder and sets the default values.
func NewGroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation restore a document set version.
func (m *GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post restore a document set version.
func (m *GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilder) Post(ctx context.Context, requestConfiguration *GroupsItemSitesItemListsItemItemsItemDocumentSetVersionsItemRestoreRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
