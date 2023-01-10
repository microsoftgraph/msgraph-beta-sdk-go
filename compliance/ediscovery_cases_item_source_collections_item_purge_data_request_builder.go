package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder provides operations to call the purgeData method.
type EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilderInternal instantiates a new PurgeDataRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder) {
    m := &EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/sourceCollections/{sourceCollection%2Did}/microsoft.graph.ediscovery.purgeData";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder instantiates a new PurgeDataRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilderInternal(urlParams, requestAdapter)
}
// Post permanently delete Microsoft Teams messages contained in a sourceCollection. You can collect and purge the following categories of Teams content:- **Teams 1:1 chats** - Chat messages, posts, and attachments shared in a Teams conversation between two people. Teams 1:1 chats are also called *conversations*.- **Teams group chats** - Chat messages, posts, and attachments shared in a Teams conversation between three or more people. Also called *1:N* chats or *group conversations*.- **Teams channels** - Chat messages, posts, replies, and attachments shared in a standard Teams channel.- **Private channels** - Message posts, replies, and attachments shared in a private Teams channel.- **Shared channels** - Message posts, replies, and attachments shared in a shared Teams channel. For more information about purging Teams messages, see:- eDiscovery solution series: Data spillage scenario - Search and purge- Advanced eDiscovery workflow for content in Microsoft Teams 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/ediscovery-sourcecollection-purgedata?view=graph-rest-1.0
func (m *EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder) Post(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation permanently delete Microsoft Teams messages contained in a sourceCollection. You can collect and purge the following categories of Teams content:- **Teams 1:1 chats** - Chat messages, posts, and attachments shared in a Teams conversation between two people. Teams 1:1 chats are also called *conversations*.- **Teams group chats** - Chat messages, posts, and attachments shared in a Teams conversation between three or more people. Also called *1:N* chats or *group conversations*.- **Teams channels** - Chat messages, posts, replies, and attachments shared in a standard Teams channel.- **Private channels** - Message posts, replies, and attachments shared in a private Teams channel.- **Shared channels** - Message posts, replies, and attachments shared in a shared Teams channel. For more information about purging Teams messages, see:- eDiscovery solution series: Data spillage scenario - Search and purge- Advanced eDiscovery workflow for content in Microsoft Teams 
func (m *EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EdiscoveryCasesItemSourceCollectionsItemPurgeDataRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
