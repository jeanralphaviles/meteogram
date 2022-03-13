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

[Try it!](https://xenon-antonym-213304.uc.r.appspot.com/meteogram/24.5465,-81.7974)

```csv
Time,Temperature,RelativeHumidity,Dewpoint,HeatIndex,WindChill,WindSpeed,WindDirection,WindGust,SkyCover,ProbabilityOfPrecipitation
2020-09-25T22:00:00-04:00,27.2222,88.0000,25.0000,31.1111,0.0000,11.1120,170.0000,22.2240,76.0000,24.0000
2020-09-25T23:00:00-04:00,27.2222,85.0000,24.4444,30.5556,0.0000,11.1120,170.0000,18.5200,74.0000,23.0000
2020-09-26T00:00:00-04:00,27.2222,88.0000,25.0000,31.1111,0.0000,9.2600,180.0000,20.3720,74.0000,22.0000
2020-09-26T01:00:00-04:00,27.2222,85.0000,24.4444,30.5556,0.0000,9.2600,180.0000,20.3720,74.0000,21.0000
2020-09-26T02:00:00-04:00,27.2222,88.0000,25.0000,31.1111,0.0000,7.4080,180.0000,20.3720,72.0000,20.0000
2020-09-26T03:00:00-04:00,26.6667,88.0000,24.4444,30.0000,0.0000,9.2600,190.0000,20.3720,68.0000,19.0000
...
```

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
