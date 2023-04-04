package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters the sections in the section group. Read-only. Nullable.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetQueryParameters
}
// ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderInternal instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    m := &ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}/sections/{onenoteSection%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder instantiates a new OnenoteSectionItemRequestBuilder and sets the default values.
func NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToNotebook provides operations to call the copyToNotebook method.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) CopyToNotebook()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CopyToSectionGroup provides operations to call the copyToSectionGroup method.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) CopyToSectionGroup()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemCopyToSectionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property sections for groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Get the sections in the section group. Read-only. Nullable.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable), nil
}
// Pages provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) Pages()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PagesById provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) PagesById(id string)(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenotePage%2Did"] = id
    }
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenoteSection entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) ParentNotebook()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemParentNotebookRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemParentNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentSectionGroup provides operations to manage the parentSectionGroup property of the microsoft.graph.onenoteSection entity.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) ParentSectionGroup()(*ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemParentSectionGroupRequestBuilder) {
    return NewItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsItemParentSectionGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property sections in groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenoteSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable), nil
}
// ToDeleteRequestInformation delete navigation property sections for groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation the sections in the section group. Read-only. Nullable.
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property sections in groups
func (m *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenoteSectionable, requestConfiguration *ItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
