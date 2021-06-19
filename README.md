# aquarium-lights
A simple GO based scheduler that turns "on" a relay attached to a Pi3 microcontroller.

An example `configuration.json`:

```json
{
  "schedules": [
    {
      "name": "sump",
      "pin": 23,
      "periods": [
        {
          "start_time": "2021-01-01T18:00:00.000-0300",
          "end_time": "2021-12-31T23:59:00.000-0300"
        },
        {
          "start_time": "2021-01-01T00:00:00.000-0300",
          "end_time": "2021-12-31T10:00:00.000-0300"
        }
      ]
    },
    {
      "name": "display",
      "pin": 22,
      "periods": [
        {
          "start_time":"2021-01-01T10:00:00.000-0300",
          "end_time": "2021-12-31T18:00:00.000-0300"
        }
      ]
    }
  ]
}
```