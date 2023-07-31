# GoURIparser

GoURIparser is a Go library that provides multiple utility functions for dealing with cloud storage URL/URIs of
the following cloud providers / services:
* Amazon Web Services: Simple Storage Service (AWS S3)
* Google Cloud Platform: Cloud Storage (GCP CS)

It uses Go version `1.20` and has no dependencies.

## Features

The following features have been implemented:
* Determine if given URL is on AWS or GCP
* Parse storage URL into its separate components (region, bucket, key)