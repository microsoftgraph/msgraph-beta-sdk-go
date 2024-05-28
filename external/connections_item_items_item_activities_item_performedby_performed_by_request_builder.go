package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder provides operations to manage the performedBy property of the microsoft.graph.externalConnectors.externalActivity entity.
type ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetQueryParameters represents an identity used to identify who is responsible for the activity.
type ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetQueryParameters
}
// NewConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderInternal instantiates a new ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder and sets the default values.
func NewConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) {
    m := &ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/connections/{externalConnection%2Did}/items/{externalItem%2Did}/activities/{externalActivity%2Did}/performedBy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder instantiates a new ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder and sets the default values.
func NewConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents an identity used to identify who is responsible for the activity.
// returns a Identityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) Get(ctx context.Context, requestConfiguration *ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable), nil
}
// ToGetRequestInformation represents an identity used to identify who is responsible for the activity.
// returns a *RequestInformation when successful
func (m *ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder when successful
func (m *ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) WithUrl(rawUrl string)(*ConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) {
    return NewConnectionsItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
