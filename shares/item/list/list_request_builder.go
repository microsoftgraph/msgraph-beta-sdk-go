package list

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i4d901644c467d6da1064f9218d2a51302b7e04affba7010edb0f37beacafb773 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/subscriptions"
    i53109b7297c409f46286c656c3579c1bb0aff65ea692d18c6998e4eca7e7cae1 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items"
    ib3ee4819f4c735d8a63e343645326309cf0357281fcf420c964b4b39d6bdf9bb "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/drive"
    ic1183cba85b0a8992427e9c7d45ad198b411eb949fdcf8bf5fa204e3fc691a5f "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/columns"
    icf939f29457e323541cf147d31bf6cecf5cb130c9d7c8639ec566b8175d1d5b1 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes"
    iec4aa5d59e58579afdbecfba73d6bd2372d1b99f1d8b3ee465c20b9750d2cabd "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/operations"
    if45c6d4c98602b6c309b1994f425724fef00fd5a63d747a5169844dea93fc758 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/activities"
    i55bc9c6f863d9539e65f9d6af8a6bfd60d606e59691d0de7981bc7a5fd1dad6c "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/operations/item"
    ib48ab6e9d0b3b2e8132b00fed04b5ed61b4e70dfa330c315456e7b72365c013e "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/items/item"
    icc49cea450b5f3038e85b1a338cc21ce61681050bc96578d052ebb2a0cf235e4 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/contenttypes/item"
    id75517f929096ab464028a948416a44d3bdb482418ed4adcbaa37d952f45ac81 "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/subscriptions/item"
    ie6a239cd17154be4c0edbc83f7fe1c284635d98323d32d5dcce6210dfd425f2b "github.com/microsoftgraph/msgraph-beta-sdk-go/shares/item/list/columns/item"
)

// ListRequestBuilder provides operations to manage the list property of the microsoft.graph.sharedDriveItem entity.
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
// ListRequestBuilderGetQueryParameters used to access the underlying list
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
func (m *ListRequestBuilder) Activities()(*if45c6d4c98602b6c309b1994f425724fef00fd5a63d747a5169844dea93fc758.ActivitiesRequestBuilder) {
    return if45c6d4c98602b6c309b1994f425724fef00fd5a63d747a5169844dea93fc758.NewActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Columns the columns property
func (m *ListRequestBuilder) Columns()(*ic1183cba85b0a8992427e9c7d45ad198b411eb949fdcf8bf5fa204e3fc691a5f.ColumnsRequestBuilder) {
    return ic1183cba85b0a8992427e9c7d45ad198b411eb949fdcf8bf5fa204e3fc691a5f.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.columns.item collection
func (m *ListRequestBuilder) ColumnsById(id string)(*ie6a239cd17154be4c0edbc83f7fe1c284635d98323d32d5dcce6210dfd425f2b.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return ie6a239cd17154be4c0edbc83f7fe1c284635d98323d32d5dcce6210dfd425f2b.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewListRequestBuilderInternal instantiates a new ListRequestBuilder and sets the default values.
func NewListRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ListRequestBuilder) {
    m := &ListRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem%2Did}/list{?%24select,%24expand}";
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
func (m *ListRequestBuilder) ContentTypes()(*icf939f29457e323541cf147d31bf6cecf5cb130c9d7c8639ec566b8175d1d5b1.ContentTypesRequestBuilder) {
    return icf939f29457e323541cf147d31bf6cecf5cb130c9d7c8639ec566b8175d1d5b1.NewContentTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentTypesById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.contentTypes.item collection
func (m *ListRequestBuilder) ContentTypesById(id string)(*icc49cea450b5f3038e85b1a338cc21ce61681050bc96578d052ebb2a0cf235e4.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did"] = id
    }
    return icc49cea450b5f3038e85b1a338cc21ce61681050bc96578d052ebb2a0cf235e4.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property list for shares
func (m *ListRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property list for shares
func (m *ListRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ListRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation used to access the underlying list
func (m *ListRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration used to access the underlying list
func (m *ListRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ListRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property list in shares
func (m *ListRequestBuilder) CreatePatchRequestInformation(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property list in shares
func (m *ListRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property list for shares
func (m *ListRequestBuilder) Delete(ctx context.Context, requestConfiguration *ListRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *ListRequestBuilder) Drive()(*ib3ee4819f4c735d8a63e343645326309cf0357281fcf420c964b4b39d6bdf9bb.DriveRequestBuilder) {
    return ib3ee4819f4c735d8a63e343645326309cf0357281fcf420c964b4b39d6bdf9bb.NewDriveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get used to access the underlying list
func (m *ListRequestBuilder) Get(ctx context.Context, requestConfiguration *ListRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
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
func (m *ListRequestBuilder) Items()(*i53109b7297c409f46286c656c3579c1bb0aff65ea692d18c6998e4eca7e7cae1.ItemsRequestBuilder) {
    return i53109b7297c409f46286c656c3579c1bb0aff65ea692d18c6998e4eca7e7cae1.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.items.item collection
func (m *ListRequestBuilder) ItemsById(id string)(*ib48ab6e9d0b3b2e8132b00fed04b5ed61b4e70dfa330c315456e7b72365c013e.ListItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["listItem%2Did"] = id
    }
    return ib48ab6e9d0b3b2e8132b00fed04b5ed61b4e70dfa330c315456e7b72365c013e.NewListItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations the operations property
func (m *ListRequestBuilder) Operations()(*iec4aa5d59e58579afdbecfba73d6bd2372d1b99f1d8b3ee465c20b9750d2cabd.OperationsRequestBuilder) {
    return iec4aa5d59e58579afdbecfba73d6bd2372d1b99f1d8b3ee465c20b9750d2cabd.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.operations.item collection
func (m *ListRequestBuilder) OperationsById(id string)(*i55bc9c6f863d9539e65f9d6af8a6bfd60d606e59691d0de7981bc7a5fd1dad6c.RichLongRunningOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["richLongRunningOperation%2Did"] = id
    }
    return i55bc9c6f863d9539e65f9d6af8a6bfd60d606e59691d0de7981bc7a5fd1dad6c.NewRichLongRunningOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property list in shares
func (m *ListRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.Listable, requestConfiguration *ListRequestBuilderPatchRequestConfiguration)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
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
// Subscriptions the subscriptions property
func (m *ListRequestBuilder) Subscriptions()(*i4d901644c467d6da1064f9218d2a51302b7e04affba7010edb0f37beacafb773.SubscriptionsRequestBuilder) {
    return i4d901644c467d6da1064f9218d2a51302b7e04affba7010edb0f37beacafb773.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-beta-sdk-go/.shares.item.list.subscriptions.item collection
func (m *ListRequestBuilder) SubscriptionsById(id string)(*id75517f929096ab464028a948416a44d3bdb482418ed4adcbaa37d952f45ac81.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription%2Did"] = id
    }
    return id75517f929096ab464028a948416a44d3bdb482418ed4adcbaa37d952f45ac81.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
