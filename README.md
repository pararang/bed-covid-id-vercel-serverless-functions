# API Hospital Bed Availability for COVID-19 Patient In Indonesia

## Data Source

The data are scrapped (on the fly per API http request) from MoH's official website named [SIRANAP V 3.0](https://yankes.kemkes.go.id/app/siranap/).

## Available Endpoint  

- Check Available  
     Get data bed available for each hospital in specific location
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
