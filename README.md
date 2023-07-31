# GoURIparser

GoURIparser is a Go library that provides multiple utility functions for dealing with cloud storage URL/URIs of
the following cloud providers / services:
* Amazon Web Services: Simple Storage Service (AWS S3)
* Google Cloud Platform: Cloud Storage (GCP CS)

It uses Go version `1.20` and has no dependencies.

## Features

The following features have been implemented:
* Determine if given URL/URI is on AWS S3 or GCP CS
* Parse storage URL into its separate components (region, bucket, key)

## References

This library is used by the following projects:
* [GoStorage](https://github.com/FaaSTools/GoStorage)
* [GoText2Speech](https://github.com/FaaSTools/GoText2Speech)
* [GoSpeech2Text](https://github.com/FaaSTools/GoSpeech2Text)