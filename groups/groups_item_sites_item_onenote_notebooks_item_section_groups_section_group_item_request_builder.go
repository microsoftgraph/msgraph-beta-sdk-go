package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder provides operations to manage the sectionGroups property of the microsoft.graph.notebook entity.
type GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters the section groups in the notebook. Read-only. Nullable.
type GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetQueryParameters
}
// GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderInternal instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) {
    m := &GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/onenote/notebooks/{notebook%2Did}/sectionGroups/{sectionGroup%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder instantiates a new SectionGroupItemRequestBuilder and sets the default values.
func NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property sectionGroups for groups
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the section groups in the notebook. Read-only. Nullable.
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property sectionGroups in groups
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupable, requestConfiguration *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property sectionGroups for groups
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the section groups in the notebook. Read-only. Nullable.
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) Get(ctx context.Context, requestConfiguration *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSectionGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupable), nil
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.sectionGroup entity.
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) ParentNotebook()(*GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemParentNotebookRequestBuilder) {
    return NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemParentNotebookRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParentSectionGroup provides operations to manage the parentSectionGroup property of the microsoft.graph.sectionGroup entity.
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) ParentSectionGroup()(*GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemParentSectionGroupRequestBuilder) {
    return NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemParentSectionGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property sectionGroups in groups
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupable, requestConfiguration *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSectionGroupFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SectionGroupable), nil
}
// SectionGroups provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) SectionGroups()(*GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilder) {
    return NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionGroupsById provides operations to manage the sectionGroups property of the microsoft.graph.sectionGroup entity.
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) SectionGroupsById(id string)(*GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["sectionGroup%2Did1"] = id
    }
    return NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionGroupsSectionGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Sections provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) Sections()(*GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsRequestBuilder) {
    return NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SectionsById provides operations to manage the sections property of the microsoft.graph.sectionGroup entity.
func (m *GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsSectionGroupItemRequestBuilder) SectionsById(id string)(*GroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["onenoteSection%2Did"] = id
    }
    return NewGroupsItemSitesItemOnenoteNotebooksItemSectionGroupsItemSectionsOnenoteSectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
