package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters the sections in the section group. Read-only. Nullable.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters
}
// GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderInternal instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    m := &GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToNotebook provides operations to call the copyToNotebook method.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) CopyToNotebook()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CopyToSectionGroup provides operations to call the copyToSectionGroup method.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) CopyToSectionGroup()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property sections for groups
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the sections in the section group. Read-only. Nullable.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sections in groups
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sections for groups
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the sections in the section group. Read-only. Nullable.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable), nil
}
// Pages provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) Pages()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PagesById provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) PagesById(id string)(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenoteSection entity.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) ParentNotebook()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemParentNotebookRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSectionGroup provides operations to manage the parentSectionGroup property of the microsoft.graph.onenoteSection entity.
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) ParentSectionGroup()(*GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemParentSectionGroupRequestBuilder) {
    return NewGroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsItemParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sections in groups
func (m *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, requestConfiguration *GroupsItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable), nil
}
