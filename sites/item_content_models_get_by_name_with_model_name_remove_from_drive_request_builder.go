package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder provides operations to call the removeFromDrive method.
type ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilderInternal instantiates a new ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder and sets the default values.
func NewItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder) {
    m := &ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/contentModels/getByName(modelName='{modelName}')/removeFromDrive", pathParameters),
    }
    return m
}
// NewItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder instantiates a new ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder and sets the default values.
func NewItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove a contentModel from a SharePoint document library.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contentmodel-removefromdrive?view=graph-rest-beta
func (m *ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder) Post(ctx context.Context, body ItemContentModelsGetByNameWithModelNameRemoveFromDrivePostRequestBodyable, requestConfiguration *ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation remove a contentModel from a SharePoint document library.
// returns a *RequestInformation when successful
func (m *ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemContentModelsGetByNameWithModelNameRemoveFromDrivePostRequestBodyable, requestConfiguration *ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder when successful
func (m *ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder) WithUrl(rawUrl string)(*ItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder) {
    return NewItemContentModelsGetByNameWithModelNameRemoveFromDriveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
