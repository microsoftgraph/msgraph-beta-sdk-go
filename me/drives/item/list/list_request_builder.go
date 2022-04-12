package list

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i072ddba6d43643cb4e0a492cfc6003149b0f11a1bbb94fcbbdcf71f0770d542a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items"
    i18a4f9c70e9242b75dfae66cc04b68ef7851da26b5ea1fbcef0a5f8835a49e7e "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/activities"
    i1f28b50836f57b5d58bbaf6536fcecef381f989c4e9cb3ea42add732311ae102 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/subscriptions"
    i261ceb17d8b68c77c74467f613cdf5532cd5e5ea820c84d683a39d39c13c4ebe "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/columns"
    ib0080b16f4fd1f64bad0fc40faa8191f83bc5af05b27dabb5054539619462021 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/drive"
    ic1ed073c992d668b0fff838202aaefd614de5739048c54fa321f3c0cf2913074 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/operations"
    icf07d83821e028c339f3d32b2cc131377b2d116e69b7fd99e7db53d773436c38 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes"
    i40668099437af2cc778494c140ef105574b3d11b369844ae96e6da4e56ed7dbb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/columns/item"
    i69531fc3cfd71461eb64582da70e123676f0660112f5d4521198ea126071edb4 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/operations/item"
    i8d0b9819c6c73c68982134065361fc5eb557175052e7f6f62cb6465cea9691ad "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/items/item"
    i9afbcf26a085bae33bec61cc4b40f70ee21e783407b414e87aa6707a9555ac2a "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/contenttypes/item"
    ibdce946586319f7ecb5c54b6cd60e9e1e6b2a7b464a4783217eb2d2c5fbc3414 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/drives/item/list/subscriptions/item"
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
// ListRequestBuilderDeleteOptions options for Delete
type ListRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ListRequestBuilderGetOptions options for Get
type ListRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ListRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ListRequestBuilderGetQueryParameters for drives in SharePoint, the underlying document library list. Read-only. Nullable.
type ListRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ListRequestBuilderPatchOptions options for Patch
type ListRequestBuilderPatchOptions struct {
    // 
    Body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Activities the activities property
func (m *ListRequestBuilder) Activities()(*i18a4f9c70e9242b75dfae66cc04b68ef7851da26b5ea1fbcef0a5f8835a49e7e.ActivitiesRequestBuilder) {
    return i18a4f9c70e9242b75dfae66cc04b68ef7851da26b5ea1fbcef0a5f8835a49e7e.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *ListRequestBuilder) Columns()(*i261ceb17d8b68c77c74467f613cdf5532cd5e5ea820c84d683a39d39c13c4ebe.ColumnsRequestBuilder) {
    return i261ceb17d8b68c77c74467f613cdf5532cd5e5ea820c84d683a39d39c13c4ebe.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*i40668099437af2cc778494c140ef105574b3d11b369844ae96e6da4e56ed7dbb.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i40668099437af2cc778494c140ef105574b3d11b369844ae96e6da4e56ed7dbb.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive%2Did}/list{?%24select,%24expand}";
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
func (m *ListRequestBuilder) ContentTypes()(*icf07d83821e028c339f3d32b2cc131377b2d116e69b7fd99e7db53d773436c38.ContentTypesRequestBuilder) {
    return icf07d83821e028c339f3d32b2cc131377b2d116e69b7fd99e7db53d773436c38.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*i9afbcf26a085bae33bec61cc4b40f70ee21e783407b414e87aa6707a9555ac2a.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return i9afbcf26a085bae33bec61cc4b40f70ee21e783407b414e87aa6707a9555ac2a.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for me
func (m *ListRequestBuilder) CreateDeleteRequestInformation(options *ListRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) CreateGetRequestInformation(options *ListRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property list in me
func (m *ListRequestBuilder) CreatePatchRequestInformation(options *ListRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property list for me
func (m *ListRequestBuilder) Delete(options *ListRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Drive the drive property
func (m *ListRequestBuilder) Drive()(*ib0080b16f4fd1f64bad0fc40faa8191f83bc5af05b27dabb5054539619462021.DriveRequestBuilder) {
    return ib0080b16f4fd1f64bad0fc40faa8191f83bc5af05b27dabb5054539619462021.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get for drives in SharePoint, the underlying document library list. Read-only. Nullable.
func (m *ListRequestBuilder) Get(options *ListRequestBuilderGetOptions)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable), nil
}
// Items the items property
func (m *ListRequestBuilder) Items()(*i072ddba6d43643cb4e0a492cfc6003149b0f11a1bbb94fcbbdcf71f0770d542a.ItemsRequestBuilder) {
    return i072ddba6d43643cb4e0a492cfc6003149b0f11a1bbb94fcbbdcf71f0770d542a.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*i8d0b9819c6c73c68982134065361fc5eb557175052e7f6f62cb6465cea9691ad.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem%2Did"] = id
    }
    return i8d0b9819c6c73c68982134065361fc5eb557175052e7f6f62cb6465cea9691ad.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ListRequestBuilder) Operations()(*ic1ed073c992d668b0fff838202aaefd614de5739048c54fa321f3c0cf2913074.OperationsRequestBuilder) {
    return ic1ed073c992d668b0fff838202aaefd614de5739048c54fa321f3c0cf2913074.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*i69531fc3cfd71461eb64582da70e123676f0660112f5d4521198ea126071edb4.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return i69531fc3cfd71461eb64582da70e123676f0660112f5d4521198ea126071edb4.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in me
func (m *ListRequestBuilder) Patch(options *ListRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Subscriptions the subscriptions property
func (m *ListRequestBuilder) Subscriptions()(*i1f28b50836f57b5d58bbaf6536fcecef381f989c4e9cb3ea42add732311ae102.SubscriptionsRequestBuilder) {
    return i1f28b50836f57b5d58bbaf6536fcecef381f989c4e9cb3ea42add732311ae102.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.me.drives.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*ibdce946586319f7ecb5c54b6cd60e9e1e6b2a7b464a4783217eb2d2c5fbc3414.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return ibdce946586319f7ecb5c54b6cd60e9e1e6b2a7b464a4783217eb2d2c5fbc3414.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
