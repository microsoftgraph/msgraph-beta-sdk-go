package getteamsteamactivitydetailwithdate

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetTeamsTeamActivityDetailWithDateRequestBuilder provides operations to call the getTeamsTeamActivityDetail method.
type GetTeamsTeamActivityDetailWithDateRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// GetTeamsTeamActivityDetailWithDateRequestBuilderGetOptions options for Get
type GetTeamsTeamActivityDetailWithDateRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal instantiates a new GetTeamsTeamActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, date *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.DateOnly)(*GetTeamsTeamActivityDetailWithDateRequestBuilder) {
    m := &GetTeamsTeamActivityDetailWithDateRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsTeamActivityDetail(date={date})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if date != nil {
        urlTplParams["date"] = (*date).String()
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGetTeamsTeamActivityDetailWithDateRequestBuilder instantiates a new GetTeamsTeamActivityDetailWithDateRequestBuilder and sets the default values.
func NewGetTeamsTeamActivityDetailWithDateRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetTeamsTeamActivityDetailWithDateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsTeamActivityDetailWithDateRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function getTeamsTeamActivityDetail
func (m *GetTeamsTeamActivityDetailWithDateRequestBuilder) CreateGetRequestInformation(options *GetTeamsTeamActivityDetailWithDateRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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