package external

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
    id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a "github.com/microsoftgraph/msgraph-beta-sdk-go/models/industrydata"
)

// IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder provides operations to call the getStatistics method.
type IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewIndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal instantiates a new MicrosoftGraphIndustryDataGetStatisticsRequestBuilder and sets the default values.
func NewIndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    m := &IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/external/industryData/runs/microsoft.graph.industryData.getStatistics()", pathParameters),
    }
    return m
}
// NewIndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder instantiates a new MicrosoftGraphIndustryDataGetStatisticsRequestBuilder and sets the default values.
func NewIndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function getStatistics
func (m *IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) Get(ctx context.Context, requestConfiguration *IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration)(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunStatisticsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.CreateIndustryDataRunStatisticsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(id2b46acaed365d10a0a4cc89e0aa6f2f76ad54e2147428aee709d25e554da66a.IndustryDataRunStatisticsable), nil
}
// ToGetRequestInformation invoke function getStatistics
func (m *IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/json;q=1")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) WithUrl(rawUrl string)(*IndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder) {
    return NewIndustryDataRunsMicrosoftGraphIndustryDataGetStatisticsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
