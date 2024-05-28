package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder provides operations to call the getExpiringVppTokenCount method.
type DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal instantiates a new DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewDeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, expiringBeforeDateTime *string)(*DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    m := &DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/depOnboardingSettings/getExpiringVppTokenCount(expiringBeforeDateTime='{expiringBeforeDateTime}')", pathParameters),
    }
    if expiringBeforeDateTime != nil {
        m.BaseRequestBuilder.PathParameters["expiringBeforeDateTime"] = *expiringBeforeDateTime
    }
    return m
}
// NewDeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder instantiates a new DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder and sets the default values.
func NewDeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function getExpiringVppTokenCount
// Deprecated: This method is obsolete. Use GetAsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponse instead.
// returns a DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) Get(ctx context.Context, requestConfiguration *DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeResponseable), nil
}
// GetAsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponse invoke function getExpiringVppTokenCount
// returns a DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) GetAsGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponse(ctx context.Context, requestConfiguration *DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateDeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeGetResponseable), nil
}
// ToGetRequestInformation invoke function getExpiringVppTokenCount
// returns a *RequestInformation when successful
func (m *DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder when successful
func (m *DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) WithUrl(rawUrl string)(*DeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder) {
    return NewDeponboardingsettingsGetexpiringvpptokencountwithexpiringbeforedatetimeGetExpiringVppTokenCountWithExpiringBeforeDateTimeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
