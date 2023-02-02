package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder provides operations to call the provisionOnDemand method.
type ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilderInternal instantiates a new ProvisionOnDemandRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder) {
    m := &ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/synchronization/jobs/{synchronizationJob%2Did}/microsoft.graph.provisionOnDemand";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams
    m.requestAdapter = requestAdapter
    return m
}
// NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder instantiates a new ProvisionOnDemandRequestBuilder and sets the default values.
func NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilderInternal(urlParams, requestAdapter)
}
// Post select a user and provision the account on-demand. The rate limit for this API is 5 requests per 10 seconds. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/synchronization-synchronizationjob-provision-on-demand?view=graph-rest-1.0
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder) Post(ctx context.Context, body ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.StringKeyStringValuePairable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateStringKeyStringValuePairFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.StringKeyStringValuePairable), nil
}
// ToPostRequestInformation select a user and provision the account on-demand. The rate limit for this API is 5 requests per 10 seconds. 
func (m *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandPostRequestBodyable, requestConfiguration *ItemSynchronizationJobsItemMicrosoftGraphProvisionOnDemandProvisionOnDemandRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
