## Solution

- Concurrent fetching for restaurants, parking-facility and petrol-station is done
- The main program has default api_key and categories it works without giving any additional data..
- Example call can be http://localhost:50053/v1/maps/places?loc=52.5159,13.3777 assuming application is running in docker and port is enabled.
- tests and caching is not yet developed due to time constraint. I can complete it by tomorrow if required.
- For Caching go-cashe is picked so that no dependency like Redis or Memcached but due to time constraint I did not develop it ..
 
## Docker commands

- I choose very light weight alpine based image with this application binary file
- Any modifications in the code sudo run make build 
- run the following command from the project directory sudo docker build -t assignment. 
- sudo docker run -it --name=app -p 50053:50053 assignment

## Fetch Commands 

- To fetch data without with default api_key use the following URL  http://localhost:50053/v1/maps/places?loc=52.5159,13.3777

- To fetch data along with query parameters, use the following URL https://places.ls.hereapi.com/places/v1/discover/explore?at=52.5159,13.3777&cat=parking-facility&apiKey=NQLeBf6xcolqAFhQyex0sHeAILpgHqSdTT45i1ahPdI&size=1

- categories(cat) can be comma separated.

