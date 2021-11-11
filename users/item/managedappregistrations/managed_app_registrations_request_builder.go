package managedappregistrations

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0bed9af6d22cdccd5d56d0f3fecc3210ae4934f3d784bcdca6f4cbe039fa70b2 "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/managedappregistrations/getuseridswithflaggedappregistration"
    ib50922c2d2a9787b10c876f2621661bd1366e024d3d5bf7c716b4df66780742d "github.com/microsoftgraph/msgraph-beta-sdk-go/users/item/managedappregistrations/ref"
)

// Builds and executes requests for operations under \users\{user-id}\managedAppRegistrations
type ManagedAppRegistrationsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type ManagedAppRegistrationsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedAppRegistrationsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Zero or more managed app registrations that belong to the user.
type ManagedAppRegistrationsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool;
    // Expand related entities
    Expand []string;
    // Filter items by property values
    Filter *string;
    // Order items by property values
    Orderby []string;
    // Search items by search phrases
    Search *string;
    // Select properties to be returned
    Select_escaped []string;
    // Skip the first n items
    Skip *int32;
    // Show only the first n items
    Top *int32;
}
// Instantiates a new ManagedAppRegistrationsRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedAppRegistrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationsRequestBuilder) {
    m := &ManagedAppRegistrationsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/managedAppRegistrations{?top,skip,search,filter,count,orderby,select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ManagedAppRegistrationsRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedAppRegistrationsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedAppRegistrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedAppRegistrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Zero or more managed app registrations that belong to the user.
// Parameters:
//  - options : Options for the request
func (m *ManagedAppRegistrationsRequestBuilder) CreateGetRequestInformation(options *ManagedAppRegistrationsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
// Zero or more managed app registrations that belong to the user.
// Parameters:
//  - options : Options for the request
func (m *ManagedAppRegistrationsRequestBuilder) Get(options *ManagedAppRegistrationsRequestBuilderGetOptions)(*ManagedAppRegistrationsResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppRegistrationsResponse() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagedAppRegistrationsResponse), nil
}
// Builds and executes requests for operations under \users\{user-id}\managedAppRegistrations\microsoft.graph.getUserIdsWithFlaggedAppRegistration()
func (m *ManagedAppRegistrationsRequestBuilder) GetUserIdsWithFlaggedAppRegistration()(*i0bed9af6d22cdccd5d56d0f3fecc3210ae4934f3d784bcdca6f4cbe039fa70b2.GetUserIdsWithFlaggedAppRegistrationRequestBuilder) {
    return i0bed9af6d22cdccd5d56d0f3fecc3210ae4934f3d784bcdca6f4cbe039fa70b2.NewGetUserIdsWithFlaggedAppRegistrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedAppRegistrationsRequestBuilder) Ref()(*ib50922c2d2a9787b10c876f2621661bd1366e024d3d5bf7c716b4df66780742d.RefRequestBuilder) {
    return ib50922c2d2a9787b10c876f2621661bd1366e024d3d5bf7c716b4df66780742d.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
