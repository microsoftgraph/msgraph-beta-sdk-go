package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder provides operations to call the createMigrationReport method.
type GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewGrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilderInternal instantiates a new GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder) {
    m := &GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/groupPolicyMigrationReports/createMigrationReport", pathParameters),
    }
    return m
}
// NewGrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder instantiates a new GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder and sets the default values.
func NewGrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createMigrationReport
// Deprecated: This method is obsolete. Use PostAsCreateMigrationReportPostResponse instead.
// returns a GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder) Post(ctx context.Context, body GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportPostRequestBodyable, requestConfiguration *GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilderPostRequestConfiguration)(GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportResponseable), nil
}
// PostAsCreateMigrationReportPostResponse invoke action createMigrationReport
// returns a GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder) PostAsCreateMigrationReportPostResponse(ctx context.Context, body GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportPostRequestBodyable, requestConfiguration *GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilderPostRequestConfiguration)(GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateGrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportPostResponseable), nil
}
// ToPostRequestInformation invoke action createMigrationReport
// returns a *RequestInformation when successful
func (m *GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder) ToPostRequestInformation(ctx context.Context, body GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportPostRequestBodyable, requestConfiguration *GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder when successful
func (m *GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder) WithUrl(rawUrl string)(*GrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder) {
    return NewGrouppolicymigrationreportsCreatemigrationreportCreateMigrationReportRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
