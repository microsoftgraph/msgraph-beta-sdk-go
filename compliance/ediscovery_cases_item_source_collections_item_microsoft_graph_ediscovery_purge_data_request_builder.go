package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder provides operations to call the purgeData method.
type EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilderInternal instantiates a new MicrosoftGraphEdiscoveryPurgeDataRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) {
    m := &EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}/microsoft.graph.ediscovery.purgeData", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder instantiates a new MicrosoftGraphEdiscoveryPurgeDataRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post permanently delete Microsoft Teams messages contained in a sourceCollection. You can collect and purge the following categories of Teams content:- Teams 1:1 chats - Chat messages, posts, and attachments shared in a Teams conversation between two people. Teams 1:1 chats are also called *conversations*.- Teams group chats - Chat messages, posts, and attachments shared in a Teams conversation between three or more people. Also called *1:N* chats or *group conversations*.- Teams channels - Chat messages, posts, replies, and attachments shared in a standard Teams channel.- Private channels - Message posts, replies, and attachments shared in a private Teams channel.- Shared channels - Message posts, replies, and attachments shared in a shared Teams channel. For more information about purging Teams messages, see:- eDiscovery solution series: Data spillage scenario - Search and purge- Advanced eDiscovery workflow for content in Microsoft Teams  This API is available in the following national cloud deployments.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-sourcecollection-purgedata?view=graph-rest-1.0
func (m *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation permanently delete Microsoft Teams messages contained in a sourceCollection. You can collect and purge the following categories of Teams content:- Teams 1:1 chats - Chat messages, posts, and attachments shared in a Teams conversation between two people. Teams 1:1 chats are also called *conversations*.- Teams group chats - Chat messages, posts, and attachments shared in a Teams conversation between three or more people. Also called *1:N* chats or *group conversations*.- Teams channels - Chat messages, posts, replies, and attachments shared in a standard Teams channel.- Private channels - Message posts, replies, and attachments shared in a private Teams channel.- Shared channels - Message posts, replies, and attachments shared in a shared Teams channel. For more information about purging Teams messages, see:- eDiscovery solution series: Data spillage scenario - Search and purge- Advanced eDiscovery workflow for content in Microsoft Teams  This API is available in the following national cloud deployments.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace on 2022-12-05 and will be removed 2023-02-01
func (m *EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) {
    return NewEdiscoveryCasesItemSourceCollectionsItemMicrosoftGraphEdiscoveryPurgeDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
