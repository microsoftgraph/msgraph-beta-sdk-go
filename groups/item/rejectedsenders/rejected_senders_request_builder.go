package rejectedsenders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    i00437f87924367e7c57f188661cf8fb5c73bd7f445ee688226aee02471eb14b5 "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/rejectedsenders/ref"
    i4cd946ba6fa87ef9076b08165c8d362f1f54b4ae7f7949a0783809a0f18bbbae "github.com/microsoftgraph/msgraph-beta-sdk-go/groups/item/rejectedsenders/count"
)

// RejectedSendersRequestBuilder provides operations to manage the rejectedSenders property of the microsoft.graph.group entity.
type RejectedSendersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RejectedSendersRequestBuilderGetQueryParameters the list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
type RejectedSendersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// RejectedSendersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RejectedSendersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RejectedSendersRequestBuilderGetQueryParameters
}
// NewRejectedSendersRequestBuilderInternal instantiates a new RejectedSendersRequestBuilder and sets the default values.
func NewRejectedSendersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RejectedSendersRequestBuilder) {
    m := &RejectedSendersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/rejectedSenders{?%24top,%24skip,%24filter,%24count,%24orderby,%24select}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRejectedSendersRequestBuilder instantiates a new RejectedSendersRequestBuilder and sets the default values.
func NewRejectedSendersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RejectedSendersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRejectedSendersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count the Count property
func (m *RejectedSendersRequestBuilder) Count()(*i4cd946ba6fa87ef9076b08165c8d362f1f54b4ae7f7949a0783809a0f18bbbae.CountRequestBuilder) {
    return i4cd946ba6fa87ef9076b08165c8d362f1f54b4ae7f7949a0783809a0f18bbbae.NewCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
func (m *RejectedSendersRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
func (m *RejectedSendersRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *RejectedSendersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get the list of users or groups that are not allowed to create posts or calendar events in this group. Nullable
func (m *RejectedSendersRequestBuilder) Get(ctx context.Context, requestConfiguration *RejectedSendersRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.DirectoryObjectCollectionResponseable), nil
}
// Ref the Ref property
func (m *RejectedSendersRequestBuilder) Ref()(*i00437f87924367e7c57f188661cf8fb5c73bd7f445ee688226aee02471eb14b5.RefRequestBuilder) {
    return i00437f87924367e7c57f188661cf8fb5c73bd7f445ee688226aee02471eb14b5.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
