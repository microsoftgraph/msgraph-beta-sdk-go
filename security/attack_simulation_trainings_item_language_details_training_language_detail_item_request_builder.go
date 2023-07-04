package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be "github.com/microsoftgraph/msgraph-beta-sdk-go/models"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder provides operations to manage the languageDetails property of the microsoft.graph.training entity.
type AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderGetQueryParameters get languageDetails from security
type AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderGetQueryParameters
}
// NewAttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderInternal instantiates a new TrainingLanguageDetailItemRequestBuilder and sets the default values.
func NewAttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder) {
    m := &AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/trainings/{training%2Did}/languageDetails/{trainingLanguageDetail%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewAttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder instantiates a new TrainingLanguageDetailItemRequestBuilder and sets the default values.
func NewAttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get languageDetails from security
func (m *AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderGetRequestConfiguration)(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrainingLanguageDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.CreateTrainingLanguageDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ie233ee762e29b4ba6970aa2a2efce4b7fde11697ca9ea81099d0f8269309c1be.TrainingLanguageDetailable), nil
}
// ToGetRequestInformation get languageDetails from security
func (m *AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttackSimulationTrainingsItemLanguageDetailsTrainingLanguageDetailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
