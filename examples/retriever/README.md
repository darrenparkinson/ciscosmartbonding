# Cisco Smartbonding Retriever

A simple application for retrieving Cisco Smart Bonding pull updates.

At present, it only displays some of the fields of the received update to standard output.

Future updates could include sending the updates to the service desk or to a message queue.

You will need the following environment variables:

* `SMART_BONDING_USERNAME` - the user name given to access the service
* `SMART_BONDING_PASSWORD` - the associated password for the username

> Note: these can be provided in a `.env` file.