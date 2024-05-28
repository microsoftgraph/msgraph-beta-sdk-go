package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder provides operations to call the purgeData method.
type EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilderInternal instantiates a new EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) {
    m := &EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}/microsoft.graph.ediscovery.purgeData", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder instantiates a new EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post permanently delete Microsoft Teams messages contained in a sourceCollection. You can collect and purge the following categories of Teams content:- Teams 1:1 chats - Chat messages, posts, and attachments shared in a Teams conversation between two people. Teams 1:1 chats are also called *conversations*.- Teams group chats - Chat messages, posts, and attachments shared in a Teams conversation between three or more people. Also called *1:N* chats or *group conversations*.- Teams channels - Chat messages, posts, replies, and attachments shared in a standard Teams channel.- Private channels - Message posts, replies, and attachments shared in a private Teams channel.- Shared channels - Message posts, replies, and attachments shared in a shared Teams channel. For more information about purging Teams messages, see:- eDiscovery solution series: Data spillage scenario - Search and purge- Advanced eDiscovery workflow for content in Microsoft Teams 
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/ediscovery-sourcecollection-purgedata?view=graph-rest-beta
func (m *EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation permanently delete Microsoft Teams messages contained in a sourceCollection. You can collect and purge the following categories of Teams content:- Teams 1:1 chats - Chat messages, posts, and attachments shared in a Teams conversation between two people. Teams 1:1 chats are also called *conversations*.- Teams group chats - Chat messages, posts, and attachments shared in a Teams conversation between three or more people. Also called *1:N* chats or *group conversations*.- Teams channels - Chat messages, posts, replies, and attachments shared in a standard Teams channel.- Private channels - Message posts, replies, and attachments shared in a private Teams channel.- Shared channels - Message posts, replies, and attachments shared in a shared Teams channel. For more information about purging Teams messages, see:- eDiscovery solution series: Data spillage scenario - Search and purge- Advanced eDiscovery workflow for content in Microsoft Teams 
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder when successful
func (m *EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder) {
    return NewEdiscoveryCasesItemSourcecollectionsItemMicrosoftgraphediscoverypurgedataMicrosoftGraphEdiscoveryPurgeDataRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
