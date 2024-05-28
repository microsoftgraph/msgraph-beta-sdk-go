package compliance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder provides operations to call the removeHold method.
type EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal instantiates a new EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    m := &EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/compliance/ediscovery/cases/{case%2Did}/custodians/microsoft.graph.ediscovery.removeHold", pathParameters),
    }
    return m
}
// NewEdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder instantiates a new EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder and sets the default values.
func NewEdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action removeHold
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) Post(ctx context.Context, body EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdRemoveHoldPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation invoke action removeHold
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *RequestInformation when successful
func (m *EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) ToPostRequestInformation(ctx context.Context, body EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdRemoveHoldPostRequestBodyable, requestConfiguration *EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Deprecated: The ediscovery Apis are deprecated under /compliance and will stop returning data from February 01, 2023. Please use the new ediscovery Apis under /security. as of 2022-12/ediscoveryNamespace
// returns a *EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder when successful
func (m *EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) WithUrl(rawUrl string)(*EdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder) {
    return NewEdiscoveryCasesItemCustodiansMicrosoftgraphediscoveryremoveholdMicrosoftGraphEdiscoveryRemoveHoldRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
