package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/externalconnectors"
)

// ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder provides operations to manage the performedBy property of the microsoft.graph.externalConnectors.externalActivity entity.
type ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderGetQueryParameters represents an identity used to identify who is responsible for the activity.
type ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderGetQueryParameters
}
// NewExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderInternal instantiates a new PerformedByRequestBuilder and sets the default values.
func NewExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder) {
    m := &ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/external/connections/{externalConnection%2Did}/items/{externalItem%2Did}/activities/{externalActivity%2Did}/performedBy{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder instantiates a new PerformedByRequestBuilder and sets the default values.
func NewExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation represents an identity used to identify who is responsible for the activity.
func (m *ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get represents an identity used to identify who is responsible for the activity.
func (m *ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilder) Get(ctx context.Context, requestConfiguration *ExternalConnectionsItemItemsItemActivitiesItemPerformedByRequestBuilderGetRequestConfiguration)(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.CreateIdentityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie98116770ca9f5eee835504331ccb9976e822c2f776cca356ee95c843b4cce86.Identityable), nil
}
