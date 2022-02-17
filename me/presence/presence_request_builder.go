package presence

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0f150f1f9267a2e06d61f15acad792bf666470435cf6f393557754329e9a5016 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/presence/clearpresence"
    i1dce1fc04294c4b323f327df6407bbf2068d5dd2350874d0d1ec7b7c666ce041 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/presence/setuserpreferredpresence"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i8c0fe1e1d104ef823cb402ec5f6d90d2100c551425ae497db0269870429ae0eb "github.com/microsoftgraph/msgraph-beta-sdk-go/me/presence/clearuserpreferredpresence"
    i8c94f05fe8662460db4ce1c0b8e6f88c23101dabbf4e6277d005beceff32c273 "github.com/microsoftgraph/msgraph-beta-sdk-go/me/presence/setpresence"
)

// PresenceRequestBuilder builds and executes requests for operations under \me\presence
type PresenceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PresenceRequestBuilderDeleteOptions options for Delete
type PresenceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PresenceRequestBuilderGetOptions options for Get
type PresenceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PresenceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PresenceRequestBuilderGetQueryParameters get presence from me
type PresenceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PresenceRequestBuilderPatchOptions options for Patch
type PresenceRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Presence;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PresenceRequestBuilder) ClearPresence()(*i0f150f1f9267a2e06d61f15acad792bf666470435cf6f393557754329e9a5016.ClearPresenceRequestBuilder) {
    return i0f150f1f9267a2e06d61f15acad792bf666470435cf6f393557754329e9a5016.NewClearPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PresenceRequestBuilder) ClearUserPreferredPresence()(*i8c0fe1e1d104ef823cb402ec5f6d90d2100c551425ae497db0269870429ae0eb.ClearUserPreferredPresenceRequestBuilder) {
    return i8c0fe1e1d104ef823cb402ec5f6d90d2100c551425ae497db0269870429ae0eb.NewClearUserPreferredPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPresenceRequestBuilderInternal instantiates a new PresenceRequestBuilder and sets the default values.
func NewPresenceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PresenceRequestBuilder) {
    m := &PresenceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/presence{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPresenceRequestBuilder instantiates a new PresenceRequestBuilder and sets the default values.
func NewPresenceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PresenceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresenceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property presence for me
func (m *PresenceRequestBuilder) CreateDeleteRequestInformation(options *PresenceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get presence from me
func (m *PresenceRequestBuilder) CreateGetRequestInformation(options *PresenceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
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
// CreatePatchRequestInformation update the navigation property presence in me
func (m *PresenceRequestBuilder) CreatePatchRequestInformation(options *PresenceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property presence for me
func (m *PresenceRequestBuilder) Delete(options *PresenceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get presence from me
func (m *PresenceRequestBuilder) Get(options *PresenceRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Presence, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewPresence() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Presence), nil
}
// Patch update the navigation property presence in me
func (m *PresenceRequestBuilder) Patch(options *PresenceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *PresenceRequestBuilder) SetPresence()(*i8c94f05fe8662460db4ce1c0b8e6f88c23101dabbf4e6277d005beceff32c273.SetPresenceRequestBuilder) {
    return i8c94f05fe8662460db4ce1c0b8e6f88c23101dabbf4e6277d005beceff32c273.NewSetPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PresenceRequestBuilder) SetUserPreferredPresence()(*i1dce1fc04294c4b323f327df6407bbf2068d5dd2350874d0d1ec7b7c666ce041.SetUserPreferredPresenceRequestBuilder) {
    return i1dce1fc04294c4b323f327df6407bbf2068d5dd2350874d0d1ec7b7c666ce041.NewSetUserPreferredPresenceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
