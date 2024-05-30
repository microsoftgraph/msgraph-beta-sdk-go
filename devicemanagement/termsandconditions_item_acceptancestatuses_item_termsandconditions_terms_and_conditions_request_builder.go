package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder provides operations to manage the termsAndConditions property of the microsoft.graph.termsAndConditionsAcceptanceStatus entity.
type TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderGetQueryParameters navigation link to the terms and conditions that are assigned.
type TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderGetQueryParameters
}
// NewTermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderInternal instantiates a new TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder and sets the default values.
func NewTermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder) {
    m := &TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions%2Did}/acceptanceStatuses/{termsAndConditionsAcceptanceStatus%2Did}/termsAndConditions{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder instantiates a new TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder and sets the default values.
func NewTermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get navigation link to the terms and conditions that are assigned.
// returns a TermsAndConditionsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTermsAndConditionsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TermsAndConditionsable), nil
}
// ToGetRequestInformation navigation link to the terms and conditions that are assigned.
// returns a *RequestInformation when successful
func (m *TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder when successful
func (m *TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder) WithUrl(rawUrl string)(*TermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder) {
    return NewTermsandconditionsItemAcceptancestatusesItemTermsandconditionsTermsAndConditionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
