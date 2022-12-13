package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/security"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TriggersRetentionEventsItemRetentionEventTypeRequestBuilder provides operations to manage the retentionEventType property of the microsoft.graph.security.retentionEvent entity.
type TriggersRetentionEventsItemRetentionEventTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetQueryParameters specifies the event that will start the retention period for labels that use this event type when an event is created.
type TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetQueryParameters
}
// NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilderInternal instantiates a new RetentionEventTypeRequestBuilder and sets the default values.
func NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggersRetentionEventsItemRetentionEventTypeRequestBuilder) {
    m := &TriggersRetentionEventsItemRetentionEventTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/security/triggers/retentionEvents/{retentionEvent%2Did}/retentionEventType{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilder instantiates a new RetentionEventTypeRequestBuilder and sets the default values.
func NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TriggersRetentionEventsItemRetentionEventTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTriggersRetentionEventsItemRetentionEventTypeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *TriggersRetentionEventsItemRetentionEventTypeRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get specifies the event that will start the retention period for labels that use this event type when an event is created.
func (m *TriggersRetentionEventsItemRetentionEventTypeRequestBuilder) Get(ctx context.Context, requestConfiguration *TriggersRetentionEventsItemRetentionEventTypeRequestBuilderGetRequestConfiguration)(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.CreateRetentionEventTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(i084fa7ab3bba802bf5cc3b408e230cc64c167a57976e0d42c37e17154afd5b78.RetentionEventTypeable), nil
}
