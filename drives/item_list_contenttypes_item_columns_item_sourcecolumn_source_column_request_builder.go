package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder provides operations to manage the sourceColumn property of the microsoft.graph.columnDefinition entity.
type ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters the source column for content type column.
type ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetQueryParameters
}
// NewItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderInternal instantiates a new ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder and sets the default values.
func NewItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    m := &ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/list/contentTypes/{contentType%2Did}/columns/{columnDefinition%2Did}/sourceColumn{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder instantiates a new ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder and sets the default values.
func NewItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the source column for content type column.
// returns a ColumnDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ColumnDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateColumnDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ColumnDefinitionable), nil
}
// ToGetRequestInformation the source column for content type column.
// returns a *RequestInformation when successful
func (m *ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder when successful
func (m *ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder) WithUrl(rawUrl string)(*ItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder) {
    return NewItemListContenttypesItemColumnsItemSourcecolumnSourceColumnRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
