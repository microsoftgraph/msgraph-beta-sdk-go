package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder provides operations to manage the details property of the microsoft.graph.landingPage entity.
type AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderGetQueryParameters get details from security
type AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderGetQueryParameters
}
// NewAttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderInternal instantiates a new LandingPageDetailItemRequestBuilder and sets the default values.
func NewAttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder) {
    m := &AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/landingPages/{landingPage%2Did}/details/{landingPageDetail%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder instantiates a new LandingPageDetailItemRequestBuilder and sets the default values.
func NewAttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get details from security
func (m *AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LandingPageDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLandingPageDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LandingPageDetailable), nil
}
// ToGetRequestInformation get details from security
func (m *AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttackSimulationLandingPagesItemDetailsLandingPageDetailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
