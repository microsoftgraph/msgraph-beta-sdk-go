package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
type CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderGetQueryParameters get a list of credentialUserRegistrationDetails objects for a given tenant.
type CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderGetQueryParameters struct {
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
// CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderGetQueryParameters
}
// CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCredentialUserRegistrationDetailsId provides operations to manage the credentialUserRegistrationDetails property of the microsoft.graph.reportRoot entity.
// Deprecated: The Reporting credentialUserRegistrationDetails API is deprecated and will stop returning data on June 30, 2024. Please use the new userRegistrationDetails API. as of 2023-06/credentialUserRegistrationDetails
// returns a *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsItemRequestBuilder when successful
func (m *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) ByCredentialUserRegistrationDetailsId(credentialUserRegistrationDetailsId string)(*CredentialuserregistrationdetailsCredentialUserRegistrationDetailsItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if credentialUserRegistrationDetailsId != "" {
        urlTplParams["credentialUserRegistrationDetails%2Did"] = credentialUserRegistrationDetailsId
    }
    return NewCredentialuserregistrationdetailsCredentialUserRegistrationDetailsItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderInternal instantiates a new CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder and sets the default values.
func NewCredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) {
    m := &CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/credentialUserRegistrationDetails{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder instantiates a new CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder and sets the default values.
func NewCredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CredentialuserregistrationdetailsCountRequestBuilder when successful
func (m *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) Count()(*CredentialuserregistrationdetailsCountRequestBuilder) {
    return NewCredentialuserregistrationdetailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of credentialUserRegistrationDetails objects for a given tenant.
// Deprecated: The Reporting credentialUserRegistrationDetails API is deprecated and will stop returning data on June 30, 2024. Please use the new userRegistrationDetails API. as of 2023-06/credentialUserRegistrationDetails
// returns a CredentialUserRegistrationDetailsCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/reportroot-list-credentialuserregistrationdetails?view=graph-rest-beta
func (m *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CredentialUserRegistrationDetailsCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCredentialUserRegistrationDetailsCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CredentialUserRegistrationDetailsCollectionResponseable), nil
}
// Post create new navigation property to credentialUserRegistrationDetails for reports
// Deprecated: The Reporting credentialUserRegistrationDetails API is deprecated and will stop returning data on June 30, 2024. Please use the new userRegistrationDetails API. as of 2023-06/credentialUserRegistrationDetails
// returns a CredentialUserRegistrationDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) Post(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CredentialUserRegistrationDetailsable, requestConfiguration *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderPostRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CredentialUserRegistrationDetailsable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateCredentialUserRegistrationDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CredentialUserRegistrationDetailsable), nil
}
// ToGetRequestInformation get a list of credentialUserRegistrationDetails objects for a given tenant.
// Deprecated: The Reporting credentialUserRegistrationDetails API is deprecated and will stop returning data on June 30, 2024. Please use the new userRegistrationDetails API. as of 2023-06/credentialUserRegistrationDetails
// returns a *RequestInformation when successful
func (m *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to credentialUserRegistrationDetails for reports
// Deprecated: The Reporting credentialUserRegistrationDetails API is deprecated and will stop returning data on June 30, 2024. Please use the new userRegistrationDetails API. as of 2023-06/credentialUserRegistrationDetails
// returns a *RequestInformation when successful
func (m *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CredentialUserRegistrationDetailsable, requestConfiguration *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The Reporting credentialUserRegistrationDetails API is deprecated and will stop returning data on June 30, 2024. Please use the new userRegistrationDetails API. as of 2023-06/credentialUserRegistrationDetails
// returns a *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder when successful
func (m *CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) WithUrl(rawUrl string)(*CredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder) {
    return NewCredentialuserregistrationdetailsCredentialUserRegistrationDetailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
