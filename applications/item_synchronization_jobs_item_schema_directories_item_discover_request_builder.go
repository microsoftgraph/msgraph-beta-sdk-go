package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder provides operations to call the discover method.
type ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilderInternal instantiates a new ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder) {
    m := &ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/schema/directories/{directoryDefinition%2Did}/discover", pathParameters),
    }
    return m
}
// NewItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder instantiates a new ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilderInternal(urlParams, requestAdapter)
}
// Post discover the latest schema definition for provisioning to an application. 
// returns a DirectoryDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/synchronization-directorydefinition-discover?view=graph-rest-beta
func (m *ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryDefinitionable), nil
}
// ToPostRequestInformation discover the latest schema definition for provisioning to an application. 
// returns a *RequestInformation when successful
func (m *ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder when successful
func (m *ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder) WithUrl(rawUrl string)(*ItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder) {
    return NewItemSynchronizationJobsItemSchemaDirectoriesItemDiscoverRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
