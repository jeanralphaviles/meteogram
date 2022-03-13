# Meteogram

Meteogram returns detailed weather forecasts for a given location. US only.

## API

Meteogram is deployed as an [App Engine](https://cloud.google.com/appengine)
app for public use.

https://xenon-antonym-213304.uc.r.appspot.com

### /meteogram

Retrieves 48 hours of weather forecasts for the location given its
[WGS 84](https://en.wikipedia.org/wiki/World_Geodetic_System) coordinate pair.
Only works for locations within the United States.

#### Request

```http
GET /meteogram/{latitude},{longitude}
```

#### Response

[Try it!](https://xenon-antonym-213304.uc.r.appspot.com/meteogram/40.6099,-111.5532)

| Time                 | Temperature | Dewpoint | MaxTemperature | MinTemperature | RelativeHumidity | ApparentTemperature | HeatIndex | WindChill | SkyCover | WindDirection | WindSpeed | WindGust | ProbabilityOfPrecipitation | QuantitativePrecipitation | Coverage | Weather      | Intensity |
|----------------------|-------------|----------|----------------|----------------|------------------|---------------------|-----------|-----------|----------|---------------|-----------|----------|----------------------------|---------------------------|----------|--------------|-----------|
| 2022-03-13T21:00:00Z | -3.8889     | -8.3333  | -3.3333        |                | 71.0000          | -11.1111            | 0.0000    | -11.1111  | 96.0000  | 260.0000      | 27.7800   | 48.1520  | 93.0000                    | 4.5720                    | definite | snow_showers | light     |
| 2022-03-13T22:00:00Z | -4.4444     | -8.8889  | -3.3333        |                | 72.0000          | -11.6667            | 0.0000    | -11.6667  | 86.0000  | 270.0000      | 27.7800   | 53.7080  | 78.0000                    | 4.5720                    | definite | snow_showers | light     |
| 2022-03-13T23:00:00Z | -5.0000     | -8.3333  | -3.3333        |                | 76.0000          | -12.7778            | 0.0000    | -12.7778  | 86.0000  | 270.0000      | 27.7800   | 53.7080  | 78.0000                    | 4.5720                    | definite | snow_showers | light     |
| 2022-03-14T00:00:00Z | -5.5556     | -8.3333  | -3.3333        |                | 80.0000          | -13.3333            | 0.0000    | -13.3333  | 86.0000  | 270.0000      | 27.7800   | 53.7080  | 78.0000                    | 4.5720                    | definite | snow_showers | light     |
| 2022-03-14T01:00:00Z | -5.5556     | -7.7778  | -3.3333        |                | 85.0000          | -14.4444            | 0.0000    | -14.4444  | 77.0000  | 300.0000      | 31.4840   | 46.3000  | 57.0000                    | 3.3020                    | likely   | snow_showers | light     |
| 2022-03-14T02:00:00Z | -5.0000     | -7.2222  | -3.3333        | -10.0000       | 85.0000          | -13.3333            | 0.0000    | -13.3333  | 77.0000  | 300.0000      | 31.4840   | 46.3000  | 57.0000                    | 3.3020                    | likely   | snow_showers | light     |

## Usage

```bash
Usage of ./meteogram:
  -address string
        Address to listen on. (default ":8080")
  -forecast_duration duration
        How far into the future should a forecast be. (default 48h0m0s)
```

## Docker

* Run the Docker container.

  ```bash
  docker run --rm --name meteogram --publish 8080:8080 jraviles/meteogram:latest
  ```

* Build your own Docker image.

  ```bash
  docker build --tag jraviles/meteogram:latest .
  ```

* Publish image to Docker Hub.

  ```bash
  docker push jraviles/meteogram:latest
  ```

## App Engine

* Deploy

  ```bash
  gcloud app deploy
  ```

## See Also

* [weather.gov API](https://weather-gov.github.io/api/)
* [github.com/icodealot/noaa library](https://github.com/icodealot/noaa)

## Author

[Jean-Ralph Aviles](http://jr.expert)
