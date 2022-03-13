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

| Time                 | Temperature | Dewpoint | MaxTemperature | MinTemperature | RelativeHumidity | ApparentTemperature | HeatIndex | WindChill | SkyCover | WindDirection | WindSpeed | WindGust | ProbabilityOfPrecipitation | QuantitativePrecipitation |
|----------------------|-------------|----------|----------------|----------------|------------------|---------------------|-----------|-----------|----------|---------------|-----------|----------|----------------------------|---------------------------|
| 2022-03-14T02:00:00Z | -5.0000     | -7.2222  | -3.3333        | -10.0000       | 85.0000          | -13.3333            | 0.0000    | -13.3333  | 77.0000  | 300.0000      | 31.4840   | 46.3000  | 57.0000                    | 3.3020                    |
| 2022-03-14T03:00:00Z | -6.1111     | -7.7778  | -3.3333        | -10.0000       | 90.0000          | -14.4444            | 0.0000    | -14.4444  | 77.0000  | 300.0000      | 31.4840   | 46.3000  | 57.0000                    | 3.3020                    |
| 2022-03-14T04:00:00Z | -6.1111     | -7.7778  |                | -10.0000       | 86.0000          | -14.4444            | 0.0000    | -14.4444  | 78.0000  | 310.0000      | 29.6320   | 46.3000  | 36.0000                    | 3.3020                    |
| 2022-03-14T05:00:00Z | -6.1111     | -7.7778  |                | -10.0000       | 86.0000          | -13.8889            | 0.0000    | -13.8889  | 78.0000  | 310.0000      | 29.6320   | 46.3000  | 36.0000                    | 3.3020                    |
| 2022-03-14T06:00:00Z | -6.6667     | -8.8889  |                | -10.0000       | 84.0000          | -15.0000            | 0.0000    | -15.0000  | 78.0000  | 310.0000      | 29.6320   | 46.3000  | 36.0000                    | 3.3020                    |
| 2022-03-14T07:00:00Z | -6.1111     | -8.3333  |                | -10.0000       | 86.0000          | -14.4444            | 0.0000    | -14.4444  | 81.0000  | 310.0000      | 29.6320   | 44.4480  | 19.0000                    | 0.0000                    |

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
