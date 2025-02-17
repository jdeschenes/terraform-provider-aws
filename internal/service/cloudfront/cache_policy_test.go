package cloudfront_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudfront"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccCloudFrontCachePolicy_basic(t *testing.T) {
	rInt := sdkacctest.RandInt()
	resourceName := "aws_cloudfront_cache_policy.example"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(cloudfront.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, cloudfront.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckCloudFrontPublicKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachePolicyConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
					resource.TestCheckResourceAttr(resourceName, "default_ttl", "50"),
					resource.TestCheckResourceAttr(resourceName, "min_ttl", "1"),
					resource.TestCheckResourceAttr(resourceName, "max_ttl", "100"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.cookies_config.0.cookie_behavior", "whitelist"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.cookies_config.0.cookies.0.items.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.headers_config.0.header_behavior", "whitelist"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.headers_config.0.headers.0.items.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.query_strings_config.0.query_string_behavior", "whitelist"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.query_strings_config.0.query_strings.0.items.0", "test"),
				),
			},
			{
				ResourceName:            "aws_cloudfront_cache_policy.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccCloudFrontCachePolicy_update(t *testing.T) {
	rInt := sdkacctest.RandInt()
	resourceName := "aws_cloudfront_cache_policy.example"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(cloudfront.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, cloudfront.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckCloudFrontPublicKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachePolicyConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
					resource.TestCheckResourceAttr(resourceName, "default_ttl", "50"),
					resource.TestCheckResourceAttr(resourceName, "min_ttl", "1"),
					resource.TestCheckResourceAttr(resourceName, "max_ttl", "100"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.cookies_config.0.cookie_behavior", "whitelist"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.cookies_config.0.cookies.0.items.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.headers_config.0.header_behavior", "whitelist"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.headers_config.0.headers.0.items.0", "test"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.query_strings_config.0.query_string_behavior", "whitelist"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.query_strings_config.0.query_strings.0.items.0", "test"),
				),
			},
			{
				Config: testAccCachePolicyUpdateConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment updated"),
					resource.TestCheckResourceAttr(resourceName, "default_ttl", "51"),
					resource.TestCheckResourceAttr(resourceName, "min_ttl", "2"),
					resource.TestCheckResourceAttr(resourceName, "max_ttl", "101"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.cookies_config.0.cookie_behavior", "allExcept"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.cookies_config.0.cookies.0.items.0", "test2"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.headers_config.0.header_behavior", "none"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.headers_config.0.headers.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.query_strings_config.0.query_string_behavior", "allExcept"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.query_strings_config.0.query_strings.0.items.0", "test2"),
				),
			},
			{
				ResourceName:            "aws_cloudfront_cache_policy.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func TestAccCloudFrontCachePolicy_noneBehavior(t *testing.T) {
	rInt := sdkacctest.RandInt()
	resourceName := "aws_cloudfront_cache_policy.example"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { acctest.PreCheck(t); acctest.PreCheckPartitionHasService(cloudfront.EndpointsID, t) },
		ErrorCheck:   acctest.ErrorCheck(t, cloudfront.EndpointsID),
		Providers:    acctest.Providers,
		CheckDestroy: testAccCheckCloudFrontPublicKeyDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCachePolicyNoneBehaviorConfig(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
					resource.TestCheckResourceAttr(resourceName, "default_ttl", "50"),
					resource.TestCheckResourceAttr(resourceName, "min_ttl", "1"),
					resource.TestCheckResourceAttr(resourceName, "max_ttl", "100"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.cookies_config.0.cookie_behavior", "none"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.cookies_config.0.cookies.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.headers_config.0.header_behavior", "none"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.headers_config.0.headers.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.query_strings_config.0.query_string_behavior", "none"),
					resource.TestCheckResourceAttr(resourceName, "parameters_in_cache_key_and_forwarded_to_origin.0.query_strings_config.0.query_strings.#", "0"),
				),
			},
			{
				ResourceName:            "aws_cloudfront_cache_policy.example",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccCachePolicyConfig(rInt int) string {
	return fmt.Sprintf(`
resource "aws_cloudfront_cache_policy" "example" {
  name        = "test-policy%[1]d"
  comment     = "test comment"
  default_ttl = 50
  max_ttl     = 100
  min_ttl     = 1
  parameters_in_cache_key_and_forwarded_to_origin {
    cookies_config {
      cookie_behavior = "whitelist"
      cookies {
        items = ["test"]
      }
    }
    headers_config {
      header_behavior = "whitelist"
      headers {
        items = ["test"]
      }
    }
    query_strings_config {
      query_string_behavior = "whitelist"
      query_strings {
        items = ["test"]
      }
    }
  }
}
`, rInt)
}

func testAccCachePolicyUpdateConfig(rInt int) string {
	return fmt.Sprintf(`
resource "aws_cloudfront_cache_policy" "example" {
  name        = "test-policy-updated%[1]d"
  comment     = "test comment updated"
  default_ttl = 51
  max_ttl     = 101
  min_ttl     = 2
  parameters_in_cache_key_and_forwarded_to_origin {
    cookies_config {
      cookie_behavior = "allExcept"
      cookies {
        items = ["test2"]
      }
    }
    headers_config {
      header_behavior = "none"
    }
    query_strings_config {
      query_string_behavior = "allExcept"
      query_strings {
        items = ["test2"]
      }
    }
  }
}
`, rInt)
}

func testAccCachePolicyNoneBehaviorConfig(rInt int) string {
	return fmt.Sprintf(`
resource "aws_cloudfront_cache_policy" "example" {
  name        = "test-policy-updated%[1]d"
  comment     = "test comment"
  default_ttl = 50
  max_ttl     = 100
  min_ttl     = 1
  parameters_in_cache_key_and_forwarded_to_origin {
    cookies_config {
      cookie_behavior = "none"
    }
    headers_config {
      header_behavior = "none"
    }
    query_strings_config {
      query_string_behavior = "none"
    }
  }
}
`, rInt)
}
