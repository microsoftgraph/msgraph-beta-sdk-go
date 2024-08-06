package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder provides operations to call the getByName method.
type ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddToDrive provides operations to call the addToDrive method.
// returns a *ItemSitesItemContentModelsGetByNameWithModelNameAddToDriveRequestBuilder when successful
func (m *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) AddToDrive()(*ItemSitesItemContentModelsGetByNameWithModelNameAddToDriveRequestBuilder) {
    return NewItemSitesItemContentModelsGetByNameWithModelNameAddToDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemContentModelsGetByNameWithModelNameRequestBuilderInternal instantiates a new ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder and sets the default values.
func NewItemSitesItemContentModelsGetByNameWithModelNameRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, modelName *string)(*ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) {
    m := &ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentModels/getByName(modelName='{modelName}')", pathParameters),
    }
    if modelName != nil {
        m.BaseRequestBuilder.PathParameters["modelName"] = *modelName
    }
    return m
}
// NewItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder instantiates a new ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder and sets the default values.
func NewItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContentModelsGetByNameWithModelNameRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get read the properties and relationships of a contentModel object by its model name. The name should be the full model filename, including the file extension; for example, exampleModel.classifier.
// returns a ContentModelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contentmodel-getbyname?view=graph-rest-beta
func (m *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentModelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateContentModelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ContentModelable), nil
}
// GetAppliedDrives provides operations to call the getAppliedDrives method.
// returns a *ItemSitesItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder when successful
func (m *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) GetAppliedDrives()(*ItemSitesItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilder) {
    return NewItemSitesItemContentModelsGetByNameWithModelNameGetAppliedDrivesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveFromDrive provides operations to call the removeFromDrive method.
// returns a *ItemSitesItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder when successful
func (m *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) RemoveFromDrive()(*ItemSitesItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder) {
    return NewItemSitesItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read the properties and relationships of a contentModel object by its model name. The name should be the full model filename, including the file extension; for example, exampleModel.classifier.
// returns a *RequestInformation when successful
func (m *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder when successful
func (m *ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder) {
    return NewItemSitesItemContentModelsGetByNameWithModelNameRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
