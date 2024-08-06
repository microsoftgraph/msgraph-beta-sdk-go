package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemContentModelsItemGetAppliedDrivesRequestBuilder provides operations to call the getAppliedDrives method.
type ItemContentModelsItemGetAppliedDrivesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContentModelsItemGetAppliedDrivesRequestBuilderGetQueryParameters list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
type ItemContentModelsItemGetAppliedDrivesRequestBuilderGetQueryParameters struct {
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
// ItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContentModelsItemGetAppliedDrivesRequestBuilderGetQueryParameters
}
// NewItemContentModelsItemGetAppliedDrivesRequestBuilderInternal instantiates a new ItemContentModelsItemGetAppliedDrivesRequestBuilder and sets the default values.
func NewItemContentModelsItemGetAppliedDrivesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContentModelsItemGetAppliedDrivesRequestBuilder) {
    m := &ItemContentModelsItemGetAppliedDrivesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/contentModels/{contentModel%2Did}/getAppliedDrives(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemContentModelsItemGetAppliedDrivesRequestBuilder instantiates a new ItemContentModelsItemGetAppliedDrivesRequestBuilder and sets the default values.
func NewItemContentModelsItemGetAppliedDrivesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContentModelsItemGetAppliedDrivesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContentModelsItemGetAppliedDrivesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// Deprecated: This method is obsolete. Use GetAsGetAppliedDrivesGetResponse instead.
// returns a ItemContentModelsItemGetAppliedDrivesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contentmodel-getapplieddrives?view=graph-rest-beta
func (m *ItemContentModelsItemGetAppliedDrivesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration)(ItemContentModelsItemGetAppliedDrivesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContentModelsItemGetAppliedDrivesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContentModelsItemGetAppliedDrivesResponseable), nil
}
// GetAsGetAppliedDrivesGetResponse list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// returns a ItemContentModelsItemGetAppliedDrivesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contentmodel-getapplieddrives?view=graph-rest-beta
func (m *ItemContentModelsItemGetAppliedDrivesRequestBuilder) GetAsGetAppliedDrivesGetResponse(ctx context.Context, requestConfiguration *ItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration)(ItemContentModelsItemGetAppliedDrivesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContentModelsItemGetAppliedDrivesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContentModelsItemGetAppliedDrivesGetResponseable), nil
}
// ToGetRequestInformation list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// returns a *RequestInformation when successful
func (m *ItemContentModelsItemGetAppliedDrivesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContentModelsItemGetAppliedDrivesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemContentModelsItemGetAppliedDrivesRequestBuilder when successful
func (m *ItemContentModelsItemGetAppliedDrivesRequestBuilder) WithUrl(rawUrl string)(*ItemContentModelsItemGetAppliedDrivesRequestBuilder) {
    return NewItemContentModelsItemGetAppliedDrivesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
