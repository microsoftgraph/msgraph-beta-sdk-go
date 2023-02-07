package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// SharedDriveItemItemRequestBuilder provides operations to manage the collection of sharedDriveItem entities.
type SharedDriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SharedDriveItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharedDriveItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SharedDriveItemItemRequestBuilderGetQueryParameters access a shared DriveItem or a collection of shared items by using a **shareId** or sharing URL. To use a sharing URL with this API, your app needs to transform the URL into a sharing token.
type SharedDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SharedDriveItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharedDriveItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SharedDriveItemItemRequestBuilderGetQueryParameters
}
// SharedDriveItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SharedDriveItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSharedDriveItemItemRequestBuilderInternal instantiates a new SharedDriveItemItemRequestBuilder and sets the default values.
func NewSharedDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharedDriveItemItemRequestBuilder) {
    m := &SharedDriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewSharedDriveItemItemRequestBuilder instantiates a new SharedDriveItemItemRequestBuilder and sets the default values.
func NewSharedDriveItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SharedDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharedDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *SharedDriveItemItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DriveItem provides operations to manage the driveItem property of the microsoft.graph.sharedDriveItem entity.
func (m *SharedDriveItemItemRequestBuilder) DriveItem()(*ItemDriveItemRequestBuilder) {
    return NewItemDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get access a shared DriveItem or a collection of shared items by using a **shareId** or sharing URL. To use a sharing URL with this API, your app needs to transform the URL into a sharing token.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/shares-get?view=graph-rest-1.0
func (m *SharedDriveItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *SharedDriveItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharedDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.sharedDriveItem entity.
func (m *SharedDriveItemItemRequestBuilder) Items()(*ItemItemsRequestBuilder) {
    return NewItemItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ItemsById provides operations to manage the items property of the microsoft.graph.sharedDriveItem entity.
func (m *SharedDriveItemItemRequestBuilder) ItemsById(id string)(*ItemItemsDriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem%2Did"] = id
    }
    return NewItemItemsDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter)
}
// List provides operations to manage the list property of the microsoft.graph.sharedDriveItem entity.
func (m *SharedDriveItemItemRequestBuilder) List()(*ItemListRequestBuilder) {
    return NewItemListRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ListItem provides operations to manage the listItem property of the microsoft.graph.sharedDriveItem entity.
func (m *SharedDriveItemItemRequestBuilder) ListItem()(*ItemListItemRequestBuilder) {
    return NewItemListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Patch update entity in shares
func (m *SharedDriveItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable, requestConfiguration *SharedDriveItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateSharedDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable), nil
}
// Permission provides operations to manage the permission property of the microsoft.graph.sharedDriveItem entity.
func (m *SharedDriveItemItemRequestBuilder) Permission()(*ItemPermissionRequestBuilder) {
    return NewItemPermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Root provides operations to manage the root property of the microsoft.graph.sharedDriveItem entity.
func (m *SharedDriveItemItemRequestBuilder) Root()(*ItemRootRequestBuilder) {
    return NewItemRootRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Site provides operations to manage the site property of the microsoft.graph.sharedDriveItem entity.
func (m *SharedDriveItemItemRequestBuilder) Site()(*ItemSiteRequestBuilder) {
    return NewItemSiteRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ToDeleteRequestInformation delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SharedDriveItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation access a shared DriveItem or a collection of shared items by using a **shareId** or sharing URL. To use a sharing URL with this API, your app needs to transform the URL into a sharing token.
func (m *SharedDriveItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SharedDriveItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
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
// ToPatchRequestInformation update entity in shares
func (m *SharedDriveItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.SharedDriveItemable, requestConfiguration *SharedDriveItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
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
