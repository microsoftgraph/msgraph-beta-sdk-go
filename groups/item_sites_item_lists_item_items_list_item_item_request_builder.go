package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemListsItemItemsListItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.list entity.
type ItemSitesItemListsItemItemsListItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSitesItemListsItemItemsListItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsListItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemListsItemItemsListItemItemRequestBuilderGetQueryParameters all items contained in the list.
type ItemSitesItemListsItemItemsListItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemListsItemItemsListItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsListItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemListsItemItemsListItemItemRequestBuilderGetQueryParameters
}
// ItemSitesItemListsItemItemsListItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemListsItemItemsListItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) Activities()(*ItemSitesItemListsItemItemsItemActivitiesRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemActivitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// ActivitiesById provides operations to manage the activities property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) ActivitiesById(id string)(*ItemSitesItemListsItemItemsItemActivitiesItemActivityOLDItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemListsItemItemsItemActivitiesItemActivityOLDItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) Analytics()(*ItemSitesItemListsItemItemsItemAnalyticsRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// NewItemSitesItemListsItemItemsListItemItemRequestBuilderInternal instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, listItemId *string)(*ItemSitesItemListsItemItemsListItemItemRequestBuilder) {
    m := &ItemSitesItemListsItemItemsListItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if listItemId != nil {
        urlTplParams["listItem%2Did"] = *listItemId
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSitesItemListsItemItemsListItemItemRequestBuilder instantiates a new ListItemItemRequestBuilder and sets the default values.
func NewItemSitesItemListsItemItemsListItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemListsItemItemsListItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemListsItemItemsListItemItemRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property items for groups
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsListItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DocumentSetVersions provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) DocumentSetVersions()(*ItemSitesItemListsItemItemsItemDocumentSetVersionsRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemDocumentSetVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// DocumentSetVersionsById provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) DocumentSetVersionsById(id string)(*ItemSitesItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemListsItemItemsItemDocumentSetVersionsDocumentSetVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
// DriveItem provides operations to manage the driveItem property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) DriveItem()(*ItemSitesItemListsItemItemsItemDriveItemRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Fields provides operations to manage the fields property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) Fields()(*ItemSitesItemListsItemItemsItemFieldsRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemFieldsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// Get all items contained in the list.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsListItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// MicrosoftGraphCreateLink provides operations to call the createLink method.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) MicrosoftGraphCreateLink()(*ItemSitesItemListsItemItemsItemMicrosoftGraphCreateLinkCreateLinkRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemMicrosoftGraphCreateLinkCreateLinkRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// MicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) MicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemSitesItemListsItemItemsItemMicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemMicrosoftGraphGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.pathParameters, m.requestAdapter, endDateTime, interval, startDateTime)
}
// Patch update the navigation property items in groups
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ItemSitesItemListsItemItemsListItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// ToDeleteRequestInformation delete navigation property items for groups
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsListItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToGetRequestInformation all items contained in the list.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemListsItemItemsListItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property items in groups
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ItemSitesItemListsItemItemsListItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Versions provides operations to manage the versions property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) Versions()(*ItemSitesItemListsItemItemsItemVersionsRequestBuilder) {
    return NewItemSitesItemListsItemItemsItemVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter)
}
// VersionsById provides operations to manage the versions property of the microsoft.graph.listItem entity.
func (m *ItemSitesItemListsItemItemsListItemItemRequestBuilder) VersionsById(id string)(*ItemSitesItemListsItemItemsItemVersionsListItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    idPtr := &id
    return NewItemSitesItemListsItemItemsItemVersionsListItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter, idPtr)
}
