package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder provides operations to manage the landingPage property of the microsoft.graph.simulation entity.
type AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderGetQueryParameters the landing page associated with a simulation during its creation.
type AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderGetQueryParameters
}
// NewAttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderInternal instantiates a new AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder and sets the default values.
func NewAttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder) {
    m := &AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/simulations/{simulation%2Did}/landingPage{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder instantiates a new AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder and sets the default values.
func NewAttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the landing page associated with a simulation during its creation.
// returns a LandingPageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LandingPageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateLandingPageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.LandingPageable), nil
}
// ToGetRequestInformation the landing page associated with a simulation during its creation.
// returns a *RequestInformation when successful
func (m *AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder when successful
func (m *AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder) {
    return NewAttacksimulationSimulationsItemLandingpageLandingPageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
