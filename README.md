## Available Endpoint

  -  Check Available
     
     Get data bed available for each hospital in specific location
     - endpoint: BASEURL/api/available
	 - parameter:
	   - province [string]: province name
	 - example: `BASEURL/api/available?province=aceh`

  -  Detail Hospital
  
     Get data detail bed available for each room types in specific hospital
	 -  endpoint: BASEURL/api/detail
	 -  parameter:
	    - code [string]: hospital code, can be obtained from the response of API Check Available
	 - example: BASEURL/api/detail?code=aceh
	
	
	Repo: https://github.com/pararang/bed-covid-id-vercel-serverless-functions
	