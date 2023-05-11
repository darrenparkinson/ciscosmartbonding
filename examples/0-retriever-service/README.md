# Cisco Smartbonding Retriever

A simple application for retrieving Cisco Smart Bonding pull updates.

At present, it only displays some of the fields of the received update to standard output.

Future updates could include sending the updates to the service desk or to a message queue.

You will need the following environment variables:

* `SMART_BONDING_CLIENTID` - the client id given to access the service
* `SMART_BONDING_SECRET` - the associated secret for the client id

> Note: these can be provided in a `.env` file.