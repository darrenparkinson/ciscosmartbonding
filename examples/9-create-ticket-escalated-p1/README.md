# Create Ticket

This example demonstrates creating a new ticket with the minimum mandatory fields.  You will need the following environment variables:

* `SMART_BONDING_CLIENTID` - the client id given to access the service
* `SMART_BONDING_SECRET` - the associated secret for the client id
* `SMART_BONDING_CONTRACT_VALUE` - contract value provided by Cisco
* `SMART_BONDING_CONTRACT_ELEMENT_VALUE` element value provided by Cisco
* `SMART_BONDING_FIRSTNAME` - partner account first name, e.g. technical
* `SMART_BONDING_LASTNAME` - partner account last name, e.g. support
* `SMART_BONDING_CCOID` - partner CCO ID
* `SMART_BONDING_TELEPHONE` - partner telephone number
* `SMART_BONDING_EMAIL` - partner email address, e.g. support@partner.com
* `SMART_BONDING_PRODUCTID` - valid entry relating to the product
* `SMART_BONDING_SERIALNO` - valid entry relating to the product
* `SMART_BONDING_CONTRACTID` - valid entry relating to the product
* `SMART_BONDING_INSTALLSITEID` - valid entry relating to the product

> Note: these can be provided in a `.env` file.

As per the documentation, you will need to perform a pull operation in order to identify if the operation was a success or not.