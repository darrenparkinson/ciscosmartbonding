# Pull

This example demonstrates pulling updates.  You will need the following environment variables:

* `SMART_BONDING_USERNAME` - the user name given to access the service
* `SMART_BONDING_PASSWORD` - the associated password for the username

> Note: these can be provided in a `.env` file.

By default, this will connect to the test endpoint, and then:
* loop until a `204 No Content` is received;
* print out the object returned on each loop

Be aware that this operation is not idempotent and so any returned objects will no longer be accessible.