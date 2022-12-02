package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters
}
// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    m := &GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder instantiates a new OnenotePageItemRequestBuilder and sets the default values.
func NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the group entity.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Content()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CopyToSection provides operations to call the copyToSection method.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CopyToSection()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemCopyToSectionRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemCopyToSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property pages for groups
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property pages in groups
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property pages for groups
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the collection of pages in the section.  Read-only. Nullable.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable), nil
}
// OnenotePatchContent provides operations to call the onenotePatchContent method.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) OnenotePatchContent()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemOnenotePatchContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentNotebook()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSection provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) ParentSection()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemParentSectionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property pages in groups
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable), nil
}
// Preview provides operations to call the preview method.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) Preview()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemPreviewRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesItemPreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
