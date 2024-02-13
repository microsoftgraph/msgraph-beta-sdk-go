package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder provides operations to manage the pages property of the microsoft.graph.onenoteSection entity.
type ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters the collection of pages in the section.  Read-only. Nullable.
type ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderGetQueryParameters
}
// ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderInternal instantiates a new ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder and sets the default values.
func NewItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) {
    m := &ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/onenote/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder instantiates a new ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder and sets the default values.
func NewItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the user entity.
// returns a *ItemOnenoteSectionsItemPagesItemContentRequestBuilder when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) Content()(*ItemOnenoteSectionsItemPagesItemContentRequestBuilder) {
    return NewItemOnenoteSectionsItemPagesItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CopyToSection provides operations to call the copyToSection method.
// returns a *ItemOnenoteSectionsItemPagesItemCopyToSectionRequestBuilder when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) CopyToSection()(*ItemOnenoteSectionsItemPagesItemCopyToSectionRequestBuilder) {
    return NewItemOnenoteSectionsItemPagesItemCopyToSectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property pages for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of pages in the section.  Read-only. Nullable.
// returns a OnenotePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable), nil
}
// OnenotePatchContent provides operations to call the onenotePatchContent method.
// returns a *ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) OnenotePatchContent()(*ItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilder) {
    return NewItemOnenoteSectionsItemPagesItemOnenotePatchContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentNotebook provides operations to manage the parentNotebook property of the microsoft.graph.onenotePage entity.
// returns a *ItemOnenoteSectionsItemPagesItemParentNotebookRequestBuilder when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) ParentNotebook()(*ItemOnenoteSectionsItemPagesItemParentNotebookRequestBuilder) {
    return NewItemOnenoteSectionsItemPagesItemParentNotebookRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ParentSection provides operations to manage the parentSection property of the microsoft.graph.onenotePage entity.
// returns a *ItemOnenoteSectionsItemPagesItemParentSectionRequestBuilder when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) ParentSection()(*ItemOnenoteSectionsItemPagesItemParentSectionRequestBuilder) {
    return NewItemOnenoteSectionsItemPagesItemParentSectionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property pages in users
// returns a OnenotePageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, requestConfiguration *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateOnenotePageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable), nil
}
// Preview provides operations to call the preview method.
// returns a *ItemOnenoteSectionsItemPagesItemPreviewRequestBuilder when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) Preview()(*ItemOnenoteSectionsItemPagesItemPreviewRequestBuilder) {
    return NewItemOnenoteSectionsItemPagesItemPreviewRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property pages for users
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/users/{user%2Did}/onenote/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}", m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of pages in the section.  Read-only. Nullable.
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property pages in users
// returns a *RequestInformation when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.OnenotePageable, requestConfiguration *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, "{+baseurl}/users/{user%2Did}/onenote/sections/{onenoteSection%2Did}/pages/{onenotePage%2Did}", m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder when successful
func (m *ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) WithUrl(rawUrl string)(*ItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder) {
    return NewItemOnenoteSectionsItemPagesOnenotePageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
