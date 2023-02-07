package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder provides operations to call the discover method.
type ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderInternal instantiates a new MicrosoftGraphDiscoverRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) {
    m := &ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/jobs/{synchronizationJob%2Did}/schema/directories/{directoryDefinition%2Did}/microsoft.graph.discover";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder instantiates a new MicrosoftGraphDiscoverRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderInternal(urlParams, requestAdapter)
}
// Post discover the latest schema definition for provisioning to an application. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/directorydefinition-discover?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryDefinitionable), nil
}
// ToPostRequestInformation discover the latest schema definition for provisioning to an application. 
func (m *ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemSchemaDirectoriesItemMicrosoftGraphDiscoverRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
