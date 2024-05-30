package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder provides operations to call the getOffice365ActivationCounts method.
type Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGetoffice365activationcountsGetOffice365ActivationCountsRequestBuilderInternal instantiates a new Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder and sets the default values.
func NewGetoffice365activationcountsGetOffice365ActivationCountsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder) {
    m := &Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/getOffice365ActivationCounts()", pathParameters),
    }
    return m
}
// NewGetoffice365activationcountsGetOffice365ActivationCountsRequestBuilder instantiates a new Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder and sets the default values.
func NewGetoffice365activationcountsGetOffice365ActivationCountsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetoffice365activationcountsGetOffice365ActivationCountsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getOffice365ActivationCounts
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder) Get(ctx context.Context, requestConfiguration *Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation invoke function getOffice365ActivationCounts
// returns a *RequestInformation when successful
func (m *Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder when successful
func (m *Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder) WithUrl(rawUrl string)(*Getoffice365activationcountsGetOffice365ActivationCountsRequestBuilder) {
    return NewGetoffice365activationcountsGetOffice365ActivationCountsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
