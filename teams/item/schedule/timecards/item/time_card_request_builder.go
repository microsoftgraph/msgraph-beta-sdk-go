package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i5127ef050fde02b8b8f1a5f5fe72b8694bf708f24b9f1e0b8352fb67d87e9683 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item/endbreak"
    i643863d7379eb2ab6adf49e281ac4151461d6911ca95ffa366ad2d8c5e9f4ce2 "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item/startbreak"
    i68e4a9fcef652132c7bebba9ebe08cb4fd8be16b5de9cd2ac842a8f4d53cc5ce "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item/confirm"
    i928cffe1d44a7c917c1e83d9a4105e827fa065d83a7a505a4f7efa3c64266fdb "github.com/microsoftgraph/msgraph-beta-sdk-go/teams/item/schedule/timecards/item/clockout"
)

// Builds and executes requests for operations under \teams\{team-id}\schedule\timeCards\{timeCard-id}
type TimeCardRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type TimeCardRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type TimeCardRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TimeCardRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get timeCards from teams
type TimeCardRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type TimeCardRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeCard;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TimeCardRequestBuilder) ClockOut()(*i928cffe1d44a7c917c1e83d9a4105e827fa065d83a7a505a4f7efa3c64266fdb.ClockOutRequestBuilder) {
    return i928cffe1d44a7c917c1e83d9a4105e827fa065d83a7a505a4f7efa3c64266fdb.NewClockOutRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *TimeCardRequestBuilder) Confirm()(*i68e4a9fcef652132c7bebba9ebe08cb4fd8be16b5de9cd2ac842a8f4d53cc5ce.ConfirmRequestBuilder) {
    return i68e4a9fcef652132c7bebba9ebe08cb4fd8be16b5de9cd2ac842a8f4d53cc5ce.NewConfirmRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new TimeCardRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTimeCardRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TimeCardRequestBuilder) {
    m := &TimeCardRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/beta/teams/{team_id}/schedule/timeCards/{timeCard_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new TimeCardRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewTimeCardRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TimeCardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTimeCardRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property timeCards for teams
// Parameters:
//  - options : Options for the request
func (m *TimeCardRequestBuilder) CreateDeleteRequestInformation(options *TimeCardRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// Get timeCards from teams
// Parameters:
//  - options : Options for the request
func (m *TimeCardRequestBuilder) CreateGetRequestInformation(options *TimeCardRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
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
// Update the navigation property timeCards in teams
// Parameters:
//  - options : Options for the request
func (m *TimeCardRequestBuilder) CreatePatchRequestInformation(options *TimeCardRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
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
// Delete navigation property timeCards for teams
// Parameters:
//  - options : Options for the request
func (m *TimeCardRequestBuilder) Delete(options *TimeCardRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TimeCardRequestBuilder) EndBreak()(*i5127ef050fde02b8b8f1a5f5fe72b8694bf708f24b9f1e0b8352fb67d87e9683.EndBreakRequestBuilder) {
    return i5127ef050fde02b8b8f1a5f5fe72b8694bf708f24b9f1e0b8352fb67d87e9683.NewEndBreakRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get timeCards from teams
// Parameters:
//  - options : Options for the request
func (m *TimeCardRequestBuilder) Get(options *TimeCardRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeCard, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewTimeCard() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.TimeCard), nil
}
// Update the navigation property timeCards in teams
// Parameters:
//  - options : Options for the request
func (m *TimeCardRequestBuilder) Patch(options *TimeCardRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *TimeCardRequestBuilder) StartBreak()(*i643863d7379eb2ab6adf49e281ac4151461d6911ca95ffa366ad2d8c5e9f4ce2.StartBreakRequestBuilder) {
    return i643863d7379eb2ab6adf49e281ac4151461d6911ca95ffa366ad2d8c5e9f4ce2.NewStartBreakRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
