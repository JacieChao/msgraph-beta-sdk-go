package getteamsdeviceusagedistributiontotalusercountswithperiod

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// Builds and executes requests for operations under \reports\microsoft.graph.getTeamsDeviceUsageDistributionTotalUserCounts(period='{period}')
type GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - period : Usage: period={period}
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, period *string)(*GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    m := &GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/reports/microsoft.graph.getTeamsDeviceUsageDistributionTotalUserCounts(period='{period}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if period != nil {
        urlTplParams["period"] = *period
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Invoke function getTeamsDeviceUsageDistributionTotalUserCounts
// Parameters:
//  - options : Options for the request
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) CreateGetRequestInformation(options *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Invoke function getTeamsDeviceUsageDistributionTotalUserCounts
// Parameters:
//  - options : Options for the request
func (m *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilder) Get(options *GetTeamsDeviceUsageDistributionTotalUserCountsWithPeriodRequestBuilderGetOptions)([]byte, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(*requestInfo, "byte", nil)
    if err != nil {
        return nil, err
    }
    return res.([]byte), nil
}
