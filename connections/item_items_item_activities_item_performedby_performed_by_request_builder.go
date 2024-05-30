package connections

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder provides operations to manage the performedBy property of the microsoft.graph.externalConnectors.externalActivity entity.
type ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetQueryParameters represents an identity used to identify who is responsible for the activity.
type ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetQueryParameters
}
// NewItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderInternal instantiates a new ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder and sets the default values.
func NewItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) {
    m := &ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/connections/{externalConnection%2Did}/items/{externalItem%2Did}/activities/{externalActivity%2Did}/performedBy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder instantiates a new ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder and sets the default values.
func NewItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents an identity used to identify who is responsible for the activity.
// returns a Identityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, error) {
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
func (m *ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder when successful
func (m *ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder) {
    return NewItemItemsItemActivitiesItemPerformedbyPerformedByRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
