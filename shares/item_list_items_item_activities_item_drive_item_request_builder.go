package shares

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListItemsItemActivitiesItemDriveItemRequestBuilder provides operations to manage the driveItem property of the microsoft.graph.itemActivityOLD entity.
type ItemListItemsItemActivitiesItemDriveItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListItemsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters get driveItem from shares
type ItemListItemsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListItemsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListItemsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListItemsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters
}
// NewItemListItemsItemActivitiesItemDriveItemRequestBuilderInternal instantiates a new ItemListItemsItemActivitiesItemDriveItemRequestBuilder and sets the default values.
func NewItemListItemsItemActivitiesItemDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListItemsItemActivitiesItemDriveItemRequestBuilder) {
    m := &ItemListItemsItemActivitiesItemDriveItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/shares/{sharedDriveItem%2Did}/list/items/{listItem%2Did}/activities/{itemActivityOLD%2Did}/driveItem{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemListItemsItemActivitiesItemDriveItemRequestBuilder instantiates a new ItemListItemsItemActivitiesItemDriveItemRequestBuilder and sets the default values.
func NewItemListItemsItemActivitiesItemDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListItemsItemActivitiesItemDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListItemsItemActivitiesItemDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the sharedDriveItem entity.
// returns a *ItemListItemsItemActivitiesItemDriveItemContentRequestBuilder when successful
func (m *ItemListItemsItemActivitiesItemDriveItemRequestBuilder) Content()(*ItemListItemsItemActivitiesItemDriveItemContentRequestBuilder) {
    return NewItemListItemsItemActivitiesItemDriveItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get driveItem from shares
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListItemsItemActivitiesItemDriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListItemsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DriveItemable), nil
}
// ToGetRequestInformation get driveItem from shares
// returns a *RequestInformation when successful
func (m *ItemListItemsItemActivitiesItemDriveItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListItemsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemListItemsItemActivitiesItemDriveItemRequestBuilder when successful
func (m *ItemListItemsItemActivitiesItemDriveItemRequestBuilder) WithUrl(rawUrl string)(*ItemListItemsItemActivitiesItemDriveItemRequestBuilder) {
    return NewItemListItemsItemActivitiesItemDriveItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
