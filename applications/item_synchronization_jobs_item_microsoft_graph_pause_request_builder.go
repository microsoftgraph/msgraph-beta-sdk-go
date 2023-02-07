package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder provides operations to call the pause method.
type ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilderInternal instantiates a new MicrosoftGraphPauseRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder) {
    m := &ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/synchronization/jobs/{synchronizationJob%2Did}/microsoft.graph.pause";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder instantiates a new MicrosoftGraphPauseRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post temporarily stop a running synchronization job. All the progress, including job state, is persisted, and the job will continue from where it left off when a start call is made.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/synchronization-synchronizationjob-pause?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation temporarily stop a running synchronization job. All the progress, including job state, is persisted, and the job will continue from where it left off when a start call is made.
func (m *ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSynchronizationJobsItemMicrosoftGraphPauseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
