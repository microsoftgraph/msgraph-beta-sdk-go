package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListsItemItemsListItemItemRequestBuilder provides operations to manage the items property of the microsoft.graph.list entity.
type ItemListsItemItemsListItemItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemItemsListItemItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsListItemItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemListsItemItemsListItemItemRequestBuilderGetQueryParameters returns the metadata for an item in a list.
type ItemListsItemItemsListItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListsItemItemsListItemItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsListItemItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemItemsListItemItemRequestBuilderGetQueryParameters
}
// ItemListsItemItemsListItemItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemItemsListItemItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Activities provides operations to manage the activities property of the microsoft.graph.listItem entity.
// returns a *ItemListsItemItemsItemActivitiesRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) Activities()(*ItemListsItemItemsItemActivitiesRequestBuilder) {
    return NewItemListsItemItemsItemActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.listItem entity.
// returns a *ItemListsItemItemsItemAnalyticsRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) Analytics()(*ItemListsItemItemsItemAnalyticsRequestBuilder) {
    return NewItemListsItemItemsItemAnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemListsItemItemsListItemItemRequestBuilderInternal instantiates a new ItemListsItemItemsListItemItemRequestBuilder and sets the default values.
func NewItemListsItemItemsListItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsListItemItemRequestBuilder) {
    m := &ItemListsItemItemsListItemItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/items/{listItem%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemListsItemItemsListItemItemRequestBuilder instantiates a new ItemListsItemItemsListItemItemRequestBuilder and sets the default values.
func NewItemListsItemItemsListItemItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemItemsListItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemItemsListItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemListsItemItemsItemCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) CreatedByUser()(*ItemListsItemItemsItemCreatedbyuserCreatedByUserRequestBuilder) {
    return NewItemListsItemItemsItemCreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreateLink provides operations to call the createLink method.
// returns a *ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) CreateLink()(*ItemListsItemItemsItemCreatelinkCreateLinkRequestBuilder) {
    return NewItemListsItemItemsItemCreatelinkCreateLinkRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete removes an item from a list.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/listitem-delete?view=graph-rest-beta
func (m *ItemListsItemItemsListItemItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemListsItemItemsListItemItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DocumentSetVersions provides operations to manage the documentSetVersions property of the microsoft.graph.listItem entity.
// returns a *ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) DocumentSetVersions()(*ItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilder) {
    return NewItemListsItemItemsItemDocumentsetversionsDocumentSetVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DriveItem provides operations to manage the driveItem property of the microsoft.graph.listItem entity.
// returns a *ItemListsItemItemsItemDriveitemDriveItemRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) DriveItem()(*ItemListsItemItemsItemDriveitemDriveItemRequestBuilder) {
    return NewItemListsItemItemsItemDriveitemDriveItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Fields provides operations to manage the fields property of the microsoft.graph.listItem entity.
// returns a *ItemListsItemItemsItemFieldsRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) Fields()(*ItemListsItemItemsItemFieldsRequestBuilder) {
    return NewItemListsItemItemsItemFieldsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get returns the metadata for an item in a list.
// returns a ListItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/listitem-get?view=graph-rest-beta
func (m *ItemListsItemItemsListItemItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemItemsListItemItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval provides operations to call the getActivitiesByInterval method.
// returns a *ItemListsItemItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) GetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithInterval(endDateTime *string, interval *string, startDateTime *string)(*ItemListsItemItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewItemListsItemItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, endDateTime, interval, startDateTime)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemListsItemItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) LastModifiedByUser()(*ItemListsItemItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewItemListsItemItemsItemLastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property items in sites
// returns a ListItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemItemsListItemItemRequestBuilder) Patch(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ItemListsItemItemsListItemItemRequestBuilderPatchRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateListItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable), nil
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.listItem entity.
// returns a *ItemListsItemItemsItemPermissionsRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) Permissions()(*ItemListsItemItemsItemPermissionsRequestBuilder) {
    return NewItemListsItemItemsItemPermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation removes an item from a list.
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsListItemItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation returns the metadata for an item in a list.
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemItemsListItemItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property items in sites
// returns a *RequestInformation when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ListItemable, requestConfiguration *ItemListsItemItemsListItemItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// Versions provides operations to manage the versions property of the microsoft.graph.listItem entity.
// returns a *ItemListsItemItemsItemVersionsRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) Versions()(*ItemListsItemItemsItemVersionsRequestBuilder) {
    return NewItemListsItemItemsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemListsItemItemsListItemItemRequestBuilder when successful
func (m *ItemListsItemItemsListItemItemRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemItemsListItemItemRequestBuilder) {
    return NewItemListsItemItemsListItemItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
