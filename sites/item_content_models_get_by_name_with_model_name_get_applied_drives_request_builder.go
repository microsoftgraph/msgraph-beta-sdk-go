package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder provides operations to call the getAppliedDrives method.
type ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderGetQueryParameters list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
type ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderGetQueryParameters struct {
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
// ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderGetQueryParameters
}
// NewItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderInternal instantiates a new ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder and sets the default values.
func NewItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder) {
    m := &ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/contentModels/getByName(modelName='{modelName}')/getAppliedDrives(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder instantiates a new ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder and sets the default values.
func NewItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// Deprecated: This method is obsolete. Use GetAsGetAppliedDrivesGetResponse instead.
// returns a ItemContentModelsGetByNameWithModelNameGetAppliedDrivesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contentmodel-getapplieddrives?view=graph-rest-beta
func (m *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderGetRequestConfiguration)(ItemContentModelsGetByNameWithModelNameGetAppliedDrivesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContentModelsGetByNameWithModelNameGetAppliedDrivesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContentModelsGetByNameWithModelNameGetAppliedDrivesResponseable), nil
}
// GetAsGetAppliedDrivesGetResponse list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// returns a ItemContentModelsGetByNameWithModelNameGetAppliedDrivesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contentmodel-getapplieddrives?view=graph-rest-beta
func (m *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder) GetAsGetAppliedDrivesGetResponse(ctx context.Context, requestConfiguration *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderGetRequestConfiguration)(ItemContentModelsGetByNameWithModelNameGetAppliedDrivesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemContentModelsGetByNameWithModelNameGetAppliedDrivesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemContentModelsGetByNameWithModelNameGetAppliedDrivesGetResponseable), nil
}
// ToGetRequestInformation list all the contentModelUsage information related to a contentModel applied to a SharePoint document library.
// returns a *RequestInformation when successful
func (m *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder when successful
func (m *ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder) WithUrl(rawUrl string)(*ItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder) {
    return NewItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
