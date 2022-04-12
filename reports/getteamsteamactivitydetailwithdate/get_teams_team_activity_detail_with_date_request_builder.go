package getteamsteamactivitydetailwithdate

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetTeamsTeamActivityDetailWithDateRequestBuilder provides operations to call the getTeamsTeamActivityDetail method.
type GetTeamsTeamActivityDetailWithDateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// GetTeamsTeamActivityDetailWithDateRequestBuilderGetOptions options for Get
type GetTeamsTeamActivityDetailWithDateRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal instantiates a new GetTeamsTeamActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, date *i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.DateOnly)(*GetTeamsTeamActivityDetailWithDateRequestBuilder) {
    m := &GetTeamsTeamActivityDetailWithDateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsTeamActivityDetail(date={date})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if date != nil {
        urlTplParams[""] = (*date).String()
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetTeamsTeamActivityDetailWithDateRequestBuilder instantiates a new GetTeamsTeamActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetTeamsTeamActivityDetailWithDateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GetTeamsTeamActivityDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getTeamsTeamActivityDetail
func (m *GetTeamsTeamActivityDetailWithDateRequestBuilder) CreateGetRequestInformation(options *GetTeamsTeamActivityDetailWithDateRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function getTeamsTeamActivityDetail
func (m *GetTeamsTeamActivityDetailWithDateRequestBuilder) Get(options *GetTeamsTeamActivityDetailWithDateRequestBuilderGetOptions)(GetTeamsTeamActivityDetailWithDateResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateGetTeamsTeamActivityDetailWithDateResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(GetTeamsTeamActivityDetailWithDateResponseable), nil
}
