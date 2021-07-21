# API Hospital Bed Availability for COVID-19 Patient In Indonesia

[![Lines Of Code](https://tokei.rs/b1/github/pararang/bed-covid-id-vercel-serverless-functions?category=code)](https://github.com/pararang/bed-covid-id-vercel-serverless-functions)
[![Go Report Card](https://goreportcard.com/badge/github.com/pararang/bed-covid-id-vercel-serverless-functions)](https://goreportcard.com/report/github.com/pararang/bed-covid-id-vercel-serverless-functions)

## Data Source

The data are scrapped (on the fly per API http request) from MoH's official website named [SIRANAP V 3.0](https://yankes.kemkes.go.id/app/siranap/).

## Available Endpoint  

- List Provinces  
     Get list of province's name  
  - endpoint: BASEURL/api/province
  - parameter: `none`
  - example: `BASEURL/api/province`

- Check Available  
     Get data bed available for each hospital in specific location, for now only support filter by province
  - endpoint: BASEURL/api/available
  - parameter:
    - province `[string]`: province name
  - example: `BASEURL/api/available?province=aceh`  
- Detail Hospital  
     Get data detail bed available for each room types in specific hospital
  - endpoint: BASEURL/api/detail
  - parameter:
    - code `[number]`: hospital code, can be obtained from the response of API Check Available
  - example: `BASEURL/api/detail?code=123456`

## Infra

 Deployed using Vercel Serverless Function that integrated with DataDog for storing log.
