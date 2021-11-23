package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc "github.com/microsoftgraph/msgraph-beta-sdk-go/models/microsoft/graph"
    i258a858c85da90b801e765428af11e65bc154f922525178ab29e97c7a14c345e "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/certificateconnectordetails/item/gethealthmetrictimeseries"
    i7e1773f4927e9225f8c986534d86959f50bad1e76f4272d85d379daeb2d90bee "github.com/microsoftgraph/msgraph-beta-sdk-go/devicemanagement/certificateconnectordetails/item/gethealthmetrics"
)

// certificateConnectorDetailsRequestBuilder builds and executes requests for operations under \deviceManagement\certificateConnectorDetails\{certificateConnectorDetails-id}
type CertificateConnectorDetailsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CertificateConnectorDetailsRequestBuilderDeleteOptions options for Delete
type CertificateConnectorDetailsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CertificateConnectorDetailsRequestBuilderGetOptions options for Get
type CertificateConnectorDetailsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CertificateConnectorDetailsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// certificateConnectorDetailsRequestBuilderGetQueryParameters collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
type CertificateConnectorDetailsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// CertificateConnectorDetailsRequestBuilderPatchOptions options for Patch
type CertificateConnectorDetailsRequestBuilderPatchOptions struct {
    // 
    Body *i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CertificateConnectorDetails;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewCertificateConnectorDetailsRequestBuilderInternal instantiates a new CertificateConnectorDetailsRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CertificateConnectorDetailsRequestBuilder) {
    m := &CertificateConnectorDetailsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/certificateConnectorDetails/{certificateConnectorDetails_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCertificateConnectorDetailsRequestBuilder instantiates a new CertificateConnectorDetailsRequestBuilder and sets the default values.
func NewCertificateConnectorDetailsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CertificateConnectorDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCertificateConnectorDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *CertificateConnectorDetailsRequestBuilder) CreateDeleteRequestInformation(options *CertificateConnectorDetailsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *CertificateConnectorDetailsRequestBuilder) CreateGetRequestInformation(options *CertificateConnectorDetailsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *CertificateConnectorDetailsRequestBuilder) CreatePatchRequestInformation(options *CertificateConnectorDetailsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *CertificateConnectorDetailsRequestBuilder) Delete(options *CertificateConnectorDetailsRequestBuilderDeleteOptions)(error) {
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
// Get collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *CertificateConnectorDetailsRequestBuilder) Get(options *CertificateConnectorDetailsRequestBuilderGetOptions)(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CertificateConnectorDetails, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.NewCertificateConnectorDetails() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i535684e11b5500196ecb4b5c6634e0651fe2c2f78b6cd0fbe097d3c9029ae7bc.CertificateConnectorDetails), nil
}
func (m *CertificateConnectorDetailsRequestBuilder) GetHealthMetrics()(*i7e1773f4927e9225f8c986534d86959f50bad1e76f4272d85d379daeb2d90bee.GetHealthMetricsRequestBuilder) {
    return i7e1773f4927e9225f8c986534d86959f50bad1e76f4272d85d379daeb2d90bee.NewGetHealthMetricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CertificateConnectorDetailsRequestBuilder) GetHealthMetricTimeSeries()(*i258a858c85da90b801e765428af11e65bc154f922525178ab29e97c7a14c345e.GetHealthMetricTimeSeriesRequestBuilder) {
    return i258a858c85da90b801e765428af11e65bc154f922525178ab29e97c7a14c345e.NewGetHealthMetricTimeSeriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch collection of certificate connector details, each associated with a corresponding Intune Certificate Connector.
func (m *CertificateConnectorDetailsRequestBuilder) Patch(options *CertificateConnectorDetailsRequestBuilderPatchOptions)(error) {
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
