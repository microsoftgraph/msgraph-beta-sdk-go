package manageddeviceenrollmentabandonmentsummarywithskipwithtopwithfilterwithskiptoken

import (
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
)

// ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder builds and executes requests for operations under \reports\microsoft.graph.managedDeviceEnrollmentAbandonmentSummary(skip={skip},top={top},filter='{filter}',skipToken='{skipToken}')
type ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetOptions options for Get
type ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse union type wrapper for classes report
type ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Union type representation for type report
    report *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Report;
}
// NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse instantiates a new managedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse and sets the default values.
func NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse()(*ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse) {
    m := &ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetReport gets the report property value. Union type representation for type report
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse) GetReport()(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Report) {
    if m == nil {
        return nil
    } else {
        return m.report
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["report"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewReport() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReport(val.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Report))
        }
        return nil
    }
    return res
}
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("report", m.GetReport())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetReport sets the report property value. Union type representation for type report
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse) SetReport(value *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.Report)() {
    if m != nil {
        m.report = value
    }
}
// NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal instantiates a new ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder and sets the default values.
func NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, skip *int32, top *int32, filter *string, skipToken *string)(*ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    m := &ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.managedDeviceEnrollmentAbandonmentSummary(skip={skip},top={top},filter='{filter}',skipToken='{skipToken}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if skip != nil {
        urlTplParams["skip"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*skip), 10)
    }
    if top != nil {
        urlTplParams["top"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*top), 10)
    }
    if filter != nil {
        urlTplParams["filter"] = *filter
    }
    if skipToken != nil {
        urlTplParams["skipToken"] = *skipToken
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder instantiates a new ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder and sets the default values.
func NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil, nil)
}
// CreateGetRequestInformation metadata for Enrollment abandonment summary report
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) CreateGetRequestInformation(options *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get metadata for Enrollment abandonment summary report
func (m *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilder) Get(options *ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenRequestBuilderGetOptions)(*ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*ManagedDeviceEnrollmentAbandonmentSummaryWithSkipWithTopWithFilterWithSkipTokenResponse), nil
}
