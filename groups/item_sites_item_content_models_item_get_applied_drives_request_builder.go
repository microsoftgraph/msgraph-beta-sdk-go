package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder provides operations to call the getAppliedDrives method.
type ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderGetQueryParameters list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
type ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderGetQueryParameters
}
// NewItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderInternal instantiates a new ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder and sets the default values.
func NewItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder) {
    m := &ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentModels/{contentModel%2Did}/getAppliedDrives(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder instantiates a new ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder and sets the default values.
func NewItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// Deprecated: This method is obsolete. Use GetAsGetAppliedDrivesGetResponse instead.
// returns a ItemSitesItemContentModelsItemGetAppliedDrivesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contentmodel-getapplieddrives?view=graph-rest-beta
func (m *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration)(ItemSitesItemContentModelsItemGetAppliedDrivesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemContentModelsItemGetAppliedDrivesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemContentModelsItemGetAppliedDrivesResponseable), nil
}
// GetAsGetAppliedDrivesGetResponse list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// returns a ItemSitesItemContentModelsItemGetAppliedDrivesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contentmodel-getapplieddrives?view=graph-rest-beta
func (m *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder) GetAsGetAppliedDrivesGetResponse(ctx context.Context, requestConfiguration *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration)(ItemSitesItemContentModelsItemGetAppliedDrivesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSitesItemContentModelsItemGetAppliedDrivesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSitesItemContentModelsItemGetAppliedDrivesGetResponseable), nil
}
// ToGetRequestInformation list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// returns a *RequestInformation when successful
func (m *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder when successful
func (m *ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder) {
    return NewItemSitesItemContentModelsItemGetAppliedDrivesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
