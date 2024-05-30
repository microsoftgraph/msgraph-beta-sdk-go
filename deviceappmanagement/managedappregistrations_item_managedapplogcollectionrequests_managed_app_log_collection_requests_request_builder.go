package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder provides operations to manage the managedAppLogCollectionRequests property of the microsoft.graph.managedAppRegistration entity.
type ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetQueryParameters zero or more log collection requests triggered for the app.
type ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetQueryParameters
}
// ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedAppLogCollectionRequestId provides operations to manage the managedAppLogCollectionRequests property of the microsoft.graph.managedAppRegistration entity.
// returns a *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestItemRequestBuilder when successful
func (m *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) ByManagedAppLogCollectionRequestId(managedAppLogCollectionRequestId string)(*ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedAppLogCollectionRequestId != "" {
        urlTplParams["managedAppLogCollectionRequest%2Did"] = managedAppLogCollectionRequestId
    }
    return NewManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderInternal instantiates a new ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder and sets the default values.
func NewManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) {
    m := &ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/managedAppRegistrations/{managedAppRegistration%2Did}/managedAppLogCollectionRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder instantiates a new ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder and sets the default values.
func NewManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ManagedappregistrationsItemManagedapplogcollectionrequestsCountRequestBuilder when successful
func (m *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) Count()(*ManagedappregistrationsItemManagedapplogcollectionrequestsCountRequestBuilder) {
    return NewManagedappregistrationsItemManagedapplogcollectionrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get zero or more log collection requests triggered for the app.
// returns a ManagedAppLogCollectionRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAppLogCollectionRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestCollectionResponseable), nil
}
// Post create new navigation property to managedAppLogCollectionRequests for deviceAppManagement
// returns a ManagedAppLogCollectionRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestable, requestConfiguration *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateManagedAppLogCollectionRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestable), nil
}
// ToGetRequestInformation zero or more log collection requests triggered for the app.
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to managedAppLogCollectionRequests for deviceAppManagement
// returns a *RequestInformation when successful
func (m *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.ManagedAppLogCollectionRequestable, requestConfiguration *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder when successful
func (m *ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) WithUrl(rawUrl string)(*ManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder) {
    return NewManagedappregistrationsItemManagedapplogcollectionrequestsManagedAppLogCollectionRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
