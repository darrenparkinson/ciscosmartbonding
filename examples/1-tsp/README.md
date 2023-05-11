# TSP Codes

This example demonstrates retrieving TSP codes.  You will need the following environment variables:

* `SMART_BONDING_CLIENTID` - the client id given to access the service
* `SMART_BONDING_SECRET` - the associated secret for the client id

> Note: these can be provided in a `.env` file.

By default, this will connect to the test endpoint, and then:
* retrieve all entries and print the number of items retrieved.

You can optionally specify the `-save` and `-csv` flags to save as json and csv data respectively.