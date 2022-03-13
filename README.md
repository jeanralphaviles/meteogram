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

| Time                 | Temperature (C) | Dewpoint (C) | MaxTemperature (C) | MinTemperature (C) | RelativeHumidity (%) | ApparentTemperature (C) | HeatIndex (C) | WindChill (C) | SkyCover (%) | WindDirection (angle) | WindSpeed (km/h) | WindGust (km/h) | ProbabilityOfPrecipitation (%) | QuantitativePrecipitation (mm) | Coverage | Weather      | Intensity |
|----------------------|-----------------|--------------|--------------------|--------------------|----------------------|-------------------------|---------------|---------------|--------------|-----------------------|------------------|-----------------|--------------------------------|--------------------------------|----------|--------------|-----------|
| 2022-03-13T21:00:00Z | -1.1111         | -8.3333      | 0.0000             |                    | 57.0000              | -7.7778                 | 0.0000        | -7.7778       | 94.0000      | 260.0000              | 27.7800          | 48.1520         | 89.0000                        | 6.3500                         | definite | snow_showers | heavy     |
| 2022-03-13T22:00:00Z | -1.6667         | -7.7778      | 0.0000             |                    | 64.0000              | -8.3333                 | 0.0000        | -8.3333       | 86.0000      | 270.0000              | 27.7800          | 53.7080         | 79.0000                        | 6.3500                         | definite | snow_showers | heavy     |
| 2022-03-13T23:00:00Z | -2.2222         | -7.2222      | 0.0000             |                    | 68.0000              | -8.8889                 | 0.0000        | -8.8889       | 86.0000      | 270.0000              | 27.7800          | 53.7080         | 79.0000                        | 6.3500                         | definite | snow_showers | heavy     |
| 2022-03-14T00:00:00Z | -3.3333         | -7.7778      | 0.0000             |                    | 71.0000              | -10.5556                | 0.0000        | -10.5556      | 86.0000      | 270.0000              | 27.7800          | 53.7080         | 79.0000                        | 6.3500                         | definite | snow_showers | heavy     |
| 2022-03-14T01:00:00Z | -3.3333         | -6.6667      | 0.0000             |                    | 79.0000              | -11.1111                | 0.0000        | -11.1111      | 84.0000      | 300.0000              | 31.4840          | 46.3000         | 64.0000                        | 3.5560                         | numerous | snow_showers | moderate  |
| 2022-03-14T02:00:00Z | -3.3333         | -5.0000      | 0.0000             | -10.5556           | 88.0000              | -11.1111                | 0.0000        | -11.1111      | 84.0000      | 300.0000              | 31.4840          | 46.3000         | 64.0000                        | 3.5560                         | numerous | snow_showers | moderate  |

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
