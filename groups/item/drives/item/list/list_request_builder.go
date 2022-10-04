package list

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i1271f32d5f7d3b0586f8de93a5b0535ce223cbfd2b48e6f411f71c399da18327 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/columns"
    i43e0b5ce13128879f06e2b07e0ccd9f8dc1dafb84ac39aa8ef8923e069dc6a8b "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes"
    i7bfd5cc8c38cfc387360207e15abdf3e46f2fcf5d71bc51b310a210e16307965 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/subscriptions"
    i923c78037243955d3854651774d2282943abed938d26a8809587cc4b58252199 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/operations"
    ia7fc35da0de2bd686f2aa5a8a1fdfc88ac3a5bbd57126d8a1bd0644eceb39752 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/activities"
    ic55f9fa554d8a522a28b3fbd75b161bce003442994fcc35c64c623a12261a280 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/drive"
    idc21fe63ea818594f7f29929ecb8355791aa9fdd52478809de7eff177ca4c915 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/items"
    i649cdddb50a802c43dc3af1a06a94775bd06e4c17653045c20a5077da03ca95d "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/columns/item"
    i998c00f0ec1e2b40c5518299a0453c926896385255ff91e4bf3d2b7234a467a8 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/contenttypes/item"
    ia7e7318eb558bf8bfc7b2e73993fafd19e3292f1d7b06c2efa414823b5a839db "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/subscriptions/item"
    iae2fd07cb5857ac5712f84db0ed1eab2184e2ab848fee6addbad50e8b44d0302 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/operations/item"
    iba31399dc8f34e424c8ebc1f280f4f9e3e58432f001227720cab229ef435deaa "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/drives/item/list/items/item"
)

// ListRequestBuilder provides operations to manage the list property of the microsoft.graph.drive entity.
type ListRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ListRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ListRequestBuilderGetQueryParameters for drives in SharePoint, the underlying document library list. Read-only. Nullable.
type ListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ListRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ListRequestBuilderGetQueryParameters
}
// ListRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ListRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities the activities property
func (m *ListRequestBuilder) Activities()(*ia7fc35da0de2bd686f2aa5a8a1fdfc88ac3a5bbd57126d8a1bd0644eceb39752.ActivitiesRequestBuilder) {
    return ia7fc35da0de2bd686f2aa5a8a1fdfc88ac3a5bbd57126d8a1bd0644eceb39752.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *ListRequestBuilder) Columns()(*i1271f32d5f7d3b0586f8de93a5b0535ce223cbfd2b48e6f411f71c399da18327.ColumnsRequestBuilder) {
    return i1271f32d5f7d3b0586f8de93a5b0535ce223cbfd2b48e6f411f71c399da18327.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*i649cdddb50a802c43dc3af1a06a94775bd06e4c17653045c20a5077da03ca95d.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i649cdddb50a802c43dc3af1a06a94775bd06e4c17653045c20a5077da03ca95d.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/drives/{drive%2Did}/list{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewListRequestBuilder instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewListRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentTypes the contentTypes property
func (m *ListRequestBuilder) ContentTypes()(*i43e0b5ce13128879f06e2b07e0ccd9f8dc1dafb84ac39aa8ef8923e069dc6a8b.ContentTypesRequestBuilder) {
    return i43e0b5ce13128879f06e2b07e0ccd9f8dc1dafb84ac39aa8ef8923e069dc6a8b.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i998c00f0ec1e2b40c5518299a0453c926896385255ff91e4bf3d2b7234a467a8.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return i998c00f0ec1e2b40c5518299a0453c926896385255ff91e4bf3d2b7234a467a8.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for groups
func (m *ListRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ListRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property list in groups
func (m *ListRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property list for groups
func (m *ListRequestBuilder) Delete(ctx context.Context, requestConfiguration *ListRequestBuilderDeleteRequestConfiguration)(error) {
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
// Drive the drive property
func (m *ListRequestBuilder) Drive()(*ic55f9fa554d8a522a28b3fbd75b161bce003442994fcc35c64c623a12261a280.DriveRequestBuilder) {
    return ic55f9fa554d8a522a28b3fbd75b161bce003442994fcc35c64c623a12261a280.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Get(ctx context.Context, requestConfiguration *ListRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable), nil
}
// Items the items property
func (m *ListRequestBuilder) Items()(*idc21fe63ea818594f7f29929ecb8355791aa9fdd52478809de7eff177ca4c915.ItemsRequestBuilder) {
    return idc21fe63ea818594f7f29929ecb8355791aa9fdd52478809de7eff177ca4c915.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*iba31399dc8f34e424c8ebc1f280f4f9e3e58432f001227720cab229ef435deaa.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem%2Did"] = id
    }
    return iba31399dc8f34e424c8ebc1f280f4f9e3e58432f001227720cab229ef435deaa.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ListRequestBuilder) Operations()(*i923c78037243955d3854651774d2282943abed938d26a8809587cc4b58252199.OperationsRequestBuilder) {
    return i923c78037243955d3854651774d2282943abed938d26a8809587cc4b58252199.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*iae2fd07cb5857ac5712f84db0ed1eab2184e2ab848fee6addbad50e8b44d0302.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return iae2fd07cb5857ac5712f84db0ed1eab2184e2ab848fee6addbad50e8b44d0302.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in groups
func (m *ListRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable), nil
}
// Subscriptions the subscriptions property
func (m *ListRequestBuilder) Subscriptions()(*i7bfd5cc8c38cfc387360207e15abdf3e46f2fcf5d71bc51b310a210e16307965.SubscriptionsRequestBuilder) {
    return i7bfd5cc8c38cfc387360207e15abdf3e46f2fcf5d71bc51b310a210e16307965.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.groups.item.drives.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*ia7e7318eb558bf8bfc7b2e73993fafd19e3292f1d7b06c2efa414823b5a839db.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ia7e7318eb558bf8bfc7b2e73993fafd19e3292f1d7b06c2efa414823b5a839db.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
