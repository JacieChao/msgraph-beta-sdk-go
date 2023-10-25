package organization

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459 "github.com/microsoftgraph/msgraph-beta-sdk-go/models/odataerrors"
)

// ItemBrandingBannerLogoRequestBuilder provides operations to manage the media for the organization entity.
type ItemBrandingBannerLogoRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemBrandingBannerLogoRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingBannerLogoRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemBrandingBannerLogoRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemBrandingBannerLogoRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemBrandingBannerLogoRequestBuilderInternal instantiates a new BannerLogoRequestBuilder and sets the default values.
func NewItemBrandingBannerLogoRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingBannerLogoRequestBuilder) {
    m := &ItemBrandingBannerLogoRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/organization/{organization%2Did}/branding/bannerLogo", pathParameters),
    }
    return m
}
// NewItemBrandingBannerLogoRequestBuilder instantiates a new BannerLogoRequestBuilder and sets the default values.
func NewItemBrandingBannerLogoRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemBrandingBannerLogoRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemBrandingBannerLogoRequestBuilderInternal(urlParams, requestAdapter)
}
// Get a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/organizationalbranding-get?view=graph-rest-1.0
func (m *ItemBrandingBannerLogoRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemBrandingBannerLogoRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *ItemBrandingBannerLogoRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ItemBrandingBannerLogoRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
        "5XX": i20a3050780ee0b0cde0a884a4f35429a20d60067e3bcda382ec5400079147459.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToGetRequestInformation a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *ItemBrandingBannerLogoRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemBrandingBannerLogoRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation a banner version of your company logo that appears on the sign-in page. The allowed types are PNG or JPEG not larger than 36 × 245 pixels. We recommend using a transparent image with no padding around the logo.
func (m *ItemBrandingBannerLogoRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemBrandingBannerLogoRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT
    requestInfo.Headers.TryAdd("Accept", "application/json, application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemBrandingBannerLogoRequestBuilder) WithUrl(rawUrl string)(*ItemBrandingBannerLogoRequestBuilder) {
    return NewItemBrandingBannerLogoRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
