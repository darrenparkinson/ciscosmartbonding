# Pull

This example demonstrates pulling updates.  You will need the following environment variables:

* `SMART_BONDING_CLIENTID` - the client id given to access the service
* `SMART_BONDING_SECRET` - the associated secret for the client id
  
> Note: these can be provided in a `.env` file.

By default, this will connect to the test endpoint, and then:
* loop until a status of "200" and a message of "no messages available to send" is received;
* print out the object returned on each loop

Be aware that this operation is not idempotent and so any returned objects will no longer be accessible.