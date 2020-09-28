# api2html

Go server that uses go HTML templates to front an API

## Usage

Setup the config file as per the example setup, and put your templates in a folder called `templates`:

```
{
    "pages": [
        {
            "api_url": "http://api.example.com/path/to/the/things",
            "uri": "/things",
            "template": "things",
            "cache_expiry": 30
        },
        {
            "api_url": "http://api.example.com/path/to/the/thing/:id",
            "uri": "/things/:id",
            "template": "thing",
            "cache_expiry": 30
        }
    ]
}
```

This is what each page object in the config file represents:

|Field|Purpose|
|---|---|
| api_url      | the URL to call, can use URL `:params` in the URI segments |
| uri          | the local URI to that would be served locally, `:params` with the same name are passed to the API URL |
| template     | The name of the template to use from the `templates` directory |
| cache_expiry | What to set the `max-age` value for the `Cache-Control` header to avoid hammering your APIs |

Then just run the binary like so:

    api2html -c config.json
