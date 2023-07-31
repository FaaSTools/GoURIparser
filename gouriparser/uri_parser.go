package gouriparser

import "strings"

// IsGoogleCloudStorageUrl returns true if given URL is a valid Google Cloud Storage URL.
// CS Object URL example: gs://test-bucket/newfile.png
// CS Object URL example: https://storage.cloud.google.com/test-bucket/newfile.png
//
// Taken from GoStorage
func IsGoogleCloudStorageUrl(urlString string) bool {
	return strings.HasPrefix(urlString, "gs://") || strings.Contains(urlString, "storage.cloud.google.com")
}

// IsAwsS3Url returns true if given URL is a valid Amazon Web Services S3 URL.
// S3 Object URL without region example: https://test-bucket.s3.amazonaws.com/newfile.png
// S3 Object URL with region example: https://test-bucket.s3.eu-central-1.amazonaws.com/newfile.png
//
// Taken from GoStorage
func IsAwsS3Url(urlString string) bool {
	return strings.HasPrefix(urlString, "https://") && strings.Contains(urlString, "s3")
}

// IsAwsS3Uri returns true if given URI is a valid Amazon Web Services S3 URI.
// AWS Object URI example: s3://test-bucket/newfile.png
//
// Taken from GoText2Speech
func IsAwsS3Uri(uriString string) bool {
	return strings.HasPrefix(uriString, "S3://")
}

// IsAwsS3UrlOrUri returns true if given URI/URL is a valid Amazon Web Services S3 URL or URI.
// S3 Object URL without region example: https://test-bucket.s3.amazonaws.com/newfile.png
// S3 Object URL with region example: https://test-bucket.s3.eu-central-1.amazonaws.com/newfile.png
// AWS Object URI example: s3://test-bucket/newfile.png
//
// Taken from GoText2Speech
func IsAwsS3UrlOrUri(urlOrUriString string) bool {
	return IsAwsS3Url(urlOrUriString) || IsAwsS3Uri(urlOrUriString)
}

// ParseGoogleCloudStorageUrl returns the bucket and key of the file stored on the given Google Cloud Storage URL.
// First return value is the bucket, second return value is the key (both as strings).
// CS Object URL example: gs://test-bucket/newfile.png
// CS Object URL example: https://storage.cloud.google.com/test-bucket/newfile.png
//
// Adapted from GoStorage
func ParseGoogleCloudStorageUrl(urlString string) (string, string) {
	var bucket string
	var key string

	if strings.HasPrefix(urlString, "gs://") {
		urlString = urlString[strings.Index(urlString, "gs://")+len("gs://"):]
	} else if strings.HasPrefix(urlString, "https://storage.cloud.google.com/") {
		urlString = urlString[strings.Index(urlString, "https://storage.cloud.google.com/")+len("https://storage.cloud.google.com/"):]
	}
	if strings.Contains(urlString, "/") {
		bucket = urlString[:strings.Index(urlString, "/")]
		key = urlString[strings.Index(urlString, "/")+1:]
	} else {
		bucket = urlString
	}

	return bucket, key
}

// ParseAwsS3Url returns the bucket, key, and region (if exists) of the file stored on the given AWS S3 URL.
// First return value is the bucket, second return value is the key, third return value is region (all as strings).
// If the region is not specified in the given URL, region is empty string (i.e. "").
// S3 Object URL without region example: https://test-bucket.s3.amazonaws.com/newfile.png
// S3 Object URL with region example: https://test-bucket.s3.eu-central-1.amazonaws.com/newfile.png
//
// Adapted from GoStorage
func ParseAwsS3Url(urlString string) (string, string, string) {
	var bucket string
	var key string
	var region string

	urlString = urlString[strings.Index(urlString, "https://")+len("https://"):]
	bucket = urlString[:strings.Index(urlString, ".")]
	urlString = urlString[strings.Index(urlString, ".")+len(".s3."):]
	if strings.HasPrefix(urlString, "amazonaws.com") { //No region specified
		region = ""
	} else {
		region = urlString[:strings.Index(urlString, ".")]
		urlString = urlString[strings.Index(urlString, ".")+1:]
	}
	urlString = urlString[strings.Index(urlString, "amazonaws.com")+len("amazonaws.com"):]
	if strings.HasPrefix(urlString, "/") {
		urlString = urlString[1:]
	}
	key = urlString

	return bucket, key, region
}
